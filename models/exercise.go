package models

type Exercise struct {
	ExerciseId int64  `json:"id"`
	Name string `json:"name"`
	Repetitions int64  `json:"repetitions"`
	Time int64  `json:"time"`
	Complexity string `json:"complexity"`
}