# Connect Four AI
A self-learning AI for Connect Four. Additionally, player vs player functionality is available.

## Training Method
This AI is trained using a genetic algorithm, complete with mutations and sexual reproduction. The AI runs partially on CUDA, so an NVIDIA card is required (possible CPU implementation later).

## Usage
Ensure CUDA and its development tools are installed, as well as Go.<br>
In the src/ directory run these:
```
nvcc --ptxas-options=-v --compiler-options '-fPIC' -o libai_tools.so --shared ai_tools.cu
LD_LIBRARY_PATH=. go run main.go
```

## TODO
- Add a temporary dummy AI implementation
- Add AI genetics
- Add Architect which handles generations of AIs, their reproduction, competition, and game playing
- Add an ncurses menu and game
- Add CUDA neural network implementation
- Add optional CPU neural network implementation


