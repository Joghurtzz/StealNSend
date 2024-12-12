# StealNSend
StealNSend is a PoC program designed to extract credentials from various web browser and software, take screenshot and send them to a FTP server. As of 3 December 2024, this program is undetectable by common antivirus software.

*This code is designed for Windows environments due to the use of platform-specific paths and environment variables.*

*Watch out! The FTP server IP and credential are stored in cleartext and the connection is not encrypted.*

## Description
This program allows you to identify and upload files from specific directories and a screenshot to a remote FTP server. It is designed to automate data transfer from predefined locations, such as browser data or application configurations, using FTP connections.

### Key Features
- **Recursive Directory Scanning** - Identifies and collects all files within specified directories while ignoring subdirectories
- **Screenshot** - Takes a screenshot the moment the program is started
- **FTP Server Connection** - Manages connection to a remote FTP server with configurable credentials (host, port, username, password)
- **Modular Organization** - The code is structured into separate files for better readability and maintainability: `main.go`: Coordinates the overall workflow. `utility.go`: Defines target directories, handles FTP connection, file uploads and other.

## Requirements
- **Go** - Version 1.18+
- **Dependencies**
```
github.com/jlaffaye/ftp
github.com/kbinani/screenshot
```

## How To Use
- **Clone the Repository** - `git clone https://github.com/Joghurtzz`
- **Configure the FTP Server parameters**
```
ftpServer := "<FTP_SERVER>"
ftpPort := <PORT>
ftpUser := "<USERNAME>"
ftpPassword := "<PASSWORD>"
```
- **Inizialize Go module** -
```
go mod init <PROJECT_NAME>
go mod tidy
```
- **Build the Program** - `GOOS=windows GOARCH=amd64 go build -o <OUTPUT_NAME>.exe main.go utility.go`

## Future improvements
- Upload files to their respective folders (i.e. Chrome files are uploaded to the Chrome folder)
- Get hardware info and network info (CPU, GPU, MAC, IP, ProductKey, saved WiFi connections...)
- Use a more secure way to upload data

## Disclaimer
This program is intended strictly for educational and research purposes. Any use of StealNSend for unauthorized access, malicious activities, or data theft is strictly prohibited. The author assumes **no** responsability for misuse of this program.

**This project is currently under development.**
