package api

import (
	"io"
	"net/http"

	"cazoodle.com/model"
	"cazoodle.com/service"
	"cazoodle.com/utils"
	"github.com/labstack/echo/v4"
)

func (a *API) GetSurvey(c echo.Context) error {
	survey_id, err := utils.ConvertStringToInt(c.Param("survey_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
			Message: err,
		})
	}
	form_id, err := utils.ConvertStringToInt(c.Param("form_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
			Message: err,
		})
	}
	data, err := a.FormSerivce.GetFormData(survey_id, form_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.SurveyReponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &model.SurveyReponse{
		Message: data,
	})
}

func (a *API) PostSurvey(c echo.Context) error {
	survey_id, err := utils.ConvertStringToInt(c.Param("survey_id"))
	if err != nil {
		return err
	}
	form_id, err := utils.ConvertStringToInt(c.Param("form_id"))
	if err != nil {
		return err
	}
	// var data []map[string]interface{}
	if service.ValidateForm(survey_id, form_id) {
	}
	abc, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	return c.String(http.StatusAccepted, string(abc))
	// x, err := utils.ConvertMapToStruct(c.Request().Body)
	// service.SaveData(survey_id, form_id, data)
	// return c.String(http.StatusOK, "Hello, World!")
}
