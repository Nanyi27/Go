package main

import "fmt"
/*
	如果中断了iota那么必须显式的恢复， 后续会自动递增
	自增类型默认是int类型
	iota能简化const类型的定义
	每次出现const的时候， iota初始化为0
*/
func main1() {
	//iota, 特殊常量， 可以认为是一个可以被编译器修改的常量
	const (
		ERR1 = iota + 1 //iota内部仍然会增加计数器
		ERR2
		ERR25 = "ha"
		ERR3
		ERR31
		ERR32
		ERR33
		ERR34 = 100
		ERR4 = iota + 1
	)
	const (
		ERRNEW1 = iota
	)
	fmt.Println(ERR1, ERR2, ERR25, ERR3, ERR34, ERR4)
	fmt.Println(ERRNEW1)
}

func a() (int, bool){
	return 0, false
}

func main2(){
	//匿名变量, 变量的作用域
	var _ int
	_, ok := a()
	if ok{
		fmt.Println("it is ok")
	}
	fmt.Println(ok)
}

var name = "caididi" // 全局变量
func main(){
	// 变量的作用域
	mainName := "局部变量"
	fmt.Println(mainName)
	fmt.Printf("全局变量 %s\n", name)

	{
		fmt.Println(mainName)
		var inta int = 6
		fmt.Println(inta);
	}
	// fmt.Println(inta) // 拿不到 {} 局部变量
}