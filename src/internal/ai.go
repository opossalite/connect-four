package internal

/*
void feed_forward(float* board, float* layer0, float* bias0, float* layer1, float* bias1, float* layer2, float* bias2, float* output);
#cgo LDFLAGS: -L. -L../ -lneuralnet
*/
import "C"

// Run inputs on feed-forward neural network (via CUDA).
func FeedForward(board []C.float,
        layer0 []C.float, bias0 []C.float,
        layer1 []C.float, bias1 []C.float,
        layer2 []C.float, bias2 []C.float,
        output []C.float) {
    C.feed_forward(&board[0], &layer0[0], &bias0[0], &layer1[0], &bias1[0], &layer2[0], &bias2[0], &output[0])
}

