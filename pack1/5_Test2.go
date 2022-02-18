package pack1

import (
	"fmt"
	"strconv"
)

// 1. magic test
type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}
func (self *Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func MagicTest() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}

// 2. 类型 String() Test
type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {
	return strconv.Itoa(t.a) + " / " + fmt.Sprintf("%f", t.b) + " / " + t.c
}

func TypeString() {
	t := &T{7, -2.35, "abc\tqwe"}
	fmt.Println(t.String())
}

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%f", c) + "°C"
}

func CelsiusTest() {
	var c Celsius
	c = 36.5
	fmt.Println(c)
}

type Day int

func (d Day) String() string {
	switch fmt.Sprintf("%d", d) {
	case "1":
		return "Mon"
	case "2":
		return "Tue"
	case "3":
		return "Wen"
	case "4":
		return "Thu"
	case "5":
		return "Fri"
	case "6":
		return "Sat"
	case "7":
		return "Sun"
	default:
		return "Error"
	}
}

func Days() {
	var d Day = 1
	fmt.Println(d)
}

type Stack []int

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}
func (s *Stack) Pop() {
	if len(*s) == 0 {
		fmt.Println("Stack is empty...")
	} else {
		fmt.Println("Stack pop is ", (*s)[0])
		*s = (*s)[1:]
	}

}
func StackOperate() {
	s := &Stack{1}
	s.Push(2)
	s.Push(5)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
}
