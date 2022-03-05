package main

import "fmt"

//质数练习
func zhisuh(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func main() {
	for i := 2; i < 100; i++ {
		if zhisuh(i) == true {
			fmt.Println(i)
		}
	}
}
