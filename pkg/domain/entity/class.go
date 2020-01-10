package entity

type ClassInterface interface {
	GetID() int64
	GetLabel() string
}

// This map contains classes size in the training dataset.
// The key of map is a Class ID and values is a total tokens count in the each class.
type ClassSizeMap map[int64]int64