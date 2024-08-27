package reverseimagegenerator

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/truly-indian/reverseImageSearch/internal/config"
	"github.com/truly-indian/reverseImageSearch/internal/logger"
	"github.com/truly-indian/reverseImageSearch/internal/types"
)

type Handler struct {
	config  *config.Config
	service Service
	logger  logger.Logger
}

func NewHandler(c *config.Config, s Service, logger logger.Logger) *Handler {
	return &Handler{
		config:  c,
		service: s,
		logger:  logger,
	}
}

func (h *Handler) HellowWorldHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello from reverse image generator")
}

func (h *Handler) ReverseImageGenerator(ctx *gin.Context) {
	var reverseImageGenerator types.ReverseImageGeneratorRequest
	if bindingErr := ctx.ShouldBindJSON(&reverseImageGenerator); bindingErr != nil {
		h.logger.LogError("error while binding request body for reverse image generator: ", bindingErr)
		ctx.JSON(http.StatusBadRequest, buildErrorResponse(types.BadRequestError(bindingErr)))
	}

	resp, err := h.service.ReverseImageGenerator(reverseImageGenerator)

	if err != nil {
		h.logger.LogError("error while generating reverse iamge: ", err)
		ctx.JSON(http.StatusInternalServerError, buildErrorResponse(types.InternalServerError(err)))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func buildErrorResponse(err *types.StatusError) types.ErrorResponse {
	return types.ErrorResponse{
		Error: types.Error{
			DisplayMessage: err.DisplayMessage,
			Message:        err.Message,
			Code:           fmt.Sprint(err.HTTPCode),
		},
	}
}
