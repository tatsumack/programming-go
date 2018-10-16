package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// URLからパラメータ値を読み取れるようにする

var palette = []color.Color{color.White, color.Black}

const (
	whilteIndex = 0
	blackIndex  = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// クエリを渡す
		lissajous(w, r.URL.Query())
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, query url.Values) {
	var (
		cycles  = 5.0
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	// クエリがセットされていた場合は上書きする
	for k, vals := range query {
		v := vals[0]
		switch k {
		case "cycles":
			cycles, _ = strconv.ParseFloat(v, 64)
		case "res":
			res, _ = strconv.ParseFloat(v, 64)
		case "size":
			size, _ = strconv.Atoi(v)
		case "nframes":
			nframes, _ = strconv.Atoi(v)
		case "delay":
			delay, _ = strconv.Atoi(v)
		}
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
