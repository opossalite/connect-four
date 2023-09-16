package internal

/*
void feed_forward(float* board, float* layer0, float* bias0, float* layer1, float* bias1, float* layer2, float* bias2, float* output);
#cgo LDFLAGS: -L. -L../ -lneuralnet
*/
import "C"

func FeedForward(board []C.float,
        layer0 []C.float, bias0 []C.float,
        layer1 []C.float, bias1 []C.float,
        layer2 []C.float, bias2 []C.float,
        output []C.float) {
    C.feed_forward(&board[0], &layer0[0], &bias0[0], &layer1[0], &bias1[0], &layer2[0], &bias2[0], &output[0])

    /*
    void feed_forward(float* board, //42 nodes
            float* layer0, float* bias0, //84 nodes
            float* layer1, float* bias1, //21 nodes
            float* layer2, float* bias2, //7 nodes
            float* output) {
                */


}

func Bridge() {
	//in := []C.float{1.23, 4.56}
    //C.test(&in[0]) // C 1.230000 4.560000
	//a := []C.float{-1,2,4,0,5,3,6,2,1}
	//b := []C.float{3,0,2,3,4,5,4,7,2}
	//var c []C.float = make([]C.float, 9)
	//Maxmul(a,b,c,3)
	//fmt.Println(c)
}



