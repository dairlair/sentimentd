package model

type TrainingClass struct {
	Model
	BrainID int64
	TrainingID int64
	ClassID int64
	SamplesCount int64
}