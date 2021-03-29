package campaign

import "time"

type Campaign struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	BackerCount      int
	GoalAmmount      int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdateAt         time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	Id         int
	CampaignId int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdateAt   time.Time
}
