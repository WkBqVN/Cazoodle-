package model

type FormTemplate struct {
	ID           int    `json:"formId" gorm:"primaryKey"`
	FormTemplate string `gorm:"form_template"`
}

type Data struct {
	ID          int         `json:"Id" gorm:"primaryKey"`
	FormData    interface{} `json:"data" gorm:"form_data"`
	FormId      int         `json:"formId" gorm:"form_id"`
	FormTempate int         `josn:"formTemplate" gorm:"form_template_id"`
}
