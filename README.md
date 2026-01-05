
<img width="1769" height="1077" alt="image 1" src="https://github.com/user-attachments/assets/7320d65f-f550-4a8a-b8f3-6d74cf3ae28c" />

## About the project

Hello to everyone who reads this text!

This is a **small educational project** created to explore the basic syntax of the Go programming language and several parts of the Go standard library.

The project was designed with **simplicity in mind** and intentionally avoids unnecessary abstractions or advanced CLI features.

> This repository is not intended to be a production-ready tool.  
> Its main purpose is **learning and experimentation**.

At the moment, **further development is not planned**.

---

## Requirements

To use this project, you need **one** of the following:

- **Go installed on your system [Download Golang](https://go.dev/dl/)**
- **A precompiled binary for your operating system from [GitHub Releases](https://github.com/mrdyuke/deplify/releases)**

You can verify that Go is installed by running:
`go version`

---

## Installation

### Using Go

If Go is installed, install the project directly from the repository:
`go install github.com/mrdyuke/deplify@latest`

The binary will be installed to `$HOME/go/bin/deplify`.

#### To run the program:

**Option 1 (recommended) - Add to PATH:**
`export PATH="$HOME/go/bin:$PATH"`  
*(add to ~/.bashrc, ~/.zshrc, etc. for persistence)*

**Option 2 - Run with full path:**
`~/go/bin/deplify`

**Option 3 - Copy to system binaries:**
`sudo cp ~/go/bin/deplify /usr/local/bin/`

After setting up the PATH, you can run:
`deplify --help`

---

### Using a precompiled binary

If you do not want to install Go:

1. Download the binary for your operating system from the **[GitHub Releases](https://github.com/mrdyuke/deplify/releases)** page
2. Rename the file if necessary
3. Add the binary to your system `PATH`

Once this is done, the program can be executed directly:
`deplify`

---

## Notes

* The project is created strictly for **educational purposes**
* The codebase is intentionally minimal
* No additional setup or configuration is required
