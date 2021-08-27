package main

import (
	"runtime"
	// "fmt"
	// "time"
	"github.com/kopjenmbeng/kanggo_rest_test/cmd"
)

func main() {
	// fmt.Println(time.Now().Unix())
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
