package main

import "fmt"

func main() {
	var int_arr =[9]int{1,2,8,1,8,9}  //让编译器根据实际的元素个数自行推断数组的长度 [1 2 8 1 8 9 0 0 0]
	var hobby = [2]string{} //数组
	var personnel = make([]string, 8)  //切片，长度为8
	hobby[0] = "basketball"
	hobby[1] = "football"

	c := map[string]int{"Hello":1, "World":2}
	for k, v := range c{
		fmt.Println(k, v)
	}

	//hobby[2] = "football"
	personnel[0] = "sunny"
	personnel[1] = "kind"
	fmt.Println(int_arr)
	fmt.Println(hobby)
	fmt.Println(personnel)
	for i, n := range hobby{
		fmt.Println(i, n)
	}
}