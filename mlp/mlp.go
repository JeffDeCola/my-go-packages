package mlp

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Neural Network Parameters
type NeuralNetworkParameters struct {
	InputNodes              int
	InputNodeLabels         []string
	HiddenLayers            int
	HiddenNodesPerLayer     []int
	OutputNodes             int
	OutputNodeLabels        []string
	LearningRate            float64
	Epochs                  int
	DatasetCSVFile          string
	Initialization          string // "random" or "file"
	WeightsAndBiasesCSVFile string
	NormalizeInputData      bool
	NormalizeMethod         string // "zero-to-one" or "minus-one-to-one
	ActivationFunction      string // "sigmoid" or "tanh"
	LossFunction            string // "mean-squared-error" or "cross-entropy"
}

// Neural network structure with up to 4 hidden layers
type NeuralNetwork struct {
	inputNodes              int           // INPUT NODES (USER ADDED)
	inputNodeLabels         []string      // INPUT NODE LABELS (USER ADDED)
	hiddenLayers            int           // HIDDEN LAYERS (USER ADDED)
	hiddenNodesPerLayer     []int         // HIDDEN NODES per [hiddenLayer] (USER ADDED)
	hiddenWeights           [][][]float64 // - weights per [hiddenLayer][hiddenNode][inputNode]
	hiddenBias              [][]float64   // - bias per [hiddenLayer][hiddenNode]
	outputNodes             int           // OUTPUT NODES (USER ADDED)
	outputNodeLabels        []string      // OUTPUT NODE LABELS (USER ADDED)
	outputWeights           [][]float64   // - weights per [outputNode][hiddenNode]
	outputBias              []float64     // - bias per [outputNode]
	epochs                  int           // EPOCHS (USER ADDED)
	datasetCSVFile          string        // DATASET CSV FILE (USER ADDED)
	minInput                []float64     // MIN INPUT VALUE
	maxInput                []float64     // MAX INPUT VALUE
	initialization          string        // INITIALIZATION (USER ADDED)
	weightsAndBiasesCSVFile string        // WEIGHTS AND BIAS CVS FILE (USER ADDED)
	normalizeInputData      bool          // NORMALIZE INPUT DATA (USER ADDED)
	normalizeMethod         string        // NORMALIZE METHOD (USER ADDED)
	activationFunction      string        // ACTIVATION FUNCTION (USER ADDED)
	lossFunction            string        // LOSS FUNCTION (USER ADDED)
	learningRate            float64       // LEARNING RATE (USER ADDED)
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
		inputNodes:              nnp.InputNodes,              // USER PROVIDED
		inputNodeLabels:         nnp.InputNodeLabels,         // USER PROVIDED
		hiddenLayers:            nnp.HiddenLayers,            // USER PROVIDED
		hiddenNodesPerLayer:     nnp.HiddenNodesPerLayer,     // USER PROVIDED
		hiddenWeights:           hiddenWeights,               // - created here
		hiddenBias:              hiddenBias,                  // - created here
		outputNodes:             nnp.OutputNodes,             // USER PROVIDED
		outputNodeLabels:        nnp.OutputNodeLabels,        // USER PROVIDED
		outputWeights:           outputWeights,               // - created here
		outputBias:              outputBias,                  // - created here
		epochs:                  nnp.Epochs,                  // USER PROVIDED
		datasetCSVFile:          nnp.DatasetCSVFile,          // USER PROVIDED
		minInput:                minInput,                    // - created here
		maxInput:                maxInput,                    // - created here
		initialization:          nnp.Initialization,          // USER PROVIDED
		weightsAndBiasesCSVFile: nnp.WeightsAndBiasesCSVFile, // USER PROVIDED
		normalizeInputData:      nnp.NormalizeInputData,      // USER PROVIDED
		normalizeMethod:         nnp.NormalizeMethod,         // USER PROVIDED
		activationFunction:      nnp.ActivationFunction,      // USER PROVIDED
		lossFunction:            nnp.LossFunction,            // USER PROVIDED
		learningRate:            nnp.LearningRate,            // USER PROVIDED
	}

	return nn
}

