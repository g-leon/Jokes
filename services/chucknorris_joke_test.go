package actions

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/g-leon/Jokes/models"
	"encoding/json"
)

var inputJoke *models.ChuckNorrisJoke = &models.ChuckNorrisJoke{
	Value: struct {
		Id         int32
		Joke       string
		Categories []string
	}{
		123,
		"Best joke ever",
		[]string{"usa", "president"},
	},
}

type TestCNJokeHandler struct {}
func (th *TestCNJokeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(inputJoke)
	w.Write(resp)
}

func TestChuckNorrisJokeService_Fetch(t *testing.T) {
	server := httptest.NewServer(&TestCNJokeHandler{})
	defer server.Close()

	cnJoke, err := ChuckNorrisJokeService{server.URL}.Fetch()
	if err != nil {
		t.Fatal(err)
	}
	if cnJoke.Value.Joke != inputJoke.Value.Joke {
		t.Fatalf("Expected %s, got %s", inputJoke.Value.Joke, cnJoke.Value.Joke)
	}
}
