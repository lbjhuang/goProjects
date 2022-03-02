package main

import "fmt"

func main() {


	var personnel = make([]string, 8) //切片，长度为8
	var hobby = make([]string, 5, 6)  //切片，长度为5 容量为6

	//hobby[2] = "football"
	personnel[0] = "sunny"
	personnel[1] = "kind"
	hobby[0] = "kind"
	hobby[1] = "kind"
	hobby[2] = "kind"
	hobby[3] = "kind"
	//hobby[4] = "kind2"   //初始第5个元素为空，后面追加则从第6个元素开始
	hobby = append(hobby, "man")  //追加元素 长度增加，容量自动扩充
	hobby = append(hobby, "man")  
	hobby = append(hobby, "female")  //追加元素 长度增加，容量自动扩充

	fmt.Println(hobby)
	fmt.Println(len(hobby))  //里面元素格式
	fmt.Println(cap(hobby))  //里面元素格式


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


	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)  //空切片不等于nil

}
