package main

import "fmt"

func main() {
	var age int = 19
	if !bas_bolunyar(age) {
		fmt.Println("salam")
	}
}

func bas_bolunyar(age int) bool {
	if age%5 == 0 {
		return true
	} else {
		return false
	}
}
