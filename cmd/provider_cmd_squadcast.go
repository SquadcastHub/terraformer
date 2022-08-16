package cmd

import (
	"os"

	squadcast_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/squadcast"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

const (
	defaultSquadcastEndpoint = "http://api.squadcast.com/v3/"
)

func newCmdSquadcastImporter(options ImportOptions) *cobra.Command {
	// var accessToken string
	// accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjVkZTRkYmIyZTdiMzMwMDAxMTEzMTRjOSIsImZpcnN0TmFtZSI6IlJvc2hhbiIsImVtYWlsIjoicm9zaGFuQHNxdWFkY2FzdC5jb20iLCJ3ZWJfdG9rZW4iOnRydWUsImFwaV90b2tlbiI6ZmFsc2UsImlhdCI6MTY1NTkwODg3MiwiZXhwIjoxNjcxNzIwMDcyLCJpc3MiOiJhcGkuc3F1YWRjYXN0LmNvbSIsImp0aSI6ImMwMzFiMmNmZTNmYjlkMDE1NzNmZDQ4YmM1ODkxNTc0OWRjMGRjNTM0ZjBjMTMxNDI3MTE1ZTM4OTBkOGU2MzYifQ.37DLXLlkfLIu--LqDs3agrNTrN77H9oF6yr6dK6Pic0"
	cmd := &cobra.Command{
		Use:   "squadcast",
		Short: "Import current state to Terraform configuration from SquadCast",
		Long:  "Import current state to Terraform configuration from SquadCast",
		RunE: func(cmd *cobra.Command, args []string) error {
			accessToken := os.Getenv("SQUADCAST_ACCESS_TOKEN")
			endpoint := os.Getenv("SQUADCAST_SERVER_URL")
			if len(endpoint) == 0 {
				endpoint = defaultSquadcastEndpoint
			}
			provider := newSquadcastProvider()
			err := Import(provider, options, []string{accessToken})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newSquadcastProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "user", "")
	// cmd.PersistentFlags().StringVarP(&accessToken, "accessToken", "", "", "YOUR_SQUADCAST_ACCESS_TOKEN or env param SQUADCAST_ACCESS_TOKEN")
	return cmd
}

func newSquadcastProvider() terraformutils.ProviderGenerator {
	return &squadcast_terraforming.SquadcastProvider{}
}
