// This package contains data types for the training results representation.
// Each training for Bayes algorithm provides these data:
//   - Classes frequencies
//   - Features per classes frequencies
//   - Total number of samples processed during the training
package result

// This map contains classes frequencies in the training dataset.
// The key of map is a Class ID and values is a count of samples with this class.
type ClassFrequency map[int64]int64

// This map contains tokens (or features) frequencies for classes in the training dataset.
// The first-level key is a Class ID, the second-level key is Token ID, and the value is a number of sample
// which have the corresponding class and token.
type TokenFrequency map[int64]map[int64]int64

type TrainingResult struct {
	SamplesCount   int64 // The total number of samples seen in dataset
	ClassFrequency ClassFrequency
	TokenFrequency TokenFrequency
}

func (t *TrainingResult) IncSamplesCount() {
	t.SamplesCount++
}

func (t *TrainingResult) IncClassCount(classID int64) {
	t.ClassFrequency[classID]++
}

func (t *TrainingResult) IncTokenCount(classID, tokenID int64) {
	if t.TokenFrequency[classID] == nil {
		t.TokenFrequency[classID] = make(map[int64]int64, 1)
	}
	t.TokenFrequency[classID][tokenID]++
}

func NewTrainingResult () *TrainingResult {
	return &TrainingResult {
		SamplesCount:   0,
		ClassFrequency: map[int64]int64{},
		TokenFrequency: map[int64]map[int64]int64{},
	}
}