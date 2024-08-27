package main

import (
	"fmt"
	"os"

	"github.com/cloudraftio/olly/kubectl-olly/olly"
	"github.com/cloudraftio/olly/kubectl-olly/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	myLLM := &olly.MyLLM{}

	keyMapConfig := ui.KeyMapConfig{
		SwitchMultiline:      []string{"ctrl+m"},
		Submit:               []string{"enter"},
		Help:                 []string{"ctrl+h"},
		Quit:                 []string{"ctrl+c", "esc"},
		CopyLastAnswer:       []string{"ctrl+y"},
		PreviousQuestion:     []string{"up"},
		NextQuestion:         []string{"down"},
		NewConversation:      []string{"ctrl+n"},
		ForgetContext:        []string{"ctrl+f"},
		RemoveConversation:   []string{"ctrl+d"},
		PreviousConversation: []string{"ctrl+p"},
		NextConversation:     []string{"ctrl+x"},
		ScrollUp:             []string{"up"},
		ScrollDown:           []string{"down"},
	}

	p := tea.NewProgram(ui.InitialModel(myLLM, keyMapConfig), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/charmbracelet/bubbles/cursor"
// 	"github.com/charmbracelet/bubbles/textarea"
// 	"github.com/charmbracelet/bubbles/viewport"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// func main() {

// 	p := tea.NewProgram(initialModel())
// 	if _, err := p.Run(); err != nil {
// 		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
// 	}
// }

// type model struct {
// 	viewport    viewport.Model
// 	messages    []string
// 	textarea    textarea.Model
// 	senderStyle lipgloss.Style
// 	err         error
// }

// func initialModel() model {
// 	ta := textarea.New()
// 	ta.Placeholder = "Ask a question..."
// 	ta.Focus()
// 	ta.Prompt = "┃ "
// 	ta.CharLimit = 100
// 	ta.SetWidth(40)
// 	ta.SetHeight(3)
// 	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
// 	ta.ShowLineNumbers = false

// 	vp := viewport.New(60, 5)
// 	vp.SetContent(`Hello! I am an expert Observability Programmer - Olly!
// Ask me any questions related to Observability`)

// 	ta.KeyMap.InsertNewline.SetEnabled(false)

// 	return model{
// 		textarea:    ta,
// 		messages:    []string{},
// 		viewport:    vp,
// 		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
// 		err:         nil,
// 	}
// }

// func (m model) Init() tea.Cmd {
// 	return textarea.Blink
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		// Make the viewport and textarea dynamic based on the terminal size
// 		m.viewport.Width = msg.Width
// 		m.textarea.SetWidth(msg.Width)
// 		m.viewport.MouseWheelEnabled = true
// 		m.viewport.MouseWheelDelta = 20
// 		// Ensure content fits the new viewport size
// 		m.viewport.GotoBottom()
// 		return m, nil
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "esc", "ctrl+c":
// 			return m, tea.Quit
// 		case "enter":
// 			v := m.textarea.Value()
// 			if v == "" {
// 				return m, nil
// 			}

// 			m.messages = append(m.messages, m.senderStyle.Render("You: ")+v)
// 			m.viewport.SetContent(m.viewport.View() + "\n" + strings.Join(m.messages, "\n"))
// 			m.viewport.MouseWheelEnabled = true
// 			m.textarea.Reset()
// 			m.viewport.GotoBottom()

// 			answer, err := sendRequest(v)
// 			if err != nil {
// 				m.err = err
// 				return m, nil
// 			}

// 			m.messages = append(m.messages, m.senderStyle.Render("Olly: ")+answer)
// 			m.viewport.Width = m.viewport.Width + 0
// 			m.viewport.Height = m.viewport.Height + 1000
// 			m.viewport.MouseWheelEnabled = true
// 			m.viewport.SetContent(m.viewport.View() + "\n" + strings.Join(m.messages, "\n"))
// 			m.viewport.GotoBottom()
// 			return m, nil
// 		default:
// 			var cmd tea.Cmd
// 			m.textarea, cmd = m.textarea.Update(msg)
// 			return m, cmd
// 		}

// 	case cursor.BlinkMsg:
// 		var cmd tea.Cmd
// 		m.textarea, cmd = m.textarea.Update(msg)
// 		return m, cmd

// 	default:
// 		return m, nil
// 	}
// }

// func (m model) View() string {
// 	return fmt.Sprintf(
// 		"%s\n\n%s",
// 		m.viewport.View(),
// 		m.textarea.View(),
// 	) + "\n\n"
// }

// func sendRequest(question string) (string, error) {
// 	url := "https://ollybackend.ambitiousflower-4724f605.centralindia.azurecontainerapps.io/api/generate"
// 	payload := map[string]string{"question": question}
// 	jsonData, err := json.Marshal(payload)
// 	if err != nil {
// 		return "", err
// 	}

// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	var response map[string]string
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		return "", err
// 	}

// 	answer, ok := response["result"]
// 	if !ok {
// 		return "", fmt.Errorf("no answer found in response")
// 	}
// 	return answer, nil
// }

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/charmbracelet/bubbles/help"
// 	"github.com/charmbracelet/bubbles/spinner"
// 	"github.com/charmbracelet/bubbles/textarea"
// 	"github.com/charmbracelet/bubbles/viewport"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// var (
// 	Debug      bool
// 	DetachMode bool
// )

// type (
// 	saveMsg struct{}
// )

// func main() {

// 	p := tea.NewProgram(initialModel())
// 	if _, err := p.Run(); err != nil {
// 		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
// 	}
// }

// type model struct {
// 	width       int
// 	height      int
// 	help        help.Model
// 	answering   bool
// 	viewport    viewport.Model
// 	messages    []string
// 	textarea    textarea.Model
// 	senderStyle lipgloss.Style
// 	spin        spinner.Model
// 	err         error
// }

// func initialModel() model {
// 	ta := textarea.New()
// 	ta.Placeholder = "Ask a question..."
// 	ta.Focus()
// 	ta.Prompt = "┃ "
// 	ta.CharLimit = -1
// 	ta.SetWidth(50)
// 	ta.SetHeight(1)
// 	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
// 	ta.ShowLineNumbers = false

// 	vp := viewport.New(50, 5)
// 	// 	vp.SetContent(`Hello! I am an expert Observability Programmer - Olly!
// 	// Ask me any questions related to Observability`)

// 	ta.KeyMap.InsertNewline.SetEnabled(false)

// 	return model{
// 		textarea:    ta,
// 		messages:    []string{},
// 		viewport:    vp,
// 		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
// 		err:         nil,
// 	}
// }

// func (m model) Init() tea.Cmd {
// 	cmds := []tea.Cmd{tea.EnterAltScreen}
// 	if !Debug { // disable blink when debug
// 		cmds = append(cmds, textarea.Blink)
// 	}
// 	if !DetachMode {
// 		cmds = append(cmds, savePeriodically())
// 	}
// 	return tea.Batch(cmds...)
// }

// func savePeriodically() tea.Cmd {
// 	return tea.Tick(15*time.Second, func(time.Time) tea.Msg { return saveMsg{} })
// }
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

// 	var (
// 		cmd  tea.Cmd
// 		cmds []tea.Cmd
// 	)

// 	m.textarea, cmd = m.textarea.Update(msg)
// 	cmds = append(cmds, cmd)
// 	m.viewport, cmd = m.viewport.Update(msg)
// 	cmds = append(cmds, cmd)

// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.width = msg.Width
// 		m.height = msg.Height
// 		m.help.Width = msg.Width
// 		m.viewport.Width = msg.Width
// 		m.viewport.Height = msg.Height - m.textarea.Height()
// 		m.textarea.SetWidth(msg.Width)
// 		m.viewport.GotoBottom()
// 	case spinner.TickMsg:
// 		if m.answering {
// 			m.spin, cmd = m.spin.Update(msg)
// 		}
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "esc", "ctrl+c":
// 			return m, tea.Quit
// 		case "enter":
// 			v := m.textarea.Value()
// 			if v == "" {
// 				return m, nil
// 			}

// 			m.messages = append(m.messages, m.senderStyle.Render("You: ")+v)
// 			m.viewport.SetContent(m.viewport.View() + "\n" + strings.Join(m.messages, "\n"))
// 			m.textarea.Reset()
// 			m.viewport.GotoBottom()

// 			answer, err := sendRequest(v)
// 			if err != nil {
// 				m.err = err
// 				return m, nil
// 			}

// 			m.messages = append(m.messages, m.senderStyle.Render("Olly: "+answer))
// 			m.viewport.Width = m.viewport.Width + 0
// 			m.viewport.Height = m.viewport.Height + 1000
// 			m.viewport.MouseWheelEnabled = true
// 			m.viewport.MouseWheelDelta = 20
// 			m.viewport.SetContent(m.viewport.View() + "\n" + strings.Join(m.messages, "\n"))
// 			m.viewport.GotoBottom()
// 			return m, nil
// 		default:
// 			var cmd tea.Cmd
// 			m.textarea, cmd = m.textarea.Update(msg)
// 			return m, cmd
// 		}
// 	default:
// 		return m, nil
// 	}
// }

// func (m model) View() string {
// 	return fmt.Sprintf(
// 		"%s\n\n%s",
// 		m.viewport.View(),
// 		m.textarea.View(),
// 	) + "\n\n"
// }

// func sendRequest(question string) (string, error) {
// 	url := "https://ollybackend.ambitiousflower-4724f605.centralindia.azurecontainerapps.io/api/generate"
// 	payload := map[string]string{"question": question}

// 	jsonData, err := json.Marshal(payload)
// 	if err != nil {
// 		return "", err
// 	}

// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}

// 	var response map[string]string
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		return "", err
// 	}

// 	answer, ok := response["result"]
// 	if !ok {
// 		return "", fmt.Errorf("no answer found in response")
// 	}
// 	return answer, nil
// }
