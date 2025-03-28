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

	logger "my-go-packages/golang/logger"
)

// Default logging level
var logmlp = logger.CreateLogger(logger.Warning, "jeffs_noTime")

// Neural Network Configuration Parameters
type NeuralNetworkConfiguration struct {
	Mode                         string // "training", "testing" or "predicting"
	InputNodes                   int
	InputNodeLabels              []string
	HiddenLayers                 int // Also update HiddenNodesPerLayer
	HiddenNodesPerLayer          []int
	OutputNodes                  int
	OutputNodeLabels             []string
	Epochs                       int
	LearningRate                 float64
	ActivationFunction           string // "sigmoid" or "tanh"
	LossFunction                 string // "mean-squared-error"
	InitWeightsBiasesMethod      string // "file" or "random"
	InitWeightsBiasesJSONFile    string
	MinMaxInputMethod            string // "file" or "calculate" from dataset File
	MinMaxOutputMethod           string // "file" or "calculate" from dataset File
	MinMaxJSONFile               string
	TrainingDatasetCSVFile       string
	NormalizeInputData           bool
	NormalizeOutputData          bool
	NormalizeMethod              string // "zero-to-one" or "minus-one-to-one
	TrainedWeightsBiasesJSONFile string
	TestingDatasetCSVFile        string
	PredictingDatasetCSVFile     string
}

// Neural network architecture and parameters
type neuralNetwork struct {
	mode                         string        // OPERATION MODE - "training", "testing" or "predicting" (USER ADDED)
	inputNodes                   int           // INPUT NODES (USER ADDED)
	inputNodeLabels              []string      // INPUT NODE LABELS (USER ADDED)
	hiddenLayers                 int           // HIDDEN LAYERS (USER ADDED)
	hiddenNodesPerLayer          []int         // HIDDEN NODES per [hiddenLayer] (USER ADDED)
	hiddenWeights                [][][]float64 // - weights per [hiddenLayer][hiddenNode][inputNode]
	hiddenBias                   [][]float64   // - bias per [hiddenLayer][hiddenNode]
	outputNodes                  int           // OUTPUT NODES (USER ADDED)
	outputNodeLabels             []string      // OUTPUT NODE LABELS (USER ADDED)
	outputWeights                [][]float64   // - weights per [outputNode][hiddenNode]
	outputBias                   []float64     // - bias per [outputNode]
	epochs                       int           // EPOCHS (USER ADDED)
	learningRate                 float64       // LEARNING RATE (USER ADDED)
	activationFunction           string        // ACTIVATION FUNCTION (USER ADDED)
	lossFunction                 string        // LOSS FUNCTION (USER ADDED)
	initWeightsBiasesMethod      string        // INITIALIZATION (USER ADDED)
	initWeightsBiasesJSONFile    string        // INIT WEIGHTS AND BIAS CVS FILE (USER ADDED)
	MinMaxInputMethod            string        // INPUT MIN MAX - "file" or "calculate" (USER ADDED)
	MinMaxOutputMethod           string        // OUTPUT MIN MAX - "file" or "calculate" (USER ADDED)
	MinMaxJSONFile               string        // MIN MAX JSON FILE (USER ADDED)
	minInput                     []float64     // - min input value
	maxInput                     []float64     // - max input value
	minOutput                    []float64     // - min output value
	maxOutput                    []float64     // - max output value
	trainingDatasetCSVFile       string        // TRAINING DATASET CSV FILE (USER ADDED)
	normalizeInputData           bool          // NORMALIZE INPUT DATA (USER ADDED)
	normalizeOutputData          bool          // NORMALIZE OUTPUT DATA (USER ADDED)
	normalizeMethod              string        // NORMALIZE METHOD (USER ADDED)
	statStartLoss                []float64     // - start Loss
	statEndLoss                  []float64     // - end Loss
	trainedWeightsBiasesJSONFile string        // TRAINED WEIGHTS AND BIAS CVS FILE (USER ADDED)
	testingDatasetCSVFile        string        // TESTING DATASET CSV FILE (USER ADDED)
	predictingDatasetCSVFile     string        // PREDICTING DATASET CSV FILE (USER ADDED)
}

// Data structure for training data IO
type trainingData struct {
	i []float64
	z []float64
}

// ShowDebug - Set logging level for this package
func ShowDebug() {
	logmlp.ChangeLogLevel(logger.Debug)
}

// ShowInfo - Set logging level for this package
func ShowInfo() {
	logmlp.ChangeLogLevel(logger.Info)
}

