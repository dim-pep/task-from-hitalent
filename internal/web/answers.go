package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dim-pep/task-from-hitalent/internal/db"
	"github.com/dim-pep/task-from-hitalent/internal/models"
)

func postAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	trim := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(trim, "/")
	if len(parts) < 3 || parts[0] != "questions" || parts[2] != "answers" { // ["questions", "id", "answers"]
		http.Error(w, "неправильный путь", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil || id <= 0 {
		http.Error(w, "неправильный id вопроса", http.StatusBadRequest)
		return
	}

	exists, err := db.QuestionExists(id)
	if err != nil {
		log.Println("Ошибка при попытке найти вопрос в БД:", err)
		http.Error(w, "Ошибка при попытке найти вопрос в БД", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "вопрос не существует!", http.StatusNotFound)
		return
	}

	var body struct {
		UserID string `json:"user_id"`
		Text   string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "ошибка сериализации", http.StatusBadRequest)
		return
	}
	body.UserID = strings.TrimSpace(body.UserID)
	body.Text = strings.TrimSpace(body.Text)

	if body.UserID == "" || body.Text == "" {
		http.Error(w, "user_id и text обязательны", http.StatusBadRequest)
		return
	}

	a := models.Answer{
		QuestionID: id,
		UserID:     body.UserID,
		Text:       body.Text,
	}

	if err := db.CreateAnswer(&a); err != nil {
		log.Println("Ошибка создания ответа в БД:", err)
		http.Error(w, "ошибка при создании ответа", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(a); err != nil {
		log.Println("ошибка сериализации:", err)
	}
}

func getDeleteAnswerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trim := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(trim, "/")
	if len(parts) != 2 || parts[0] != "answers" {
		http.Error(w, "неправильный путь", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil || id <= 0 {
		http.Error(w, "неправильный id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		var a models.Answer
		if err := db.GetAnswerByID(id, &a); err != nil {
			log.Println("Ошибка при поиске ответа в БД:", err)
			http.Error(w, "ответ не найден", http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(a); err != nil {
			log.Println("ошибка сериализации:", err)
			http.Error(w, "ошибка сериализации", http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		if err := db.DeleteAnswerByID(id); err != nil {
			log.Println("Ошибка удаления ответа в БД:", err)
			http.Error(w, "ошибка удаления ответа", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, DELETE")
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
