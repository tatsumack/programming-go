package main

import (
	"bufio"
	"fmt"
	"os"
)

// 重複行が含まれているファイル名を出力する

// keyが行で、valueがファイル名のsetとなるmapを用意する。
// setはGo言語に存在しないので、mapのフィールドに空のstructを使ってsetを実現する。ちなみに空のstructでメモリを確保することはないらしい。
// 参考: https://groups.google.com/forum/#!msg/golang-nuts/lb4xLHq7wug/MhrSLkyS4F8J

func main() {
	counts := make(map[string]int)
	fileMap := make(map[string]map[string]struct{})
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileMap)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for file, _:= range fileMap[line] {
				fmt.Println(file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileMap map[string]map[string]struct{}) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if fileMap[input.Text()] == nil {
			fileMap[input.Text()] = make(map[string]struct{})
		}
		fileMap[input.Text()][f.Name()] = struct{}{}
	}
}
