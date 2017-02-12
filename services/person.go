package actions

import (
	"github.com/g-leon/Jokes/models"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type PersonFetcher interface {
	Fetch() (*models.Person, error)
}

type PersonService struct {
	Url string
}

func (ps PersonService) Fetch() (*models.Person, error) {
	resp, err := http.Get(ps.Url)
	if err != nil {
		return nil, err
	}

	p := &models.Person{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, p)

	return p, nil
}