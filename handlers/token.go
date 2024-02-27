package handlers

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/models"
)

type TokenHandler struct{}

type TokenResponse struct {
	TokenMhs   string `json:"token_mhs"`
	TokenDosen string `json:"token_dosen"`
}

func NewTokenHandler(db *gorm.DB, logger *zap.Logger) *TokenHandler {
	return &TokenHandler{}
}

// @Summary		Get Token
// @Description	Get Token
// @Tags			Token
// @Accept			json
// @Produce		json
// @Success		200			{object}	models.BaseResponse[TokenResponse]
// @Router			/token [get]
func (h *TokenHandler) GetToken(c *gin.Context) {
	c.JSON(200, models.BaseResponse[TokenResponse]{
		Message: "success",
		Data: TokenResponse{
			TokenMhs:   os.Getenv("TOKEN_MHS"),
			TokenDosen: os.Getenv("TOKEN_DOSEN"),
		},
	})
}
