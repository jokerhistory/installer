package ext

//https://studygolang.com/articles/2073
import (
//	"bufio" //缓存IO
	//"fmt"   //io 工具包
	"os"
	"io/ioutil" //io 工具包
	//"io"
	//"bytes"
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

func ReadFile(readf string) ([]byte){
		// fmt.Printf("写入 %d 个字节n", n);
		b,err := ioutil.ReadFile(readf)
		// actionb,err := ioutil.ReadFile(tpaction)
		// footb,err := ioutil.ReadFile(tpfoot)
		check(err)
		
		return b
}
	
func WriteFile(outputb []byte,outputf string) {

	
	//  var f *os.File
	// // 	/*****************************  第三种方式:  使用 File(Write,WriteString) 写入文件 ***********************************************/
	// // //var d1 = []byte("")
	// isExist := checkFileIsExist(filename)
	// if(isExist){
	// 	f, err := os.OpenFile (filename, os.O_APPEND, 0666) //创建文件	
	// 	check(err)
	// }else
	// {
	// 	f, err := os.Create(filename) //创建文件	
	// 	check(err)
	// }



	
	//buf := bytes.NewBuffer(outputb)
	// buf.Write(actionb)
	// buf.Write(footb)
	
	//var d1 = []byte(content);

	err := ioutil.WriteFile(outputf,outputb, 0666)  //写入文件(字节数组)
	check(err)
	
	//----------------
	// defer f.Close()
	// n2, err3 := f.Write(d1) //写入文件(字节数组)
	// check(err3)
	// fmt.Printf("写入 %d 个字节n", n2)
	// n3, err3 := f.WriteString("----------writesn") //写入文件(字节数组)
	// fmt.Printf("写入 %d 个字节n", n3)
	// f.Sync()

	/***************************** 第四种方式:  使用 bufio.NewWriter 写入文件 ***********************************************/
	// w := bufio.NewWriter(f) //创建新的 Writer 对象
	// n4, err := w.WriteString(content)
	// fmt.Printf("写入 %d 个字节n", n4)
	// w.Flush()
	// f.Close()
}
