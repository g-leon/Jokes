package actions

import (
	"testing"
	"net/http/httptest"
	"github.com/g-leon/Jokes/models"
	"encoding/json"
	"net/http"
)

var inputPerson *models.Person = &models.Person{Name: "Donald", Surname: "Trump", Gender: "male", Region: "USA"}

type TestHandler struct {}
func (th *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(inputPerson)
	w.Write(resp)
}

func TestPersonService_Fetch(t *testing.T) {
	server := httptest.NewServer(&TestHandler{})
	defer server.Close()

	p, err := PersonService{server.URL}.Fetch()
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != inputPerson.Name {
		t.Fatalf("Expected %s, got %s", inputPerson.Name, p.Name)
	}
	if p.Surname != inputPerson.Surname {
		t.Fatalf("Expected %s, got %s", inputPerson.Surname, p.Surname)
	}
}

