package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	errorh "github.com/temp-go-dev/sample-batch/error"
)

func main() {
	// 引数の取得
	flag.Parse()
	args := flag.Args()

	for _, s := range args {
		fmt.Println(s)
	}

	err := run(args)
	// 終了コード判定を行い、処理終了
	os.Exit(errorh.HandleExit(err))
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("too few arguments")
	}

	if args[0] == "01234" {
		return errorh.NewExitError(errorh.ExitCodeWarn, "E001", errors.New("01234 error"))
	}

	return nil
}
