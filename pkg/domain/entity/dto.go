package entity

// This map contains classes frequencies in the training dataset.
// The key of map is a Class ID and values is a count of samples with this class.
type classFrequencies map[int64]int64

// This map contains tokens (or features) frequencies for classes in the training dataset.
// The first-level key is a Class ID, the second-level key is Token ID, and the value is a number of sample
// which have the corresponding class and token.
type tokenFrequencies map[int64]map[int64]int64

type trainingResult struct {
	samplesSeen int64 // The total number of samples seen in dataset
	classFrequencies classFrequencies
	tokenFrequencies tokenFrequencies
}