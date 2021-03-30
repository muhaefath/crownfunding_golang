package campaign

import "golang_project/user"

type GetCampaignInput struct {
	Id int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmmount      int    `json:"goal_ammount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CreateCampaignImageInput struct {
	CampaignId int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}
