/*	解析数据格式: JSON XML gob Google缓冲协议等等

	数据结构 --> 指定格式 = 序列化 或 编码（传输之前）
	指定格式 --> 数据格式 = 反序列化 或 解码（传输之后）
	序列化是在内存中把数据转换成指定格式（data -> string），反之亦然（string -> data structure）

	编码也是一样的，只是输出一个数据流（实现了 io.Writer 接口）
	解码是从一个数据流（实现了 io.Reader）输出到一个数据结构。
	我们都比较熟悉XML格式 但有些时候JSON（JavaScript Object Notation被作为首选，主要是由于其格式上非常简洁
	通常 JSON 被用于 web 后端和浏览器之间的通讯，但是在其它场景也同样的有用。
*/
package pack7

import (
	"bytes"
	"crypto/sha1"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

type Address struct {
	Type, City, Country string
}

type VCard struct {
	FirstName, LastName string
	Addresses           []*Address
	Remark              string
}

//JSON数据格式
func JsonDataFormat() {
	pa := &Address{"private", "Shanghai", "China"}
	wa := &Address{"work", "Hangzhou", "China"}
	vc := VCard{"MG", "LM", []*Address{pa, wa}, "none"}
	fmt.Printf("%v: \n", vc)
	// JSON Format: 输出vc的json数据格式
	js, _ := json.Marshal(vc) //出于安全考虑,在web应用中最好使用json.MarshalforHTML() 其对数据执行HTML转码 所以文本可以被安全地嵌在 HTML <script> 标签中
	fmt.Printf("JSON format: %s\n", js)

	//序列化 | Encoder 编码器  将vc struct以json格式写入到文件中
	fmt.Println("Encoder...")
	file, _ := os.OpenFile("data/7_vcard.json", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc) //将数据对象v的json编码写入io.Writer w中
	if err != nil {
		log.Println("Error in encoding json")
	}
	fmt.Println("已将vc struct以json格式写入到文件中")

	//反序列化 | Decoder 解码器
	fmt.Println("Decoder...")
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err2 := json.Unmarshal(b, &f)
	if err2 != nil {
		return
	}
	//f指向的值是一个map: key是一个string value是自身存储作为空接口类型的值
	fmt.Println(f)
	//使用断言查看f类型
	m := f.(map[string]interface{})
	fmt.Printf("%T ", m)           //查看f类型
	fmt.Println(m)                 //查看f的值
	fmt.Println(reflect.TypeOf(f)) //reflect查看Type
	//用于处理未知的json数据时 保证类型安全
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println("Unkown")
		}
	}
	//先了解JSON数据结构 再对其反序列化
	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
	}
	var fm FamilyMember
	err = json.Unmarshal(b, &fm)
	fmt.Println(fm)

	//编码和解码流
	//json包提供Decoder和Encoder类型来支持常用JSON数据流读写
	//NewDecoder和NewEncoder函数分别封装了io.Reader和io.Writer 接口
}

//XML数据格式
func XMLDataFormat() {
	//跟JSON一样有Marshal()和UnMarshal() 从XML中编码和解码数据
	var t xml.Token
	var err error
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement: //xml开始标签
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement: //xml结束标签
			fmt.Println("End of token")
		case xml.CharData: //开始标签到结束标签之间的实际文本 内容是[]byte
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
			fmt.Println("Default")
		}
	}
}

//Gob(Go binary) Go自己的以二进制形式序列化和反序列化数据的格式 在encoding包中
//通常用于远程方法调用(RPCs)参数和结果的传输 以及应用程序和机器之间的数据传输
type P struct {
	X, Y, Z int
	Name    string
}
type Q struct {
	X, Y *int32
	Name string
}

func Gob1() {
	//初始化encoder decoder 都建立在网络(network)连接上
	var network bytes.Buffer        //网络连接
	enc := gob.NewEncoder(&network) //写入网络
	dec := gob.NewDecoder(&network) //从网络中读取
	//Encoder发送值value
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encoder error:", err)
	}
	//Decoder接收值value
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decoder error:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}

func Gob2() {
	//struct 存为 gob文件
	pa := &Address{"private", "Shanghai", "China"}
	wa := &Address{"work", "Hangzhou", "China"}
	vc := VCard{"MG", "LM", []*Address{pa, wa}, "none"}
	file, _ := os.OpenFile("data/7_vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc) //Encode(需要传输的接口)
	if err != nil {
		log.Println("Error in encoding gob file")
	}
}

//解码gob文件并打印
func Degob() {
	//读取文件
	file, err := os.OpenFile("data/7_vcard.gob", os.O_RDONLY, 0664)
	if err != nil {
		log.Fatal("Open file error:", err)
		return
	}
	//解码
	dec := gob.NewDecoder(file)
	var vc VCard
	err = dec.Decode(&vc)
	if err != nil {
		log.Fatal("Decoder error:", err)
		return
	}
	//print
	fmt.Printf("%s_%s: ", vc.FirstName, vc.LastName)
	for _, i := range vc.Addresses {
		fmt.Printf("%s ", *i)
	}
	fmt.Printf("%s", vc.Remark)
}

func GobHash_sha1() {
	fmt.Println("sha1加密: ")
	hasher := sha1.New() //创建一个新的hash.Hash对象 用来计算SHA1校验值
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result1: %x\n", hasher.Sum(b))
	fmt.Printf("Result2: %d\n", hasher.Sum(b))
	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result data: %x\n", checksum)
}
