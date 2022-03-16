/*
结构体与方法
在Go中，结构体是类的一种简化形式
*/
package utils

import (
	"fmt"
	"go_test/pack1"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type T struct{ a, b int }
type struct1 struct {
	i1  int
	f1  float32
	str string
}
type Person struct{ firstName, lastName string }
type number struct{ f float32 }

// 可以通过引用自身来定义 定义在链表或二叉树的节点中
type Node struct {
	data float64
	su   *Node // 后继节点
}
type Tree struct {
	le   *Tree // left node
	data float64
	ri   *Tree // right node
}

// 结构体工厂所用struct
type File struct {
	fd   int
	name string
}

// 带有tags的struct
type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

// 方法所用struct
type TwoInts struct {
	a int
	b int
}

func (ti *TwoInts) String() string {
	return "(" + strconv.Itoa(ti.a) + "/" + strconv.Itoa(ti.b) + ")"
}

type IntVector []int

type myTime struct {
	time.Time // 匿名字段
}

// 内嵌类型
type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamePoint struct {
	Point // 父类
	name  string
}

func (np *NamePoint) Abs() float64 {
	return np.Point.Abs() * 100.
}

func StructAndFunction() {
	fmt.Println("====结构体字段赋值====")
	var s T
	s.a = 5 // 点号符. 在Go中叫选择器(selector)
	s.b = 8
	// 使用new()
	var t *T   // 也可以写成t:=new(T) 直接分配内存了
	t = new(T) // 变量t是指向T的指针 此时结构体字段的值是它们所属类型的零值
	t.a, t.b = 1, 1

	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "MG"
	// ms := &struct1{10, 15.5, "MG"}  // 与上面四行等价 底层仍会调用new() 即: new(Type)<=>&Type{}
	fmt.Println(ms.i1, ms.f1, ms.str)
	fmt.Println(ms)

	// 1. struct as a value type:  结构体作为值类型
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	//fmt.Println(pers1.firstName, pers1.lastName)
	upPerson(&pers1)
	fmt.Println(pers1.firstName, pers1.lastName)
	// 2. struct as a pointer:  结构体作为一个指针
	pers2 := new(Person)
	pers2.firstName = "Chris" // 没有像C++中那样使用->操作符，Go会自动转化 但也可以解指针(*pers2).lastName="Woodward"
	pers2.lastName = "Woodward"
	//fmt.Println(pers2.firstName, pers2.lastName)
	upPerson(pers2)
	fmt.Println(pers2.firstName, pers2.lastName)
	// 3. struct as a literal: 混合方法
	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Println(pers3.firstName, pers3.lastName)

	fmt.Println("====结构体转换====")

	a := number{5.0}
	type nr number    // alias type 起别名
	b := nr{5.0}      // nr这里是number的别名 alias type
	var c = number(b) // 需要转换
	fmt.Println(a, b, c)

	fmt.Println("====结构体Test====")
	vc := pack1.VCardFunc()
	fmt.Printf("%v, %v\n", &vc, vc.Addr)

	fmt.Println("====结构体工厂====")
	f := NewFile(10, "./test.txt") // 工厂实例化类型的一个对象
	fmt.Printf("fd: %v, name: %v\n", f.fd, f.name)
	fmt.Println(unsafe.Sizeof(f)) // f所占内存

	fmt.Println("====new/make struct的编译问题====")
	type Foo map[string]string
	type Bar struct {
		thingOne string
		thingTwo int
	}

	// compile ok 能够编译成功
	y := new(Bar) // new 用于值类型(array struct)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1
	fmt.Println("用new实例化数组和结构体")
	// compile fail
	//z := make(Bar)  // make 用于引用类型(slice map channel)
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	// compile ok
	x := make(Foo)
	x["x"] = "hello"
	x["y"] = "world"
	fmt.Println("用make实例化slice map channel")
	// compile fail
	//y := new(Foo)
	//(*y)["x"]="hello"
	//(*y)["y"]="world"

	fmt.Println("====使用自定义包中的结构体====")
	expStruct1 := new(pack1.ExpStruct)
	expStruct1.Mi1 = 10
	expStruct1.Mf1 = 16.
	fmt.Printf("Mi1 = %d, Mf1 = %f\n", expStruct1.Mi1, expStruct1.Mf1)

	fmt.Println("====带标签的结构体====")
	tt := TagType{true, "Phone", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

	fmt.Println("====匿名字段和内嵌结构体====")
	type innerS struct {
		in1 int
		in2 int
	}
	type outerS struct {
		b int
		c float32
		// anonymous field 匿名字段
		int //在一个结构体中对于每一种数据类型只能有一个匿名字段
		innerS
	}
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60 // 通过类型 outer.int 的名字来获取存储在匿名字段中的数据
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)

	fmt.Println("====方法====")
	//Go方法是作用在接收者(receiver)上的一个函数，接收者是某种类型的变量，所以方法是一种特殊类型的函数
	// 结构体类型上的方法例子
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("The sum is: %v\n", two1.AddThem())
	fmt.Printf("Add them to the param: %v\n", two1.AddToParam(20))
	two2 := TwoInts{3, 4}
	fmt.Println(two2.AddThem())

	// 非结构体类型上方法例子
	fmt.Println(IntVector{1, 2, 3, 4, 5, 6}.IntVectorSum())
	fmt.Println(pack1.Employee{300.0}.GiveRaise(0.3))
	m := myTime{time.Now()}
	fmt.Println("Full time:", m.String())
	fmt.Println("First 3 chars:", m.first3Chars())

	// 指针或值作为接收者 => 只有指针作为接收者的时候才可以改变起内部成员的值
	pack1.Pointer_value()

	// 指针方法和值方法都可以在指针或非指针上被调用
	pack1.Print_List()

	// 方法与未导出字段: 利用方法修改或调用未导出字段
	p := new(pack1.Person)
	// p.firstName undefined
	p.SetFirstName("Eric")
	fmt.Println(p.GetFirstName())

	//内嵌类型的方法和继承
	n := &NamePoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Point.Abs()) // 5 内嵌类型的方法
	fmt.Println(n.Abs())       // 500 n.Abs() 外层类型的方法覆写内嵌类型对应的方法

	// 在类型中嵌入功能
	// 实现方法1: 聚合(组合): 包含一个所需功能类型的具名字段
	pack1.Embed1()
	// 实现方法2: 内嵌: 内嵌(匿名地)所需功能类型，如“内嵌类型的方法和继承”所示
	pack1.Embed2()

	// 多重继承
	pack1.MultiInherit()

	// 类型的String()方法和格式化描述符
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)  // %T 会给出类型的完全规格
	fmt.Printf("two1 is: %#v\n", two1) // %#v 会给出实例的完整输出，包括它的字段（在程序自动生成 Go 代码时也很有用）。

	// Test
	pack1.MagicTest()
	pack1.TypeString()
	pack1.CelsiusTest()
	pack1.Days()
	pack1.StackOperate()

}

// 结构体字段自定义操作 这里是strings.ToUpper()
func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

/*	结构体类型对应的工厂方法
	return 一个指向结构体实例的指针
*/
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

// 输出结构体标签
func refTag(tt TagType, ix int) {
	// reflect包的TypeOf()可以获取变量的正确类型
	// 另外变量如果是一个结构体类型 就可以通过Field来索引结构体的字段 然后就可以使用Tag属性
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

// 带receiver的方法
// 接收者类型可以是(几乎)任何类型，不仅仅是结构体类型：任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。
// 但是接收者不能是一个接口类型（参考 第 11 章），因为接口是一个抽象定义，但是方法却是具体实现；如果这样做会引发一个编译错误
func (recv *TwoInts) AddThem() int {
	return recv.a + recv.b
}
func (recv *TwoInts) AddToParam(param int) int {
	return recv.a + recv.b + param
}
func (i IntVector) IntVectorSum() (s int) {
	for _, x := range i {
		s += x
	}
	return
}

// 不能直接调用其他包里的类型
// 解决方法：将别的包的类型作为匿名类型嵌入在一个本地的新结构体中
func (mt myTime) first3Chars() string {
	return mt.Time.String()[0:3]
}
