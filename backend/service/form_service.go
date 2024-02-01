package service

import (
	"encoding/json"
	"errors"
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

func (s *FormService) GetFormTemplate(form_template_id int) ([]interface{}, error) {
	if form_template_id == 0 {
		return nil, nil
	}
	var f model.FormTemplate
	output := s.Repo.DB.First(&f, form_template_id)
	if output.Error != nil {
		return nil, output.Error
	}
	var formData []map[string]interface{}
	if err := json.Unmarshal([]byte(f.FormTemplate), &formData); err != nil {
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

func (s *FormService) SaveData(client_id, survey_id, form_id int, formData string) error {
	var c model.Client
	output := s.Repo.DB.First(&c, client_id)
	if output.Error != nil {
		return output.Error
	}
	var svey model.Survey
	output = s.Repo.DB.First(&svey, survey_id)
	if output.Error != nil {
		return output.Error
	}
	if !ValidateSurveyId(c, survey_id) {
		return errors.New("survey Id not found")
	} else {
		var d model.Data
		output = s.Repo.DB.First(&d, form_id)
		if output.Error != nil {
			return output.Error
		}
		d.FormData = formData
		output = s.Repo.DB.Save(&d)
		if output.Error != nil {
			return output.Error
		}
	}
	return nil
}

func (s *FormService) SaveTemplate(input string) error {
	var f model.FormTemplate
	f.FormTemplate = input
	output := s.Repo.DB.Save(&f)
	if output.Error != nil {
		return output.Error
	}
	return nil
}

func ValidateForm(survey_id, formid int) bool {
	return true
}

func ValidateSurveyId(c model.Client, survey_id int) bool {
	isFound := false
	for _, value := range c.Survey_ids {
		if value == survey_id {
			isFound = true
		}
	}
	return isFound
}

// func GetSurveyById()
