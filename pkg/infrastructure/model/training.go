package model

type Training struct {
	Model
	BrainID int64
	Comment string
	SamplesCount int64
}