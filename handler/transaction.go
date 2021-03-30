package handler

import (
	"golang_project/helper"
	"golang_project/transaction"
	"golang_project/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {

	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {

	var input transaction.GetTransactionDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {

		response := helper.APIRespone("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {

		response := helper.APIRespone("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespone("Succsess to get campaign transaction", http.StatusOK, "succsess", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
