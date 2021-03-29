package campaign

type Service interface {
	FindCampaigns(user_id int) ([]Campaign, error)
	FindCampaign(input GetCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {

	return &service{repository}
}

func (s *service) FindCampaigns(user_id int) ([]Campaign, error) {

	if user_id != 0 {
		campaigns, err := s.repository.FindByUserId(user_id)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) FindCampaign(input GetCampaignInput) (Campaign, error) {

	campaigns, err := s.repository.FindById(input.Id)
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
