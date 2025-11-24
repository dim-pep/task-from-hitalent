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

func getPostQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		log.Println("MethodPost")
		var body struct {
			Text string `json:"text"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "неправильный JSON", http.StatusBadRequest)
			return
		}

		body.Text = strings.TrimSpace(body.Text)
		if body.Text == "" {
			http.Error(w, "В вопросе нет текста", http.StatusBadRequest)
			return
		}

		q := models.Question{
			Text: body.Text,
		}

		err := db.CreateQuestion(q)
		if err != nil {
			log.Println("Ошибка создания вопроса в БД:", err)
			http.Error(w, "ошибка при создание вопроса в БД", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		var qs []models.Question

		err := db.GetQuestions(&qs)
		if err != nil {
			log.Println("Ошибка получения вопросов из БД:", err)
			http.Error(w, "Ошибка получения вопросов из БД", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(qs); err != nil {
			log.Println("ошибка сириализации:", err)
			http.Error(w, "ошибка сириализации:", http.StatusInternalServerError)
			return
		}
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
}

func getDeleteQuestionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trim := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(trim, "/")
	if len(parts) != 2 || parts[0] != "questions" {
		http.Error(w, "Неправильный путь", http.StatusBadRequest)
		return
	}

	idStr := parts[1]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "неправильный id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:

		var q models.Question

		// err = db.GetQuestionByID(id, &q)
		err = db.GetQuestionWithAnswers(id, &q)
		if err != nil {
			log.Println("Ошибка получения вопроса:", err)
			http.Error(w, "вопрос не найден", http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(q); err != nil {
			log.Println("ошибка сериализации:", err)
			http.Error(w, "ошибка сериализации", http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		err = db.DeleteQuestionByID(id)
		if err != nil {
			log.Println("Ошибка удаления вопроса:", err)
			http.Error(w, "ошибка удаления вопроса", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	default:
		w.Header().Set("Allow", "GET, DELETE")
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

//getDeleteIDQuestions
