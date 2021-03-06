// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

//var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.Black,
	color.RGBA{0x00, 0x80, 0x00, 0xFF},
	color.RGBA{0x00, 0x81, 0x00, 0xFF},
	color.RGBA{0x00, 0x82, 0x00, 0xFF},
	color.RGBA{0x00, 0x83, 0x00, 0xFF},
	color.RGBA{0x00, 0x84, 0x00, 0xFF},
	color.RGBA{0x00, 0x85, 0x00, 0xFF},
	color.RGBA{0x00, 0x86, 0x00, 0xFF},
	color.RGBA{0x00, 0x87, 0x00, 0xFF},
	color.RGBA{0x00, 0x88, 0x00, 0xFF},
	color.RGBA{0x00, 0x89, 0x00, 0xFF},
	color.RGBA{0x00, 0x8A, 0x00, 0xFF},
	color.RGBA{0x00, 0x8B, 0x00, 0xFF},
	color.RGBA{0x00, 0x8C, 0x00, 0xFF},
	color.RGBA{0x00, 0x8D, 0x00, 0xFF},
	color.RGBA{0x00, 0x8E, 0x00, 0xFF},
	color.RGBA{0x00, 0x8F, 0x00, 0xFF},
	color.RGBA{0x00, 0x90, 0x00, 0xFF},
	color.RGBA{0x00, 0x91, 0x00, 0xFF},
	color.RGBA{0x00, 0x92, 0x00, 0xFF},
	color.RGBA{0x00, 0x93, 0x00, 0xFF},
	color.RGBA{0x00, 0x94, 0x00, 0xFF},
	color.RGBA{0x00, 0x95, 0x00, 0xFF},
	color.RGBA{0x00, 0x96, 0x00, 0xFF},
	color.RGBA{0x00, 0x97, 0x00, 0xFF},
	color.RGBA{0x00, 0x98, 0x00, 0xFF},
	color.RGBA{0x00, 0x99, 0x00, 0xFF},
	color.RGBA{0x00, 0x9A, 0x00, 0xFF},
	color.RGBA{0x00, 0x9B, 0x00, 0xFF},
	color.RGBA{0x00, 0x9C, 0x00, 0xFF},
	color.RGBA{0x00, 0x9D, 0x00, 0xFF},
	color.RGBA{0x00, 0x9E, 0x00, 0xFF},
	color.RGBA{0x00, 0x9F, 0x00, 0xFF},
	color.RGBA{0x00, 0xA0, 0x00, 0xFF},
	color.RGBA{0x00, 0xA1, 0x00, 0xFF},
	color.RGBA{0x00, 0xA2, 0x00, 0xFF},
	color.RGBA{0x00, 0xA3, 0x00, 0xFF},
	color.RGBA{0x00, 0xA4, 0x00, 0xFF},
	color.RGBA{0x00, 0xA5, 0x00, 0xFF},
	color.RGBA{0x00, 0xA6, 0x00, 0xFF},
	color.RGBA{0x00, 0xA7, 0x00, 0xFF},
	color.RGBA{0x00, 0xA8, 0x00, 0xFF},
	color.RGBA{0x00, 0xA9, 0x00, 0xFF},
	color.RGBA{0x00, 0xAA, 0x00, 0xFF},
	color.RGBA{0x00, 0xAB, 0x00, 0xFF},
	color.RGBA{0x00, 0xAC, 0x00, 0xFF},
	color.RGBA{0x00, 0xAD, 0x00, 0xFF},
	color.RGBA{0x00, 0xAE, 0x00, 0xFF},
	color.RGBA{0x00, 0xAF, 0x00, 0xFF},
	color.RGBA{0x00, 0xB0, 0x00, 0xFF},
	color.RGBA{0x00, 0xB1, 0x00, 0xFF},
	color.RGBA{0x00, 0xB2, 0x00, 0xFF},
	color.RGBA{0x00, 0xB3, 0x00, 0xFF},
	color.RGBA{0x00, 0xB4, 0x00, 0xFF},
	color.RGBA{0x00, 0xB5, 0x00, 0xFF},
	color.RGBA{0x00, 0xB6, 0x00, 0xFF},
	color.RGBA{0x00, 0xB7, 0x00, 0xFF},
	color.RGBA{0x00, 0xB8, 0x00, 0xFF},
	color.RGBA{0x00, 0xB9, 0x00, 0xFF},
	color.RGBA{0x00, 0xBA, 0x00, 0xFF},
	color.RGBA{0x00, 0xBB, 0x00, 0xFF},
	color.RGBA{0x00, 0xBC, 0x00, 0xFF},
	color.RGBA{0x00, 0xBD, 0x00, 0xFF},
	color.RGBA{0x00, 0xBE, 0x00, 0xFF},
	color.RGBA{0x00, 0xBF, 0x00, 0xFF},
	color.RGBA{0x00, 0xC0, 0x00, 0xFF},
	color.RGBA{0x00, 0xC1, 0x00, 0xFF},
	color.RGBA{0x00, 0xC2, 0x00, 0xFF},
	color.RGBA{0x00, 0xC3, 0x00, 0xFF},
	color.RGBA{0x00, 0xC4, 0x00, 0xFF},
	color.RGBA{0x00, 0xC5, 0x00, 0xFF},
	color.RGBA{0x00, 0xC6, 0x00, 0xFF},
	color.RGBA{0x00, 0xC7, 0x00, 0xFF},
	color.RGBA{0x00, 0xC8, 0x00, 0xFF},
	color.RGBA{0x00, 0xC9, 0x00, 0xFF},
	color.RGBA{0x00, 0xCA, 0x00, 0xFF},
	color.RGBA{0x00, 0xCB, 0x00, 0xFF},
	color.RGBA{0x00, 0xCC, 0x00, 0xFF},
	color.RGBA{0x00, 0xCD, 0x00, 0xFF},
	color.RGBA{0x00, 0xCE, 0x00, 0xFF},
	color.RGBA{0x00, 0xCF, 0x00, 0xFF},
	color.RGBA{0x00, 0xD0, 0x00, 0xFF},
	color.RGBA{0x00, 0xD1, 0x00, 0xFF},
	color.RGBA{0x00, 0xD2, 0x00, 0xFF},
	color.RGBA{0x00, 0xD3, 0x00, 0xFF},
	color.RGBA{0x00, 0xD4, 0x00, 0xFF},
	color.RGBA{0x00, 0xD5, 0x00, 0xFF},
	color.RGBA{0x00, 0xD6, 0x00, 0xFF},
	color.RGBA{0x00, 0xD7, 0x00, 0xFF},
	color.RGBA{0x00, 0xD8, 0x00, 0xFF},
	color.RGBA{0x00, 0xD9, 0x00, 0xFF},
	color.RGBA{0x00, 0xDA, 0x00, 0xFF},
	color.RGBA{0x00, 0xDB, 0x00, 0xFF},
	color.RGBA{0x00, 0xDC, 0x00, 0xFF},
	color.RGBA{0x00, 0xDD, 0x00, 0xFF},
	color.RGBA{0x00, 0xDE, 0x00, 0xFF},
	color.RGBA{0x00, 0xDF, 0x00, 0xFF},
	color.RGBA{0x00, 0xE0, 0x00, 0xFF},
	color.RGBA{0x00, 0xE1, 0x00, 0xFF},
	color.RGBA{0x00, 0xE2, 0x00, 0xFF},
	color.RGBA{0x00, 0xE3, 0x00, 0xFF},
	color.RGBA{0x00, 0xE4, 0x00, 0xFF},
	color.RGBA{0x00, 0xE5, 0x00, 0xFF},
	color.RGBA{0x00, 0xE6, 0x00, 0xFF},
	color.RGBA{0x00, 0xE7, 0x00, 0xFF},
	color.RGBA{0x00, 0xE8, 0x00, 0xFF},
	color.RGBA{0x00, 0xE9, 0x00, 0xFF},
	color.RGBA{0x00, 0xEA, 0x00, 0xFF},
	color.RGBA{0x00, 0xEB, 0x00, 0xFF},
	color.RGBA{0x00, 0xEC, 0x00, 0xFF},
	color.RGBA{0x00, 0xED, 0x00, 0xFF},
	color.RGBA{0x00, 0xEE, 0x00, 0xFF},
	color.RGBA{0x00, 0xEF, 0x00, 0xFF},
	color.RGBA{0x00, 0xF0, 0x00, 0xFF},
	color.RGBA{0x00, 0xF1, 0x00, 0xFF},
	color.RGBA{0x00, 0xF2, 0x00, 0xFF},
	color.RGBA{0x00, 0xF3, 0x00, 0xFF},
	color.RGBA{0x00, 0xF4, 0x00, 0xFF},
	color.RGBA{0x00, 0xF5, 0x00, 0xFF},
	color.RGBA{0x00, 0xF6, 0x00, 0xFF},
	color.RGBA{0x00, 0xF7, 0x00, 0xFF},
	color.RGBA{0x00, 0xF8, 0x00, 0xFF},
	color.RGBA{0x00, 0xF9, 0x00, 0xFF},
	color.RGBA{0x00, 0xFA, 0x00, 0xFF},
	color.RGBA{0x00, 0xFB, 0x00, 0xFF},
	color.RGBA{0x00, 0xFC, 0x00, 0xFF},
	color.RGBA{0x00, 0xFD, 0x00, 0xFF},
	color.RGBA{0x00, 0xFE, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0xFE, 0x00, 0xFF},
	color.RGBA{0x00, 0xFD, 0x00, 0xFF},
	color.RGBA{0x00, 0xFC, 0x00, 0xFF},
	color.RGBA{0x00, 0xFB, 0x00, 0xFF},
	color.RGBA{0x00, 0xFA, 0x00, 0xFF},
	color.RGBA{0x00, 0xF9, 0x00, 0xFF},
	color.RGBA{0x00, 0xF8, 0x00, 0xFF},
	color.RGBA{0x00, 0xF7, 0x00, 0xFF},
	color.RGBA{0x00, 0xF6, 0x00, 0xFF},
	color.RGBA{0x00, 0xF5, 0x00, 0xFF},
	color.RGBA{0x00, 0xF4, 0x00, 0xFF},
	color.RGBA{0x00, 0xF3, 0x00, 0xFF},
	color.RGBA{0x00, 0xF2, 0x00, 0xFF},
	color.RGBA{0x00, 0xF1, 0x00, 0xFF},
	color.RGBA{0x00, 0xF0, 0x00, 0xFF},
	color.RGBA{0x00, 0xEF, 0x00, 0xFF},
	color.RGBA{0x00, 0xEE, 0x00, 0xFF},
	color.RGBA{0x00, 0xED, 0x00, 0xFF},
	color.RGBA{0x00, 0xEC, 0x00, 0xFF},
	color.RGBA{0x00, 0xEB, 0x00, 0xFF},
	color.RGBA{0x00, 0xEA, 0x00, 0xFF},
	color.RGBA{0x00, 0xE9, 0x00, 0xFF},
	color.RGBA{0x00, 0xE8, 0x00, 0xFF},
	color.RGBA{0x00, 0xE7, 0x00, 0xFF},
	color.RGBA{0x00, 0xE6, 0x00, 0xFF},
	color.RGBA{0x00, 0xE5, 0x00, 0xFF},
	color.RGBA{0x00, 0xE4, 0x00, 0xFF},
	color.RGBA{0x00, 0xE3, 0x00, 0xFF},
	color.RGBA{0x00, 0xE2, 0x00, 0xFF},
	color.RGBA{0x00, 0xE1, 0x00, 0xFF},
	color.RGBA{0x00, 0xE0, 0x00, 0xFF},
	color.RGBA{0x00, 0xDF, 0x00, 0xFF},
	color.RGBA{0x00, 0xDE, 0x00, 0xFF},
	color.RGBA{0x00, 0xDD, 0x00, 0xFF},
	color.RGBA{0x00, 0xDC, 0x00, 0xFF},
	color.RGBA{0x00, 0xDB, 0x00, 0xFF},
	color.RGBA{0x00, 0xDA, 0x00, 0xFF},
	color.RGBA{0x00, 0xD9, 0x00, 0xFF},
	color.RGBA{0x00, 0xD8, 0x00, 0xFF},
	color.RGBA{0x00, 0xD7, 0x00, 0xFF},
	color.RGBA{0x00, 0xD6, 0x00, 0xFF},
	color.RGBA{0x00, 0xD5, 0x00, 0xFF},
	color.RGBA{0x00, 0xD4, 0x00, 0xFF},
	color.RGBA{0x00, 0xD3, 0x00, 0xFF},
	color.RGBA{0x00, 0xD2, 0x00, 0xFF},
	color.RGBA{0x00, 0xD1, 0x00, 0xFF},
	color.RGBA{0x00, 0xD0, 0x00, 0xFF},
	color.RGBA{0x00, 0xCF, 0x00, 0xFF},
	color.RGBA{0x00, 0xCE, 0x00, 0xFF},
	color.RGBA{0x00, 0xCD, 0x00, 0xFF},
	color.RGBA{0x00, 0xCC, 0x00, 0xFF},
	color.RGBA{0x00, 0xCB, 0x00, 0xFF},
	color.RGBA{0x00, 0xCA, 0x00, 0xFF},
	color.RGBA{0x00, 0xC9, 0x00, 0xFF},
	color.RGBA{0x00, 0xC8, 0x00, 0xFF},
	color.RGBA{0x00, 0xC7, 0x00, 0xFF},
	color.RGBA{0x00, 0xC6, 0x00, 0xFF},
	color.RGBA{0x00, 0xC5, 0x00, 0xFF},
	color.RGBA{0x00, 0xC4, 0x00, 0xFF},
	color.RGBA{0x00, 0xC3, 0x00, 0xFF},
	color.RGBA{0x00, 0xC2, 0x00, 0xFF},
	color.RGBA{0x00, 0xC1, 0x00, 0xFF},
	color.RGBA{0x00, 0xC0, 0x00, 0xFF},
	color.RGBA{0x00, 0xBF, 0x00, 0xFF},
	color.RGBA{0x00, 0xBE, 0x00, 0xFF},
	color.RGBA{0x00, 0xBD, 0x00, 0xFF},
	color.RGBA{0x00, 0xBC, 0x00, 0xFF},
	color.RGBA{0x00, 0xBB, 0x00, 0xFF},
	color.RGBA{0x00, 0xBA, 0x00, 0xFF},
	color.RGBA{0x00, 0xB9, 0x00, 0xFF},
	color.RGBA{0x00, 0xB8, 0x00, 0xFF},
	color.RGBA{0x00, 0xB7, 0x00, 0xFF},
	color.RGBA{0x00, 0xB6, 0x00, 0xFF},
	color.RGBA{0x00, 0xB5, 0x00, 0xFF},
	color.RGBA{0x00, 0xB4, 0x00, 0xFF},
	color.RGBA{0x00, 0xB3, 0x00, 0xFF},
	color.RGBA{0x00, 0xB2, 0x00, 0xFF},
	color.RGBA{0x00, 0xB1, 0x00, 0xFF},
	color.RGBA{0x00, 0xB0, 0x00, 0xFF},
	color.RGBA{0x00, 0xAF, 0x00, 0xFF},
	color.RGBA{0x00, 0xAE, 0x00, 0xFF},
	color.RGBA{0x00, 0xAD, 0x00, 0xFF},
	color.RGBA{0x00, 0xAC, 0x00, 0xFF},
	color.RGBA{0x00, 0xAB, 0x00, 0xFF},
	color.RGBA{0x00, 0xAA, 0x00, 0xFF},
	color.RGBA{0x00, 0xA9, 0x00, 0xFF},
	color.RGBA{0x00, 0xA8, 0x00, 0xFF},
	color.RGBA{0x00, 0xA7, 0x00, 0xFF},
	color.RGBA{0x00, 0xA6, 0x00, 0xFF},
	color.RGBA{0x00, 0xA5, 0x00, 0xFF},
	color.RGBA{0x00, 0xA4, 0x00, 0xFF},
	color.RGBA{0x00, 0xA3, 0x00, 0xFF},
	color.RGBA{0x00, 0xA2, 0x00, 0xFF},
	color.RGBA{0x00, 0xA1, 0x00, 0xFF},
	color.RGBA{0x00, 0xA0, 0x00, 0xFF},
	color.RGBA{0x00, 0x9F, 0x00, 0xFF},
	color.RGBA{0x00, 0x9E, 0x00, 0xFF},
	color.RGBA{0x00, 0x9D, 0x00, 0xFF},
	color.RGBA{0x00, 0x9C, 0x00, 0xFF},
	color.RGBA{0x00, 0x9B, 0x00, 0xFF},
	color.RGBA{0x00, 0x9A, 0x00, 0xFF},
	color.RGBA{0x00, 0x99, 0x00, 0xFF},
	color.RGBA{0x00, 0x98, 0x00, 0xFF},
	color.RGBA{0x00, 0x97, 0x00, 0xFF},
	color.RGBA{0x00, 0x96, 0x00, 0xFF},
	color.RGBA{0x00, 0x95, 0x00, 0xFF},
	color.RGBA{0x00, 0x94, 0x00, 0xFF},
	color.RGBA{0x00, 0x93, 0x00, 0xFF},
	color.RGBA{0x00, 0x92, 0x00, 0xFF},
	color.RGBA{0x00, 0x91, 0x00, 0xFF},
	color.RGBA{0x00, 0x90, 0x00, 0xFF},
	color.RGBA{0x00, 0x8F, 0x00, 0xFF},
	color.RGBA{0x00, 0x8E, 0x00, 0xFF},
	color.RGBA{0x00, 0x8D, 0x00, 0xFF},
	color.RGBA{0x00, 0x8C, 0x00, 0xFF},
	color.RGBA{0x00, 0x8B, 0x00, 0xFF},
	color.RGBA{0x00, 0x8A, 0x00, 0xFF},
	color.RGBA{0x00, 0x89, 0x00, 0xFF},
	color.RGBA{0x00, 0x88, 0x00, 0xFF},
	color.RGBA{0x00, 0x87, 0x00, 0xFF},
	color.RGBA{0x00, 0x86, 0x00, 0xFF},
	color.RGBA{0x00, 0x85, 0x00, 0xFF},
	color.RGBA{0x00, 0x84, 0x00, 0xFF},
	color.RGBA{0x00, 0x83, 0x00, 0xFF},
	color.RGBA{0x00, 0x82, 0x00, 0xFF},
	color.RGBA{0x00, 0x81, 0x00, 0xFF},
}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		//		img := image.NewPaletted(rect, palette)
		img := image.NewPaletted(rect, palette)
		var index uint8
		index = 1
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				index)
			index++
			if len(palette) == int(index) {
				index = 1
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
