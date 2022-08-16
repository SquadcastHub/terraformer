package squadcast

import (
	"encoding/json"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type UserGenerator struct {
	SquadcastService
}

type OncallReminderRule struct {
	Type         string `json:"type" tf:"type"`
	DelayMinutes int    `json:"time" tf:"delay_minutes"`
}

type Ability struct {
	ID string `json:"id" tf:"-"`
	// Name    string `json:"name" tf:"-"`
	Slug string `json:"slug" tf:"-"`
	// Default bool   `json:"default" tf:"-"`
}

type Contact struct {
	DialCode    string `json:"dial_code" tf:"-"`
	PhoneNumber string `json:"phone_number" tf:"-"`
}

type PersonalNotificationRule struct {
	Type         string `json:"type" tf:"type"`
	DelayMinutes int    `json:"time" tf:"delay_minutes"`
}

type User struct {
	// Name string `json:"name"`
	AbilitiesSlugs            []string                    `json:"-" tf:"abilities"`
	Name                      string                      `json:"-" tf:"name"`
	PhoneNumber               string                      `json:"-" tf:"phone"`
	ID                        string                      `json:"id" tf:"id"`
	Abilities                 []*Ability                  `json:"abilities" tf:"-"`
	Bio                       string                      `json:"bio" tf:"-"`
	Contact                   Contact                     `json:"contact" tf:"-"`
	Email                     string                      `json:"email" tf:"email"`
	FirstName                 string                      `json:"first_name" tf:"first_name"`
	IsEmailVerified           bool                        `json:"email_verified" tf:"is_email_verified"`
	IsInGracePeriod           bool                        `json:"in_grace_period" tf:"-"`
	IsOverrideDnDEnabled      bool                        `json:"is_override_dnd_enabled" tf:"is_override_dnd_enabled"`
	IsPhoneVerified           bool                        `json:"phone_verified" tf:"is_phone_verified"`
	IsTrialSignup             bool                        `json:"is_trial_signup" tf:"-"`
	LastName                  string                      `json:"last_name" tf:"last_name"`
	OncallReminderRules       []*OncallReminderRule       `json:"oncall_reminder_rules" tf:"-"`
	PersonalNotificationRules []*PersonalNotificationRule `json:"notification_rules" tf:"-"`
	Role                      string                      `json:"role" tf:"role"`
	TimeZone                  string                      `json:"time_zone" tf:"time_zone"`
	Title                     string                      `json:"title" tf:"-"`
}

var response struct {
	Data *[]User `json:"data"`
	// *Meta
}

type Users []User

var UserAllowEmptyValues = []string{}

func (g UserGenerator) createResources(users Users) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, user := range users {
		fmt.Println(user.ID,
			"user_"+(user.ID),
			"squadcast_user",
			"squadcast")
		resources = append(resources, terraformutils.NewSimpleResource(
			user.ID,
			"user_"+(user.ID),
			"squadcast_user",
			"squadcast",
			[]string{},
		))
	}
	return resources
}

func (g *UserGenerator) InitResources() error {
	body, err := g.generateRequest("https://api.squadcast.com/v3/organizations/5a4262733c36823b3ed91fb9/users")
	if err != nil {
		return err
	}
	// var users Users
	// Unmarshal the body to the response struct
	err = json.Unmarshal(body, &response)
	// err = json.Unmarshal(body, &users)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(*response.Data)
	return nil
}
