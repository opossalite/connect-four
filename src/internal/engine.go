package internal

import (
	"unsafe"
)

// Return state of the engine after each turn.
type State int32
const (
    Ok State = iota //dropped piece
    Invalid //column full
    Win //win
)


// Stores information for one win tile.
type WinTile struct {
    down int32 //down
    right_asc int32 //top right
    left_asc int32 //bottom left
    left_desc int32 //top left
    right_desc int32 //bottom right
}

// Stores the Connect Four board.
//
// Uses float32 instead of int8 to remove type casting for neural networks.
type Board struct {
    TilesRed *[42]float32
    TilesYellow *[42]float32
    WinTiles *[42]WinTile //horizontal, vertical, upwards slope, downwards slope
}


// Return a new empty Board.
func NewBoard() Board {
    var board_block = [294]int32{} //allocates enough contiguous memory for our board
    var tiles_red = (*[42]float32)(unsafe.Pointer(&board_block[0]))
    var tiles_yellow = (*[42]float32)(unsafe.Pointer(&board_block[42]))
    var tiles_win = (*[42]WinTile)(unsafe.Pointer(&board_block[84]))
    return Board{TilesRed: tiles_red, TilesYellow: tiles_yellow, WinTiles: tiles_win}
}


// Drop a red piece onto the board at the specified column.
func (board Board) DropRed(column int32) (State, int32) {
    if board.TilesRed[column*7] != 0 { //column completely filled
        return Invalid, 0
    }
    var tile = drop(board.TilesRed, board.TilesYellow, column)
    var res = win_detection(board.TilesRed, board.WinTiles, tile)
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
    var res = win_detection(board.TilesRed, board.WinTiles, tile)
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


// Compare the sign of the integer.
func match(num int32, sign bool) bool {
    if sign && num > 0 {
        return true
    } else if !sign && num < 0 {
        return true
    } else {
        return false
    }
}


// Will update the game's win tiles and detect a win. Takes the tile where the last move was played.
func win_detection(board *[42]float32, win_tiles *[42]WinTile, tile int32) State {
    //left is tile-6, right is tile+6, up is tile-1, down is tile+1
    
    // down
    if tile % 6 < 5 && board[tile] == board[tile+1] { //not the bottom row
        if win_tiles[tile+1].down >= 3 {
            return Win
        }
        win_tiles[tile].down = win_tiles[tile+1].down + 1
    } 

    var left_val int32 = 0
    var right_val int32 = 0
    var edge int32 = 0 //0 = not an edge, 1 = left edge, 2 = right edge

    // ascending group (0-indexing for testing)
    if tile % 6 < 5 && tile > 5 && board[tile] == board[tile-5] { //not the bottom row nor left side
        left_val = win_tiles[tile-5].right_asc //grab the bottom left's group count
    } else { //this tile is the left edge of the ascending group
        left_val = 0
        edge = 1
    }
    if tile % 6 > 0 && tile < 36 && board[tile] == board[tile+5] { //not the top row nor right side
        right_val = win_tiles[tile+5].left_asc //grab the top right's group count
    } else { //this tile is the right edge of the ascending group
        right_val = 0
        edge = 2
    }
    if edge == 0 { //add 2 and push to both edges
        edge = left_val + right_val + 2
        if edge >= 3 {
            return Win
        }
        win_tiles[tile - (6 * (left_val+1)) + (left_val+1)].right_asc = edge
        win_tiles[tile + (6 * (right_val+1)) - (right_val+1)].left_asc = edge
    } else if edge == 1 { //add 1 and push to right edge
        edge = right_val + 1
        if edge >= 3 {
            return Win
        }
        win_tiles[tile].right_asc = right_val + 1
        win_tiles[tile + (6 * (right_val+1)) - (right_val+1)].left_asc = edge
    } else if edge == 2 { //add 1 and push to left edge
        edge = left_val + 1
        if edge >= 3 {
            return Win
        }
        win_tiles[tile - (6 * (left_val+1)) + (left_val+1)].right_asc = edge
        win_tiles[tile].left_asc = left_val + 1
    }

    left_val = 0
    right_val = 0

    // descending group (1-indexing for testing)
    if tile > 5 && tile % 6 > 0 && board[tile] == board[tile-7] { //not the left side nor top row
        left_val = win_tiles[tile-7].right_desc //grab the top left's group count
    } else { //this tile is the left edge of the descending group
        left_val = 0
    }
    if tile < 36 && tile % 6 < 5 && board[tile] == board[tile+7] { //not the right side nor bottom row
        right_val = win_tiles[tile+7].left_asc //grab the bottom right's group count
    } else { //this tile is the right edge of the descending group
        right_val = 0
    }
    if left_val == 0 { //push to right edge
        edge = right_val + 1
        if edge >= 3 {
            return Win
        }
        win_tiles[tile].right_asc = edge
        win_tiles[tile + (6*right_val) + right_val].left_asc = edge
    } else if right_val == 0 { //push to right edge
        edge = left_val + 1
        if edge >= 3 {
            return Win
        }
        win_tiles[tile - (6*left_val) - left_val].right_asc = edge
        win_tiles[tile].left_asc = edge
    } else { //push to both edges
        edge = left_val + right_val
        if edge >= 3 {
            return Win
        }
        win_tiles[tile - (6*left_val) - left_val].right_asc = edge
        win_tiles[tile + (6*right_val) + right_val].left_asc = edge
    }

    left_val = 0
    right_val = 0

    // horizontal group (1-indexing)
    if tile > 5 && board[tile] == board[tile-6] { //not left side
        left_val = win_tiles[tile-6].right_asc //grab the left's group count
    } else { //this tile is the left edge of the horizontal group
        left_val = 0
    }
    if tile < 36 && board[tile] == board[tile+6] { //not the right side
        right_val = win_tiles[tile+6].left_asc //grab the right's group count
    } else { //this tile is the right edge of the horizontal group
        right_val = 0
    }
    if left_val == 0 { //push to right edge
        edge = right_val + 1
        if edge >= 3 {
            return Win
        }
        win_tiles[tile].right_asc = edge
        win_tiles[tile + (6*right_val)].left_asc = edge
    } else if right_val == 0 { //push to right edge
        edge = left_val + 1
        if edge >= 3 {
            return Win
        }
        win_tiles[tile - (6*left_val)].right_asc = edge
        win_tiles[tile].left_asc = edge
    } else { //push to both edges
        edge = left_val + right_val + 1
        if edge >= 3 {
            return Win
        }
        win_tiles[tile - (6*left_val)].right_asc = edge
        win_tiles[tile + (6*right_val)].left_asc = edge
    }

    return Ok
}


/*

+---+---+---+---+---+---+---+
|   |   |   |   |   |   |   |
+---+---+---+---+---+---+---+
|   |   |   |   |   |   |   |
+---+---+---+---+---+---+---+
|   |   |   |   |   |   |   |
+---+---+---+---+---+---+---+
|   |   |   |   |   |   |   |
+---+---+---+---+---+---+---+
|   |   |   |   |   |   |   |
+---+---+---+---+---+---+---+
|   |   |   |   |   |   |   |
+---+---+---+---+---+---+---+

●  ○

*/






