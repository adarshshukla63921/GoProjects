package main

import (
	"fmt"
	"os"
)

func readFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "Error reading file:" + fmt.Sprint(err)
	}

	return string(file)
}

func writeFile(){
	content := []byte("Hello, Go!\nThis is a new file.")
    // Create/overwrite file with permissions 0644
    err := os.WriteFile("output.txt", content, 0644)
    if err != nil {
        panic(err)
    }
}

func checkFileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return  err==nil
}

func removeFile(fileName string) {
	err := os.Remove(fileName)

	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
	fmt.Println("File removed successfully")
}

func main() {
	writeFile()
	fmt.Println(readFile("output.txt"))
	fmt.Println("Does output.txt exist?", checkFileExists("output.txt"))
	//removeFile("unwanted.txt")
}