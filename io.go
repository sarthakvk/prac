package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	if scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println(scanner.Bytes())
		fmt.Println(scanner.Text())

	}
	
}