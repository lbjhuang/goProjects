package main

import "fmt"

func main() {

	var personnel = make([]string, 8) //切片，长度为8
	var hobby = make([]string, 5, 6)  //切片，长度为4 容量为6

	//hobby[2] = "football"
	personnel[0] = "sunny"
	personnel[1] = "kind"
	hobby[0] = "kind"
	hobby[1] = "kind"
	hobby[2] = "kind"
	hobby[3] = "kind"
	hobby = append(hobby, "man")  //追加元素  长度增加，容量自动扩充
	hobby = append(hobby, "man")  //追加元素  长度增加，容量自动扩充


	fmt.Println(hobby)
	fmt.Println(personnel)
	fmt.Println(len(personnel))
	fmt.Println(cap(personnel))
	fmt.Println(personnel[1:2])

	var int_ar = []int{1, 8, 9, 8, 9, 6}

	//数组切片
	fmt.Println(int_ar[0:2])
	fmt.Println(int_ar[:2])
	fmt.Println(int_ar[2:])

	fmt.Println(len(int_ar[:2]))
	fmt.Println(cap(int_ar[:2]))

	fmt.Printf("%T", int_ar[0:2])

}
