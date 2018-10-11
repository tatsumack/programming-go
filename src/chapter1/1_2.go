package main

import (
	"fmt"
	"os"
)

// echoプログラムを修正して、引数のインデックスと値の組を出力する
// Printlnはオペランド間に空白を挿入し、最後に改行文字を出力してくれる

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
