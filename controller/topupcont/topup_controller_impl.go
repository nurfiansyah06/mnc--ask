// package topupcont

// import (
// 	"golang-test/controller"
// 	"golang-test/model/topup"
// 	"golang-test/usecase"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func NewTopUpCont(topUpUsecase usecase.TopUpUsecase) controller.TopUpController {
// 	return &TopUpContImpl{topUpUsecase: topUpUsecase}
// }

// func (t *TopUpContImpl) TopUp(c *gin.Context) {
// 	var (
// 		request topup.TopUpRequest
// 	)

// 	// Bind JSON request body to struct
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	createdTopUp, err := t.topUpUsecase.TopUp(request)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	response := topup.TopUpResponse{
// 		Status: "SUCCESS",
// 		Result: struct {
// 			TopUpID       string    `json:"top_up_id"`
// 			AmountTopUp   int       `json:"amount_top_up"`
// 			BalanceBefore int       `json:"balance_before"`
// 			BalanceAfter  int       `json:"balance_after"`
// 			CreatedDate   time.Time `json:"created_date"`
// 		}{
// 			TopUpID:       createdTopUp.TopUpID,
// 			AmountTopUp:   createdTopUp.AmountTopUp,
// 			BalanceBefore: createdTopUp.BalanceBefore,
// 			BalanceAfter:  createdTopUp.BalanceAfter,
// 			CreatedDate:   createdTopUp.CreatedDate,
// 		},
// 	}

//		c.JSON(http.StatusOK, response)
//	}
package topupcont

import (
	"golang-test/controller"
	"golang-test/model/topup"
	"golang-test/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopUpContImpl struct {
	useCase usecase.AccountUsecase
}

func NewTopUpController(topUpusesase usecase.AccountUsecase) controller.TopupController {
	return &TopUpContImpl{useCase: topUpusesase}
}
func (t *TopUpContImpl) TopUp(c *gin.Context) {

	var (
		request topup.TopUpRequest
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestTopup := topup.TopUpRequest{
		Amount: request.Amount,
	}

	userID := c.MustGet("user_id").(string)

	createdTopUp, err := t.useCase.TopUp(requestTopup, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := topup.TopUpResponse{
		Status: "SUCCESS",
		Result: topup.TopUpResult{
			TopUpID:       createdTopUp.Result.TopUpID,
			AmountTopUp:   createdTopUp.Result.AmountTopUp,
			BalanceBefore: createdTopUp.Result.BalanceBefore,
			BalanceAfter:  createdTopUp.Result.BalanceAfter,
			CreatedDate:   createdTopUp.Result.CreatedDate,
		},
	}

	c.JSON(http.StatusOK, response)
}
