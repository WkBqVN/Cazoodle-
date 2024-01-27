package model

type Survey struct {
	ID      int
	FormsID string
}

type Forms struct {
	ID       int    `json:"formId" gorm:"primaryKey"`
	FormData string `gorm:"FormData"`
}

type client struct{}

type SurveyReponse struct {
	Message interface{}
}
