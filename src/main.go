package main

import "fmt"
import "connect-four/internal"

func main() {
    internal.Bridge()

    y := make([]float32, 42)
    z := make([]float32, 42)
    x := internal.Board{y, z}
    fmt.Println(x.TilesRed[0])
}


