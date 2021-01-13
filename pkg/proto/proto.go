package proto

import (
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/proto"
)

func ProtoRead(path string) {
	file, fileErr := os.Open(path) //打开文件
	checkError(fileErr)

	defer file.Close() //关闭文件 ,defer 会在程序最后运行
	fs, fsErr := file.Stat()
	checkError(fsErr)
	buffer := make([]byte, fs.Size()) //创建 byte切片
	//把file文件内容读取到buffer
	_, readErr := io.ReadFull(file, buffer)
	checkError(readErr)

	//初始化pb结构体对象并将buffer中的文件内容读取到pb结构体中
	msg := &stProto.HelloWorld{}
	pbErr := proto.Unmarshal(buffer, msg) //反序列化数据
	checkError(pbErr)
}

//检查错误
func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
