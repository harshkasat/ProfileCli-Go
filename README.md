# Go CLI Application

This is a simple Go-based CLI application that uses the Charm library's `tea` model to create an interactive command-line interface. The application includes functionalities such as opening URLs in a web browser, copying URLs to the clipboard, and more.

```bash 
    /app
    │   main.go          # The main Go application file
    │   go.mod           # Go module file
    │   go.sum           # Go dependencies file
    │   Dockerfile       # Dockerfile for building the image
    │   docker-compose.yml  # Docker Compose configuration
    └─── README.md       # This README file
```


## Features

- Interactive CLI for selecting and running different commands
- Open URLs in the default web browser
- Copy multiple URLs to the clipboard
- Cross-platform support (Windows, macOS, Linux)

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/your-repo.git
cd your-repo
