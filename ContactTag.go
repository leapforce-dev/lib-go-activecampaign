package activecampaign

type ContactTag struct {
	Contact          string `json:"contact"`
	Tag              string `json:"tag"`
	CreatedDate      string `json:"cdate"`
	CreatedTimestamp string `json:"created_timestamp"`
	UpdatedTimestamp string `json:"updated_timestamp"`
	CreatedBy        string `json:"created_by"`
	UpdatedBy        string `json:"updated_by"`
	ID               string `json:"id"`
	Links            *Links `json:"links"`
}
