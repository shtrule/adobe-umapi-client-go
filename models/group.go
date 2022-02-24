package models

//GroupRoot Adobe CC portal user groups
type GroupRoot struct {
	LastPage bool    `json:"lastPage"`
	Result   string  `json:"result"` //TODO: Add error handling based on result
	Groups   []Group `json:"groups"`
}

//Group struct's GroupId, GroupName, MemberCount and Type should always be returned by the API.
//The other properties are included depending on the product type
type Group struct {
	Type               string `json:"type"`
	GroupID            int64  `json:"groupId"`
	GroupName          string `json:"groupName"`
	MemberCount        int    `json:"memberCount"`
	UserGroupName      string `json:"userGroupName,omitempty"`
	AdminGroupName     string `json:"adminGroupName,omitempty"`
	LicenseQuota       string `json:"licenseQuota,omitempty"`
	ProductName        string `json:"productName,omitempty"`
	ProductProfileName string `json:"productProfileName,omitempty"`
}
