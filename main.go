package main

import (
	"flag"
	"fmt"
	"time"

	g "github.com/ashah360/kigen/generator"
)

// The MIT License (MIT)
// Copyright (c) 2020 ashah360
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Usage: kigen -qty=<quantity> -score=<score> -delay=<delay>

func main() {

	var qty = flag.Int("qty", 1, "quantity of codes to generate")
	var score = flag.Int("score", 1000000, "game score")
	var delay = flag.Int("delay", 800, "delay in ms between requests")

	flag.Parse()

	for n := 0; n < *qty; n++ {
		code := g.Generate(*score)
		fmt.Println(code)
		time.Sleep(time.Duration(*delay) * time.Millisecond)
	}
}
