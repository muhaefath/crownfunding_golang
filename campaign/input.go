package campaign

type GetCampaignInput struct {
	Id int `uri:"id" binding:"required"`
}
