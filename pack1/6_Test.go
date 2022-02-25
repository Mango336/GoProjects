/* 6_interface_reflect 例子 */
package pack1

import (
	"fmt"
	"reflect"
)

/* 例子: 使用Sorter接口 排序 */
type IntArray []int

func (ia IntArray) Len() int           { return len(ia) }
func (ia IntArray) Less(i, j int) bool { return ia[i] > ia[j] }
func (ia IntArray) Swap(i, j int)      { ia[i], ia[j] = ia[j], ia[i] }

type StringArray []string

func (sa StringArray) Len() int           { return len(sa) }
func (sa StringArray) Less(i, j int) bool { return sa[i] > sa[j] }
func (sa StringArray) Swap(i, j int)      { sa[i], sa[j] = sa[j], sa[i] }

type Dayss struct {
	Num                 int
	ShortName, LongName string
}
type DayArray struct {
	Data []*Dayss
}

func (da *DayArray) Len() int           { return len(da.Data) }
func (da *DayArray) Less(i, j int) bool { return da.Data[i].Num > da.Data[j].Num }
func (da *DayArray) Swap(i, j int)      { da.Data[i], da.Data[j] = da.Data[j], da.Data[i] }

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for i := 0; i < data.Len(); i++ {
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(i, j) {
				data.Swap(i, j)
			}
		}
	}
}

/* reflect struct */
type T2 struct {
	A int
	B string
}

func ReflectStruct() {
	t := T2{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem() // 传递地址 可修改结构体内容
	typeT2 := s.Type()
	fmt.Println("the type of T2 is: ", typeT2)
	for i := 0; i < typeT2.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("Field#%02d: %s %s = %v\n",
			i, typeT2.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("The reset T2 is: ", t)
}

/* cars Test */
type Any interface{}
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}
type Cars []*Car

func CarsTest() {
	// 定义cars
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	//allCars2 := Cars{ford, bmw, merc, bmw2}
	allNewBWMs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("AllNewBMWs: ", allNewBWMs)

	manufactures := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufactures)
	allCars.Process(sortedAppender)
	fmt.Println(sortedCars)

}

// 定义一个通用的Process()函数 接收一个作用与每一辆car的f函数作为参数
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// 实现一个查找函数来获取子集合 并在Process()中传入一个闭包执行(这样就可以访问局部切片cars)
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// 实现Map功能 去掉除 car 对象以外的东西
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

//
func MakeSortedAppender(manufactures []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufactures {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}
