package main

import (
	"flag"
	"fmt"
	"os"
)

/*
	java命令集中常用方式

	java [-options] class [args]
	java [-options] -jar jarfile [args]
	javaw [-options] class [args]
	javaw [-options] -jar jarfile [args]

*/

// 定义java 命令的 options结构 struct
type Cmd struct {
	// help 信息
	helpFlag bool
	// 版本信息
	versionFlag bool
	// 类路径
	cpOption string
	// main class 名称
	class string
	// main class 参数
	args [] string
}

/*
	解析选项参数
	-version 		版本信息
	-？/ -help 		帮助信息
	-cp/-classpath  置顶用户类路径
	-Dpropery=value 设置java系统属性
	-Xms<size>		初始堆空间大小
	-Xmx<size>		最大堆空间大小
	-Xss<size>		最小堆空间大小
*/
func parseCmd() *Cmd {

	cmd := &Cmd{}
	// 终端打印
	flag.Usage = printUsage

	// options
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")

	flag.BoolVar(&cmd.helpFlag, "?", false, "pring help message")

	flag.BoolVar(&cmd.versionFlag, "version", false, "pring version and exit")

	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")

	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")

	flag.Parse()

	// 捕获没有被解析的参数
	args := flag.Args()

	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

// 打印提示
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...] \n", os.Args[0])
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath: %s class: %s  args: %v \n", cmd.cpOption, cmd.class, cmd.args)
}
