package db

import "github.com/dim-pep/task-from-hitalent/internal/models"

func CreateAnswer(a *models.Answer) error {
	if err := DbConn.Create(a).Error; err != nil {
		return err
	}
	return nil
}

func GetAnswerByID(id int, a *models.Answer) error {
	if err := DbConn.First(a, id).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAnswerByID(id int) error {
	if err := DbConn.Delete(&models.Answer{}, id).Error; err != nil {
		return err
	}
	return nil
}
