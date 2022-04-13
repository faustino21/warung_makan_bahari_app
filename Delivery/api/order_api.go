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

type OrderApi struct {
	api.BaseApi
	order usecase.OrderUseCase
}

func (o *OrderApi) Ordering() gin.HandlerFunc {
	funcName := "api.orderApi.Ordering"

	return func(c *gin.Context) {
		var orderReq httpReq.OrderReq
		err := o.ParseRequestBody(c, &orderReq)
		if err != nil {
			log.Error().Msgf(funcName+" : %w", err)
			o.ParsingError(c, err)
			return
		}

		result, err := o.order.OrderProduct(orderReq)
		if err != nil {
			resoponse.NewAppHttpResponse(c).FailedResp(http.StatusBadRequest, resoponse.NewFailedMessage(err.Error()))
			return
		}
		resoponse.NewAppHttpResponse(c).SuccessResp(http.StatusOK, resoponse.NewSuccessMessage("Order Success", result))
	}
}

func OrderApiRoute(orderRoute *gin.RouterGroup, order usecase.OrderUseCase) *OrderApi {
	orderApi := OrderApi{
		order: order,
	}

	orderRoute.POST("/Order", orderApi.Ordering())
	return &orderApi
}
