package resoponse

import "github.com/gin-gonic/gin"

type AppHttpResponse interface {
	SuccessResp(httpCode int, successMessage *SuccessMessage)
	FailedResp(httpCode int, failedMessage *FailedMessage)
}

type JsonResponse struct {
	ctx *gin.Context
}

func (j *JsonResponse) SuccessResp(httpCode int, successMessage *SuccessMessage) {
	successMessage.HttpResponse = httpCode
	j.ctx.JSON(httpCode, successMessage)
	j.ctx.Abort()
}

func (j JsonResponse) FailedResp(httpCode int, failedMessage *FailedMessage) {
	failedMessage.HttpResponse = httpCode
	j.ctx.JSON(httpCode, failedMessage)
	j.ctx.Abort()
}

func NewAppHttpResponse(ctx *gin.Context) AppHttpResponse {
	return &JsonResponse{
		ctx: ctx,
	}
}
