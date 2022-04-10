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

type TableApi struct {
	api.BaseApi
	table usecase.TableUseCase
}

func (t *TableApi) GetAllTable() gin.HandlerFunc {
	funcName := "api.tableApi.GetAllTable"

	return func(c *gin.Context) {
		var page httpReq.PageReq
		err := t.ParseRequestBody(c, &page)
		if err != nil {
			log.Error().Msgf(funcName+" : %w", err)
			t.ParsingError(c, err)
			return
		}

		list, err := t.table.GetAllTable(page.PageReq)
		if err != nil {
			resoponse.NewAppHttpResponse(c).FailedResp(http.StatusInternalServerError, resoponse.NewFailedMessage(err.Error()))
			return
		}
		resoponse.NewAppHttpResponse(c).SuccessResp(http.StatusOK, resoponse.NewSuccessMessage("Get Table Success", list))
	}
}

func TableApiRoute(tableRoute *gin.RouterGroup, table usecase.TableUseCase) *TableApi {
	tableApi := TableApi{
		table: table,
	}
	tableGroup := tableRoute.Group("/Table")
	tableGroup.GET("/getAllTable", tableApi.GetAllTable())
	return &tableApi
}
