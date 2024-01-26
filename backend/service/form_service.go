package service

import "cazoodle.com/utils"

type Record struct {
	Title string
	Value interface{}
}

func GetFormData(survey_id, form_id int) (interface{}, error) {
	form := make(map[string]Record)
	form["email"] = Record{
		Title: "hello",
		Value: "myemail",
	}
	form["age"] = Record{
		Title: "this is my age",
		Value: 20,
	}
	result, err := utils.ConvertMapToStruct(form)
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
