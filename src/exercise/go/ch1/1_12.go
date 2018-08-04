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
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	// cyclesStr, ok := r.URL.Query()["cycles"]

	// if !ok || len(cyclesStr[0]) < 1 {
	// 	log.Println("Url Param 'cycles' is missing")
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cyclesStr := r.Form["cycles"]

	log.Println("by ParseForm(), cycles: ", cyclesStr)

	cycles, err := strconv.Atoi(cyclesStr[0])

	if err != nil {
		fmt.Fprintf(w, "invalid cycles format, cycles: %s", cyclesStr[0])
		log.Fatal("invalid cycles format")
		return
	} else {
		// cycle
		lissajous(w, cycles)
	}

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
