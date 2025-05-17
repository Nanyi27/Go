package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包作用域
// 使用 括号 可以避免重复写 var
var (
	aa = 3
	bb = "sssss"
	cc = false
)

func variableZeroValue() {
	// 作用域 函数内
	var a int // 默认是 0
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	// Go语言非常严格，变量定义了就必须要使用
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// Go 语言类型自动推导，类似于C++中的auto
	var a, b, c, d = 3, 4, false, "def" // 不同类型变量可以一起声明
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, false, "def" // 可以使用 ：定义
	b = 5                            // 第二次需要使用 = 赋值
	fmt.Println(a, b, c, d)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f",
		cmplx.Exp(1i*math.Pi)+1)
	// cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println()
}

func triangle() {
	// 类型转换是强制的
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("Hello, world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	euler()
	triangle()
}
