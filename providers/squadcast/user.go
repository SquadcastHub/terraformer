package squadcast

import (
	"encoding/json"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type UserGenerator struct {
	SquadcastService
}

type User struct {
	Name string `json:"name"`
}

type Users []User

var UserAllowEmptyValues = []string{}

func (g UserGenerator) createResources(users Users) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range users {
		resources = append(resources, terraformutils.NewSimpleResource(
			user.Name,
			"user_"+(user.Name),
			"rabbitmq_user",
			"rabbitmq",
			UserAllowEmptyValues,
		))
	}
	return resources
}

func (g *UserGenerator) InitResources() error {
	body, err := g.generateRequest("https://api.squadcast.com/v3/users")
	if err != nil {
		return err
	}
	var users Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(users)
	return nil
}
