package db

import (
	"github.com/dim-pep/task-from-hitalent/internal/models"
)

func CreateQuestion(q models.Question) error {
	if err := DbConn.Create(&q).Error; err != nil {
		return err
	}
	return nil
}

func GetQuestions(q *[]models.Question) error {
	if err := DbConn.Order("created_at DESC").Find(&q).Error; err != nil {
		return err
	}
	return nil
}

func GetQuestionWithAnswers(id int, q *models.Question) error {
	if err := DbConn.Preload("Answers").First(q, id).Error; err != nil {
		return err
	}
	return nil
}

func DeleteQuestionByID(id int) error {
	if err := DbConn.Delete(&models.Question{}, id).Error; err != nil {
		return err
	}
	return nil
}

func QuestionExists(id int) (bool, error) {
	var count int64
	if err := DbConn.Model(&models.Question{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
