package main

import (
	"fmt"
	"log"
	"os"
	"golang.design/x/clipboard"
	"os/exec"
	"runtime"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    choices  []string
    cursor   int  
    selected map[int]struct{}
}
const linkedinUrl = "https://www.linkedin.com/in/harshkasat/"
const githubUrl = "https://github.com/harshkasat"
const twitterUrl = "https://twitter.com/harsh__kasat"
const portfolioUrl = "https://machoharsh-tech.vercel.app/"


func initialModel() model {
	return model{
		choices: []string{"Linkedin", "Twitter", "Github", "Portfolio", "CopyClipboard", "Url"},
		selected: make(map[int]struct{}),
	}
}


func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
		case "ctrl+c", "q":
        	return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
			_, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
				switch m.choices[m.cursor] {
				case "Linkedin":
					openBrowser(linkedinUrl)
				case "Twitter":
					openBrowser(twitterUrl)
				case "Github":
					openBrowser(githubUrl)
				case "Portfolio":
					openBrowser(portfolioUrl)
				case "CopyClipboard":
					copyClipboard()
				case "Url":
					redirectUrl()
				}
				delete(m.selected, m.cursor)

            }
            // Execute the function based on the selected option
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    // The header
    s := "Select an option:\n\n"

    // Iterate over our choices
    for i, choice := range m.choices {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected!
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}


func main() {
	p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

func copyClipboard(){
	if errors := clipboard.Init(); errors != nil {
		log.Fatal(errors)
	}

	var url = fmt.Sprintf("LinkedIn %v \nGithub %v \nTwiiter %v \nPortfolio %v",linkedinUrl, githubUrl, twitterUrl, portfolioUrl)

	clipboard.Write(clipboard.FmtText, []byte(url))
}

func redirectUrl(){
	var urls = []string{linkedinUrl, githubUrl, twitterUrl, portfolioUrl}
	for _, url := range urls{
		openBrowser(url)
		// time.Sleep(10*time.Second)
	}
}

func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{"-na", "Google Chrome", url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		fmt.Println("Unsupported platform")
		return
	}

	err := exec.Command(cmd, args...).Start()
	if err != nil {
		log.Fatalf("Failed to start browser: %w", err)
	}
}
