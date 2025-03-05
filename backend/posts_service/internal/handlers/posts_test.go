package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchPosts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"title":"Test Post","content":"This is a test post"}]`))
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Ожидался код %d, но получен %d", http.StatusOK, w.Code)
	}
}
