package engine 

// Stores the Connect Four board.
//
// Uses float32 instead of int8 to remove type casting for neural networks.
type Board struct {
    TilesRed []float32
    TilesYellow []float32
}

//func test() {
//}






