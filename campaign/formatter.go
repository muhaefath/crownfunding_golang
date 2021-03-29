package campaign

import "strings"

type CampaignFormatter struct {
	Id               int    `json:id`
	UserId           int    `json:user_id`
	Name             string `json:name`
	ShortDescription string `json:short_description`
	Description      string `json:description`
	ImageUrl         string `json:image_url`
	GoalAmmount      int    `json:goal_ammount`
	CurrentAmount    int    `json:current_ammount`
	Slug             string `json:slug`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Id = campaign.Id
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.Description = campaign.Description
	campaignFormatter.GoalAmmount = campaign.GoalAmmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {

		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaingDetailFormatter struct {
	Id               int                      `json:id`
	UserId           int                      `json:user_id`
	Name             string                   `json:name`
	ShortDescription string                   `json:short_description`
	Description      string                   `json:description`
	ImageUrl         string                   `json:image_url`
	GoalAmmount      int                      `json:goal_ammount`
	CurrentAmount    int                      `json:current_ammount`
	Slug             string                   `json:slug`
	Perks            []string                 `json:perks`
	User             CampaignUserFormatter    `json:user`
	Images           []CampaignImageFormatter `json:images`
}

type CampaignUserFormatter struct {
	Name     string `json:name`
	ImageUrl string `json:image_url`
}

type CampaignImageFormatter struct {
	ImageUrl  string `json:image_url`
	IsPrimary bool   `json:is_primary`
}

func FormatDetailCampaign(campaign Campaign) CampaingDetailFormatter {

	campaingDetailFormatter := CampaingDetailFormatter{}
	campaingDetailFormatter.Id = campaign.Id
	campaingDetailFormatter.UserId = campaign.UserId
	campaingDetailFormatter.Name = campaign.Name
	campaingDetailFormatter.ShortDescription = campaign.ShortDescription
	campaingDetailFormatter.Description = campaign.Description
	campaingDetailFormatter.GoalAmmount = campaign.GoalAmmount
	campaingDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaingDetailFormatter.Slug = campaign.Slug
	campaingDetailFormatter.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		campaingDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, perk)
	}
	campaingDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.Avatar

	campaingDetailFormatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageUrl = image.FileName
		is_primary := false
		if image.IsPrimary == 1 {
			is_primary = true
		}
		campaignImageFormatter.IsPrimary = is_primary

		images = append(images, campaignImageFormatter)
	}

	campaingDetailFormatter.Images = images

	return campaingDetailFormatter
}
