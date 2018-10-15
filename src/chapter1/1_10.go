package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// キャッシュされているか調査するために、出力をファイルに保存する

// 出力結果を見るとあまり時間変わらないのでキャッシュされていない気がする

// 0.81s  464317 https://www.amazon.co.jp/
// 0.81s elapsed

// 1.02s  486025 https://www.amazon.co.jp/
// 1.02s elapsed

func main() {
	start := time.Now()

	// ファイルの生成
	fileName := os.Args[1]
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("can't create file: %s\n", fileName)
		os.Exit(1)
	}

	ch := make(chan string)
	for _, url := range os.Args[2:] {
		go fetch(url, ch)
	}
	for range os.Args[2:] {
		// Fprintlnでファイルに出力
		fmt.Fprintln(f, <-ch)
	}
	// Fprintfでファイルに出力
	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}