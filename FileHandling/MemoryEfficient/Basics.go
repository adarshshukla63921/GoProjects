package main

import (
	"bufio" // for efficient reading and writing
	"fmt"
	"os"
	"time"
)

func readLineByLine(filePath string) {
	// you open the file you want to read
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file.")
		return
	}
	// make sure to close the file before leaving
	defer file.Close()

	// you create the a new scanner to read the file
	// Scanner provides efficient reading
	scanner := bufio.NewScanner(file)
	lineNum :=1

	// you read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(lineNum,line)
		lineNum++
	}
}

func writeLineByLine(fileName string){

	// you create a file
	file, err := os.Create(fileName)

	// you check there was no issue creating the file
	if err!=nil {
		fmt.Println("Error :",err)
		return
	}

	// make sure you always close the file.
	defer file.Close()
	
    // you get the content you want to put in the file
	lines := []string{
		"line 1 \n",
		"line 2 \n",
		"line 3 \n",
		"line 4 \n",
		"line 5 \n"}
	
	// you get a Writer
	// again , using bufio because it improves performance for large files.
	writer := bufio.NewWriter(file)
	
	// you write the content in the file using the writer we created eariler
	for _, line := range lines {
		_,err = writer.WriteString(line)
		if(err!=nil){
			fmt.Println("some error happended")
			return
		}
	}
	
	// you make sure to flush the file, to avoid the case where data stays in memory
	err = writer.Flush()
	if(err != nil){
		fmt.Println("Error : ",err)
		return
	}
}

func appendToFile(fileName string){

	// you open the file you want to append to.
	// os.O_APPEND , os.O_WRONLY [write only] , os.O_CREATE are flags
	// 0644 is a permission 
	file,err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE,0644)
	if err != nil {
		fmt.Println("Error :",err)
		return
	}
	// close the file before leaving
	defer file.Close()

	// getting something to append in our file
	timeStamp := time.Now()
	logEntry := fmt.Sprintln(timeStamp) + "\n"

	// appending something in our file
	_ , err = file.WriteString(logEntry)
	if err != nil {
		fmt.Println("error : ",err)
		return
	}
}

func deleteFile(fileName string){
	err := os.Remove(fileName)
	if err!=nil{
		fmt.Println("unable to remove the file :",err)
		return
	}
}
func main() {
	//readLineByLine("example.txt")
	writeLineByLine("bufferedIo.txt")
	appendToFile("bufferedIo.txt")
	// ensure correct removal of file.
	deleteFile("remove.txt")
}