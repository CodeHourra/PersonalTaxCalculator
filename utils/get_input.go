package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetCommandString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(input, "\n", "", -1), err
}

func GetCommandFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.Replace(input, "\n", "", -1)
	inputFloat64, _ := strconv.ParseFloat(input, 64)
	return inputFloat64, err
}

// ExecShell 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func ExecShell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)

	return out.String(), err
}

// 错误处理函数
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func StrEmpty(str string) bool {
	if len(str) != 0 && str != "" {
		return true
	}
	return false
}

func StrNotEmpty(str string) bool {
	return !StrEmpty(str)
}
