package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"

	L "americanas.teste/lists" 
	S "americanas.teste/structs"
)

func TestSaveListsRoute(t *testing.T) {
	r := gin.Default()
	r.POST("/saveLists", L.SaveLists)

	t.Run("Salvar Listas com Sucesso", func(t *testing.T) {
		lists := S.RequestSave{List1: []int{1, 3, 5}, List2: []int{2, 4, 6}}
		payload, _ := json.Marshal(lists)

		req, _ := http.NewRequest("POST", "/saveLists", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Erro ao Processar JSON", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/saveLists", bytes.NewBuffer([]byte("invalid JSON")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestMergeRoute(t *testing.T) {
	r := gin.Default()
	r.GET("/merge", L.Merge)

	t.Run("Obter Lista Mesclada", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/merge", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}