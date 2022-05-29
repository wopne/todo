package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	/**
	0x200 好像是十六进制
	0010 0000 0000

	0010 0000 0000	0x200 = 2^9 = 512
	0000 0000 1000	0x8   = 2^3 = 8
	0000 0000 0010	0x2   = 2^1 = 2

	0010 0000 1010  0x20a = 2^9 + 2^3 + 2^1 = 522

	os.O_CREATE|os.O_APPEND|os.O_RDWR

	按位或运算符（|）
	运算规则：0|0=0；   0|1=1；   1|0=1；    1|1=1；

	更多: https://blog.csdn.net/21aspnet/article/details/7387373
	*/

	fmt.Println(os.O_CREATE, os.O_APPEND, os.O_RDWR, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.O_CREATE|os.O_APPEND)

	logFileLocation, _ := os.OpenFile("./docs/work/20220519/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)

	fmt.Println(logFileLocation)

	//参考: https://learnku.com/go/wikis/23398
	//ex, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}d
	//exPath := filepath.Dir(ex)
	//fmt.Println(exPath)

	dir, _ := os.Getwd()
	fmt.Println(dir)

	//fileInfo, _ := ioutil.ReadDir("./")
	//
	//for index, info := range fileInfo {
	//	fmt.Println(index, info.Name())
	//}

	_, file, line, ok := runtime.Caller(0)
	fmt.Println(file, line, ok)

	fmt.Println(os.Args[0], len(os.Args))
}
