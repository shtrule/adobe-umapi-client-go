package clients

import (
	"encoding/json"
	"fmt"

	adobeCcModels "github.com/shtrule/adobe-umapi-client-go/models"
)

type AdobeClient interface {
	GetUsers() ([]adobeCcModels.User, error)
	GetGroups() ([]adobeCcModels.Group, error)
}

type UmapiClient struct {
	ClientId       string
	OrganizationId string
	Token          string
}

func NewUmapiClient(clientId string, organizationId string, token string) *UmapiClient {
	return &UmapiClient{
		ClientId:       clientId,
		OrganizationId: organizationId,
		Token:          token,
	}
}

//TODO: Extract a GetRequest method to a class with builder pattern
func (client UmapiClient) GetGroups() ([]adobeCcModels.Group, error) {
	var (
		pageNumber = 0
		groups     []adobeCcModels.Group
		groupRoot  adobeCcModels.GroupRoot
	)

	for {
		url := fmt.Sprintf("https://usermanagement.adobe.io/v2/usermanagement/groups/%v/%v", client.OrganizationId, pageNumber)

		responseBody, err := GetRequest(url, client.Token, client.ClientId)

		if err != nil {
			return groups, err
		}

		if err := json.Unmarshal(responseBody, &groupRoot); err != nil {
			return groups, fmt.Errorf("failed to convert response body to Groups, error: %s", err.Error())
		}

		groups = append(groups, groupRoot.Groups...)

		if groupRoot.LastPage {
			break
		}

		pageNumber++
	}

	return groups, nil
}

func (client UmapiClient) GetUsers() ([]adobeCcModels.User, error) {
	var (
		pageNumber = 0
		users      []adobeCcModels.User
		userRoot   adobeCcModels.UserRoot
	)

	for {
		url := fmt.Sprintf("https://usermanagement.adobe.io/v2/usermanagement/users/%v/%v?directOnly=false", client.OrganizationId, pageNumber)

		responseBody, err := GetRequest(url, client.Token, client.ClientId)

		if err != nil {
			return users, err
		}

		if err := json.Unmarshal(responseBody, &userRoot); err != nil {
			return users, fmt.Errorf("unable to convert response body to Users, error: %s", err.Error())
		}

		users = append(users, userRoot.Users...)

		if userRoot.LastPage {
			break
		}

		pageNumber++
	}

	return users, nil
}
