package model

type Class struct {
	Model
	BrainID int64
	Label string
}

func (class *Class) GetID() int64 {
	return class.Model.ID
}

func (class *Class) GetLabel() string {
	return class.Label
}