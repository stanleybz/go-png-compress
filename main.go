package main

import (
	"image"
  "image/jpeg"
	"os"
  "flag"
)

func main() {

  inputFile := os.Args[1]
  outputFile := os.Args[2]
  quality := flag.Int("q", 50, "an int")

  fImg1, _ := os.Open(inputFile)
  defer fImg1.Close()
  img1, _, _ := image.Decode(fImg1)

  toimg, _ := os.Create(outputFile)
  defer toimg.Close()

  jpeg.Encode(toimg, img1, &jpeg.Options{*quality})
}
