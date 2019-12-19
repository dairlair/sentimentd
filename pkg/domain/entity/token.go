package entity

type TokenInterface interface {
	GetID() int64
	GetText() string
}