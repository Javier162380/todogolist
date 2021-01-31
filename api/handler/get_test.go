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

func TestGetHandler(t *testing.T) {
	repository := repository.NewRepository("memory", ".")

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/gettask", GetHandler(repository))

	t.Run("Test Get Handler invalid request return 400", func(t *testing.T) {

		b, err := json.Marshal(map[string]interface{}{"task_name": "pepe", "task_label": 345})
		require.NoError(t, err)

		req, err := http.NewRequest("GET", "/gettask", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("Test Get Handler valid request return 200", func(t *testing.T) {

		gr := getRequest{TaskName: "12423", TaskLabel: "workf"}

		b, err := json.Marshal(gr)
		require.NoError(t, err)

		req, err := http.NewRequest("GET", "/gettask", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

	})

}
