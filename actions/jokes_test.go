package actions

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/markbates/willie"
	"log"
)

func Test_JokesHandler(t *testing.T) {
	r := require.New(t)

	w := willie.New(App())
	res := w.Request("/").Get()

	log.Print(res.Body)
	r.Equal(200, res.Code)
}
