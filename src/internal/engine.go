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
// Uses float32 to simplify neural networks.
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


// Will update the game's win tiles and detect a win. Takes the tile where the last move was played.
func win_detection(board *[42]float32, win_tiles *[42]WinTile, tile int32) State {
    //left is tile-6, right is tile+6, up is tile-1, down is tile+1
    var res = win_detection_down(board, win_tiles, tile) ||
        win_detection_ascending(board, win_tiles, tile) ||
        win_detection_descending(board, win_tiles, tile) ||
        win_detection_horizontal(board, win_tiles, tile)
    if res {
        return Win
    }
    return Ok
}


// Win detection for tiles below.
func win_detection_down(board *[42]float32, win_tiles *[42]WinTile, tile int32) bool {
    if tile % 6 < 5 && board[tile] == board[tile+1] { //not the bottom row
        if win_tiles[tile+1].down >= 3 {
            return true 
        }
        win_tiles[tile].down = win_tiles[tile+1].down + 1
    } 
    return false
}


// Win detection for tiles in the ascending group.
func win_detection_ascending(board *[42]float32, win_tiles *[42]WinTile, tile int32) bool {
    var left_val int32  //num of tiles in this direction, 0 implies the new tile is the edge
    var right_val int32
    var edge int32 //simply stores the value to place in both edges

    if tile > 5 && tile % 6 < 5 && board[tile] == board[tile-5] { //not the left side nor bottom row
        left_val = win_tiles[tile-5].right_asc //grab the bottom left's group count
    } else { //this tile is the left edge of the ascending group
        left_val = 0
    }
    if tile < 36 && tile % 6 > 0 && board[tile] == board[tile+5] { //not the right side nor top row
        right_val = win_tiles[tile+5].left_asc //grab the top right's group count
    } else { //this tile is the right edge of the ascending group
        right_val = 0
    }
    if left_val == 0 { //push to right edge
        edge = right_val + 1
        if edge >= 3 {
            return true 
        }
        win_tiles[tile].right_asc = edge
        win_tiles[tile + (6*right_val) - right_val].left_asc = edge
    } else if right_val == 0 { //push to left edge
        edge = left_val + 1
        if edge >= 3 {
            return true 
        }
        win_tiles[tile - (6*left_val) + left_val].right_asc = edge
        win_tiles[tile].left_asc = edge
    } else { //push to both edges
        edge = left_val + right_val
        if edge >= 3 {
            return true 
        }
        win_tiles[tile - (6*left_val) + left_val].right_asc = edge
        win_tiles[tile + (6*right_val) - right_val].left_asc = edge
    }
    return false
}


// Win detection for tiles in the descending group.
func win_detection_descending(board *[42]float32, win_tiles *[42]WinTile, tile int32) bool {
    var left_val int32  //num of tiles in this direction, 0 implies the new tile is the edge
    var right_val int32
    var edge int32 //simply stores the value to place in both edges

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
            return true
        }
        win_tiles[tile].right_asc = edge
        win_tiles[tile + (6*right_val) + right_val].left_asc = edge
    } else if right_val == 0 { //push to left edge
        edge = left_val + 1
        if edge >= 3 {
            return true
        }
        win_tiles[tile - (6*left_val) - left_val].right_asc = edge
        win_tiles[tile].left_asc = edge
    } else { //push to both edges
        edge = left_val + right_val
        if edge >= 3 {
            return true
        }
        win_tiles[tile - (6*left_val) - left_val].right_asc = edge
        win_tiles[tile + (6*right_val) + right_val].left_asc = edge
    }
    return false
}


// Win detection for tiles in the horizontal group.
func win_detection_horizontal(board *[42]float32, win_tiles *[42]WinTile, tile int32) bool {
    var left_val int32  //num of tiles in this direction, 0 implies the new tile is the edge
    var right_val int32
    var edge int32 //simply stores the value to place in both edges

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
            return true
        }
        win_tiles[tile].right_asc = edge
        win_tiles[tile + (6*right_val)].left_asc = edge
    } else if right_val == 0 { //push to left edge
        edge = left_val + 1
        if edge >= 3 {
            return true
        }
        win_tiles[tile - (6*left_val)].right_asc = edge
        win_tiles[tile].left_asc = edge
    } else { //push to both edges
        edge = left_val + right_val + 1
        if edge >= 3 {
            return true
        }
        win_tiles[tile - (6*left_val)].right_asc = edge
        win_tiles[tile + (6*right_val)].left_asc = edge
    }
    return false
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






