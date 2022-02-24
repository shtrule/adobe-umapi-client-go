package models

//UserRoot Adobe CC portal users
type UsersRoot struct {
	LastPage bool   `json:"lastPage"`
	Result   string `json:"result"`
	Users    []User `json:"users"`
}

type UserRoot struct {
	Result string `json:"result"`
	User   User   `json:"user"`
}

//User Adobe CC user type
type User struct {
	Id        string   `json:"id"`
	Email     string   `json:"email"`
	Status    string   `json:"status"`   //one of: { "active", "disabled", "locked", "removed" }. currently returns only active users
	Username  string   `json:"username"` //applicable for Enterprise and Federated users. For most AdobeID users, this value is the same as the email address
	Domain    string   `json:"domain"`
	Country   string   `json:"country"`
	Type      string   `json:"type"` //one of: { "adobeID", "enterpriseID", "federatedID", "unknown" }
	Groups    []string `json:"groups,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
}
