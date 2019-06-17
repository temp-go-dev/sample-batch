package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 引数の取得
	flag.Parse()
	args := flag.Args()

	for _, s := range args {
		fmt.Println(s)
	}

	// 終了コードの設定
	os.Exit(1)
}
