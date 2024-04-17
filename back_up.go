package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Error creating...", err)
		return
	}
	defer f.Close()

	currentTime := time.Now()
	backupFileName := fmt.Sprintf("file_backup_%s.txt", currentTime.Format("2006-01-02_150405"))

	_, err = os.Create(backupFileName)
	if err != nil {
		fmt.Println("Error creating file..", err)
		return
	}
	fmt.Println("Successfully created backup file")
}
