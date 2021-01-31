package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todogolist/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTaskHandler(t *testing.T) {
	repository := repository.NewRepository("memory", ".")

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/createtask", CreateHandler(repository))

	t.Run("Invalid Request returns a 400", func(t *testing.T) {
		createtaskrequests := taskRequest{TaskLabel: "work"}

		b, err := json.Marshal(createtaskrequests)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/createtask", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("Correct Request returns a 200", func(t *testing.T) {
		createtaskrequests := taskRequest{TaskLabel: "work", TaskName: "124"}

		b, err := json.Marshal(createtaskrequests)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/createtask", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

	})

}
