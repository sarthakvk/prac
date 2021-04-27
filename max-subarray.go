package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err  := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to accept input")
		return
	}

	arr := SliceAtoi(strings.Fields(str))
	best, sum := 0, 0

	for i := range arr {
		sum = max(arr[i], sum + arr[i])
		best = max(best, sum)
	}

	fmt.Println(best)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func SliceAtoi(str []string) []int {
	si := make([]int, 0, len(str))

	for _, v := range str {
		i, _ := strconv.Atoi(v)
		si = append(si, i)
	}
	return si
}