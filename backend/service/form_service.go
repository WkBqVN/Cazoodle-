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

type Record struct {
	Title string
	Value interface{}
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

func (s *FormService) GetFormData(survey_id, form_id int) (interface{}, error) {
	// surveyData := s.DB.Find(&survey_id)
	var f model.Forms
	output := s.Repo.DB.First(&f, form_id)
	if output.Error != nil {
		return nil, output.Error
	}
	formData := make(map[string]Record)
	if err := json.Unmarshal([]byte(f.FormData), &formData); err != nil {
		return nil, err
	}
	result, err := utils.ConvertMapToStruct(formData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SaveData(survey_id, form_id int, data map[string]interface{}) {

}

func GetFormById(id int) {

}

func ValidateForm(survey_id, formid int) bool {
	return true
}

// func GetSurveyById()
