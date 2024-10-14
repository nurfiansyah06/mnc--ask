package transfercont

import (
	"golang-test/controller"
	"golang-test/model/transfer"
	transferModel "golang-test/model/transfer"
	"golang-test/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferControllerImpl struct {
	transferUsecase usecase.TransferUsecase
}

func NewTransferController(transferUsecase usecase.TransferUsecase) controller.TransferController {
	return &TransferControllerImpl{
		transferUsecase: transferUsecase,
	}
}

func (t *TransferControllerImpl) Transfer(c *gin.Context) {
	var transferRequest transfer.TransferRequest
	if err := c.ShouldBindJSON(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userID := c.MustGet("user_id").(string)

	transfer, err := t.transferUsecase.Transfer(userID, transferRequest.TargetUser, transferRequest.Remarks, transferRequest.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := transferModel.TransferResponse{
		Status: "SUCCESS",
		Result: transferModel.TransferDetails{
			TransferID:    transfer.TransferID,
			Amount:        transfer.Amount,
			Remarks:       transfer.Remarks,
			BalanceBefore: transfer.BalanceBefore,
			BalanceAfter:  transfer.BalanceAfter,
			CreatedDate:   transfer.CreatedDate,
		},
	}

	c.JSON(http.StatusOK, response)
}
