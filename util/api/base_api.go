package api

import (
	"Revisi_WBH/util/resoponse"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type BaseApi struct {
}

func (b *BaseApi) ParseRequestBody(c *gin.Context, body interface{}) error {
	err := c.ShouldBindJSON(body)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) ParsingError(c *gin.Context, err error) {
	log.Error().Msg("Parsing Error")
	resoponse.NewAppHttpResponse(c).FailedResp(http.StatusInternalServerError, resoponse.NewFailedMessage(err.Error()))
}
