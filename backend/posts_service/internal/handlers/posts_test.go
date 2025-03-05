package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFetchPosts(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка создания мока базы данных: %v", err)
	}
	defer db.Close()

	// Ожидаем вызов FetchPosts и возвращаем тестовые данные
	rows := sqlmock.NewRows([]string{"id", "title", "content"}).
		AddRow(1, "Test Post", "This is a test post")
	mock.ExpectQuery("SELECT id, title, content FROM posts").WillReturnRows(rows)

	// Создаем HTTP-запрос
	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	w := httptest.NewRecorder()

	// Вызываем обработчик
	handler := FetchPosts(db)
	handler(w, req)

	// Проверяем код ответа
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался код %d, но получен %d", http.StatusOK, w.Code)
	}

	// Проверяем, что запрос в базу был ожидаемым
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Не все ожидания были выполнены: %v", err)
	}
}
