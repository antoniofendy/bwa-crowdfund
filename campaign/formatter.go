package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserID           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

type DetailCampaignUser struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type DetailCampaignImage struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type DetailCampaignFormatter struct {
	ID               int                   `json:"id"`
	Name             string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	ImageURL         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	UserID           int                   `json:"user_id"`
	Slug             string                `json:"slug"`
	Description      string                `json:"description"`
	User             DetailCampaignUser    `json:"user"`
	Perks            []string              `json:"perks"`
	Images           []DetailCampaignImage `json:"images"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignsFormatter = append(campaignsFormatter, FormatCampaign(campaign))
	}

	return campaignsFormatter
}

func FormatDetailCampaign(campaign Campaign) DetailCampaignFormatter {
	detailCampaignFormatter := DetailCampaignFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		Description:      campaign.Description,
		User: DetailCampaignUser{
			Name:      campaign.User.Name,
			AvatarURL: campaign.User.AvatarFileName,
		},
		Perks:  []string{""},
		Images: []DetailCampaignImage{},
	}

	if len(campaign.CampaignImages) > 0 {
		detailCampaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	// assign perks data
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	detailCampaignFormatter.Perks = perks

	// assign images data
	for _, campaignImage := range campaign.CampaignImages {

		isPrimary := false

		if campaignImage.IsPrimary == 1 {
			isPrimary = true
		}

		newCampaignImage := DetailCampaignImage{
			ImageURL:  campaignImage.FileName,
			IsPrimary: isPrimary,
		}
		detailCampaignFormatter.Images = append(detailCampaignFormatter.Images, newCampaignImage)
	}

	return detailCampaignFormatter
}
