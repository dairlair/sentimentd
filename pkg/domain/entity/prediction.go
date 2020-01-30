package entity

import (
	"encoding/json"
	"time"
)

// Prediction represents results of predictor
type Prediction struct {
	probabilities map[int64]float64
}

// NewPrediction is a constructor
func NewPrediction(probabilities map[int64]float64) Prediction {
	return Prediction{
		probabilities: probabilities,
	}
}

// GetClassIDs Returns identifiers list of predicted classes
func (p *Prediction) GetClassIDs() []int64 {
	ids := make([]int64, len(p.probabilities))
	var i int64 = 0
	for id := range p.probabilities {
		ids[i] = id
		i++
	}

	return ids
}

// GetClassProbability returns identifiers list of predicted classes
func (p *Prediction) GetClassProbability(classID int64) float64 {
	if probability, ok := p.probabilities[classID]; ok {

		return probability
	}

	return 0
}

// HumanizedPrediction describes structure which contains class labels and their probabilities
type HumanizedPrediction struct {
	Probabilities map[string]float64 `json:"probabilities"`
	Duration      float64            `json:"duration"`
}

// NewHumanizedPrediction is a constructor
func NewHumanizedPrediction(probabilities map[string]float64, duration time.Duration) HumanizedPrediction {
	return HumanizedPrediction{
		Probabilities: probabilities,
		Duration:      duration.Seconds(),
	}
}

// JSON helps to convert structure to string
func (hp HumanizedPrediction) JSON() (string, error) {
	str, err := json.Marshal(hp)
	if err != nil {
		return "", err
	}

	return string(str), nil
}
