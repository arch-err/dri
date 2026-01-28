package styles

import "github.com/charmbracelet/lipgloss"

var (
	AppStyle = lipgloss.NewStyle().Padding(1, 2)

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("63")).
			MarginBottom(1)

	StatusSuccess = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	StatusFailure = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	StatusRunning = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	StatusPending = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	StatusKilled  = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))

	ActiveTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("63")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1)

	InactiveTabStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("244")).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("238")).
				Padding(0, 1)

	TabGapStyle = lipgloss.NewStyle().
			Border(lipgloss.Border{Bottom: "─"}).
			BorderForeground(lipgloss.Color("238"))

	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))

	HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

func StatusIcon(status string) string {
	switch status {
	case "success":
		return StatusSuccess.Render("✓")
	case "failure", "error":
		return StatusFailure.Render("✗")
	case "running":
		return StatusRunning.Render("●")
	case "pending":
		return StatusPending.Render("○")
	case "killed":
		return StatusKilled.Render("✗")
	default:
		return StatusPending.Render("?")
	}
}

func StatusStyle(status string) lipgloss.Style {
	switch status {
	case "success":
		return StatusSuccess
	case "failure", "error":
		return StatusFailure
	case "running":
		return StatusRunning
	case "killed":
		return StatusKilled
	default:
		return StatusPending
	}
}
