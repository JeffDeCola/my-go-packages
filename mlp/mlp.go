package mlp

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Neural Network Parameters
type NeuralNetworkParameters struct {
	InputNodes          int
	InputNodeLabels     []string
	HiddenLayers        int
	HiddenNodesPerLayer []int
	OutputNodes         int
	OutputNodeLabels    []string
	LearningRate        float64
	Epochs              int
	DatasetCSVFile      string
}

// Neural network structure with up to 4 hidden layers
type NeuralNetwork struct {
	inputNodes          int           // INPUT NODES (USER ADDED)
	inputNodeLabels     []string      // INPUT NODE LABELS (USER ADDED)
	hiddenLayers        int           // HIDDEN LAYERS (USER ADDED)
	hiddenNodesPerLayer []int         // HIDDEN NODES per [hiddenLayer] (USER ADDED)
	hiddenWeights       [][][]float64 // - weights per [hiddenLayer][hiddenNode][inputNode]
	hiddenBias          [][]float64   // - bias per [hiddenLayer][hiddenNode]
	outputNodes         int           // OUTPUT NODES (USER ADDED)
	outputNodeLabels    []string      // OUTPUT NODE LABELS (USER ADDED)
	outputWeights       [][]float64   // - weights per [outputNode][hiddenNode]
	outputBias          []float64     // - bias per [outputNode]
	learningRate        float64       // LEARNING RATE (USER ADDED)
	epochs              int           // EPOCHS (USER ADDED)
	minInput            []float64     // MIN INPUT VALUE
	maxInput            []float64     // MAX INPUT VALUE
	datasetCSVFile      string        // DATASET CSV FILE
}

// Data structure for training data IO
type trainingData struct {
	i []float64
	z []float64
}

// Create the MLP neural network
func (nnp NeuralNetworkParameters) CreateNeuralNetwork() *NeuralNetwork {

	// Initialize hiddenWeights slice
	// hiddenWeights[hiddenLayers#][hiddenNodesPerLayer#][inputNodesNumber or hiddenNodesPerLayerNumber]
	hiddenWeights := make([][][]float64, nnp.HiddenLayers)
	for i := range hiddenWeights {
		hiddenWeights[i] = make([][]float64, nnp.HiddenNodesPerLayer[i])
		for j := range hiddenWeights[i] {
			if i == 0 {
				hiddenWeights[i][j] = make([]float64, nnp.InputNodes)
			} else {
				hiddenWeights[i][j] = make([]float64, nnp.HiddenNodesPerLayer[i-1])
			}
		}
	}

	// Initialize hiddenBias slice
	// hiddenBias[hiddenLayersNumber][hiddenNodesPerLayerNumber]
	hiddenBias := make([][]float64, nnp.HiddenLayers)
	for i := range hiddenBias {
		hiddenBias[i] = make([]float64, nnp.HiddenNodesPerLayer[i])
	}

	// Initialize outputWeights slice
	// outputWeights[outputNodesNumber][hiddenNodesPerLayerNumber of last hidden layer]
	outputWeights := make([][]float64, nnp.OutputNodes)
	for i := range outputWeights {
		outputWeights[i] = make([]float64, nnp.HiddenNodesPerLayer[nnp.HiddenLayers-1])
	}

	// Initialize outputBias slice
	// outputBias[outputNodesNumber]
	outputBias := make([]float64, nnp.OutputNodes)

	//Initialize the min and max input values
	minInput := make([]float64, nnp.InputNodes)
	maxInput := make([]float64, nnp.InputNodes)

	// Create the neural network
	nn := &NeuralNetwork{
		inputNodes:          nnp.InputNodes,          // USER PROVIDED
		inputNodeLabels:     nnp.InputNodeLabels,     // USER PROVIDED
		hiddenLayers:        nnp.HiddenLayers,        // USER PROVIDED
		hiddenNodesPerLayer: nnp.HiddenNodesPerLayer, // USER PROVIDED
		hiddenWeights:       hiddenWeights,           // - created here
		hiddenBias:          hiddenBias,              // - created here
		outputNodes:         nnp.OutputNodes,         // USER PROVIDED
		outputNodeLabels:    nnp.OutputNodeLabels,    // USER PROVIDED
		outputWeights:       outputWeights,           // - created here
		outputBias:          outputBias,              // - created here
		learningRate:        nnp.LearningRate,        // UER PROVIDED
		epochs:              nnp.Epochs,              // USER PROVIDED
		minInput:            minInput,                // - created here
		maxInput:            maxInput,                // - created here
		datasetCSVFile:      nnp.DatasetCSVFile,      // USER PROVIDED
	}

	return nn
}

