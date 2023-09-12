package main

import "fmt"
import "connect-four/internal"

func main() {
    internal.Bridge()

    board := internal.NewBoard()

    var state, tile = board.DropRed(0)
    if state == internal.Invalid {
        fmt.Println("Invalid drop!")
        return
    }
    fmt.Printf("Dropped into %d\n", tile)

    state, tile = board.DropRed(0)
    if state == internal.Invalid {
        fmt.Println("Invalid drop!")
        return
    }
    fmt.Printf("Dropped into %d\n", tile)

    state, tile = board.DropRed(0)
    if state == internal.Invalid {
        fmt.Println("Invalid drop!")
        return
    }
    fmt.Printf("Dropped into %d\n", tile)

    state, tile = board.DropRed(0)
    if state == internal.Invalid {
        fmt.Println("Invalid drop!")
        return
    }
    fmt.Printf("Dropped into %d\n", tile)

    state, tile = board.DropRed(0)
    if state == internal.Invalid {
        fmt.Println("Invalid drop!")
        return
    }
    fmt.Printf("Dropped into %d\n", tile)

    state, tile = board.DropRed(0)
    if state == internal.Invalid {
        fmt.Println("Invalid drop!")
        return
    }
    fmt.Printf("Dropped into %d\n", tile)

    state, tile = board.DropRed(0)
    if state == internal.Invalid {
        fmt.Println("Invalid drop!")
        return
    }
    fmt.Printf("Dropped into %d\n", tile)

    ////fmt.Println(fmt.Parse("Dropped into {{.tile}}"))
    //fmt.Printf("Red: %f\n", board.TilesRed)
    //fmt.Println("Red: ", board.TilesRed)
    //fmt.Println("Red1: ", board.TilesRed[5])
    //fmt.Printf("Yellow: %f\n", board.TilesYellow)

    //b := [10]float32{}
    //c := &b
    //d := c[2:5]
    ////(*c)[0] = 1
    //d[0] = 1
    //fmt.Printf("Test: %f\n", b)
}


