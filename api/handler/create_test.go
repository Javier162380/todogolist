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
	r.POST("createtask/", CreateHandler(repository))

	t.Run("Invalid Request returns a 400", func(t *testing.T) {
		createtaskrequests := taskRequest{TaskLabel: "work", TaskDate: "2020-12-12"}

		b, err := json.Marshal(createtaskrequests)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/creaatetask", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

}
