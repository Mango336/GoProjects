package pack1

import (
	"fmt"
)

type VCard struct {
	name  string
	Addr  *Address // 使用指针类型 占用内存少
	birth string
	image int
}

type Address struct {
	Addrs string
}

func VCardFunc() *VCard {
	vc := new(VCard)
	vc.name = "ylm"
	vc.Addr = &Address{
		"abc",
	}
	vc.birth = "1998-03-23"
	vc.image = 1
	return vc
}

type ExpStruct struct {
	Mi1 int
	Mf1 float32
}

type Employee struct {
	Salary float32
}

func (e Employee) GiveRaise(f float32) float32 {
	return e.Salary * (1 + f)
}

type B struct {
	thing int
}

func (b *B) change() { // 接收指向B的指针 并改变它内部的成员
	b.thing = 1
}
func (b B) write() string { // 接收通过拷贝B的值 并只输出B的内容
	return fmt.Sprint(b)
}

func Pointer_value() {
	// 指针方法和值方法都可以在指针或非指针上被调用
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())
	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())
}

type List []int

// 类型List在值上有一个方法Len() 在指针上有方法Append()
// 但是这两个方法都可以在两种类型的变量上被调用

func (l List) Len() int {
	return len(l)
}
func (l *List) Append(val int) {
	*l = append(*l, val)
}
func Print_List() {
	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len()) // [1] (len: 1)
	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)\n", plst, plst.Len()) // &[2] (len: 1)
}

// Person 被导出 但其字段没有被导出 那么在调用该包的程序中不能修改或调用其字段
// 使用方法来修改或调用未导出字段
type Person struct {
	firstName string
	lastName  string
}

func (p *Person) GetFirstName() string {
	return p.firstName
}
func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

type Log struct {
	msg string
}
type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}
func (c *Customer) Log() *Log {
	return c.log
}
func (l *Log) String() string {
	return l.msg
}
func Embed1() {
	//c := new(Customer)
	//c.Name = "Barak Obama"
	//c.log = new(Log)
	//c.log.msg = "1 - Yes we can!"
	c := &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}

	fmt.Println(c)
	c.log.Add("2 - The world will be a better place!")
	fmt.Println(c.log)
	fmt.Println(c.Log())
	fmt.Println(c.log.String())
}

type Customer2 struct {
	Name string
	Log  // 内嵌 内嵌类型不需要指针
}

func (c *Customer2) String() string {
	return c.Name + "\nLogs: " + fmt.Sprintln(c.Log)
}

func Embed2() {
	fmt.Println("====Embed2====")
	c := &Customer2{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Log.Add("2 - The world will be a better place!")
	fmt.Println(c)
}

// 多重继承
type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func MultiInherit() { // 多重继承
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It exhibits behavior of a Phone: ", cp.Call())
}

