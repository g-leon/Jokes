package actions

import (
	"github.com/g-leon/Jokes/models"
	"strings"
)

// RandomJokeService is a service that generates a random person
// and a random joke and then applies the joke to that person
type RandomJokeService struct {}

func (rjs RandomJokeService) Fetch(cnJokeFetcher ChuckNorrisJokeFetcher, pFetcher PersonFetcher) (
	*models.RandomJoke, error) {

	cnJoke, err := cnJokeFetcher.Fetch()
	if err != nil {
		return nil, err
	}

	person, err := pFetcher.Fetch()
	if err != nil {
		return nil, err
	}

	joke := strings.Replace(cnJoke.Value.Joke, "John Doe", person.Name + " " + person.Surname, -1)
	randomJoke := &models.RandomJoke{Value: joke}

	return randomJoke, nil
}
