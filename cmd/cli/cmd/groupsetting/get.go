package groupsetting

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/theopenlane/core/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get group settings",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	command.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "group setting id to retrieve")
}

// get group settings from the platform
func get(ctx context.Context) error {
	// setup http client
	client, err := cmd.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer cmd.StoreSessionCookies(client)

	// filter options
	id := cmd.Config.String("id")

	// if setting ID is not provided, get settings which will automatically filter by group id
	if id == "" {
		o, err := client.GetAllGroupSettings(ctx)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	o, err := client.GetGroupSettingByID(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(o)
}