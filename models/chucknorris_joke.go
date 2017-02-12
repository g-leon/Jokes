package models

type ChuckNorrisJoke struct {
	Type string
	Value struct {
		Id int32
		Joke string
		Categories []string
	}
}
