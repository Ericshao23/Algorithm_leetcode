// https://codeforces.com/contest/1717/problem/A
package main

import "fmt"

func main() {
	newFunction()
	fmt.Println("hello world")
}

func newFunction() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int32
		fmt.Scan(&n)
		fmt.Println(n + 2*(n/2+n/3))
	}
}

// # Go 1.13 及以上（推荐）

// # Windows执行
// go env -w GO111MODULE=on
// go env -w GOPROXY=https://goproxy.io,direct

// # Windows PowerShell 执行
// $env:GO111MODULE = "on"
// $env:GOPROXY = "https://goproxy.cn"

// # macOS 或 Linux 执行
// export GO111MODULE=on
// export GOPROXY=https://goproxy.cn

// # 或者  macOS 或 Linux 执行
// echo "export GO111MODULE=on" >> ~/.profile
// echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
// source ~/.profile

// // 手动安装
// go get -u -v github.com/mdempsky/gocode
// go get -u -v github.com/uudashr/gopkgs/v2/cmd/gopkgs
// go get -u -v github.com/ramya-rao-a/go-outline
// go get -u -v github.com/acroca/go-symbols
// go get -u -v golang.org/x/tools/cmd/guru
// go get -u -v golang.org/x/tools/cmd/gorename
// go get -u -v github.com/cweill/gotests/...
// go get -u -v github.com/fatih/gomodifytags
// go get -u -v github.com/josharian/impl
// go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
// go get -u -v github.com/haya14busa/goplay/cmd/goplay
// go get -u -v github.com/godoctor/godoctor
// go get -u -v github.com/go-delve/delve/cmd/dlv
// go get -u -v github.com/stamblerre/gocode
// go get -u -v github.com/rogpeppe/godef
// go get -u -v github.com/sqs/goreturns
// go get -u -v golang.org/x/lint/golint
