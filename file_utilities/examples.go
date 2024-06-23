package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//create a file
	fileName := "sampleFile.txt"
//	newFile, err:= os.Create(fileName)
//	if err != nil {
//		log.Fatal(err)
//	}

//	defer newFile.Close()

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("Size in bytes: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Last Modified: ", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Println("System interface type: %T\n", fileInfo.Sys())
	fmt.Println("System info: %+v\n\n", fileInfo.Sys())
}
