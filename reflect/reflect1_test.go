package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func Test1013(t *testing.T) {
	var Num float64 = 3.14

	v := reflect.ValueOf(Num)
	t1 := reflect.TypeOf(Num)

	fmt.Println("Reflect : Num.Value = ", v)
	fmt.Println("Reflect : Num.Type  = ", t1)
}

func Test102519(t *testing.T) {
	var Num = 3.14
	v := reflect.ValueOf(Num)
	t1 := reflect.TypeOf(Num)
	fmt.Println(v)
	fmt.Println(t1)

	origin := v.Interface().(float64)
	fmt.Println(origin)

}

func Test102531(t *testing.T) {
	var Num float64 = 3.14
	v := reflect.ValueOf(Num)
	v.SetFloat(6.18) //panic

}

func Test102538(t *testing.T) {
	var Num float64 = 3.14
	v := reflect.ValueOf(Num)
	fmt.Println("v的可写性:", v.CanSet())

}

func Test102545(t *testing.T) { //通过反射修改数值
	var f float64 = 3.41
	fmt.Println(f)
	p := reflect.ValueOf(&f)
	v := p.Elem()
	v.SetFloat(6.18)
	fmt.Println(f)

}

func Test102555(t *testing.T) { //在Go 语言中，函数是第一类对象
	hl := hello
	fv := reflect.ValueOf(hl) //反射的是函数
	fv.Call(nil)
	h2 := hello2
	fv2 := reflect.ValueOf(h2)
	args := []reflect.Value{reflect.ValueOf("wudebao")}
	fv2.Call(args)
}
func hello() {
	fmt.Println("Hello world!")
}
func hello2(a string) {
	fmt.Println("Hello world2!" + a)
}

func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input) //想要的type类型
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input) //想要的具体的值
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func Test10251449(t *testing.T) {
	user := User{1, "Allen.Wu", 25}

	DoFiledAndMethod(user)

}

func Test10251505(t *testing.T) {
	user := User{1, "Allen.Wu", 25}

	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
	getValue := reflect.ValueOf(user)

	// 一定要指定参数为正确的方法名
	// 2. 先看看带有参数的调用方法
	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs") //这是对象的方法
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	methodValue.Call(args)

	// 一定要指定参数为正确的方法名
	// 3. 再看看无参数的调用方法
	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)

}
