package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Creating reader for taking input
	// as a string which have space seprated
	// Integers
	fmt.Println("Enter space seperated integers to sort: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to accept input")
		return
	}

	// Converting string to slice of
	// unsorted Integers, if the string
	// contains not integer then it will
	// cause panic
	arr := SliceAtoi(strings.Fields(str))

	// Calling MergeSort
	// wich takes slice and staring and
	// end of the slice indexes
	MergeSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

// MergeSort Implementation
func MergeSort(arr []int, l, r int) {
	// l is index of left side; r is index of right side
	// If r == l, then sub array is alredy sorted i.e it has only one element
	// otherwise split array in half and call MergeSort on both half then
	// merge those sorted halfs
	if r > l {
		m := (l + r) / 2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)

		merge(arr, l, m, r)
	}

}

func merge(arr []int, l, m, r int) {
	// i & j are the indexes used to walk the copy of left and right subarrays
	// while k is the starting index of the orignal array
	i, j, k := 0, 0, l

	// Copy of left and right subarray
	L := make([]int, m-l+1)
	R := make([]int, r-m)

	copy(L, arr[l:m+1])
	copy(R, arr[m+1:r+1])

	// Merging left and right sorted sub arrays
	for i < len(L) && j < len(R) {
		switch {
		case L[i] <= R[j]:
			arr[k] = L[i]
			i++
			k++
		case L[i] > R[j]:
			arr[k] = R[j]
			j++
			k++
		}
	}

	// added the remains of the Left/Right
	// Subarrays
	for i < len(L) {
		arr[k] = L[i]
		i++
		k++
	}

	for j < len(R) {
		arr[k] = R[j]
		j++
		k++
	}

}

// Converts String of integer seperated by spcae to a
// Slice of intgers for formatting input
func SliceAtoi(str []string) []int {
	si := make([]int, 0, len(str))

	for _, v := range str {
		i, _ := strconv.Atoi(v)
		si = append(si, i)
	}
	return si
}