// Initialize the neural network (weights and bias) from values -1 to 1
func (nn *NeuralNetwork) InitializeNeuralNetwork() {

	// Random number generator from 0-1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// INIT HIDDEN LAYER(s)
	// l is the hidden layer number
	for l := 0; l < nn.hiddenLayers; l++ {
		// h is the hidden node number
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			// in is the input/hidden node number
			if l == 0 {
				for in := 0; in < nn.inputNodes; in++ {
					nn.hiddenWeights[l][hn][in] = r.Float64()*2 - 1
				}
			} else {
				for in := 0; in < nn.hiddenNodesPerLayer[l-1]; in++ {
					nn.hiddenWeights[l][hn][in] = r.Float64()*2 - 1
				}
			}
			nn.hiddenBias[l][hn] = r.Float64()*2 - 1
		}
	}

	// INIT OUTPUT LAYER
	// o is the output node number
	for on := 0; on < nn.outputNodes; on++ {
		// n is the hidden node number
		for hn := 0; hn < nn.hiddenNodesPerLayer[nn.hiddenLayers-1]; hn++ {
			nn.outputWeights[on][hn] = r.Float64()*2 - 1
		}
		nn.outputBias[on] = r.Float64()*2 - 1
	}

}

// Print the neural network
func (nn *NeuralNetwork) PrintNeuralNetwork() {

	// NOTE: i is the node number
	fmt.Println("Input nodes:", nn.inputNodes)
	fmt.Println("Hidden layers:", nn.hiddenLayers)

	// Print weights and bias for each hidden layer
	for l := 0; l < nn.hiddenLayers; l++ {
		fmt.Println("HIDDEN LAYER", l, "NODES:", nn.hiddenNodesPerLayer[l])

		// Print weights for each nodes in this hidden layer
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			fmt.Printf("    weights node: %d %.3f\n", hn, nn.hiddenWeights[l][hn])
		}
		// Print bias for each node in this hidden layer
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			fmt.Printf("    bias node:    %d %.3f\n", hn, nn.hiddenBias[l][hn])
		}
	}

	// Print the output layer
	fmt.Println("OUTPUT NODES:", nn.outputNodes)
	// Print weights for each nodes in the output layer
	for on := 0; on < nn.outputNodes; on++ {
		fmt.Printf("    weights node: %d %.3f\n", on, nn.outputWeights[on])
	}
	// Print bias for each node in the output layer
	for on := 0; on < nn.outputNodes; on++ {
		fmt.Printf("    bias node:    %d %.3f\n", on, nn.outputBias[on])
	}

	// Print the learning rate
	fmt.Println("Learning rate    ", nn.learningRate)

	// Print the min and max input values
	fmt.Println("Min Input Value: ", nn.minInput)
	fmt.Println("Max Input Value: ", nn.maxInput)

}

// Read csv file and Get the min and max values from the dataset - for normalization function
func (nn *NeuralNetwork) GetInputMinMaxFromCSV() error {

	// Open the CSV file
	file, err := os.Open(nn.datasetCSVFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the header row
	_, err = reader.Read()
	if err != nil {
		return err
	}

	// Initialize min and max values with the first data row
	record, err := reader.Read()
	if err != nil {
		return err
	}

	for i := 0; i < nn.inputNodes; i++ {
		value, err := strconv.ParseFloat(record[i], 64)
		if err != nil {
			return err
		}
		nn.minInput[i] = value
		nn.maxInput[i] = value
	}

	// Iterate over the rest of the records to find min and max values
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		for i := 0; i < nn.inputNodes; i++ {
			value, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				return err
			}
			if value < nn.minInput[i] {
				nn.minInput[i] = value
			}
			if value > nn.maxInput[i] {
				nn.maxInput[i] = value
			}
		}
	}

	return nil
}

