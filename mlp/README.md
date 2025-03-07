# MULTI-LAYER PERCEPTRON (MLP) PACKAGE

  _A package to implement a 'single example' scalable multi-layer
  perceptron (MLP) neural network._

TL;DR,

```go
nn := nnp.CreateNeuralNetwork()
nn.PrintNeuralNetwork()
err := nn.GetInputMinMaxFromCSV()
nn.PrintInputMinMax()
err = nn.InitializeNeuralNetwork()
err := nn.TrainNeuralNetwork()
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#overview)
* [CONFIGURE YOUR NEURAL NETWORK](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#configure-your-neural-network)
* [CREATE NEURAL NETWORK](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#create-neural-network)
* [CREATE YOUR TRAINING DATASET FILE](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#create-your-training-dataset-file)
* [GET INPUT MID MAX VALUES OF YOUR DATASET](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#get-input-mid-max-values-of-your-dataset)
* [STEP 1 - INITIALIZATION](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#step-1---initialization)
* [THE TRAINING LOOP](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#the-training-loop)
  * [READING THE CVS DATASET FILE](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#reading-the-cvs-dataset-file)
  * [STEP 2 - NORMALIZATION](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#step-2---normalization)
  * [STEP 3 - FORWARD PASS](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#step-3---forward-pass)
  * [STEP 4 - BACKWARD PASS](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#step-4---backward-pass)
  * [STEP 5 - UPDATE WEIGHTS & BIASES](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#step-5---update-weights--biases)
* [STEP 6 - SAVE WEIGHTS & BIASES](https://github.com/JeffDeCola/my-go-packages/tree/master/mlp#step-6---save-weights--biases)

Documentation and Reference

* [artificial intelligence](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet)
cheat sheet
* [neural networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md)
cheat sheet
* [my-neural-networks](https://github.com/JeffDeCola/my-neural-networks/tree/main)
  * [perceptron-simple-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/perceptron-simple-example)
  * [mlp-classification-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-classification-example)
  * [mlp-image-recognition-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-regression-example)
  * [mlp-regression-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-image-recognition-example)
* [the-math-behind-training-mlp-neural-networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/the-math-behind-training-mlp-neural-networks.md)

## OVERVIEW

A  multi-layer perceptron (MLP) neural network has the following structure,

* The input layer
* The hidden layer(s)
* The output layer

As an example, a neural network with 3 input nodes, 1 hidden layer
with 4 nodes and 2 output nodes would look like,

![IMAGE multi-layer-perceptron-neural-network-scalable IMAGE](../docs/pics/multi-layer-perceptron-neural-network-scalable.svg)

A 'single example' means you train the neural network with one row of data
at a time. This is not the most efficient way to train a neural network,
but it is a good way to understand the process, as opposed to training
with a batch or mini-batch of data.

## CONFIGURE YOUR NEURAL NETWORK

In this package you may configure your multi-layer
perceptron (MLP) neural network by setting the following parameters,

* The number of input nodes
* The input node labels
* The number of hidden layers
* The number of nodes in each hidden layer
* The number of output nodes
* The output node labels
* The number of epochs $E$
* The dataset CSV file
* How to Initialize the weights and biases _random or file_
* The weights and biases CSV file
* Normalize the input data _true or false_
* Normalize the input data _zero-to-one or minus-one-to-one_
* The activation function _sigmoid or tanh_
* The loss function _mean-squared-error or cross-entropy_
* The learning rate $\eta$

You would do this by creating a
`NeuralNetworkParameters` struct.  For example,

```go
nnp := mlp.NeuralNetworkParameters{
  InputNodes:              2,
  InputNodeLabels:         []string{"midterm-grade", "hours-studied", "last-test-grade"},
  HiddenLayers:            1,
  HiddenNodesPerLayer:     []int{3},
  OutputNodes:             1,
  OutputNodeLabels:        []string{"predicted-percentage-passing-final", "predicted-final-grade"},
  Epochs:                  100,
  DatasetCSVFile:          "dataset.csv",
  Initialization:          "file",               // or "random"
  WeightsAndBiasesCSVFile: "weights-and-biases.csv",
  NormalizeInputData:      true,                 // or false
  NormalizeMethod:         "zero-to-one",        // or "minus-one-to-one"
  ActivationFunction:      "sigmoid",            // or "tanh"
  LossFunction:            "mean-squared-error", // or "cross-entropy"
  LearningRate:            0.1,
}
```

## CREATE NEURAL NETWORK

To create a neural network, you take your parameters and feed them
into the `CreateNeuralNetwork` method which will return a
`NeuralNetwork` struct.

```go
nn := nnp.CreateNeuralNetwork()
```

You can also print out the neural network structure if you would like,

```go
nn.PrintNeuralNetwork()
```

## CREATE YOUR TRAINING DATASET FILE

You will use a standard csv file with the first row being the labels
and the rest of the rows being the input and target output data.
For example, a dataset could look something like this,

```csv
midterm-grade,hours-studied,last-test-grade,pred-perc-passing-final,pred-final-grade
89,48,79,80,82
75,23,85,70,78
etc...
```

## GET INPUT MID MAX VALUES OF YOUR DATASET

Before you start training, we need to find the min and max values
of your dataset (your csv file). The min and max values will be
used to normalize your dataset.
This is done by calling the `GetInputMinMaxFromCSV` method,

```go
err := nn.GetInputMinMaxFromCSV()
```

You can print out the min and max values if you want,

```go
nn.PrintInputMinMax()
```

You chose what the filename is in the `NeuralNetworkParameters` struct,

```go
DatasetCSVFile: "filename.csv",
```

## STEP 1 - INITIALIZATION

The first step in training a neural network is to initialize the weights
and bias. You can read the weights and bias from a json file or initialize them
randomly.

```go
err = nn.InitializeNeuralNetwork()
```

You chose how to initialize the weights and biases in the
`NeuralNetworkParameters` struct,

```go
Initialization:          "file" |  "random"
```

## THE TRAINING LOOP

Now that out neural network is configured and we have our dataset,
we can train our neural network.
To put it simple, training a neural network is
the process of adjusting the weights
of the network in order to minimize the loss in the output from
the network.
To achieve this we use a optimization technique called **Stochastic Gradient Descent**.
We calculate loss using a loss function and calculate the derivate and we update
the weights during backpropagation. The main goal is to minimize
the difference(loss) between predicted output and actual output.
MLP uses a supervised learning technique called **backpropagation** for training.
In our case, we will call the `TrainNeuralNetwork` method,

```go
err := nn.TrainNeuralNetwork()
```

This one method does a lot of heavy lifting so let's break it down.
If you're not interested in these details, you can skip to step 7.

There are two loops in the `TrainNeuralNetwork` method,

1. The outer loop is the number of epochs $E$.
2. The inner loop is the number of rows in the dataset.

```go
nn.epochLoop()
```

which calls the datasetLoop method,

```go
nn.datasetLoop()
```

You can set the number of epochs in the `NeuralNetworkParameters` struct,

```go
Epochs:       #
```

### READING THE CVS DATASET FILE

We will not store the csv file in memory, but rather read it line by line.
This is because the csv file could be very large.
But it will take a little longer to train. This is the trade off.

```go
ch := nn.readCSVFileLineByLine()
```

The channel `ch` will contain each line of the csv file.

You chose what the filename is in the `NeuralNetworkParameters` struct,

```go
DatasetCSVFile: "filename.csv",
```

### STEP 2 - NORMALIZATION

Normalization, also called min-max scaling, changes the values of
input data set to occupy a range of [0, 1] or [-1, 1],
reducing the influence of unusual values of out model.
We will normalize the input data between 0 and 1.
This is done by the `normalizeInputData` method.

```go
data = nn.normalizeInputData(data)
```

You may turn normalization on/off and chose the method
in the `NeuralNetworkParameters` struct,

```go
NormalizeInputData: true | false
NormalizeMethod:    "zero-to-one" | "minus-one-to-one"
```

### STEP 3 - FORWARD PASS

Giving our normalized $x_{[0]}$, $x_{[1]}$ and $x_{[2]}$ input training data,
**compute the output for each layer and
propagate through layers to obtain the outputs
$y_{[0]}$ and $y_{[1]}$**.

The method `forwardPass` does this,

```go
yOutput, yHidden := nn.forwardPass(x)
```

You may chose the activation function in the `NeuralNetworkParameters` struct,

```go
ActivationFunction: "sigmoid" | "tanh"
```

### STEP 4 - BACKWARD PASS

Now  that we have the outputs $y$, calculate the error (delta **$\delta$**)
between target data ($z$) and actual output ($y$)
and propagate backwards.

### STEP 5 - UPDATE WEIGHTS & BIASES

## STEP 6 - SAVE WEIGHTS & BIASES
