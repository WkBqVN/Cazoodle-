package api

import (
	"io"
	"net/http"

	"cazoodle.com/model"
	"cazoodle.com/service"
	"cazoodle.com/utils"
	"github.com/labstack/echo/v4"
)

func (a *API) GetFormById(c echo.Context) error {
	form_id, err := utils.ConvertStringToInt(c.Param("form_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
			Message: err,
		})
	}
	data, err := a.FormSerivce.GetFormTemplate(form_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.SurveyReponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &model.SurveyReponse{
		Message: data,
	})
}

func (a *API) SaveFormTemplate(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	err = a.FormSerivce.SaveTemplate(string(body))
	if err != nil {
		return err
	}
	return nil
}

func (a *API) PostForms(c echo.Context) error {
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
	client_id, err := utils.ConvertStringToInt(c.Param("client_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
			Message: err,
		})
	}
	// var data []map[string]interface{}
	if service.ValidateForm(survey_id, form_id) {
		// return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
		// 	Message: "Bad request",
		// })
	}
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
			Message: "body" + err.Error(),
		})
	}
	err = a.FormSerivce.SaveData(client_id, survey_id, form_id, string(body))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.SurveyReponse{
			Message: err,
		})
	}
	return c.JSON(http.StatusAccepted, &model.SurveyReponse{
		Message: string(body),
	})
}
