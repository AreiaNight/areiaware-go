# Panic Button - Educational Ransomware System

## About

This is an educational project, I created this because as a kid I wanted a bait for possible thiefs (in my country is usual) and make my computer unreachable. So this is a kind of ransomware based on that idea because this isn't for malicious purposes but a way for me to learn Go and make that idea I had as a little girl true.

**DISCLAIMER:** This project is purely educational and should only be used in controlled environments for learning purposes. Do not use this software on systems you don't own or without explicit permission.

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
- [ ] GUI interface
- [ ] Scheduled auto-lock functionality
- [ ] Email alert integration
- [ ] Screenshot capture on activation
