package clients

import (
	"encoding/json"
	"fmt"

	models "github.com/shtrule/adobe-umapi-client-go/models"
)

type AdobeClient interface {
	GetUsers() ([]models.User, error)
	GetUser(userEmail string) (models.User, error)
	GetGroups() ([]models.Group, error)
	GetUsersByGroup(groupName string) ([]models.User, error)
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
func (client UmapiClient) GetGroups() ([]models.Group, error) {
	var (
		pageNumber = 0
		groups     []models.Group
		groupRoot  models.GroupRoot
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

func (client UmapiClient) GetUsers() ([]models.User, error) {
	var (
		pageNumber = 0
		users      []models.User
		userRoot   models.UsersRoot
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

func (client UmapiClient) GetUser(userEmail string) (models.User, error) {
	var userRoot models.UserRoot

	url := fmt.Sprintf("https://usermanagement.adobe.io/v2/usermanagement/organizations/%v/users/%v", client.OrganizationId, userEmail)

	responseBody, err := GetRequest(url, client.Token, client.ClientId)

	if err != nil {
		return models.User{}, err
	}

	if err := json.Unmarshal(responseBody, &userRoot); err != nil {
		return models.User{}, fmt.Errorf("unable to convert response body to Users, error: %s", err.Error())
	}

	return userRoot.User, nil
}

func (client UmapiClient) GetUsersByGroup(groupName string) ([]models.User, error) {
	var (
		pageNumber = 0
		users      []models.User
		userRoot   models.UsersRoot
	)

	for {
		url := fmt.Sprintf("https://usermanagement.adobe.io/v2/usermanagement/users/%v/%v/%v", client.OrganizationId, pageNumber, groupName)

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
