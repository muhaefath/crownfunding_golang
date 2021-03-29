package handler

import (
	"golang_project/campaign"
	"golang_project/helper"
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
