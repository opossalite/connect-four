#include "cuda_runtime.h"
#include "device_launch_parameters.h"


/// Matrix multiplication for a layer.
__global__ void layer(float* input, float* matrix, float* output, int input_size) {
    int i = threadIdx.x;
    output[i / input_size] += input[i % input_size] * matrix[i];
}

/// Bias and normalization for a hidden layer.
__global__ void hidden_normalization(float* output, float* bias) {
    int i = threadIdx.x;
    output[i] = tanhf(output[i] + bias[i]);
}

/// Bias and normalization for output layer.
__global__ void output_normalization(float* output, float* bias) {
    int i = threadIdx.x;
    output[i] = (tanhf(output[i] + bias[i]) + 1) / 2; //sigmoid
}


extern "C" {
    /// Runs feed-forward propogation on the given input.
    void feed_forward(float* board, //42 nodes
            float* layer0, float* bias0, //84 nodes
            float* layer1, float* bias1, //21 nodes
            float* layer2, float* bias2, //7 nodes
            float* output) {
        // board (42) -> (84) -> (21) -> output (7)
        // board[42], layer0[3528], bias0[84], layer1[1764], bias1[21], layer2[147], bias2[7], output[7]
        
        float* cudaVec; //holds the left operands
        float* cudaMat; //holds the right operands
        float* cudaOut; //holds the intermediate outputs
        float* cudaBias; //holds the biases

        // move inputs into memory
        cudaMalloc(&cudaVec, 42 * sizeof(float));
        cudaMalloc(&cudaMat, 3528 * sizeof(float));
        cudaMalloc(&cudaOut, 84 * sizeof(float));
        cudaMalloc(&cudaBias, 84 * sizeof(float));
        cudaMemcpy(cudaVec, board, 42 * sizeof(float), cudaMemcpyHostToDevice);
        cudaMemcpy(cudaMat, layer0, 3528 * sizeof(float), cudaMemcpyHostToDevice);
        cudaMemcpy(cudaBias, bias0, 84 * sizeof(float), cudaMemcpyHostToDevice);

        // first layer
        layer<<<1, 3528>>>(cudaVec, cudaMat, cudaOut, 42); 
        hidden_normalization<<<1, 84>>>(cudaOut, cudaBias);

        // deallocation
        cudaFree(cudaVec);
        cudaFree(cudaMat);
        cudaFree(cudaBias);

        // prepare inputs for second layer
        cudaVec = cudaOut;
        cudaMalloc(&cudaMat, 1764 * sizeof(float));
        cudaMalloc(&cudaOut, 21 * sizeof(float));
        cudaMalloc(&cudaBias, 21 * sizeof(float));
        cudaMemcpy(cudaMat, layer1, 1764 * sizeof(float), cudaMemcpyHostToDevice);
        cudaMemcpy(cudaBias, bias1, 21 * sizeof(float), cudaMemcpyHostToDevice);

        // second layer
        layer<<<1, 1764>>>(cudaVec, cudaMat, cudaOut, 84); 
        hidden_normalization<<<1, 21>>>(cudaOut, cudaBias);

        // deallocation
        cudaFree(cudaVec);
        cudaFree(cudaMat);
        cudaFree(cudaBias);

        // prepare inputs for third layer
        cudaVec = cudaOut;
        cudaMalloc(&cudaMat, 147 * sizeof(float));
        cudaMalloc(&cudaOut, 7 * sizeof(float));
        cudaMalloc(&cudaBias, 7 * sizeof(float));
        cudaMemcpy(cudaMat, layer2, 147 * sizeof(float), cudaMemcpyHostToDevice);
        cudaMemcpy(cudaBias, bias2, 7 * sizeof(float), cudaMemcpyHostToDevice);

        // third layer
        layer<<<1, 147>>>(cudaVec, cudaMat, cudaOut, 21); 
        output_normalization<<<1, 7>>>(cudaOut, cudaBias);

        // retrieve and deallocate
        cudaMemcpy(output, cudaOut, 7 * sizeof(float), cudaMemcpyDeviceToHost);
        cudaFree(cudaVec);
        cudaFree(cudaMat);
        cudaFree(cudaOut);
        cudaFree(cudaBias);
    }
}

