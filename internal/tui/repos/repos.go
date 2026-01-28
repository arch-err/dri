package repos

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

type repoItem struct {
	repo *drone.Repo
}

func (i repoItem) Title() string       { return i.repo.Slug }
func (i repoItem) FilterValue() string { return i.repo.Slug }
func (i repoItem) Description() string {
	if !i.repo.Active {
		return "inactive"
	}
	b := i.repo.Build
	if b.Number == 0 {
		return "no builds"
	}
	parts := []string{
		fmt.Sprintf("%s #%d", styles.StatusIcon(b.Status), b.Number),
	}
	if b.Finished > 0 {
		parts = append(parts, timeAgo(b.Finished))
	}
	return strings.Join(parts, " · ")
}

type Model struct {
	list list.Model
}

func New(repos []*drone.Repo, width, height int) Model {
	items := make([]list.Item, len(repos))
	for i, r := range repos {
		items[i] = repoItem{repo: r}
	}

	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "Repositories"
	l.DisableQuitKeybindings()
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)

	return Model{list: l}
}

func (m Model) Update(msgin tea.Msg) (Model, tea.Cmd) {
	if kmsg, ok := msgin.(tea.KeyMsg); ok && kmsg.String() == "enter" {
		if !m.IsFiltering() {
			if item, ok := m.list.SelectedItem().(repoItem); ok {
				return m, func() tea.Msg {
					return msg.RepoSelectedMsg{Repo: item.repo}
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
