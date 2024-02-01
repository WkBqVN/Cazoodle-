package model

type Survey struct {
	ID          int
	FormID      int
	FormTempate int
}

type SurveyReponse struct {
	Message interface{} `json:"message"`
}
