package handler

import (
	"fmt"
	"golang_project/campaign"
	"golang_project/helper"
	"golang_project/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {

	return &campaignHandler{service}
}

func (h *campaignHandler) FindCampaigns(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.FindCampaigns(userId)

	if err != nil {

		response := helper.APIRespone("FindCampaigns Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespone("FindCampaigns registered", http.StatusOK, "succsess", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) FindCampaign(c *gin.Context) {

	var input campaign.GetCampaignInput

	err := c.ShouldBindUri(&input)
	if err != nil {

		response := helper.APIRespone("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.FindCampaign(input)
	if err != nil {

		response := helper.APIRespone("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespone("Succsess to get detail campaign", http.StatusOK, "succsess", campaign.FormatDetailCampaign(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {

	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIRespone("Create campaign Failed input", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIRespone("Create campaign Failed", http.StatusUnprocessableEntity, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespone("Create campaign Succsess", http.StatusOK, "succsess", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {

	var inputId campaign.GetCampaignInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {

		response := helper.APIRespone("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIRespone("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCampaign, err := h.service.UpdateCampaign(inputId, inputData)
	if err != nil {

		response := helper.APIRespone("Failed to update campaign", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespone("Update campaign Succsess", http.StatusOK, "succsess", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {

	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIRespone("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	input.User = currentUser
	userId := currentUser.Id

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIRespone("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/campaign/%d-%s", userId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIRespone("Failed to upload campaign image", http.StatusBadRequest, err.Error(), data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIRespone("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIRespone("Succsess upload campaign image", http.StatusOK, "succsess", data)

	c.JSON(http.StatusBadRequest, response)
}