// Create the MLP neural network
func (nnp NeuralNetworkConfiguration) CreateNeuralNetwork() *neuralNetwork {

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
	minOutput := make([]float64, nnp.OutputNodes)
	maxOutput := make([]float64, nnp.OutputNodes)

	// Initialize the start and end loss
	statStartLoss := make([]float64, nnp.OutputNodes)
	statEndLoss := make([]float64, nnp.OutputNodes)

	// Create the neural network
	nn := &neuralNetwork{
		mode:                         nnp.Mode,                         // USER PROVIDED
		inputNodes:                   nnp.InputNodes,                   // USER PROVIDED
		inputNodeLabels:              nnp.InputNodeLabels,              // USER PROVIDED
		hiddenLayers:                 nnp.HiddenLayers,                 // USER PROVIDED
		hiddenNodesPerLayer:          nnp.HiddenNodesPerLayer,          // USER PROVIDED
		hiddenWeights:                hiddenWeights,                    // -created here
		hiddenBias:                   hiddenBias,                       // -created here
		outputNodes:                  nnp.OutputNodes,                  // USER PROVIDED
		outputNodeLabels:             nnp.OutputNodeLabels,             // USER PROVIDED
		outputWeights:                outputWeights,                    // -created here
		outputBias:                   outputBias,                       // -created here
		epochs:                       nnp.Epochs,                       // USER PROVIDED
		learningRate:                 nnp.LearningRate,                 // USER PROVIDED
		activationFunction:           nnp.ActivationFunction,           // USER PROVIDED
		lossFunction:                 nnp.LossFunction,                 // USER PROVIDED
		initWeightsBiasesMethod:      nnp.InitWeightsBiasesMethod,      // USER PROVIDED
		initWeightsBiasesJSONFile:    nnp.InitWeightsBiasesJSONFile,    // USER PROVIDED
		MinMaxInputMethod:            nnp.MinMaxInputMethod,            // USER PROVIDED
		MinMaxOutputMethod:           nnp.MinMaxOutputMethod,           // USER PROVIDED
		MinMaxJSONFile:               nnp.MinMaxJSONFile,               // USER PROVIDED
		minInput:                     minInput,                         // -created here
		maxInput:                     maxInput,                         // -created here
		minOutput:                    minOutput,                        // -created here
		maxOutput:                    maxOutput,                        // -created here
		trainingDatasetCSVFile:       nnp.TrainingDatasetCSVFile,       // USER PROVIDED
		normalizeInputData:           nnp.NormalizeInputData,           // USER PROVIDED
		normalizeOutputData:          nnp.NormalizeOutputData,          // USER PROVIDED
		normalizeMethod:              nnp.NormalizeMethod,              // USER PROVIDED
		statStartLoss:                statStartLoss,                    // -created here
		statEndLoss:                  statEndLoss,                      // -created here
		trainedWeightsBiasesJSONFile: nnp.TrainedWeightsBiasesJSONFile, // USER PROVIDED
		testingDatasetCSVFile:        nnp.TestingDatasetCSVFile,        // USER PROVIDED
		predictingDatasetCSVFile:     nnp.PredictingDatasetCSVFile,     // USER PROVIDED
	}

	return nn
}

// Print the neural network
func (nn *neuralNetwork) PrintNeuralNetwork() {

	// OPERATION MODE
	fmt.Println("OPERATION MODE")
	fmt.Println("    Mode:                    ", nn.mode)

	// INPUT LAYER ******************************************
	fmt.Println("INPUT LAYER")
	fmt.Println("    Nodes:                   ", nn.inputNodes)
	fmt.Println("    Labels:                  ", nn.inputNodeLabels)

	// HIDDEN LAYER ******************************************
	fmt.Println("HIDDEN LAYER")
	fmt.Println("    Layers:                  ", nn.hiddenLayers)
	// Print weights and bias for each hidden layer
	for l := 0; l < nn.hiddenLayers; l++ {
		fmt.Println("    Layer", l, "has", nn.hiddenNodesPerLayer[l], "Nodes")

		// Print weights for each nodes in this hidden layer
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			fmt.Printf("        weights node: %d       %.3f\n", hn, nn.hiddenWeights[l][hn])
		}
		// Print bias for each node in this hidden layer
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			fmt.Printf("        bias node:    %d       %.3f\n", hn, nn.hiddenBias[l][hn])
		}
	}

	// OUTPUT LAYER ******************************************
	fmt.Println("OUTPUT LAYER")
	// Print the output layer
	fmt.Println("    Nodes:                   ", nn.outputNodes)
	fmt.Println("    Labels:                  ", nn.outputNodeLabels)
	// Print weights for each nodes in the output layer
	for on := 0; on < nn.outputNodes; on++ {
		fmt.Printf("    weights node:     %d       %.3f\n", on, nn.outputWeights[on])
	}
	// Print bias for each node in the output layer
	for on := 0; on < nn.outputNodes; on++ {
		fmt.Printf("    bias node:        %d       %.3f\n", on, nn.outputBias[on])
	}

	// OTHER INFO
	fmt.Println("OTHER INFO")
	fmt.Println("    Epochs                   ", nn.epochs)
	fmt.Println("    Learning Rate            ", nn.learningRate)
	fmt.Println("    Activation Function      ", nn.activationFunction)
	fmt.Println("    Loss Function            ", nn.lossFunction)
	fmt.Println("    Init Weights Biases      ", nn.initWeightsBiasesMethod)
	fmt.Println("    Init Weights Biases JSON ", nn.initWeightsBiasesJSONFile)
	fmt.Println("    Min Max Input Method     ", nn.MinMaxInputMethod)
	fmt.Println("    Min Max Output Method    ", nn.MinMaxOutputMethod)
	fmt.Println("    Min Max JSON File        ", nn.MinMaxJSONFile)
	fmt.Println("    Min Input Value          ", nn.minInput)
	fmt.Println("    Max Input Value          ", nn.maxInput)
	fmt.Println("    Min Output Value         ", nn.minOutput)
	fmt.Println("    Max Output Value         ", nn.maxOutput)
	fmt.Println("    Training Dataset CSV     ", nn.trainingDatasetCSVFile)
	fmt.Println("    Normalize Input Data     ", nn.normalizeInputData)
	fmt.Println("    Normalize Output Data    ", nn.normalizeOutputData)
	fmt.Println("    Normalize Method         ", nn.normalizeMethod)
	fmt.Println("    Start Loss               ", nn.statStartLoss)
	fmt.Println("    End Loss                 ", nn.statEndLoss)
	fmt.Println("    Trained W & B JSON       ", nn.trainedWeightsBiasesJSONFile)
	fmt.Println("    Testing Dataset CSV      ", nn.testingDatasetCSVFile)
	fmt.Println("    Predicting Dataset CSV   ", nn.predictingDatasetCSVFile)

}

