package entity

type Prediction struct {
	probabilities map[int64]float64
}

func NewPrediction(probabilities map[int64]float64) Prediction {
	return Prediction{
		probabilities:probabilities,
	}
}

// Returns identifiers list of predicted classes
func (p *Prediction) GetClassIDs() []int64 {
	ids := make([]int64, len(p.probabilities))
	for id, _ := range p.probabilities {
		ids = append(ids, id)
	}

	return ids
}

// Returns identifiers list of predicted classes
func (p *Prediction) GetClassProbability(classID int64) float64 {
	if probability, ok := p.probabilities[classID]; ok {

		return probability
	}

	return 0
}