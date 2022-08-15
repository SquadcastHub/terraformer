package squadcast

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type SquadcastProvider struct {
	terraformutils.Provider
	APIkey string
}

func (p *SquadcastProvider) Init(args []string) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if apiKey := os.Getenv("SQUADCAST_API_KEY"); apiKey != "" {
		p.APIkey = os.Getenv("SQUADCAST_API_KEY")
	}
	if args[0] != "" {
		p.APIkey = args[0]
	}
	if p.APIkey == "" {
		return errors.New("requred API Key missing")
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
		"api-key": p.APIkey,
	})

	return nil
}

func (p *SquadcastProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_key": cty.StringVal(p.APIkey),
	})
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
