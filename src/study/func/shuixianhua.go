package main

import "fmt"

//水仙花数计算
func shuixianhua(n int) bool {
	if n < 100 {
		return false
	}
	var flag = false
	one := n % 10
	two := (n / 10) % 10
	three := n / 100
	if one*one*one+two*two*two+three*three*three == n {
		flag = true
	}
	return flag
}

func main() {
	fmt.Println(118 / 100) //1 商
	fmt.Println(118 % 100) //18 余

	for i := 100; i < 1000; i++ {
		if shuixianhua(i) == true {
			fmt.Println(i)
		}
	}
}
