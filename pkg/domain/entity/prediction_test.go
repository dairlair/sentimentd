package entity

import (
	"testing"
)

func BenchmarkPrediction_GetClassIds(b *testing.B) {
	prediction := NewPrediction(map[int64]float64{
		1: 0.967,
		2: 0.54,
		3: 0.54,
		4: 0.22,
	})

	for i := 0; i < b.N; i++ {
		prediction.GetClassIDs()
	}
}