// STEP 1 - INITIALIZATION
// Initialize the neural network (weights and biases)
// Random or from file
func (nn *neuralNetwork) InitializeNeuralNetwork() error {

	logmlp.Info("STEP 1 - INITIALIZATION ------------------------------------")

	// RANDOM OR FROM FILE

	if nn.initWeightsBiasesMethod == "random" {
		nn.initializeRandom()
	} else if nn.initWeightsBiasesMethod == "file" {
		err := nn.initializeFromFile()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("invalid initialization method: %s", nn.initWeightsBiasesMethod)
	}

	return nil
}

// STEP 1 - INITIALIZATION
// random
func (nn *neuralNetwork) initializeRandom() {

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
func (nn *neuralNetwork) initializeFromFile() error {

	err := nn.loadWeightsBiasesFromJSON(nn.initWeightsBiasesJSONFile)
	if err != nil {
		return err
	}

	return nil

}

// STEP 1 - INITIALIZATION
// Load weights and biases from a json file
func (nn *neuralNetwork) loadWeightsBiasesFromJSON(filename string) error {

	// Read the JSON data from the file
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading weights and biases from file: %v", err)
	}

	// Create a struct to hold the weights and biases
	var weightsBiasesMinMax struct {
		HiddenWeights [][][]float64 `json:"hiddenWeights"`
		HiddenBias    [][]float64   `json:"hiddenBias"`
		OutputWeights [][]float64   `json:"outputWeights"`
		OutputBias    []float64     `json:"outputBias"`
	}

	// Unmarshal the JSON data
	err = json.Unmarshal(jsonData, &weightsBiasesMinMax)
	if err != nil {
		return fmt.Errorf("error unmarshaling weights and biases: %v", err)
	}

	// CHECK NUMBER HIDDEN LAYERS MATCH
	if len(weightsBiasesMinMax.HiddenWeights) != nn.hiddenLayers {
		return fmt.Errorf("mismatch in number of hidden layers: expected %d, got %d",
			nn.hiddenLayers, len(weightsBiasesMinMax.HiddenWeights))
	}

	// CHECK NUMBER NODES PER HIDDEN LAYER MATCH
	for l := 0; l < nn.hiddenLayers; l++ {
		if len(weightsBiasesMinMax.HiddenWeights[l]) != nn.hiddenNodesPerLayer[l] {
			return fmt.Errorf("mismatch in number of hidden nodes in layer %d: expected %d, got %d",
				l, nn.hiddenNodesPerLayer[l], len(weightsBiasesMinMax.HiddenWeights[l]))
		}
	}

	// CHECK OUTPUT NODES MATCH
	if len(weightsBiasesMinMax.OutputWeights) != nn.outputNodes {
		return fmt.Errorf("mismatch in number of output nodes: expected %d, got %d",
			nn.outputNodes, len(weightsBiasesMinMax.OutputWeights))
	}

	// Load the weights and biases
	nn.hiddenWeights = weightsBiasesMinMax.HiddenWeights
	nn.hiddenBias = weightsBiasesMinMax.HiddenBias
	nn.outputWeights = weightsBiasesMinMax.OutputWeights
	nn.outputBias = weightsBiasesMinMax.OutputBias

	return nil
}

// STEP 8 - SAVE WEIGHTS & BIASES TO A FILE
// Save weights and biases to a json file
func (nn *neuralNetwork) SaveWeightsBiasesToJSON() error {

	logmlp.Info("STEP 8 - SAVE WEIGHTS & BIASES TO A FILE -------------------")

	filename := nn.trainedWeightsBiasesJSONFile

	// Create a struct to hold the weights and biases amd min/max inputs and outputs
	stuff := struct {
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

	// Marshal the weights and biases min/max to JSON
	jsonData, err := json.MarshalIndent(stuff, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling weights and biases: %v", err)
	}

	// Write the JSON data to a file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing weights and biases to file: %v", err)
	}

	return nil
}

