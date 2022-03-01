package main

import "fmt"

//指针类型接受者和值类型接受者
type User struct {
	name string
	email string
}

//1. 当接受者为指针类型时，可以通过该方法改变接受者成员变量的值，即使使用了非指针类型实例调用该函数，也可以改变实例对应的成员变量的值。
func (u *User) changeName() {  //指针类型接受者
	u.name = "james"
}

func main() {
	//u := &User{"perter", "1258158@qq.com"}          //创建指针类型结构体实例
	u := User{"perter","1258158@qq.com"}  //创建非指针类型结构体实例
	fmt.Println("name: ", u.name, "email: ", u.email)
	u.changeName()
	fmt.Println("name: ", u.name, "email: ", u.email)

}
//输出：
//name:  perter email:  1258158@qq.com
//name:  james email:  1258158@qq.com



//2. 值类型接受者：当接受者不是一个指针的时候，传过来的u是对应接受者值的副本，无法改变原来u的成员变量的值
//func (u User) changeName() {  //值类型接受者，里面的方法操作内容将直接影响到接受者
//	u.name = "james"
//}
//
//func main() {
//	u := &User{"perter","1258158@qq.com"}  //创建指针类型结构体实例
//	fmt.Println("name: ",u.name,"email: ",u.email)
//	u.changeName()
//	fmt.Println("name: ",u.name,"email: ",u.email)
//
//	//输出：
//	//name:  perter email:  1258158@qq.com
//	//name:  perter email:  1258158@qq.com
//}
