# Panic Button - Educational Ransomware System

![panic](https://github.com/AreiaNight/panicButton-A-go-project/blob/main/panic.png)

## About

This is an educational project, I created this because as a kid I wanted a bait for possible thiefs (in my country is usual) and make my computer unreachable. So this is a kind of ransomware based on that idea because this isn't for malicious purposes but a way for me to learn Go and make that idea I had as a little girl true.

**DISCLAIMER:** This project is purely educational and should only be used in controlled environments for learning purposes. Do not use this software on systems you don't own or without explicit permission.

## Version 2.5

This version continues evolving the project by improving **alerting, configurability, and command management**, while keeping the educational focus.

### 1. Hermes Module (Discord Bot Integration)

- **Dedicated Alert Module:**  
  A new module named `hermes` has been introduced, responsible for handling all alerting and external communications.

- **Custom Discord Bot Support:**  
  Instead of relying solely on basic webhooks, the system can now connect to a **custom Discord bot**, allowing richer and more controlled alert messages.

- **Real-Time Notifications:**  
  Alerts are sent automatically when critical actions occur (e.g., encryption start, unlock attempts).

### 2. Configuration System

- **JSON Credentials File:**  
  Sensitive data such as Discord bot tokens, channel IDs, and webhook-related information are now stored in a configurable JSON file.

- **YAML Command Configuration:**  
  A YAML file is used to define and manage commands, making the system easier to extend and customize without modifying the source code.

- **Separation of Logic and Configuration:**  
  This change improves maintainability and allows safer experimentation when adding new features.

  ### 3. Cross-Platform Support

- **Multi-OS Compatibility:**  
  Starting with version 2.5, the core functionality is now compatible with:
  - Linux
  - Windows
  - macOS

---


## Features

- AES-256-GCM encryption for secure file locking
- Automatic file hiding with random filenames
- Recovery password system inspired by Castlevania characters
- Encrypted mapping system to track original file names
- Discord alert system (webhook and bot-based)
- Native Windows MessageBox notifications
- Full file restoration capabilities
- Modular alerting via the Hermes module
- Configurable credentials (JSON) and commands (YAML)


## Technical Stack

- **Language:** Go 1.21+
- **Encryption:** AES-256-GCM
- **Key Derivation:** SHA-256
- **Configuration:** JSON / YAML
- **Platform:** Windows (with potential cross-platform support)
- **Alerting:** Discord Bot & Webhookst)

## Prerequisites

- Go 1.21 or higher
- Windows OS (for MessageBox features)
- Discord Bot Token (for Hermes module)

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

- [x] Discord allert
- [ ] Extended Discord notifications:
    - [ ] User who executed the binary
    - [ ] Execution date and time
    - [ ] Hostname and OS information
    - [ ] Basic hardware and system metadata
- [x] GUI interface
- [x] Auto-lock functionality
- [ ] Email alert integration
- [ ] Screenshot capture on activation
- [ ] Cross-platform support refinements

## Security & Ethics

This project is designed with a strong emphasis on **ethical use, transparency, and controlled experimentation**.

### Educational Purpose Only

- This software was created exclusively for **learning and research purposes**, particularly to understand:
  - File encryption mechanisms
  - Ransomware behavioral patterns
  - Defensive alerting and monitoring techniques
  - Secure coding practices in Go

- It is **not intended for real-world deployment** against unsuspecting users or systems.

### Ownership and Consent

- The software must only be executed on:
  - Systems you personally own, or
  - Systems where you have **explicit and informed permission** from the owner.

- Running this project on third-party systems without consent may be illegal and unethical.

### No Financial Extortion

- This project does **not** implement:
  - Payment mechanisms
  - Cryptocurrency wallets
  - Extortion logic or demands

- Any resemblance to ransomware behavior exists solely to study its technical structure, not to replicate criminal activity.

### Controlled Design Decisions

- Encryption is reversible by design through:
  - A hardcoded Master Password
  - A temporary Recovery Password generated at runtime

- These safeguards exist to prevent irreversible data loss during testing and learning.

### Responsible Disclosure and Usage

- Users are encouraged to:
  - Study the code to better understand offensive techniques from a defensive standpoint
  - Use the knowledge gained to improve detection, response, and prevention strategies

- Any modifications that remove safeguards or add malicious intent are **explicitly discouraged**.

### Legal Notice

- The author assumes no responsibility for misuse of this software.
- The user is fully responsible for ensuring compliance with local laws and regulations.

---

By using or modifying this project, you acknowledge that you understand the ethical implications of ransomware-like software and agree to use it responsibly.
