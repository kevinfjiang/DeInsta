package main

import "fmt"



func main() {
	vote:=10
	// vote, present := map[int64]bool{-1: true, 0: true, 1: true}[vote]
	if vt, present := map[int]bool{-1: true, 0: true, 1: true}[vote]; !present{
		fmt.Println(vt)
	}
	fmt.Println("valid")
}

