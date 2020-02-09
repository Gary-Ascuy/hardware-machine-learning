// Author: Gary Ascuy <gary.ascuy@gmail.com>
// Example: go run main.go --histo --image assets/orange.png
// Yellow = {255, 255, 0}, Red = {255, 0, 0}

package main

import (
  "os"
  "fmt"
  "flag"
  "math"
  "image"
  _ "image/gif"
  _ "image/jpeg"
  _ "image/png"
)

// Command Line Interface - Inputs
var imagePath = flag.String("image", "./assets/orange.png", "Path to image (gif, jpeg, png)")
var histogramFlag = flag.Bool("histo", false, "Show image histogram")

// Main Entry Point
func main() {
  flag.Parse()
  fmt.Print("Loading: ", *imagePath)

  if infile, err := os.Open(*imagePath); err == nil {
    defer infile.Close()

    if img, _, imgErr := image.Decode(infile); imgErr == nil {
      fmt.Println(", Ready !!!")
      bounds := img.Bounds()

      // Histogram
      if *histogramFlag { Histogram(bounds, img) }

      // Predict
      Predict(bounds, img)
    } else { flag.Parse() }
  } else { flag.Parse() }
}

// Prints Predict information
func Predict(bounds image.Rectangle, img image.Image) {
  var yellow uint32 = 0
  var red uint32 = 0
  var half uint32 = 32767

  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      r, g, b, _ := img.At(x, y).RGBA()
      if r == g && g == b { continue }

      if r > half && g > half && b < half { yellow++ }
      if r > half && g < half && b < half { red++ }
    }
  }

  redMean := float64(red) / (float64(red) + float64(yellow))
  fmt.Println("\n\n{ RED =", red, ", YELLOW =", yellow, "}, RED =", math.Round(redMean * 100), "%")
  fmt.Println("Prediction: It is an", TernaryOperator(redMean > 0.5, "STRAWBERRY", "ORANGE"), "!\n")
}

// Prints Image Histogram
func Histogram(bounds image.Rectangle, img image.Image) {
  var histogram [16][4]int
  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      r, g, b, a := img.At(x, y).RGBA()
      histogram[r>>12][0]++
      histogram[g>>12][1]++
      histogram[b>>12][2]++
      histogram[a>>12][3]++
    }
  }

  // Print the results.
  fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
  for i, x := range histogram {
    fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
  }
}

func TernaryOperator(expresion bool, a string, b string) string {
  if (expresion) { return a }
  return b
}