// Print the min and max input values for each input
func (nn *NeuralNetwork) PrintInputMinMax() {

	// Print the min and max input values
	fmt.Println("Min Input Value: ", nn.minInput)
	fmt.Println("Max Input Value: ", nn.maxInput)
}

// Train the neural network by reading the dataset from the CSV file
func (nn *NeuralNetwork) TrainNeuralNetwork() {

	// Setup the channel to read the CSV file line by line
	ch := nn.readCSVFileLineByLine()

	// Train the neural network for the number of epochs
	for epoch := 0; epoch < nn.epochs; epoch++ {

		// print the epoch number
		fmt.Println("Epoch", epoch)

		// Read the data rows one by one until EOF
		for {

			// STEP 6.1 Read the data from the channel
			// Receive the data from the channel
			data := <-ch // BLOCKING

			if data.i == nil {
				break
			}

			// Print the input data
			fmt.Println("    Input Data:     ", data.i)

			// STEP 6.2 Normalize the input data
			x := nn.normalizeInputData(data.i)

			// STEP 6.3 Forward Pass
			yHidden, yOutput := nn.forwardPass(x)

			// Print the outputs to .3 decimal places
			fmt.Printf("        yHidden:     %.3f\n", yHidden)
			fmt.Printf("        yOutput:     %.3f\n", yOutput)

			// STEP 6.4 Backward Pass
			deltaOutput, deltaHidden := nn.backwardPass(data.z, yOutput, yHidden)

			// Print the deltas to .3 decimal places
			fmt.Printf("        deltaOutput: %.3f\n", deltaOutput)
			fmt.Printf("        deltaHidden: %.3f\n", deltaHidden)
		}

	}

}

// ReadCSVFileLineByLine reads the CSV file line by line and
// returns a channel to read the TrainingData struct.
// Instead of calling this function each time you want to open the file,
// keep the file open and just loop from the start of the file.
func (nn *NeuralNetwork) readCSVFileLineByLine() chan trainingData {

	// Create a new TrainingData struct
	data := trainingData{
		i: make([]float64, nn.inputNodes),
		z: make([]float64, nn.outputNodes),
	}

	// Create a channel to read each line and return the TrainingData struct
	ch := make(chan trainingData)

	// Open the CSV file
	file, err := os.Open(nn.datasetCSVFile)
	if err != nil {
		return nil
	}

	// Read the data rows in a separate goroutine until EOF
	go func() {

		defer file.Close()
		defer close(ch)

		// We want to keep reading the file for the number of epochs
		for epoch := 0; epoch < nn.epochs; epoch++ {

			// Make sure we're at the beginning of the file
			_, err := file.Seek(0, 0)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Create a new CSV reader
			reader := csv.NewReader(file)

			// Read the header row
			_, err = reader.Read()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Read the data rows one by one until EOF
			// Send back data.i nil to alert EOF
			for {

				// Read the data row
				dataLine, err := reader.Read()

				// EOF
				if err != nil {
					data.i = nil
					ch <- data // BLOCKING
					data.i = make([]float64, nn.inputNodes)

					break

				} else {
					// Read the input data
					for i := 0; i < nn.inputNodes; i++ {
						value, err := strconv.ParseFloat(dataLine[i], 64)
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						data.i[i] = value

					}

					// Read the output data
					for i := 0; i < nn.outputNodes; i++ {
						value, err := strconv.ParseFloat(dataLine[nn.inputNodes+i], 64)
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						data.z[i] = value
					}

					// Send the data to the channel
					ch <- data                          // BLOCKING
					time.Sleep(1000 * time.Microsecond) // wait
				}
			}

		}

	}()

	return ch
}

