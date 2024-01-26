package api

import (
	"net/http"

	"cazoodle.com/model"
	"cazoodle.com/service"
	"cazoodle.com/utils"
	"github.com/labstack/echo/v4"
)

func GetSurvey(c echo.Context) error {
	survey_id, err := utils.ConvertStringToInt(c.QueryParam("survey_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
			Message: err,
		})
	}
	form_id, err := utils.ConvertStringToInt(c.QueryParam("form_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
			Message: err,
		})
	}
	data, err := service.GetFormData(survey_id, form_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.SurveyReponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &model.SurveyReponse{
		Message: data,
	})
}

func PostSurvey(c echo.Context) error {
	survey_id, err := utils.ConvertStringToInt(c.Param("survey_id"))
	if err != nil {
		return err
	}
	form_id, err := utils.ConvertStringToInt(c.Param("form_id"))
	if err != nil {
		return err
	}
	var data map[string]interface{}
	if service.ValidateForm(survey_id, form_id) {
		return nil
	}
	service.SaveData(survey_id, form_id, data)
	return c.String(http.StatusOK, "Hello, World!")
}
