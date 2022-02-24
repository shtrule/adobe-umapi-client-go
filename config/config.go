package config

//CreativeCloudConfig Adobe CC api config params
type CreativeCloudConfig struct {
	TechnicalAccountID string `json:"TechnicalAccountId"`
	ClientID           string `json:"ClientId"`
	OrganizationID     string `json:"OrganizationId"`
	ClientSecret       string `json:"ClientSecret"`
}
