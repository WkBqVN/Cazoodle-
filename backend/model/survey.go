package model

type Survey struct {
	Email    string
	ID       int
	FilePath []string
	Data     []byte
}

type SurveyReponse struct {
	Message interface{}
}
