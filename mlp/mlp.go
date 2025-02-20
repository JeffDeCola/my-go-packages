package mlp

import (
	"encoding/csv"
	"fmt"
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
			fmt.Println("    weights node:", hn, nn.hiddenWeights[l][hn])
		}
		// Print bias for each node in this hidden layer
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			fmt.Println("    bias node:   ", hn, nn.hiddenBias[l][hn])
		}
	}

	// Print the output layer
	fmt.Println("OUTPUT NODES:", nn.outputNodes)
	// Print weights for each nodes in the output layer
	for on := 0; on < nn.outputNodes; on++ {
		fmt.Println("    weights node:", on, nn.outputWeights[on])
	}
	// Print bias for each node in the output layer
	for on := 0; on < nn.outputNodes; on++ {
		fmt.Println("    bias node:   ", on, nn.outputBias[on])
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
	for i := 0; i < nn.inputNodes; i++ {
		fmt.Println("Input Node", i, "Min:", nn.minInput[i], "Max:", nn.maxInput[i])
	}
}

// Train the neural network by reading the dataset from the CSV file
func (nn *NeuralNetwork) TrainNeuralNetwork() {

	// Create a new TrainingData struct
	data := trainingData{
		i: make([]float64, nn.inputNodes),
		z: make([]float64, nn.outputNodes),
	}

	// Setup the channel to read the CSV file line by line
	ch, err := nn.readCSVFileLineByLine()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Train the neural network for the number of epochs
	for epoch := 0; epoch < nn.epochs; epoch++ {

		// print the epoch number
		fmt.Println("Epoch:", epoch)

		// Read the data rows one by one

		// Print the input and output data
		for data = range ch {
			fmt.Println("Input/output Data: ", data.i, data.z)
		}

	}

}

// ReadCSVFileLineByLine reads the CSV file line by line and
// returns a channel to read the TrainingData struct
func (nn *NeuralNetwork) readCSVFileLineByLine() (<-chan trainingData, error) {

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
		return nil, err
	}

	// Create a new CSV reader
	reader := csv.NewReader(file)

	go func() {

		defer file.Close()
		defer close(ch)

		// Read the header row
		_, err = reader.Read()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Read the data rows one by one
		for {

			dataLine, err := reader.Read()
			if err != nil {
				break
			}

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

			fmt.Println("-----------Input/output Data: ", data.i, data.z)

			ch <- data
		}

		// print EOF
		fmt.Println("EOF")

		// reset the file pointer
		file.Seek(0, 0)

	}()

	fmt.Println("returning channelkjhagsdfkljhasdfljhasdf")

	return ch, nil
}
