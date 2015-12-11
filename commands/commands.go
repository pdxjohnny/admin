package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/admin/adduser"
)

// Commands are the commands available
var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "adduser",
		Short: "Start the adduser server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			adduser.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