// STEP 2 - MIN & MAX INPUT VALUES
// Set the min and max values for the input and output nodes
func (nn *neuralNetwork) SetMinMaxValues() error {

	logmlp.Info("STEP 2 - MIN & MAX INPUT VALUES ----------------------------")

	if nn.MinMaxInputMethod == "file" {
		// Load min and max values from a json file
		err := nn.loadMinMaxValuesFromJSON("input")
		if err != nil {
			return err
		}
	} else if nn.MinMaxInputMethod == "calculate" {
		// Use the min and max values from the dataset
		err := nn.calculateMinMaxValuesFromCSV("input")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("invalid min max input method: %s", nn.MinMaxInputMethod)
	}

	if nn.MinMaxOutputMethod == "file" {
		// Load min and max values from a json file
		err := nn.loadMinMaxValuesFromJSON("output")
		if err != nil {
			return err
		}
	} else if nn.MinMaxOutputMethod == "calculate" {
		// Use the min and max values from the dataset
		err := nn.calculateMinMaxValuesFromCSV("output")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("invalid min max output method: %s", nn.MinMaxOutputMethod)
	}

	return nil
}

// STEP -  MIN & MAX INPUT VALUES
// Load min and max values from a json file
func (nn *neuralNetwork) loadMinMaxValuesFromJSON(inout string) error {

	// Read the JSON data from the file
	jsonData, err := os.ReadFile(nn.MinMaxJSONFile)
	if err != nil {
		return fmt.Errorf("error reading min max values from file: %v", err)
	}

	// Create a struct to hold the min and max values
	var minMax struct {
		MinInput  []float64 `json:"minInput"`
		MaxInput  []float64 `json:"maxInput"`
		MinOutput []float64 `json:"minOutput"`
		MaxOutput []float64 `json:"maxOutput"`
	}

	// Unmarshal the JSON data
	err = json.Unmarshal(jsonData, &minMax)
	if err != nil {
		return fmt.Errorf("error unmarshaling min max values: %v", err)
	}

	// CHECK MIN INPUT VALUES MATCH
	if len(minMax.MinInput) != nn.inputNodes {
		return fmt.Errorf("mismatch in number of input nodes: expected %d, got %d",
			nn.inputNodes, len(minMax.MinInput))
	}

	// CHECK MAX INPUT VALUES MATCH
	if len(minMax.MaxInput) != nn.inputNodes {
		return fmt.Errorf("mismatch in number of input nodes: expected %d, got %d",
			nn.inputNodes, len(minMax.MaxInput))
	}

	// CHECK MIN OUTPUT VALUES MATCH
	if len(minMax.MinOutput) != nn.outputNodes {
		return fmt.Errorf("mismatch in number of output nodes: expected %d, got %d",
			nn.outputNodes, len(minMax.MinOutput))
	}

	// CHECK MAX OUTPUT VALUES MATCH
	if len(minMax.MaxOutput) != nn.outputNodes {
		return fmt.Errorf("mismatch in number of output nodes: expected %d, got %d",
			nn.outputNodes, len(minMax.MaxOutput))
	}

	// Load the min and max values
	if inout == "input" {
		copy(nn.minInput, minMax.MinInput)
		copy(nn.maxInput, minMax.MaxInput)
	}

	if inout == "output" {
		copy(nn.minOutput, minMax.MinOutput)
		copy(nn.maxOutput, minMax.MaxOutput)
	}

	return nil

}

// STEP 7 - SAVE MIN & MAX VALUES TO A FILE
func (nn *neuralNetwork) SaveMinMaxValuesToJSON() error {

	logmlp.Info("STEP 7 - SAVE MIN AND MAX VALUES TO A FILE -------------------")

	filename := nn.MinMaxJSONFile

	// Create a struct to hold the min and max values
	stuff := struct {
		MinInput  []float64 `json:"minInput"`
		MaxInput  []float64 `json:"maxInput"`
		MinOutput []float64 `json:"minOutput"`
		MaxOutput []float64 `json:"maxOutput"`
	}{
		MinInput:  nn.minInput,
		MaxInput:  nn.maxInput,
		MinOutput: nn.minOutput,
		MaxOutput: nn.maxOutput,
	}

	// Marshal the min and max values to JSON
	jsonData, err := json.MarshalIndent(stuff, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling min and max values: %v", err)
	}

	// Write the JSON data to a file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing min and max values to file: %v", err)
	}

	return nil

}

// STEP 2 -  MIN & MAX INPUT VALUES
// Read csv file and Get the min and max values from the dataset - for normalization function
func (nn *neuralNetwork) calculateMinMaxValuesFromCSV(inOut string) error {

	var file *os.File
	var err error

	// Open the CSV file
	if nn.mode == "training" {
		file, err = os.Open(nn.trainingDatasetCSVFile)
	} else if nn.mode == "testing" {
		file, err = os.Open(nn.testingDatasetCSVFile)
	} else if nn.mode == "predicting" {
		file, err = os.Open(nn.predictingDatasetCSVFile)
	} else {
		return fmt.Errorf("invalid mode: %s", nn.mode)
	}
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

	// READ FIRST DATASET - GET INPUT MIN MAX
	if inOut == "input" {
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
	}

	if inOut == "output" {
		// READ FIRST DATASET - GET OUTPUT MIN MAX
		for i := 0; i < nn.outputNodes; i++ {
			// Trim whitespace from the CSV field before parsing
			trimmedValue := strings.TrimSpace(record[nn.inputNodes+i])
			value, err := strconv.ParseFloat(trimmedValue, 64)
			if err != nil {
				return err
			}
			nn.minOutput[i] = value
			nn.maxOutput[i] = value
		}
	}

	// Iterate over the rest of the records to find min and max values
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if inOut == "input" {
			// GET MIN MAX INPUT
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

		if inOut == "output" {
			// GET MIN MAX OUTPUT
			for i := 0; i < nn.outputNodes; i++ {
				// Trim whitespace from the CSV field before parsing
				trimmedValue := strings.TrimSpace(record[nn.inputNodes+i])
				value, err := strconv.ParseFloat(trimmedValue, 64)
				if err != nil {
					return err
				}
				if value < nn.minOutput[i] {
					nn.minOutput[i] = value
				}
				if value > nn.maxOutput[i] {
					nn.maxOutput[i] = value
				}
			}
		}

	}

	return nil
}

