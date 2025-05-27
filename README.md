# ProfileCli-Go

A simple command-line interface (CLI) application written in Go that allows users to quickly access the developer's LinkedIn, GitHub, Twitter, and portfolio website.  It also provides the option to copy all URLs to the clipboard.

[![Go](https://github.com/harshkasat/ProfileCli-Go/actions/workflows/go.yml/badge.svg)](https://github.com/harshkasat/ProfileCli-Go/actions/workflows/go.yml)
[![License](https://img.shields.io/github/license/harshkasat/ProfileCli-Go)](https://github.com/harshkasat/ProfileCli-Go/blob/main/LICENSE)


## Project Overview

This project aims to provide a convenient and efficient way to access the developer's online profiles.  Instead of manually typing URLs, users can interact with a simple menu-driven CLI to open the desired profile in their default web browser or copy all URLs to their clipboard.  This is particularly useful for quickly sharing contact information.

## Table of Contents

* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Usage](#usage)
* [Code Overview](#code-overview)
* [Contributing](#contributing)
* [License](#license)


## Prerequisites

* Go 1.18 or higher
* A compatible terminal
* `golang.design/x/clipboard` library (installed during installation)
* A default web browser configured on your system.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/harshkasat/ProfileCli-Go.git
   cd ProfileCli-Go
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run main.go
   ```


## Usage

After running the application, a menu will appear in your terminal:

```
Select an option:

 > [ ] Linkedin
  [ ] Twitter
  [ ] Github
  [ ] Portfolio
  [ ] CopyClipboard
  [ ] Url
Press q to quit.
```

Use the up/down arrow keys or 'k'/'j' to navigate. Press Enter or space to select an option.  "CopyClipboard" will copy all URLs to your clipboard. "Url" will open all URLs sequentially in your browser.  Press 'q' or 'ctrl+c' to quit.


## Code Overview

The core logic resides in `main.go`.  The application uses the Bubble Tea framework for the UI.

**Key functions:**

* `initialModel()`: Creates the initial application state.
* `Update(msg tea.Msg)`: Handles user input (key presses).  It opens the browser using `openBrowser()` based on the selected option, or copies URLs to the clipboard using `copyClipboard()`.
* `View()`: Renders the menu to the terminal.
* `copyClipboard()`: Copies all URLs to the clipboard.
* `redirectUrl()`: Opens all URLs sequentially in the browser.
* `openBrowser(url string)`: Opens the given URL in the default web browser, handling different operating systems.

**Example (copyClipboard function):**

```go
func copyClipboard(){
    if errors := clipboard.Init(); errors != nil {
        log.Fatal(errors)
    }

    var url = fmt.Sprintf("LinkedIn %v \nGithub %v \nTwiiter %v \nPortfolio %v",linkedinUrl, githubUrl, twitterUrl, portfolioUrl)

    clipboard.Write(clipboard.FmtText, []byte(url))
}
```


## Contributing

Contributions are welcome! Please open an issue or submit a pull request.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
