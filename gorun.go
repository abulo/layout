package main

import (
	"fmt"
	"os"

	"github.com/abulo/ratel/v3/watch"
)

var (
	runcmd  string
	runArgs string
	// command os.Args
)

func init() {
	// flag.StringVar(&runcmd, "run", "", "要运行的Go文件")
	// flag.StringVar(&runArgs, "args", "", "参数")
}

func main() {
	//解析参数
	// flag.Parse()
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("路径参数不存在", "参考:./gorun cmd/server.go")
		os.Exit(1)
	}
	runcmd = args[1]
	watch.Run(runcmd, runArgs)
}
