package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type LLM interface {
	GenerateResponse(question string) (string, error)
}
type errMsg error

type Model struct {
	viewport  viewport.Model
	textarea  textarea.Model
	spinner   spinner.Model
	messages  []string
	err       error
	llm       LLM
	isLoading bool
	keymap    keyMap
	ready     bool
	maxWidth  int
}

func InitialModel(llm LLM, keyMapConfig KeyMapConfig) Model {
	ta := textarea.New()
	ta.Placeholder = "Ask a question..."
	ta.Focus()
	ta.ShowLineNumbers = false
	ta.Prompt = "Input: "
	ta.CharLimit = -1
	ta.SetHeight(1)

	vp := viewport.New(90, 20)

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := Model{
		textarea: ta,
		viewport: vp,
		spinner:  sp,
		llm:      llm,
		keymap:   newKeyMap(keyMapConfig),
		maxWidth: 90, // Set a default max width
	}
	// Add the welcome message
	welcomeMsg := "Olly: Hello! I am an expert Observability Programmer - Olly! Ask me any questions related to Observability."
	m.messages = append(m.messages, welcomeMsg)
	m.updateViewportContent()

	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, spinner.Tick)
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.isLoading {
				return m, nil
			}
			m.isLoading = true
			question := strings.TrimSpace(m.textarea.Value())
			if question == "" {
				m.isLoading = false
				return m, nil
			}
			m.messages = append(m.messages, "You: "+question)
			m.textarea.Reset()
			cmds = append(cmds, m.generateResponse(question))
			cmds = append(cmds, m.spinner.Tick)
			m.updateViewportContent()

		}

	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 4
		m.textarea.SetWidth(msg.Width)
		m.maxWidth = msg.Width
		if !m.ready {
			m.ready = true
		}
		m.updateViewportContent()

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case errMsg:
		m.err = msg
		m.isLoading = false
		return m, nil

	case string:
		m.isLoading = false
		m.messages = append(m.messages, "Olly: "+msg)
		m.updateViewportContent()
		return m, nil
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
func (m *Model) updateViewportContent() {
	var sb strings.Builder
	for i, msg := range m.messages {
		if i > 0 {
			sb.WriteString("\n\n") // Add extra line between messages
		}
		if strings.HasPrefix(msg, "You: ") {
			content := strings.TrimPrefix(msg, "You: ")
			rendered, _ := glamour.Render(content, "dark")
			sb.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Render("You: ") + rendered)
		} else if strings.HasPrefix(msg, "Olly: ") {
			content := strings.TrimPrefix(msg, "Olly: ")
			rendered, _ := glamour.Render(content, "dark")
			sb.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Render("Olly: ") + rendered)
		} else {
			sb.WriteString(msg)
		}
	}
	m.viewport.SetContent(sb.String())
	m.viewport.GotoBottom()
}
func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	var sb strings.Builder
	sb.WriteString(m.viewport.View())
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat("─", m.viewport.Width))
	sb.WriteString("\n")
	if m.isLoading {
		sb.WriteString("\n")
		sb.WriteString(m.spinner.View() + " Thinking...")
	} else {
		sb.WriteString(m.textarea.View())
	}
	return sb.String()
}

func (m Model) generateResponse(question string) tea.Cmd {
	return func() tea.Msg {
		answer, err := m.llm.GenerateResponse(question)
		if err != nil {
			return errMsg(err)
		}
		return answer
	}
}
