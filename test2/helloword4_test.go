package test3

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"testing"
)

func Test1013(t *testing.T) {
	fmt.Println("test2...")

}

func test23() {
	fmt.Println("test2...")

}

func TestName(t *testing.T) {

	vm := otto.New()
	// 使用vm在go语言中运行一段简单的javascript代码
	vm.Run(`
    abc = 2 + 2;
    console.log("The value of abc is " + abc); // 4
    `)
	// 获取javascript中的变量的值
	if value, err := vm.Get("abc"); err == nil {
		if value_int, err := value.ToInteger(); err == nil {
			fmt.Println(value_int, err) //4 <nil>
		}
	}

	// 定义了一个sayHello的javascript方法，其调用的是一段go语言代码。
	vm.Set("sayHello", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("Hello, %s.\n", call.Argument(0).String())
		return otto.Value{}
	})
	_, _ = vm.Run(`
    sayHello("Xyzzy");      // Hello, Xyzzy.
    `)

}
