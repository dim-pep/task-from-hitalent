package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dim-pep/task-from-hitalent/config"
	"github.com/dim-pep/task-from-hitalent/internal/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestQuestionsHandlers_getPostQuestions(t *testing.T) {
	godotenv.Load("../../.env")
	conf := config.LoadConfig()
	conf.DBHost = "localhost"
	db.Conn(conf)

	postBody := map[string]string{"text": "2+2?"}
	data, err := json.Marshal(postBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/questions", bytes.NewReader(data))
	rr := httptest.NewRecorder()

	getPostQuestions(rr, req)
	require.Equal(t, http.StatusCreated, rr.Code)
}
