package main

import (
	"fmt"
)

/* interfaces */
type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (e *Rectangle) Area() float32 {
	return e.length * e.width
}

/* interface 例子 valuable */
type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

// stock position的value
func (sp *stockPosition) getValue() float32 {
	return sp.sharePrice * sp.count
}

type car struct {
	make  string
	model string
	price float32
}

// car的value
func (c *car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}


func main() {
	fmt.Println("====接口====")
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1 // 接口变量包含指向Square变量的引用 通过他调用Square上的Area()方法
	// shorter
	// areaIntf := Shaper(sq1)
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	// 多态性: 根据当前的类型选择正确的方法，或者说：同一种类型在不同的实例上似乎表现出不同的行为
	// 类型Rectangle也实现了Shaper接口 针对不同的实例：调用的Area()方法也不同
	r := &Rectangle{5, 3}
	q := &Square{5}
	// 创建Shaper类型的数组 迭代它的每一个元素并在上面调用 Area() 方法，以此来展示多态行为：
	shapes := []Shaper{r, q} // shapes := []Shaper{Shaper(r), Shaper(q)}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		// 在调用 shapes[n].Area()) 这个时，只知道 shapes[n] 是一个 Shaper 对象
		// 最后它摇身一变成为了一个 Square 或 Rectangle 对象，并且表现出了相对应的行为。
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	fmt.Println("====接口例子====")
	var vi valuable = &stockPosition{"GOOG", 577.20, 4.}
	showValue(vi)
	vi = &car{"BMW", "M3", 66500}
	showValue(vi)

	fmt.Println("====接口嵌套接口====")


}