// Print the min and max input values for each input
func (nn *neuralNetwork) PrintMinMaxValues() {

	// Print the min and max input values
	fmt.Println("Min Input Value:  ", nn.minInput)
	fmt.Println("Max Input Value:  ", nn.maxInput)
	fmt.Println("Min Output Value: ", nn.minOutput)
	fmt.Println("Max Output Value: ", nn.maxOutput)
}

// TRAINING LOOP
// Train the neural network by reading the dataset from the CSV file
func (nn *neuralNetwork) TrainNeuralNetwork() error {

	logmlp.Info("THE TRAINING LOOP ------------------------------------------")

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
func (nn *neuralNetwork) epochLoop() error {

	startToken := true

	// Train the neural network for the number of epochs
	for epoch := 0; epoch < nn.epochs; epoch++ {

		// The Epoch Number
		logmlp.Info(fmt.Sprintf("EPOCH NUMBER %d", epoch))

		err := nn.datasetLoop(startToken)
		if err != nil {
			return err
		}

		startToken = false

	}

	return nil

}

// TRAINING LOOP - DATASET LOOP
// STEP 3 - NORMALIZATION
// STEP 4 - FORWARD PASS
// STEP 5 - BACKWARD PASS
// STEP 6 - UPDATE WEIGHTS AND BIASES
// Train the neural network by reading the dataset from the CSV file
func (nn *neuralNetwork) datasetLoop(startToken bool) (err error) {

	loss := make([]float64, nn.outputNodes)

	// Setup the channel to read the CSV file line by line
	ch := nn.readCSVFileLineByLine()

	// Read the data rows one by one until EOF
	for {

		// Receive the data from the channel
		data := <-ch // BLOCKING

		if data.i == nil {
			break
		}

		// STEP 3 - NORMALIZATION
		x, z, err := nn.normalization(data)
		if err != nil {
			return err
		}

		// STEP 4 - FORWARD PASS
		// Compute the output of the neural network for the current input data
		aHidden, yOutput := nn.forwardPass(x)

		// STEP 5 BACKWARD PASS
		hWNew, hBNew, oWNew, oBNew := nn.backwardPass(x, z, yOutput, aHidden)

		// STEP 6 UPDATE WEIGHTS AND BIASES
		// Copy slice into another slice
		logmlp.Info("    STEP 6 - UPDATE WEIGHTS AND BIASES ---------------------")
		copy(nn.hiddenWeights, hWNew)
		copy(nn.hiddenBias, hBNew)
		copy(nn.outputWeights, oWNew)
		copy(nn.outputBias, oBNew)

		// Calculate the loss for every output node
		// yOutput can be from normalized input data
		// z can be from normalized output data
		logmlp.Debug("        Loss")
		for i := 0; i < nn.outputNodes; i++ {
			if nn.lossFunction == "mean-squared-error" {
				loss[i] = 1.0 / 2.0 * math.Pow(yOutput[i]-z[i], 2)

			} else {
				return fmt.Errorf("invalid loss function: %s", nn.lossFunction)
			}

			logmlp.Debug(fmt.Sprintf("            Node %d:  %.5f", i, loss[i]))
		}

		// Save the startLoss if startToken is true
		if startToken {
			copy(nn.statStartLoss, loss)
			startToken = false
		}

	}

	// Copy end loss for each output node
	copy(nn.statEndLoss, loss)

	return nil
}

// ReadCSVFileLineByLine reads the CSV file line by line and
// returns a channel to read the TrainingData struct.
// Instead of calling this function each time you want to open the file,
// keep the file open and just loop from the start of the file.
func (nn *neuralNetwork) readCSVFileLineByLine() chan trainingData {

	// Create a new TrainingData struct
	data := trainingData{
		i: make([]float64, nn.inputNodes),
		z: make([]float64, nn.outputNodes),
	}

	// Create a channel to read each line and return the TrainingData struct
	ch := make(chan trainingData)

	// Open the CSV file
	file, err := os.Open(nn.trainingDatasetCSVFile)
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
				logmlp.Error("Error:", err)
				return
			}

			// Create a new CSV reader
			reader := csv.NewReader(file)

			// Read the header row
			_, err = reader.Read()
			if err != nil {
				logmlp.Error("Error:", err)
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
							logmlp.Error("Error:", err)
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
							logmlp.Error("Error:", err)
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

// STEP 3 - NORMALIZATION
func (nn *neuralNetwork) normalization(data trainingData) ([]float64, []float64, error) {

	logmlp.Info("    STEP 3 - NORMALIZATION ------------------------------------")

	var x, z []float64
	err := error(nil)

	// STEP 3.1 - NORMALIZE INPUT

	logmlp.Info("        STEP 3.1 - NORMALIZE INPUT")
	if nn.normalizeInputData {
		if nn.normalizeMethod == "zero-to-one" {
			x, err = nn.normalizeZeroToOne("input", data.i)
			if err != nil {
				return nil, nil, err
			}

		} else if nn.normalizeMethod == "minus-one-to-one" {
			x, err = nn.normalizeMinusOneToOne("input", data.i)
			if err != nil {
				return nil, nil, err
			}
		} else {
			return nil, nil, fmt.Errorf("invalid normalization method: %s", nn.normalizeMethod)
		}
	} else {
		x = data.i
	}

	logmlp.Debug(fmt.Sprintf("        Original:    %.5f", data.i))
	logmlp.Debug(fmt.Sprintf("        Normalized:  %.5f", x))

	// STEP 3.2 - NORMALIZE OUTPUT
	logmlp.Info("        STEP 3.2 - NORMALIZE OUTPUT")
	if nn.normalizeOutputData {
		if nn.normalizeMethod == "zero-to-one" {
			z, err = nn.normalizeZeroToOne("output", data.z)
			if err != nil {
				return nil, nil, err
			}
		} else if nn.normalizeMethod == "minus-one-to-one" {
			z, err = nn.normalizeMinusOneToOne("output", data.z)
			if err != nil {
				return nil, nil, err
			}
		} else {
			return nil, nil, fmt.Errorf("invalid normalization method: %s", nn.normalizeMethod)
		}
	} else {
		z = data.z
	}

	logmlp.Debug(fmt.Sprintf("        Original:    %.5f", data.z))
	logmlp.Debug(fmt.Sprintf("        Normalized:  %.5f", z))

	return x, z, nil
}

// STEP 3 - NORMALIZATION
// Normalize the input data between 0 and 1
func (nn *neuralNetwork) normalizeZeroToOne(inOut string, i []float64) ([]float64, error) {

	if inOut == "input" {
		// NORMALIZE INPUTS
		x := make([]float64, nn.inputNodes)
		for j := 0; j < nn.inputNodes; j++ {
			if nn.minInput[j] == nn.maxInput[j] {
				// If min and max are the same, set normalized value to 0.5
				x[j] = 0.5
			} else {
				x[j] = (i[j] - nn.minInput[j]) / (nn.maxInput[j] - nn.minInput[j])
			}
		}

		return x, nil

	} else if inOut == "output" {
		// NORMALIZE OUTPUTS
		z := make([]float64, nn.outputNodes)
		for j := 0; j < nn.outputNodes; j++ {
			if nn.minOutput[j] == nn.maxOutput[j] {
				// If min and max are the same, set normalized value to 0.5
				z[j] = 0.5
			} else {
				z[j] = (i[j] - nn.minOutput[j]) / (nn.maxOutput[j] - nn.minOutput[j])
			}
		}

		return z, nil

	} else {
		return nil, fmt.Errorf("invalid input/output: %s", inOut)
	}

}

// Denormalize the output data between 0 and 1
func (nn *neuralNetwork) denormalizeZeroToOne(inOut string, i []float64) []float64 {
	if inOut == "output" {
		// DENORMALIZE OUTPUTS
		z := make([]float64, nn.outputNodes)
		for j := 0; j < nn.outputNodes; j++ {
			z[j] = i[j]*(nn.maxOutput[j]-nn.minOutput[j]) + nn.minOutput[j]
		}

		return z
	}

	return nil
}

// STEP 3 - NORMALIZATION
// Normalize the input data between -1 and 1
func (nn *neuralNetwork) normalizeMinusOneToOne(inOut string, i []float64) ([]float64, error) {

	if inOut == "input" {
		// NORMALIZE INPUTS
		x := make([]float64, nn.inputNodes)
		for j := 0; j < nn.inputNodes; j++ {
			if nn.minInput[j] == nn.maxInput[j] {
				// If min and max are the same, set normalized value to 0
				x[j] = 0
			} else {
				x[j] = 2*((i[j]-nn.minInput[j])/(nn.maxInput[j]-nn.minInput[j])) - 1
			}
		}

		return x, nil

	} else if inOut == "output" {
		// NORMALIZE OUTPUTS
		z := make([]float64, nn.outputNodes)
		for j := 0; j < nn.outputNodes; j++ {
			if nn.minOutput[j] == nn.maxOutput[j] {
				// If min and max are the same, set normalized value to 0
				z[j] = 0
			} else {
				z[j] = 2*((i[j]-nn.minOutput[j])/(nn.maxOutput[j]-nn.minOutput[j])) - 1
			}
		}
		return z, nil

	} else {
		return nil, fmt.Errorf("invalid input/output: %s", inOut)
	}

}

// Denormalize the output data between -1 and 1
func (nn *neuralNetwork) denormalizeMinusOneToOne(inOut string, i []float64) []float64 {

	if inOut == "output" {
		// DENORMALIZE OUTPUTS
		z := make([]float64, nn.outputNodes)
		for j := 0; j < nn.outputNodes; j++ {
			z[j] = (i[j]+1)*(nn.maxOutput[j]-nn.minOutput[j])/2 + nn.minOutput[j]
		}

		return z
	}

	return nil
}

// STEP 4 - FORWARD PASS
// ForwardPass calculates the output of the neural network
func (nn *neuralNetwork) forwardPass(x []float64) (aHidden [][]float64, yOutput []float64) {

	logmlp.Info("    STEP 4 - FORWARD PASS ----------------------------------")

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

			// STEP 4.1 - THE SUMMATION FUNCTION FOR THE HIDDEN LAYERS
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

			// STEP 4.2 - THE ACTIVATION FUNCTION FOR THE HIDDEN LAYERS
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

		// STEP 4.3 - THE SUMMATION FUNCTION FOR THE OUTPUT LAYER
		for hn := 0; hn < nn.hiddenNodesPerLayer[nn.hiddenLayers-1]; hn++ {
			s += aHidden[nn.hiddenLayers-1][hn] * nn.outputWeights[o][hn]
		}
		s += nn.outputBias[o]

		// STEP 4.4 - THE ACTIVATION FUNCTION FOR THE OUTPUT LAYER
		if nn.activationFunction == "sigmoid" {
			yOutput[o] = sigmoid(s)
		} else if nn.activationFunction == "tanh" {
			yOutput[o] = tanh(s)
		} else {
			return nil, nil
		}

	}

	// Print the outputs
	logmlp.Debug(fmt.Sprintf("        aHidden:     %.5f", aHidden))
	logmlp.Debug(fmt.Sprintf("        yOutput:     %.5f", yOutput))

	return aHidden, yOutput
}

// STEP 4 - FORWARD PASS
// Activation function - Sigmoid
func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// STEP 4 - FORWARD PASS
// Activation function - Sigmoid Derivative
func sigmoidDerivative(x float64) float64 {
	return x * (1 - x)
}

// STEP 4 - FORWARD PASS
// Activation function - Tanh
func tanh(x float64) float64 {
	return math.Tanh(x)
}

// STEP 4 - FORWARD PASS
// Activation function - Tanh Derivative
func tanhDerivative(x float64) float64 {
	return 1 - x*x
}

// STEP 5 - BACKWARD PASS
// BackwardPass calculates the deltas for the neural network
func (nn *neuralNetwork) backwardPass(x []float64, z []float64, yOutput []float64, aHidden [][]float64) ([][][]float64, [][]float64, [][]float64, []float64) {

	logmlp.Info("    STEP 5 - BACKWARD PASS ---------------------------------")

	// STEP 5.1 - THE ERROR SIGNAL FOR THE OUTPUT LAYER
	logmlp.Info("        STEP 5.1 - THE ERROR SIGNAL FOR THE OUTPUT LAYER")
	deltaOutput := make([]float64, nn.outputNodes)
	for o := 0; o < nn.outputNodes; o++ {

		if nn.activationFunction == "sigmoid" {
			deltaOutput[o] = sigmoidDerivative(yOutput[o]) * (yOutput[o] - z[o])
		}
		if nn.activationFunction == "tanh" {
			deltaOutput[o] = tanhDerivative(yOutput[o]) * (yOutput[o] - z[o])
		}
	}

	logmlp.Debug(fmt.Sprintf("        delta O:     %.5f", deltaOutput))

	// STEP 5.2 - THE ERROR SIGNAL FOR THE HIDDEN LAYERS
	logmlp.Info("        STEP 5.2 - THE ERROR SIGNAL FOR THE HIDDEN LAYERS")
	deltaHidden := make([][]float64, nn.hiddenLayers)
	// START LAST HIDDEN LAYER FIRST
	for l := nn.hiddenLayers - 1; l >= 0; l-- {
		deltaHidden[l] = make([]float64, nn.hiddenNodesPerLayer[l])
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			// Last layer first - Use deltaOutput
			if l == nn.hiddenLayers-1 {
				sum := 0.0
				for o := 0; o < nn.outputNodes; o++ {
					sum += nn.outputWeights[o][hn] * deltaOutput[o]
				}
				if nn.activationFunction == "sigmoid" {
					deltaHidden[l][hn] = sigmoidDerivative(aHidden[l][hn]) * sum
				}
				if nn.activationFunction == "tanh" {
					deltaHidden[l][hn] = tanhDerivative(aHidden[l][hn]) * sum
				}
				// All other layers - User deltaHidden
			} else {
				sum := 0.0
				for hnNext := 0; hnNext < nn.hiddenNodesPerLayer[l+1]; hnNext++ {
					sum += nn.hiddenWeights[l+1][hnNext][hn] * deltaHidden[l+1][hnNext]
				}
				if nn.activationFunction == "sigmoid" {
					deltaHidden[l][hn] = sigmoidDerivative(aHidden[l][hn]) * sum
				}
				if nn.activationFunction == "tanh" {
					deltaHidden[l][hn] = tanhDerivative(aHidden[l][hn]) * sum
				}
			}
		}
	}

	logmlp.Debug(fmt.Sprintf("        delta H:     %.5f", deltaHidden))

	// STEP 5.3 - THE NEW WEIGHTS & BIASES FOR THE OUTPUT LAYER
	logmlp.Info("        STEP 5.3 - THE NEW WEIGHTS & BIASES FOR THE OUTPUT LAYER")
	newOutputWeights := make([][]float64, nn.outputNodes)
	for i := range newOutputWeights {
		newOutputWeights[i] = make([]float64, nn.hiddenNodesPerLayer[nn.hiddenLayers-1])
	}
	newOutputBiases := make([]float64, nn.outputNodes)

	for o := 0; o < nn.outputNodes; o++ {
		for hn := 0; hn < nn.hiddenNodesPerLayer[nn.hiddenLayers-1]; hn++ {
			newOutputWeights[o][hn] = nn.outputWeights[o][hn] - (nn.learningRate * deltaOutput[o] * aHidden[nn.hiddenLayers-1][hn])
		}
		newOutputBiases[o] = nn.outputBias[o] - (nn.learningRate * deltaOutput[o])
	}

	logmlp.Debug(fmt.Sprintf("        newOWeights: %.5f", newOutputWeights))
	logmlp.Debug(fmt.Sprintf("        newOBiases:  %.5f", newOutputBiases))

	// STEP 5.4 - THE NEW WEIGHTS & BIASES FOR THE HIDDEN LAYERS
	logmlp.Info("        STEP 5.4 - THE NEW WEIGHTS & BIASES FOR THE HIDDEN LAYERS")
	newHiddenWeights := make([][][]float64, nn.hiddenLayers)
	for i := range newHiddenWeights {
		newHiddenWeights[i] = make([][]float64, nn.hiddenNodesPerLayer[i])
		for j := range newHiddenWeights[i] {
			if i == 0 {
				newHiddenWeights[i][j] = make([]float64, nn.inputNodes)
			} else {
				newHiddenWeights[i][j] = make([]float64, nn.hiddenNodesPerLayer[i-1])
			}
		}
	}
	newHiddenBiases := make([][]float64, nn.hiddenLayers)
	for i := range newHiddenBiases {
		newHiddenBiases[i] = make([]float64, nn.hiddenNodesPerLayer[i])
	}

	// START FROM THE LAST HIDDEN LAYER
	for l := nn.hiddenLayers - 1; l >= 0; l-- {
		for hn := 0; hn < nn.hiddenNodesPerLayer[l]; hn++ {
			if l == 0 {
				// First layer - User X
				for in := 0; in < nn.inputNodes; in++ {
					newHiddenWeights[l][hn][in] = nn.hiddenWeights[l][hn][in] - (nn.learningRate * deltaHidden[l][hn] * x[in])
				}
			} else {
				// Not first layer - Use aHidden Previous Layer
				for in := 0; in < nn.hiddenNodesPerLayer[l-1]; in++ {
					newHiddenWeights[l][hn][in] = nn.hiddenWeights[l][hn][in] - (nn.learningRate * deltaHidden[l][hn] * aHidden[l-1][in])
				}
			}
			newHiddenBiases[l][hn] = nn.hiddenBias[l][hn] - (nn.learningRate * deltaHidden[l][hn])
		}
	}

	logmlp.Debug(fmt.Sprintf("        newHWeights: %.5f", newHiddenWeights))
	logmlp.Debug(fmt.Sprintf("        newHBiases:  %.5f", newHiddenBiases))

	return newHiddenWeights, newHiddenBiases, newOutputWeights, newOutputBiases

}

