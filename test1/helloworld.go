// Package main provides ...
package main
import (
	"fmt"
	"time"
)

func main() {

		s:="hello world ";
		x := 10;
		// 流程控制语句



		//for {
			fmt.Println("死循环");
		//}
		// go 语言中只有一个for循环 
		for x > 0{
				x -- ;
				fmt.Println("this is hello world");
		}

		// if 条件控制语句
		if len(s) ==  11 {
				fmt.Println("this string's len is 11");

		}else{
				
				fmt.Println("this string's len is not 11");
		}
		for i := 0; i < len(s); i++ {
		    fmt.Printf("%c\n",s[i]);
			fmt.Println(s[i]);
		}
		for i := 0; i < 10; i++ {
			fmt.Println("hello world");
		}
		fmt.Println(helloworld());
		fmt.Println(sayHelloWorld());
		var_test1();
		con_test1();
		sum,sub:=calc(1,2);
		fmt.Println(sum,sub);

		go fmt.Println(add(1,2));
		time.Sleep(time.Second * 1);
		fmt.Println("hello world")
}

func helloworld() string{
		fmt.Printf("hello world");
		return "hello world"
	
}
func add(a int , b int) int {
		return a + b
}
func calc(a int ,b int) (int , int ) {
		return a + b ,a - b
}

func sayHelloWorld() string {
	return "hello world"
}
//标识符
//关键字
// 变量
// 常量
func  var_test1() {
		var a int
		var b string
		var c bool
		var d int = 8
		var e string = "hello world"
		fmt.Println(a,b,c,d,e)
	
}

// 权限控制
func Test1(){
		fmt.Println("this is Test1");
}
func  con_test1(){
		const a int = 9;
		const b string = "hello world";
		const c bool = false;
		const d = 9;
		fmt.Println(a,b,c,d);
		const (
				e string = "hello c"
				f int = 0;
				g 
		)
		fmt.Println(e,f,g);
		const (
				h =  1<<iota
				j
				k
		)

		fmt.Println(h,j,k);
}


