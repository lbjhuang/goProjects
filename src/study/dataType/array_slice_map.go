package main

import "fmt"

func main() {
	var int_arr = [9]int{1, 2, 8, 1, 8, 9} //数组的长度固定，后面的元素自动填充0 [1 2 8 1 8 9 0 0 0]
	fmt.Println(len(int_arr))
	for i, n := range int_arr {
		fmt.Println(i, n)
	}

	var int_arr1 =[...]int{1,2,8,1,8,9}  //让编译器根据实际的元素个数自行推断得出数组的长度 [1 2 8 1 8 9]
	fmt.Println(len(int_arr1))
	for i, n := range int_arr1 {
		fmt.Println(i, n)
	}



	var hobby = [2]string{}                //数组，长度为2，元素是字符串

	hobby[0] = "basketball"
	hobby[1] = "football"
	//hobby[2] = "football"  //超出界限，会报错

	c := map[string]int{"Hello": 1, "World": 2} //map类型
	for k, v := range c {
		fmt.Println(k, v)
	}
	fmt.Println(len(c))

	var personnel = make([]string, 8)      //切片，长度为8，元素是字符串
	personnel[0] = "sunny"
	personnel[1] = "kind"
	fmt.Println(len(personnel))  //切片长度
	fmt.Println(cap(personnel)) //切片容量

	fmt.Println(personnel)
	for i, n := range hobby {
		fmt.Println(i, n)
	}


	/*删除map中的元素*/
	delete(c, "Hello")  //删除Hello键
	fmt.Println(c)
	for k, v := range c {
		fmt.Println(k, v)
	}
}
