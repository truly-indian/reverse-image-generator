package reverseimagegenerator

import (
	"encoding/json"
	"errors"
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

	ctx.Header("Content-Type", "application/json")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Transfer-Encoding", "chunked")

	flusher, ok := ctx.Writer.(http.Flusher)

	if !ok {
		h.logger.LogError("error while getting flusher type assertion", nil)
		ctx.JSON(http.StatusInternalServerError, buildErrorResponse(types.InternalServerError(errors.New("flusher not initialised"))))
		return
	}

	productBatchChan := make(chan []types.Product)
	errorChan := make(chan error)

	go func() {
		err := h.service.ReverseImageGenerator(reverseImageGenerator, productBatchChan)
		if err != nil {
			errorChan <- err
		}
		close(productBatchChan)
	}()

	for {
		select {
		case batch, ok := <-productBatchChan:
			if !ok {
				return
			}

			if len(batch) > 0 {
				batchData, err := json.Marshal(batch)

				if err != nil {
					h.logger.LogError("error while marshalling batch data: ", err)
					ctx.JSON(http.StatusInternalServerError, buildErrorResponse(types.InternalServerError(err)))
					return
				}

				_, writeErr := ctx.Writer.Write(append(batchData, '\n'))
				if writeErr != nil {
					h.logger.LogError("error while writing batch to response: ", writeErr)
					ctx.JSON(http.StatusInternalServerError, buildErrorResponse(types.InternalServerError(writeErr)))
					return
				}

				flusher.Flush()
			}
		case err := <-errorChan:
			h.logger.LogError("error while generating reverse image: ", err)
			ctx.JSON(http.StatusInternalServerError, buildErrorResponse(types.InternalServerError(err)))
			return
		}
	}
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
