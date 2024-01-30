package service

import (
	"encoding/json"
	"sync"

	"cazoodle.com/model"
	"cazoodle.com/repository"
	"cazoodle.com/utils"
)

type FormService struct {
	Repo *repository.Repository
}

var s *FormService
var once sync.Once

func GetFormServiceInstance() *FormService {
	once.Do(func() {
		s = &FormService{
			Repo: repository.GetInstance(),
		}
	})
	return s
}

func (s *FormService) GetFormData(survey_id, form_id int) ([]interface{}, error) {
	// surveyData := s.DB.Find(&survey_id)
	var f model.Forms
	output := s.Repo.DB.First(&f, form_id)
	if output.Error != nil {
		return nil, output.Error
	}
	var formData []map[string]interface{}
	if err := json.Unmarshal([]byte(f.FormData), &formData); err != nil {
		return nil, err
	}
	var result []interface{}
	for _, value := range formData {
		data, err := utils.ConvertMapToStruct(value)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}

func SaveData(survey_id, form_id int, data []map[string]interface{}) {

}

func GetFormById(id int) {

}

func ValidateForm(survey_id, formid int) bool {
	return true
}

// func GetSurveyById()