// Print the neural network
func (nn *NeuralNetwork) PrintNeuralNetwork() {

	// INPUT LAYER ******************************************
	fmt.Println("INPUT LAYER")
	fmt.Println("    Nodes:               ", nn.inputNodes)
	fmt.Println("    Labels:              ", nn.inputNodeLabels)

	// HIDDEN LAYER ******************************************
	fmt.Println("HIDDEN LAYER")
	fmt.Println("    Layers:              ", nn.hiddenLayers)
	// Print weights and bias for each hidden layer
	for l := 0; l < nn.hiddenLayers; l++ {
		fmt.Println("    Layer", l, "has", nn.hiddenNodesPerLayer[l], "Nodes")

		// Print weights for each nodes in this hidden layer
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			fmt.Printf("        weights node: %d   %.3f\n", hn, nn.hiddenWeights[l][hn])
		}
		// Print bias for each node in this hidden layer
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			fmt.Printf("        bias node:    %d   %.3f\n", hn, nn.hiddenBias[l][hn])
		}
	}

	// OUTPUT LAYER ******************************************
	fmt.Println("OUTPUT LAYER")
	// Print the output layer
	fmt.Println("    Nodes:               ", nn.outputNodes)
	fmt.Println("    Labels:              ", nn.outputNodeLabels)
	// Print weights for each nodes in the output layer
	for on := 0; on < nn.outputNodes; on++ {
		fmt.Printf("    weights node:     %d   %.3f\n", on, nn.outputWeights[on])
	}
	// Print bias for each node in the output layer
	for on := 0; on < nn.outputNodes; on++ {
		fmt.Printf("    bias node:        %d   %.3f\n", on, nn.outputBias[on])
	}

	// OTHER INFO
	fmt.Println("OTHER INFO")
	fmt.Println("    Epochs               ", nn.epochs)
	fmt.Println("    Dataset CSV          ", nn.datasetCSVFile)
	fmt.Println("    Min Input Value:     ", nn.minInput)
	fmt.Println("    Max Input Value:     ", nn.maxInput)
	fmt.Println("    Initialization       ", nn.initialization)
	fmt.Println("    W & B CSV:           ", nn.weightsAndBiasesCSVFile)
	fmt.Println("    Normalize Input Data ", nn.normalizeInputData)
	fmt.Println("    Normalize Method     ", nn.normalizeMethod)
	fmt.Println("    Activation Function  ", nn.activationFunction)
	fmt.Println("    Loss Function        ", nn.lossFunction)
	fmt.Println("    Learning rate        ", nn.learningRate)

}

// STEP 1 - INITIALIZATION
// Initialize the neural network (weights and biases) from values -1 to 1
func (nn *NeuralNetwork) InitializeNeuralNetwork() error {

	// RANDOM OR FROM FILE

	if nn.initialization == "random" {
		nn.initializeRandom()
	} else if nn.initialization == "file" {
		err := nn.initializeFromFile()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("invalid initialization method: %s", nn.initialization)
	}

	return nil
}

// STEP 1 - INITIALIZATION
// random
func (nn *NeuralNetwork) initializeRandom() {

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

// STEP 1 - INITIALIZATION
// from file
func (nn *NeuralNetwork) initializeFromFile() error {

	err := nn.loadWeightsAndBiases()
	if err != nil {
		return err
	}

	return nil

}

// Save weights and biases to a json file
func (nn *NeuralNetwork) saveWeightsAndBiases() error {

	// Create a struct to hold the weights and biases
	weightsAndBiases := struct {
		HiddenWeights [][][]float64 `json:"hiddenWeights"`
		HiddenBias    [][]float64   `json:"hiddenBias"`
		OutputWeights [][]float64   `json:"outputWeights"`
		OutputBias    []float64     `json:"outputBias"`
	}{
		HiddenWeights: nn.hiddenWeights,
		HiddenBias:    nn.hiddenBias,
		OutputWeights: nn.outputWeights,
		OutputBias:    nn.outputBias,
	}

	// Marshal the weights and biases to JSON
	jsonData, err := json.MarshalIndent(weightsAndBiases, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling weights and biases: %v", err)
	}

	filename := nn.weightsAndBiasesCSVFile
	// Write the JSON data to a file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing weights and biases to file: %v", err)
	}

	return nil
}

