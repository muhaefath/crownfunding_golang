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

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {

	currentUser := c.MustGet("current_user").(user.User)
	userId := currentUser.Id

	transactions, err := h.service.GetTransactionByUserId(userId)
	if err != nil {

		response := helper.APIRespone("Failed to get user transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespone("Succsess to get user transaction", http.StatusOK, "succsess", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {

	var input transaction.CreateTransaction

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIRespone("Create transaction Failed input", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIRespone("Create transaction Failed", http.StatusUnprocessableEntity, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespone("Create transaction Succsess", http.StatusOK, "succsess", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}
