package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	fmt.Println("initializing godan...")
	createFiles("test1", "test2")

}

func createFiles(func1, func2 string) bool {
	f1, err := os.Create("funcs/f1.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	byteSlice := []byte(func1)
	bytesWritten, err := f1.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

	f2, err := os.Create("funcs/f2.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	byteSlice = []byte(func2)
	bytesWritten, err = f2.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

	return true
}
