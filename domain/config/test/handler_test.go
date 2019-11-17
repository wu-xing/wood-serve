package test

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/wu-xing/wood-serve/domain/config"
	"github.com/wu-xing/wood-serve/domain/config/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	configJson = `{"noteTypeConfigs":[{"name":"Markdown","canExportPdf":false}]}
`
)

func TestHandler_GetAppConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockRepository(ctrl)

	m.
		EXPECT().
		GetConfig().
		Return(&config.AppConfig{
			AppConfigData: &config.AppConfigData{
				NoteTypeConfigs: []config.NoteTypeProp{
					{
						Name:         "Markdown",
						CanExportPdf: false,
					},
				},
			},
		})

	handler := &config.Handler{Repo: m}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.GetAppConfig(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, configJson, rec.Body.String())
	}
}
