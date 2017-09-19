package main

//https://studygolang.com/articles/2073
import (
	"bufio" //缓存IO
	"fmt"   //io 工具包
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
  from: http://www.isharey.com/?p=143
*/

func main() {
	var wireteString = "测试n"
	//var filename = "./output1.txt"
	var f *os.File
	//var err1 error

	/*****************************  第三种方式:  使用 File(Write,WriteString) 写入文件 ***********************************************/
	var d1 = []byte(wireteString)
	f, err3 := os.Create("./output.txt") //创建文件
	check(err3)
	defer f.Close()
	n2, err3 := f.Write(d1) //写入文件(字节数组)
	check(err3)
	fmt.Printf("写入 %d 个字节n", n2)
	n3, err3 := f.WriteString("----------writesn") //写入文件(字节数组)
	fmt.Printf("写入 %d 个字节n", n3)
	f.Sync()

	/***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************/
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n4, err3 := w.WriteString("-------bufferedn")
	fmt.Printf("写入 %d 个字节n", n4)
	w.Flush()
	f.Close()
}
