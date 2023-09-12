package internal


// Return state of the engine after each turn.
type State int32
const (
    Ok State = iota //dropped piece
    Invalid //column full
    Win //win
)


// Stores the Connect Four board.
//
// Uses float32 instead of int8 to remove type casting for neural networks.
type Board struct {
    TilesRed *[42]float32
    TilesYellow *[42]float32
    WinTiles *[168]int32
}


// Drop a red piece onto the board at the specified column.
func (board Board) DropRed(column int32) (State, int32) {
    if board.TilesRed[column*7] != 0 { //column completely filled
        return Invalid, 0
    }
    var tile = drop(board.TilesRed, board.TilesYellow, column)
    //win detection here
    return Ok, tile
}


// Drop a yellow piece onto the board at the specified column.
func (board Board) DropYellow(column int32) (State, int32) {
    if board.TilesYellow[column*7] != 0 { //column completely filled
        return Invalid, 0
    }
    var tile = drop(board.TilesYellow, board.TilesRed, column)
    //win detection here
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