// Normalize the input data
func (nn *NeuralNetwork) normalizeInputData(i []float64) []float64 {
	x := make([]float64, nn.inputNodes)
	for j := 0; j < nn.inputNodes; j++ {
		x[j] = (i[j] - nn.minInput[j]) / (nn.maxInput[j] - nn.minInput[j])
	}
	return x
}

// ForwardPass calculates the output of the neural network
func (nn *NeuralNetwork) forwardPass(x []float64) (yHidden [][]float64, yOutput []float64) {

	// Initialize the hidden outputs for each layer
	yHidden = make([][]float64, nn.hiddenLayers)
	for l := 0; l < nn.hiddenLayers; l++ {
		yHidden[l] = make([]float64, nn.hiddenNodesPerLayer[l])
	}

	// Initialize the output outputs
	yOutput = make([]float64, nn.outputNodes)

	// Calculate the output of each hidden node for each layer
	for l := 0; l < nn.hiddenLayers; l++ {
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			yHidden[l][hn] = 0.0
			s := 0.0
			// SUMMATION FUNCTION
			if l == 0 {
				for in := 0; in < nn.inputNodes; in++ {
					s += x[in] * nn.hiddenWeights[l][hn][in]
				}
				s += nn.hiddenBias[l][hn]
			} else {
				for in := 0; in < nn.hiddenNodesPerLayer[l-1]; in++ {
					s += yHidden[l-1][in] * nn.hiddenWeights[l][hn][in]
				}
				s += nn.hiddenBias[l][hn]
			}
			// ACTIVATION FUNCTION
			yHidden[l][hn] = sigmoid(s)
		}
	}

	// Calculate the output of each output node
	for o := 0; o < nn.outputNodes; o++ {
		yOutput[o] = 0.0
		s := 0.0
		// SUMMATION FUNCTION
		for hn := 0; hn < nn.hiddenNodesPerLayer[nn.hiddenLayers-1]; hn++ {
			s += yHidden[nn.hiddenLayers-1][hn] * nn.outputWeights[o][hn]
		}
		s += nn.outputBias[o]
		// ACTIVATION FUNCTION
		yOutput[o] = sigmoid(s)
	}

	return yHidden, yOutput
}

// Activation functions
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func sigmoidDerivative(x float64) float64 {
	return x * (1 - x)
}

// BackwardPass calculates the deltas for the neural network
func (nn *NeuralNetwork) backwardPass(z []float64, yOutput []float64, yHidden [][]float64) (deltaOutput []float64, deltaHidden [][]float64) {

	// Initialize the delta for the output layer
	deltaOutput = make([]float64, nn.outputNodes)

	// Initialize the delta for the hidden layers
	deltaHidden = make([][]float64, nn.hiddenLayers)
	for l := 0; l < nn.hiddenLayers; l++ {
		deltaHidden[l] = make([]float64, nn.hiddenNodesPerLayer[l])
	}

	// Calculate the delta for each output node
	for o := 0; o < nn.outputNodes; o++ {
		// ERROR FUNCTION
		deltaOutput[o] = (z[o] - yOutput[o]) * sigmoidDerivative(yOutput[o])
	}

	// Calculate the delta for each hidden node for each layer
	for l := nn.hiddenLayers - 1; l >= 0; l-- {
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			// ERROR FUNCTION
			// If this is the last hidden layer use the output weights
			if l == nn.hiddenLayers-1 {
				s := 0.0
				for o := 0; o < nn.outputNodes; o++ {
					s += deltaOutput[o] * nn.outputWeights[o][hn]
				}
				deltaHidden[l][hn] = s * sigmoidDerivative(yHidden[l][hn])
				// Use the hidden weights
			} else {
				s := 0.0
				for hn1 := 0; hn1 < nn.hiddenNodesPerLayer[l+1]; hn1++ {
					s += deltaHidden[l+1][hn1] * nn.hiddenWeights[l+1][hn1][hn]
				}
				deltaHidden[l][hn] = s * sigmoidDerivative(yHidden[l][hn])
			}
		}
	}

	return deltaOutput, deltaHidden
}
