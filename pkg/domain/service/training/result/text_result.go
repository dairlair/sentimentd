// This package contains data types for the training results representation.
// Each training for Bayes algorithm provides these data:
//   - Classes frequencies
//   - Features per classes frequencies
//   - Total number of samples processed during the training
package result

// This map contains classes frequencies in the training dataset.
// The key of map is a Class ID and values is a count of samples with this class.
type TextClassFrequency map[string]int64

// This map contains tokens (or features) frequencies for classes in the training dataset.
// The first-level key is a Class ID, the second-level key is Token ID, and the value is a number of sample
// which have the corresponding class and token.
type TextTokenFrequency map[string]map[string]int64

type TextTrainingResult struct {
	SamplesCount   int64 // The total number of samples seen in dataset
	TextClassFrequency TextClassFrequency
	TextTokenFrequency TextTokenFrequency
}

func (t *TextTrainingResult) IncSamplesCount() {
	t.SamplesCount++
}

func (t *TextTrainingResult) IncClassCount(class string) {
	t.TextClassFrequency[class]++
}

func (t *TextTrainingResult) IncTokenCount(class, token string) {
	if t.TextTokenFrequency[class] == nil {
		t.TextTokenFrequency[class] = make(map[string]int64, 1)
	}
	t.TextTokenFrequency[class][token]++
}

func NewTextTrainingResult () *TextTrainingResult {
	return &TextTrainingResult {
		SamplesCount:   0,
		TextClassFrequency: map[string]int64{},
		TextTokenFrequency: map[string]map[string]int64{},
	}
}