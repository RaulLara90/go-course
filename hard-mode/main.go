package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	buf := make([]byte, 99999)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Print(string(buf[:n]))
	}

}
