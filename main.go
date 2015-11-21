package main

import "os"
import "fmt"

func main() {
	progName := os.Args[0]
	extraArgs := os.Args[1:]

	fmt.Println("Why hello there!")
	fmt.Println(progName)
	fmt.Println(extraArgs)
}
