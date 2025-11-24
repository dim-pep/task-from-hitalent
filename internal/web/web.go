package web

import (
	"log"
	"net/http"
	"strings"

	"github.com/dim-pep/task-from-hitalent/config"
)

func StartWeb(conf config.Config) {
	http.HandleFunc("/questions", getPostQuestions)
	http.HandleFunc("/answers/", getDeleteAnswerByID)
	http.HandleFunc("/questions/", questionsSubrouter)

	log.Println("Сервер запущен на :" + conf.AppPort)
	if err := http.ListenAndServe(":"+conf.AppPort, nil); err != nil {
		log.Fatal(err)
	}
}

func questionsSubrouter(w http.ResponseWriter, r *http.Request) {
	trim := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(trim, "/")
	if len(parts) >= 3 && parts[2] == "answers" {
		postAnswer(w, r)
		return
	}
	getDeleteQuestionByID(w, r)
}
