package main

import (
	"image"
  "image/jpeg"
	"os"
  "flag"
	"fmt"
)

func main() {

  inputFile := os.Args[2]
  outputFile := os.Args[3]
  quality := flag.Int("q", 50, "an int")
	flag.Parse()
	fmt.Print(*quality)
  fImg1, _ := os.Open(inputFile)
  defer fImg1.Close()
  img1, _, _ := image.Decode(fImg1)

  toimg, _ := os.Create(outputFile)
  defer toimg.Close()

  jpeg.Encode(toimg, img1, &jpeg.Options{*quality})
}
