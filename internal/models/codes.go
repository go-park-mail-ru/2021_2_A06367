package models

type StatusCode int

const (
	Okey StatusCode = iota
	NotFound
	InternalError
	Unauthed
	Conflict
	InvalidBody
)
