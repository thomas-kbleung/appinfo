package appinfo

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo"
)

// Api handles /app_info restful api.
type Api struct {
	config ApiConfig
}

// ApiConfig stores dependancies passed from New function.
type ApiConfig struct {
	Repo Repository
}


// NewApi creates new Api handler.
func NewApi(config ApiConfig) *Api {
	return &Api{config: config}
}

// GetInfo returns application info for a specific system
func (api *Api) GetInfo(c echo.Context) (reterr error) {

	// retrieve parameters from url
	system := c.Param("system")
	if system == "" {
		logrus.Warn("system parameter cannot be blank.")
		return c.NoContent(http.StatusBadRequest)
	}

	// get information from repository
	a, err := api.config.Repo.Get(system)
	if err != nil {
		logrus.Error("Retrieving app info failured")
		return c.NoContent(http.StatusInternalServerError)
	}

	type Response struct {
		AppInfo           *AppInfo `json:"app_info" `
	}

	response := &Response{
		AppInfo:            a,
	}

	// allow cache for 4 hours
	c.Response().Header().Add("Cache-Control", "max-age=14400,public")
	return c.JSON(http.StatusOK, response)
}
