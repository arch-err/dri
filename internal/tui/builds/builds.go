package builds

import (
	"fmt"
	"strings"
	"time"

	"github.com/arch-err/dri/internal/tui/msg"
	"github.com/arch-err/dri/internal/tui/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/drone/drone-go/drone"
)

type buildItem struct {
	build *drone.Build
}

func (i buildItem) Title() string {
	return fmt.Sprintf("%s #%d %s", styles.StatusIcon(i.build.Status), i.build.Number, i.build.Message)
}

func (i buildItem) FilterValue() string {
	return fmt.Sprintf("#%d %s %s %s %s",
		i.build.Number,
		i.build.Status,
		i.build.Message,
		i.build.Event,
		i.build.Target,
	)
}

func (i buildItem) Description() string {
	parts := []string{
		i.build.Event,
		i.build.Target,
		i.build.Author,
	}
	if i.build.Finished > 0 {
		parts = append(parts, timeAgo(i.build.Finished))
	} else if i.build.Started > 0 {
		parts = append(parts, "started "+timeAgo(i.build.Started))
	}
	return strings.Join(parts, " | ")
}

type Model struct {
	list list.Model
}

func New(buildList []*drone.Build, repoSlug string, width, height int) Model {
	items := make([]list.Item, len(buildList))
	for i, b := range buildList {
		items[i] = buildItem{build: b}
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = fmt.Sprintf("Builds — %s", repoSlug)
	l.DisableQuitKeybindings()
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)

	return Model{list: l}
}

func (m Model) Update(msgin tea.Msg) (Model, tea.Cmd) {
	if kmsg, ok := msgin.(tea.KeyMsg); ok && kmsg.String() == "enter" {
		if !m.IsFiltering() {
			if item, ok := m.list.SelectedItem().(buildItem); ok {
				return m, func() tea.Msg {
					return msg.BuildSelectedMsg{Build: item.build}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msgin)
	return m, cmd
}

func (m Model) View() string {
	return m.list.View()
}

func (m Model) IsFiltering() bool {
	return m.list.FilterState() == list.Filtering
}

func (m *Model) SetSize(w, h int) {
	m.list.SetSize(w, h)
}

func timeAgo(unix int64) string {
	t := time.Unix(unix, 0)
	d := time.Since(t)

	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		m := int(d.Minutes())
		if m == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", m)
	case d < 24*time.Hour:
		h := int(d.Hours())
		if h == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", h)
	default:
		days := int(d.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}
}
