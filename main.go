package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("new.txt")
	if err != nil {
		fmt.Println("Error creating file..", err)
		return
	}

	_, err = os.ReadFile("new.txt")
	if err != nil {
		fmt.Println("Error reading file..", err)
		return
	}
	defer f.Close()
	fmt.Print("Enter the text: ")
	user_input := bufio.NewReader(os.Stdin)
	str, err := user_input.ReadString('\n')
	if err != nil {
		fmt.Println("Error with reading string... ", err)
		return
	}
	f.WriteString(str)
	f.Seek(0, 0)

	f4, err := os.Create("copy_new.txt")
	if err != nil {
		fmt.Println("Wrong reading file...", err)
		return
	}
	defer f4.Close()
	_, err = io.Copy(f4, f)
	if err != nil {
		fmt.Println("Error copying...", err)
		return
	}
	fmt.Println("Saved successfully...")

}
