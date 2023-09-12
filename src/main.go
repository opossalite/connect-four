package main

import "fmt"
import "unsafe"
import "connect-four/internal"

func main() {
    internal.Bridge()

    //y := make([]float32, 42)
    //z := make([]float32, 42)
    //y := [42]float32{}
    //z := [42]float32{}
    //a := [168]int32{}
    var board_block = [252]int32{} //allocates enough contiguous memory for our board
    var tiles_red = (*[42]float32)(unsafe.Pointer(&board_block[0]))
    var tiles_yellow = (*[42]float32)(unsafe.Pointer(&board_block[42]))
    var tiles_win = (*[168]int32)(unsafe.Pointer(&board_block[84]))
    board := internal.Board{TilesRed: tiles_red, TilesYellow: tiles_yellow, WinTiles: tiles_win}

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