// Print the min and max input values for each input
func (nn *neuralNetwork) PrintTrainingSummary() {

	// Print Summary of start and end loss for each output node
	fmt.Println("SUMMARY:")
	for i := 0; i < nn.outputNodes; i++ {
		fmt.Println("    Node", i)
		fmt.Println("        Start Loss: ", nn.statStartLoss[i])
		fmt.Println("        End Loss:   ", nn.statEndLoss[i])
		// get percentage change and make it positive
		change := 100 * math.Abs((nn.statEndLoss[i]-nn.statStartLoss[i])/nn.statStartLoss[i])
		fmt.Printf("        Change:      %.2f%%\n", change)
	}

}

// TEST NEURAL NETWORK
// Provide inputs and get outputs
func (nn *neuralNetwork) TestNeuralNetwork(data []float64) error {

	// Normalize inputs
	x, err := nn.normalizeZeroToOne("input", data)
	if err != nil {
		return err
	}

	// Print input and normalized input and min max inputs
	logmlp.Debug("Print the inputs")
	logmlp.Debug(fmt.Sprintf("    Input:  %.2f", data))
	logmlp.Debug(fmt.Sprintf("    Normalized Input:  %.2f", x))
	logmlp.Debug(fmt.Sprintf("    Min Input:  %.2f", nn.minInput))
	logmlp.Debug(fmt.Sprintf("    Max Input:  %.2f", nn.maxInput))

	// Forward pass
	_, yOutput := nn.forwardPass(x)

	// Denormalize the output yOutput
	y := make([]float64, nn.outputNodes)
	if nn.normalizeOutputData {
		if nn.normalizeMethod == "zero-to-one" {
			y = nn.denormalizeZeroToOne("output", yOutput)
		} else if nn.normalizeMethod == "minus-one-to-one" {
			y = nn.denormalizeMinusOneToOne("output", yOutput)
		} else {
			return fmt.Errorf("invalid normalization method: %s", nn.normalizeMethod)
		}
	}

	// Print the outputs to .2 decimal places
	logmlp.Debug("Print the outputs")
	logmlp.Debug(fmt.Sprintf("    Output: Expected %.2f, Got %.2f", data[2], y))

	return nil

}
