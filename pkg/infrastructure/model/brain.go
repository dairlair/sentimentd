package model

type Brain struct {
	Model
	Name string
	Description string
}

func (brain *Brain) GetID() int64 {
	return brain.Model.ID
}

func (brain *Brain) GetName() string {
	return brain.Name
}

func (brain *Brain) GetDescription() string {
	return brain.Description
}