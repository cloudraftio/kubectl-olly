package main

import (
	"fmt"
	"os"

	"github.com/cloudraftio/kubectl-olly/olly"
	"github.com/cloudraftio/kubectl-olly/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	if olly.GetAPIKey() == "" {
		fmt.Println("Error: OLLY_API_KEY not set. Please set it in your environment or in $HOME/.olly file.")
		os.Exit(1)
	}

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
