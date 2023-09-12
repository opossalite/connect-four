package internal

import "unsafe"

// Return state of the engine after each turn.
type State int32
const (
    Ok State = iota //dropped piece
    Invalid //column full
    Win //win
    WinRed
    WinYellow
)


// Stores the Connect Four board.
//
// Uses float32 instead of int8 to remove type casting for neural networks.
type Board struct {
    TilesRed *[42]float32
    TilesYellow *[42]float32
    WinTiles *[168]int32
}


// Return a new empty Board.
func NewBoard() Board {
    var board_block = [252]int32{} //allocates enough contiguous memory for our board
    var tiles_red = (*[42]float32)(unsafe.Pointer(&board_block[0]))
    var tiles_yellow = (*[42]float32)(unsafe.Pointer(&board_block[42]))
    var tiles_win = (*[168]int32)(unsafe.Pointer(&board_block[84]))
    return Board{TilesRed: tiles_red, TilesYellow: tiles_yellow, WinTiles: tiles_win}
}


// Drop a red piece onto the board at the specified column.
func (board Board) DropRed(column int32) (State, int32) {
    if board.TilesRed[column*7] != 0 { //column completely filled
        return Invalid, 0
    }
    var tile = drop(board.TilesRed, board.TilesYellow, column)
    var res = win_detection(board.TilesRed, board.WinTiles)
    if res != Ok { //not a non-win, so a player has won
        return Win, tile
    }
    return Ok, tile
}


// Drop a yellow piece onto the board at the specified column.
func (board Board) DropYellow(column int32) (State, int32) {
    if board.TilesYellow[column*7] != 0 { //column completely filled
        return Invalid, 0
    }
    var tile = drop(board.TilesYellow, board.TilesRed, column)
    var res = win_detection(board.TilesRed, board.WinTiles)
    if res != Ok { //not a non-win, so a player has won
        return Win, tile
    }
    return Ok, tile
}


// Handles piece dropping regardless of color. Returns the dropped tile's location.
func drop(board_main *[42]float32, board_secondary *[42]float32, column int32) int32 {
    var main_slice = board_main[column*6:column*6+6]
    var secondary_slice = board_secondary[column*6:column*6+6]
    for i := int32(1); i < 6; i++ {
        if main_slice[i] != 0 { //tile above is open
            main_slice[i-1] = 1
            secondary_slice[i-1] = -1
            return column*6 + i-1
        }
    }
    // column was completely open
    main_slice[5] = 1
    secondary_slice[5] = -1
    return column*6 + 5
}


// Will update the game's win tiles and detect a win.
func win_detection(board *[42]float32, win_tiles *[168]int32) State {
    return Ok
}



