# Panic Button - Educational Ransomware System

![panic](https://github.com/AreiaNight/panicButton-A-go-project/blob/main/panic.png)

## About

This is an educational project, I created this because as a kid I wanted a bait for possible thiefs (in my country is usual) and make my computer unreachable. So this is a kind of ransomware based on that idea because this isn't for malicious purposes but a way for me to learn Go and make that idea I had as a little girl true.

**DISCLAIMER:** This project is purely educational and should only be used in controlled environments for learning purposes. Do not use this software on systems you don't own or without explicit permission.

## Version 2.0

This new version introduces significant enhancements in structure, security, and user experience:

1. Clean and Modular Architecture

    Organized Codebase: The project now follows best practices for Go development. Core functions for encryption, file traversal, and utilities have been moved to the artefacts/ folder.

    Logical Modules: Responsibilities are clearly separated into distinct packages (e.g., atenea, cerberos, hecate), making the code easier to read, maintain, and expand.

2. Recovery Password Feature

    Dual Security Access: The underlying encryption still relies on the fixed Master Password. However, the decryption process now accepts either the hardcoded Master Password or the new temporary Recovery Password.

    Automatic Generation: Upon initial execution (go run .), a unique, temporary Recovery Password is automatically generated and displayed to the user.

    Single-Use Safety: Once the Recovery Password is used to unlock the files, it is designed to be discarded.

3. Simplified Execution

    Automatic Encryption: Running the program without any command-line arguments (using go run .) will now automatically initiate the file encryption process.

    This flow includes generating and displaying the Recovery Password before locking the files.

4. New Visual Warning

    Native Pop-up Window: A new native operating system pop-up window is displayed to warn the user (or intruder) immediately before the encryption process begins.

## Features

- AES-256-GCM encryption for secure file encryption
- Automatic file hiding with random names
- Recovery password system based on Castlevania characters
- Encrypted mapping system to track original file names
- Discord webhook integration for security alerts
- Windows MessageBox notifications
- Complete file restoration capabilities


## Technical Stack

- **Language:** Go 1.21+
- **Encryption:** AES-256-GCM (Advanced Encryption Standard)
- **Key Derivation:** SHA-256
- **Platform:** Windows (with potential cross-platform support)

## Prerequisites

- Go 1.21 or higher
- Windows OS (for MessageBox features)

## Installation

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/panic-button.git
cd panic-button
```

### 2. Initialize Go module
```bash
go mod init panic-button
```

### 3. Install dependencies
```bash
go get golang.org/x/term
go get golang.org/x/sys/windows
```

### 4. Build the project
```bash
go build -o panic-button.exe
```

Or compile with size optimization:
```bash
go build -ldflags "-s -w" -o panic-button.exe
```

## Usage

### Basic Commands
```bash
# Show help menu
panic-button.exe h

# Show about information
panic-button.exe a

# Encrypt files (lock)
panic-button.exe l

# Decrypt files (unlock)
panic-button.exe u

# Generate recovery password
panic-button.exe g

# Test password input
panic-button.exe gp
```

## Roadmap

Potential future features:

- [ ] Discord allert using webhook
- [x] GUI interface
- [x] Auto-lock functionality
- [ ] Email alert integration
- [ ] Screenshot capture on activation
