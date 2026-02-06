package uploads

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func copyExactFiles(srcDir string, destDir string) error{
	srcFile, err := os.Open(srcDir)
	if err != nil{
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(destDir)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
    }

	defer destFile.Close()


	// copyinh the file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
    }

	// flushing file metadata to disk
	err = destFile.Sync()
	if err != nil{
		return fmt.Errorf("failed to sync destination file: %w", err)
	}

	return  nil
}


func MoveFiles(){

	homePath, _ := os.UserHomeDir()

	srcDir := filep ath.Join(homePath, "uploadtesting/src")
	destDir := filepath.Join(homePath, "uploadtesting/dest")

	// copyExactFiles(srcDir, destDir)

	if _, err := os.Stat(destDir); os.IsNotExist(err){
		err := os.Mkdir(destDir, os.ModePerm)
		if err != nil{
			fmt.Printf("Error while creating destination directory: %v\n", err)
			log.Fatal(err)
			return
		}
	}

	var filesToCopy []string

	files, err := os.ReadDir(srcDir)

	if err!=nil{
		log.Fatal(err)
		return
	}

	for _, file := range files{
		// fileName := file.Name()
		filesToCopy = append(filesToCopy, file.Name())
	}

	fmt.Println(filesToCopy)

	for _, file := range filesToCopy{
		srcPath := filepath.Join(srcDir, file)
		destPath := filepath.Join(destDir, file)

		err = copyExactFiles(srcPath, destPath)
		if err != nil{
			fmt.Printf("Error copying file %s: %v\n", file, err)
		} else {
			fmt.Printf("Successfully copied %s to %s\n", file, destPath)
		}
	}
	
}


