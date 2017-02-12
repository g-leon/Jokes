package actions

import (
	"testing"
	"github.com/g-leon/Jokes/models"
)

type PersonServiceMock struct {}
func (p PersonServiceMock) Fetch() (*models.Person, error) {
	return &models.Person{Name: "Donald", Surname: "Trump", Gender: "male", Region: "USA"}, nil
}

type ChuckNorrisJokeServiceMock struct {}
func (c ChuckNorrisJokeServiceMock) Fetch() (*models.ChuckNorrisJoke, error) {
	return &models.ChuckNorrisJoke{
		Value: struct {
			Id int32
			Joke string
			Categories []string
		}{
			123,
			in,
			[]string{"usa", "president"},
		},
	}, nil
}

var in string = "What's the difference between God and John Doe? God doesn't think he is John Doe"
var out string = "What's the difference between God and Donald Trump? God doesn't think he is Donald Trump"

func TestRandomJokeService_Fetch(t *testing.T) {
	joke, err := RandomJokeService{}.Fetch(ChuckNorrisJokeServiceMock{}, PersonServiceMock{})
	if err != nil {
		t.Fatal(err)
	}

	if joke.Value != out {
		t.Errorf("Expected %s, got %s", out, joke.Value)
	}
}
