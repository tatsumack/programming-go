package main

import (
	"fmt"
	"os"
	"strings"
)

// echoプログラムを修正してコマンド名も出力する

// strings.Joinにos.Argsを丸っと渡せばよい

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
