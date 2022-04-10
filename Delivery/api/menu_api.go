package api

import (
	"Revisi_WBH/Delivery/httpReq"
	"Revisi_WBH/usecase"
	"Revisi_WBH/util/api"
	"Revisi_WBH/util/resoponse"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type MenuApi struct {
	api.BaseApi
	menu usecase.MenuUseCase
}

func (m *MenuApi) GetAllMenu() gin.HandlerFunc {
	funcName := "api.menuApi.GetAllMenu"

	return func(c *gin.Context) {
		var page httpReq.PageReq
		err := m.ParseRequestBody(c, &page)
		if err != nil {
			log.Error().Msgf(funcName+" : %w", err)
			m.ParsingError(c, err)
			return
		}

		list, err := m.menu.GetAllMenu(page.PageReq)
		if err != nil {
			resoponse.NewAppHttpResponse(c).FailedResp(http.StatusInternalServerError, resoponse.NewFailedMessage(err.Error()))
			return
		}
		resoponse.NewAppHttpResponse(c).SuccessResp(http.StatusOK, resoponse.NewSuccessMessage("success", list))
	}
}

func MenuApiRoute(menuRoute *gin.RouterGroup, menu usecase.MenuUseCase) *MenuApi {
	menuApi := MenuApi{
		menu: menu,
	}

	menuGroup := menuRoute.Group("/Menu")
	menuGroup.GET("/getAllMenu", menuApi.GetAllMenu())
	return &menuApi
}
