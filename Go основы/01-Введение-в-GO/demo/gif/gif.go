// Run with "web" command-line argument for web server.
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
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // Первый цвет палитры
	blackIndex = 1 // Следующий цвет палитры
)

func main() {
	//!-main
	// Последовательность изображений является детерминированной,
	// если мы не используем случайное зерно
	// Генератор псевдослучайных чисел с использованием текущего времени.
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // количество колебаний
		res     = 0.001 // Угловое разрашение
		size    = 100   // Размер изображения [-size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами
	)
	freq := rand.Float64() * 3.0 // Относительная частота колебаний
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
