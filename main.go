package main

import (
	"os"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("initializing godan...")


	func1 := `func func1() { 
		   var m map[string]string
		   k := string("x")
		   fmt.Println(m[k])
		   fmt.Println(m[k])
	   }

`

	   func2 := `func func2() {
		   var m map[string]string
		   fmt.Println(m[string("x")])
		   fmt.Println(m[string("x")])
	   }`
	createFiles(func1, func2)
	run()

}

func createFiles(func1, func2 string) bool {
	f1, err := os.Create("funcs/f1.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	header := `package main

import (
  "fmt"
  "time"
)

func main() {
   start := time.Now()
   func1()
   elapsed := time.Since(start)
   fmt.Printf("func1: %v\n", elapsed)
   start = time.Now()
   func2()
   elapsed = time.Since(start)
   fmt.Printf("func2: %v\n", elapsed)
}


`
	byteSlice := []byte(header + func1 + func2)
	bytesWritten, err := f1.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

	return true
}

func run() {
	fmt.Println("running...")
	out, err := exec.Command("go", "run", "funcs/f1.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %s\n", out)
}
