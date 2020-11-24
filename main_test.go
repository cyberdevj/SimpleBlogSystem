package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	Router().ServeHTTP(w, req)
	return w
}

// GET /articles/
// GET /articles/:id
// POST /articles/
// PATCH /articles/:id
// DELETE /articles/:id
// TestArticlesGET function ...
func TestArticlesCRUD(t *testing.T) {
	initDb()

	t.Run("Retrieving all articles", func(t *testing.T) {
		w := performRequest("GET", "/articles")
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Retrieving article with valid ID", func(t *testing.T) {
		w := performRequest("GET", "/articles/5fbc39e2ede5196ac29c2bde")
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Retrieving article with invalid ID", func(t *testing.T) {
		w := performRequest("GET", "/articles/0")
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Retrieving article with missing ID", func(t *testing.T) {
		w := performRequest("GET", "/articles/")
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Insert multiple articles", func(t *testing.T) {
		articlesList := []*Article{
			{
				Title:   "Example1",
				Content: "ExampleContent1",
				Author:  "ExampleAuthor1",
			},
			{
				Title:   "Example2",
				Content: "ExampleContent2",
				Author:  "ExampleAuthor2",
			},
			{
				Title:   "Example3",
				Content: "ExampleContent3",
				Author:  "ExampleAuthor3",
			},
			{
				Title:   "Example4",
				Content: "ExampleContent4",
				Author:  "ExampleAuthor4",
			},
			{
				Title:   "Example5",
				Content: "ExampleContent5",
				Author:  "ExampleAuthor5",
			},
		}

		for _, article := range articlesList {
			payload, _ := json.Marshal(article)
			req, err := http.NewRequest("POST", "/articles", bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			Router().ServeHTTP(w, req)

			assert.Equal(t, nil, err)
			assert.Equal(t, http.StatusOK, w.Code)
		}
	})

	t.Run("Patch request for articles", func(t *testing.T) {
		w := performRequest("PATCH", "/articles")
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})

	t.Run("Delete request for articles", func(t *testing.T) {
		w := performRequest("DELETE", "/articles")
		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})
}
