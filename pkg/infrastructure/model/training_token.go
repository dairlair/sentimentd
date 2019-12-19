package model

type TrainingToken struct {
	Model
	BrainID int64
	TrainingID int64
	ClassID int64
	TokenID int64
	SamplesCount int64
}