package result

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrainingResult_IncSamplesCount(t *testing.T) {
	counts := []int64{1, 13, 24}
	for _, count := range counts {
		result := NewTrainingResult()
		for i := int64(0); i < count ; i++ {
			result.IncSamplesCount()
		}
		assert.Equal(t, count, result.SamplesCount, "The samples counter must works properly")
	}
}

func TestTrainingResult_IncClassCount(t *testing.T) {
	// This data mean that class#1 seen 7 times and class#255 seen 8 times
	classesFrequencies := map[int64]int64{1: 7, 255: 8}
	result := NewTrainingResult()
	for classID, count := range classesFrequencies {
		for i := int64(0); i < count ; i++ {
			result.IncClassCount(classID)
		}
		assert.Equal(t, count, result.ClassFrequency[classID], "The class frequency counter must works properly")
	}
}

func TestTrainingResult_IncTokenCount(t *testing.T) {
	// This data means than token#2 seen 3 times in class#1, and token#3 seen 4 times in class#2
	frequencies := map[int64]map[int64]int64{1: {2: 3}, 2: {3: 4}}
	result := NewTrainingResult()
	for classID, tokenFrequencies := range frequencies {
		for tokenId, count := range tokenFrequencies {
			for i := int64(0); i < count ; i++ {
				result.IncTokenCount(classID, tokenId)
			}
			assert.Equal(t, count, result.TokenFrequency[classID][tokenId], "The token frequency counter must works properly")
		}
	}
}