package utils

import (
	"fmt"
	"go_test/pack1"
	"math"
	"reflect"
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

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

/* interface 例子 valuable 显示其多态性 */
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

/* 方法集与接口的使用 */
type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

/* 空接口 */
var i = 5
var str = "ABC"

type Person2 struct {
	name string
	age  int
}
type Any interface{} // 空接口

/* 通用类型的节点数据结构 */
type node struct {
	left  *node
	data  interface{}
	right *node
}

func NewNode(left, right *node) *node { // 创建空节点
	return &node{left, nil, right}
}
func (n *node) SetData(data interface{}) { // 给节点设置数据
	n.data = data
}

/* 反射结构体 */
type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func InterfaceAndReflect() {
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

	fmt.Println("====类型断言====")
	// varI.(T) 使用类型断言来测试在某个时刻varI是否包含类型T的值
	// varI为一个接口类型
	if t, ok := areaIntf.(*Square); ok {
		// ok==True 包含; 因为areaIntf=sq1 sq1是*Square
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
	fmt.Println("====类型判断====")
	// t得到了areaIntf的值和类型
	// 所有case中列举的类型(nil除外)都必须实现对应的接口(即这里的Shaper)
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with the value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with the value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Println("Unexpected type %T\n", t)
	}

	fmt.Println("====类型分类函数====")
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)

	fmt.Println("====测试某值是否实现了某个接口====")
	if v, ok := areaIntf.(Shaper); ok {
		fmt.Printf("q implements Area(): %s\n", v.Area())
	}

	fmt.Println("====使用方法集与接口====")
	// 空值
	var lst List
	// CountInto(lst, 1, 10)  编译错误
	// List没有实现Appender(Appender方法有指针接收器)
	// CountInto 需要一个Appender 而它的方法Append只定义在指针上

	// 在lst上调用LongEnough是可以的 因为Len定义在值上
	if LongEnough(lst) { // 相同的接收器类型
		fmt.Println("- lst is long enough")
	}

	// 指针类型
	plst := new(List)
	// CountInto 接收Appender Appender接收*List 所以这里用plst没问题 上面用lst不是指针 有问题
	//  在 plst 上调用 LongEnough 也是可以的，因为指针会被自动解引用。
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Println("- plst is long enough")
	}

	/* 例子: 使用Sorter接口排序 */
	// int排序
	//data := &pack1.IntArray{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	intData := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	id := pack1.IntArray(intData)
	pack1.Sort(id)
	fmt.Println(id)
	// 字符串排序
	stringData := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	sd := pack1.StringArray(stringData)
	pack1.Sort(sd)
	fmt.Println(sd)
	// days 星期排序
	Sunday := pack1.Dayss{0, "SUN", "Sunday"}
	Monday := pack1.Dayss{1, "MON", "Monday"}
	Tuesday := pack1.Dayss{2, "TUE", "Tuesday"}
	Wednesday := pack1.Dayss{3, "WED", "Wednesday"}
	Thursday := pack1.Dayss{4, "THU", "Thursday"}
	Friday := pack1.Dayss{5, "FRI", "Friday"}
	Saturday := pack1.Dayss{6, "SAT", "Saturday"}
	dayData := []*pack1.Dayss{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	dd := pack1.DayArray{dayData}
	pack1.Sort(&dd)
	for _, d := range dayData {
		fmt.Printf("%s ", d.LongName)
	}

	fmt.Println("\n====空接口====")
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person2)
	pers1.name = "Rob Pike"
	pers1.age = 25
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person2:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}

	// 空接口在type-switch中联合匿名函数(lambda函数)的用法
	TypeSwitch()

	// 二叉树: 通用定义
	root := NewNode(nil, nil)
	root.SetData("root node")
	ln := NewNode(nil, nil)
	ln.SetData("left node")
	rn := NewNode(nil, nil)
	rn.SetData(2)
	root.left = ln
	root.right = rn
	fmt.Println(root)

	fmt.Println("====反射包====")
	var x float64 = 3.4
	fmt.Println("type: ", reflect.TypeOf(x)) // 返回被检查对象的类型
	v := reflect.ValueOf(x)                  // 返回被检查对象的值
	fmt.Println("value: ", v)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind())
	fmt.Println("value: ", v.Float())
	fmt.Println(v.Interface()) // 变量v的Interface()方法可以得到还原（接口）值 所以可以这样打印v的值
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
	//v.SetFloat(3.1415) // 编译错误 因为reflect.ValueOf(x)是拷贝了x 那么更改v没有用; 需要传递地址
	fmt.Println("v能否设置值: ", v.CanSet())
	v = reflect.ValueOf(&x) // 传递x的地址给v
	fmt.Println("type of v: ", v.Type())
	fmt.Println("v能否设置值: ", v.CanSet())
	v = v.Elem() // 需要Elem()函数 间接使用指针
	fmt.Println("The elem of v is: ", v)
	fmt.Println("v能否设置值: ", v.CanSet())
	v.SetFloat(3.1425)
	fmt.Println(v.Interface())
	fmt.Println(v)

	fmt.Println("====反射结构体====")
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret) // 或 typ:=value.Type()
	fmt.Println(value, typ)
	fmt.Println(value.Type())
	knd := value.Kind()
	fmt.Println(knd)
	// 迭代遍历每个Field的值 NumField()=>返回结构体内的字段数量
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error信息: panic: reflect: reflect.Value.SetString using value obtained using unexported field
		// 结构体中只有被导出字段(首字母大写)才可被设置值
		//value.Field(i).SetString("C++")
	}
	// Method(n).Call(nil) 使用索引n来调用签名在结构体上的方法
	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results)
	fmt.Println(value.NumMethod()) // 返回签名在结构体上方法的数量
	// 结构体的反射 例子
	pack1.ReflectStruct()

	//fmt.Println("====动态方法调用====")
	pack1.CarsTest()

}

// 类型分类函数 (可变长度参数 可是任意类型的数组 根据数组元素的实际类型执行不同的动作
// 在处理来自于外部的、类型未知的数据时 如: 解析诸如JSON或XML编码的数据 类型测试和转换会非常有用
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is bool.\n", i)
		case float32, float64:
			fmt.Printf("Param #%d is float.\n", i)
		case int, int64:
			fmt.Printf("Param #%d is int.\n", i)
		case string:
			fmt.Printf("Param #%d is string.\n", i)
		case nil:
			fmt.Printf("Param #%d is nil.\n", i)
		default:
			fmt.Printf("Param #%d is unknown.\n", i)
		}
	}
}

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type\n", v)
		case int:
			fmt.Printf("any %v is an int type\n", v)
		case float32:
			fmt.Printf("any %v is a float32 type\n", v)
		case string:
			fmt.Printf("any %v is a string type\n", v)
		case specialString:
			fmt.Printf("any %v is a special String!\n", v)
		default:
			fmt.Println("unknown type!\n")
		}
	}
	testFunc(whatIsThis)
	a := 1
	testFunc(a)
	b := false
	testFunc(b)
	testFunc("asd")
}
