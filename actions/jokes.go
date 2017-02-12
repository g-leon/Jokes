package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/g-leon/Jokes/services"
	"net/http"
)

func JokesHandler(c buffalo.Context) error {
	randomJoke, err := actions.RandomJokeService{}.Fetch(
		actions.ChuckNorrisJokeService{"http://api.icndb.com/jokes/random?firstName=John&lastName=Doe"},
		actions.PersonService{"http://uinames.com/api"},
	)
	if err != nil {
		c.Logger().Error(err)
		var msg string = "It seems the server is experiencing some issues. Please retry later."
		return c.Render(500, r.JSON(map[string]interface{}{
			"Status": http.StatusInternalServerError,
			"Error": msg,
		}))
	}

	return c.Render(200, r.JSON(randomJoke))
}
