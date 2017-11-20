package main

import (
  "github.com/nfnt/resize"
  "image"
  "image/jpeg"
  "os"
  "flag"
  "fmt"
)

func main() {

  // Get input file
  inputFile := os.Args[3]

  // Set output file
  outputFile := os.Args[4]

  // Check quality
  quality := flag.Int("q", 50, "an int")
  width := flag.Int("w", 400, "an int")
  flag.Parse()

  // Debug log
  fmt.Print(inputFile, outputFile, *quality, *width)

  // Open input file and decode
  fImg1, _ := os.Open(inputFile)
  defer fImg1.Close()
  img1, _, _ := image.Decode(fImg1)

  // Create output file
  toimg, _ := os.Create(outputFile)
  defer toimg.Close()

  img2 := resize.Resize(uint(*width), 0, img1, resize.Lanczos3)

  // Encode
  jpeg.Encode(toimg, img2, &jpeg.Options{*quality})
}
