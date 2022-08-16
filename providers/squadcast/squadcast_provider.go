package squadcast

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type SquadcastProvider struct {
	terraformutils.Provider
	AccessToken string
}

func (p *SquadcastProvider) Init(args []string) error {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	if accessToken := os.Getenv("SQUADCAST_ACCESS_TOKEN"); accessToken != "" {
		p.AccessToken = os.Getenv("SQUADCAST_ACCESS_TOKEN")
	}
	// if args[0] != "" {
	// 	p.AccessToken = args[0]
	// }
	if p.AccessToken == "" {
		return errors.New("requred Access Token missing")
	}
	return nil
}

func (p *SquadcastProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"accessToken": p.AccessToken,
	})

	return nil
}

func (p *SquadcastProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"accessToken": cty.StringVal(p.AccessToken),
	})
}

func (p *SquadcastProvider) GetBasicConfig() cty.Value {
	return p.GetConfig()
}

func (p *SquadcastProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p *SquadcastProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *SquadcastProvider) GetName() string {
	return "squadcast"
}

func (p *SquadcastProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"user": &UserGenerator{},
	}
}
