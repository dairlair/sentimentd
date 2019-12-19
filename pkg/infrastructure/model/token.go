package model

type Token struct {
	Model
	BrainID int64
	Text string
}

func (class *Token) GetID() int64 {
	return class.Model.ID
}

func (class *Token) GetText() string {
	return class.Text
}