// Load weights and biases from a json file
func (nn *NeuralNetwork) loadWeightsAndBiases() error {

	filename := nn.weightsAndBiasesCSVFile

	// Read the JSON data from the file
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading weights and biases from file: %v", err)
	}

	// Create a struct to hold the weights and biases
	var weightsAndBiases struct {
		HiddenWeights [][][]float64 `json:"hiddenWeights"`
		HiddenBias    [][]float64   `json:"hiddenBias"`
		OutputWeights [][]float64   `json:"outputWeights"`
		OutputBias    []float64     `json:"outputBias"`
	}

	// Unmarshal the JSON data
	err = json.Unmarshal(jsonData, &weightsAndBiases)
	if err != nil {
		return fmt.Errorf("error unmarshaling weights and biases: %v", err)
	}

	// CHECK NUMBER HIDDEN LAYERS MATCH
	if len(weightsAndBiases.HiddenWeights) != nn.hiddenLayers {
		return fmt.Errorf("mismatch in number of hidden layers: expected %d, got %d",
			nn.hiddenLayers, len(weightsAndBiases.HiddenWeights))
	}

	// CHECK NUMBER NODES PER HIDDEN LAYER MATCH
	for l := 0; l < nn.hiddenLayers; l++ {
		if len(weightsAndBiases.HiddenWeights[l]) != nn.hiddenNodesPerLayer[l] {
			return fmt.Errorf("mismatch in number of hidden nodes in layer %d: expected %d, got %d",
				l, nn.hiddenNodesPerLayer[l], len(weightsAndBiases.HiddenWeights[l]))
		}
	}

	// CHECK OUTPUT NODES MATCH
	if len(weightsAndBiases.OutputWeights) != nn.outputNodes {
		return fmt.Errorf("mismatch in number of output nodes: expected %d, got %d",
			nn.outputNodes, len(weightsAndBiases.OutputWeights))
	}

	// Load the weights and biases
	nn.hiddenWeights = weightsAndBiases.HiddenWeights
	nn.hiddenBias = weightsAndBiases.HiddenBias
	nn.outputWeights = weightsAndBiases.OutputWeights
	nn.outputBias = weightsAndBiases.OutputBias

	return nil
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

	// Read the first record (Could contain whitespace)
	record, err := reader.Read()
	if err != nil {
		return err
	}

	// READ FIRST DATASET
	for i := 0; i < nn.inputNodes; i++ {
		// Trim whitespace from the CSV field before parsing
		trimmedValue := strings.TrimSpace(record[i])
		value, err := strconv.ParseFloat(trimmedValue, 64)
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
			// Trim whitespace from the CSV field before parsing
			trimmedValue := strings.TrimSpace(record[i])
			value, err := strconv.ParseFloat(trimmedValue, 64)
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

// TRAINING LOOP
// Train the neural network by reading the dataset from the CSV file
func (nn *NeuralNetwork) TrainNeuralNetwork() error {

	// TRAINING LOOP - EPOCH LOOP
	// Train the neural network by reading the dataset from the CSV file
	err := nn.epochLoop()
	if err != nil {
		return err
	}

	return nil

}

// TRAINING LOOP - EPOCH LOOP
// Train the neural network by reading the dataset from the CSV file
func (nn *NeuralNetwork) epochLoop() error {

	// Train the neural network for the number of epochs
	for epoch := 0; epoch < nn.epochs; epoch++ {

		// print the epoch number
		fmt.Println("Epoch", epoch)

		err := nn.datasetLoop()
		if err != nil {
			return err
		}

	}

	return nil

}

// TRAINING LOOP - DATASET LOOP
// Train the neural network by reading the dataset from the CSV file
func (nn *NeuralNetwork) datasetLoop() error {

	var x []float64

	// Setup the channel to read the CSV file line by line
	ch := nn.readCSVFileLineByLine()

	// Read the data rows one by one until EOF
	for {

		// Receive the data from the channel
		data := <-ch // BLOCKING

		if data.i == nil {
			break
		}

		// Print the input data
		fmt.Println("    Input Data:     ", data.i)

		// STEP 2 -  Normalize the input data
		if nn.normalizeInputData {
			if nn.normalizeMethod == "zero-to-one" {
				x = nn.normalizeInputDataZeroToOne(data.i)
			} else if nn.normalizeMethod == "minus-one-to-one" {
				x = nn.normalizeInputDataMinusOneToOne(data.i)
			} else {
				return fmt.Errorf("invalid normalization method: %s", nn.normalizeMethod)
			}
		} else {
			x = data.i
		}

		fmt.Printf("    Normalized Data: %.3f\n", x)

		// STEP 3 - Forward Pass
		// Compute the output of the neural network for the current input data
		aHidden, yOutput := nn.forwardPass(x)

		// Print the outputs to .3 decimal places
		fmt.Printf("        aHidden:     %.3f\n", aHidden)
		fmt.Printf("        yOutput:     %.3f\n", yOutput)

		// STEP 6.4 Backward Pass
		//deltaOutput, deltaHidden := nn.backwardPass(data.z, yOutput, aHidden)

		// Print the deltas to .3 decimal places
		//fmt.Printf("        deltaOutput: %.3f\n", deltaOutput)
		//fmt.Printf("        deltaHidden: %.3f\n", deltaHidden)

	}

	return nil
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
						// Trim whitespace before parsing
						trimmedValue := strings.TrimSpace(dataLine[i])
						value, err := strconv.ParseFloat(trimmedValue, 64)
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						data.i[i] = value

					}

					// Read the output data
					for i := 0; i < nn.outputNodes; i++ {
						// Trim whitespace before parsing
						trimmedValue := strings.TrimSpace(dataLine[nn.inputNodes+i])
						value, err := strconv.ParseFloat(trimmedValue, 64)
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

// STEP 2 - NORMALIZATION
// Normalize the input data between 0 and 1
func (nn *NeuralNetwork) normalizeInputDataZeroToOne(i []float64) []float64 {
	x := make([]float64, nn.inputNodes)
	for j := 0; j < nn.inputNodes; j++ {
		if nn.minInput[j] == nn.maxInput[j] {
			// If min and max are the same, set normalized value to 0.5
			x[j] = 0.5
		} else {
			x[j] = (i[j] - nn.minInput[j]) / (nn.maxInput[j] - nn.minInput[j])
		}
	}
	return x
}

// STEP 2 - NORMALIZATION
// Normalize the input data between -1 and 1
func (nn *NeuralNetwork) normalizeInputDataMinusOneToOne(i []float64) []float64 {
	x := make([]float64, nn.inputNodes)
	for j := 0; j < nn.inputNodes; j++ {
		if nn.minInput[j] == nn.maxInput[j] {
			// If min and max are the same, set normalized value to 0
			x[j] = 0
		} else {
			x[j] = 2*((i[j]-nn.minInput[j])/(nn.maxInput[j]-nn.minInput[j])) - 1
		}
	}
	return x
}

// STEP 3 - FORWARD PASS
// ForwardPass calculates the output of the neural network
func (nn *NeuralNetwork) forwardPass(x []float64) (aHidden [][]float64, yOutput []float64) {

	// Initialize the hidden outputs for each layer
	aHidden = make([][]float64, nn.hiddenLayers)
	for l := 0; l < nn.hiddenLayers; l++ {
		aHidden[l] = make([]float64, nn.hiddenNodesPerLayer[l])
	}

	// Initialize the output outputs
	yOutput = make([]float64, nn.outputNodes)

	// Calculate the output of each hidden node for each layer
	for l := 0; l < nn.hiddenLayers; l++ {
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			aHidden[l][hn] = 0.0
			s := 0.0

			// SUMMATION FUNCTION
			if l == 0 {
				for in := 0; in < nn.inputNodes; in++ {
					s += x[in] * nn.hiddenWeights[l][hn][in]
				}
				s += nn.hiddenBias[l][hn]
			} else {
				for in := 0; in < nn.hiddenNodesPerLayer[l-1]; in++ {
					s += aHidden[l-1][in] * nn.hiddenWeights[l][hn][in]
				}
				s += nn.hiddenBias[l][hn]
			}

			// ACTIVATION FUNCTION
			if nn.activationFunction == "sigmoid" {
				aHidden[l][hn] = sigmoid(s)
			} else if nn.activationFunction == "tanh" {
				aHidden[l][hn] = tanh(s)
			} else {
				return nil, nil
			}

		}
	}

	// Calculate the output of each output node
	for o := 0; o < nn.outputNodes; o++ {
		yOutput[o] = 0.0
		s := 0.0

		// SUMMATION FUNCTION
		for hn := 0; hn < nn.hiddenNodesPerLayer[nn.hiddenLayers-1]; hn++ {
			s += aHidden[nn.hiddenLayers-1][hn] * nn.outputWeights[o][hn]
		}
		s += nn.outputBias[o]

		// ACTIVATION FUNCTION
		if nn.activationFunction == "sigmoid" {
			yOutput[o] = sigmoid(s)
		} else if nn.activationFunction == "tanh" {
			yOutput[o] = tanh(s)
		} else {
			return nil, nil
		}

	}

	return aHidden, yOutput
}

// Activation function - Sigmoid
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// Derivative of the sigmoid function
func sigmoidDerivative(x float64) float64 {
	return x * (1 - x)
}

// Activation function - Tanh
func tanh(x float64) float64 {
	return math.Tanh(x)
}

// Derivative of the tanh function
func tanhDerivative(x float64) float64 {
	return 1 - x*x
}

// STEP 4 - BACKWARD PASS
// BackwardPass calculates the deltas for the neural network
func (nn *NeuralNetwork) backwardPass(z []float64, yOutput []float64, aHidden [][]float64) (deltaOutput []float64, deltaHidden [][]float64) {

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
				deltaHidden[l][hn] = s * sigmoidDerivative(aHidden[l][hn])
				// Use the hidden weights
			} else {
				s := 0.0
				for hn1 := 0; hn1 < nn.hiddenNodesPerLayer[l+1]; hn1++ {
					s += deltaHidden[l+1][hn1] * nn.hiddenWeights[l+1][hn1][hn]
				}
				deltaHidden[l][hn] = s * sigmoidDerivative(aHidden[l][hn])
			}
		}
	}

	return deltaOutput, deltaHidden
}
