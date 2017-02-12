package actions

import (
	"github.com/g-leon/Jokes/models"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type ChuckNorrisJokeFetcher interface {
	Fetch() (*models.ChuckNorrisJoke, error)
}

type ChuckNorrisJokeService struct {
	Url string
}

func (c ChuckNorrisJokeService) Fetch() (*models.ChuckNorrisJoke, error) {
	resp, err := http.Get(c.Url)
	if err != nil {
		return nil, err
	}

	cnJoke := &models.ChuckNorrisJoke{}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(respBody, cnJoke)

	return cnJoke, nil
}
