package commands

import (
	"github.com/spf13/cobra"
	"github.com/tliron/kutil/util"
)

var site string
var registry string

func init() {
	rootCommand.AddCommand(installCommand)
	installCommand.Flags().StringVarP(&site, "site", "s", "default", "site name")
	installCommand.Flags().BoolVarP(&cluster, "cluster", "c", false, "cluster mode")
	installCommand.Flags().StringVarP(&registry, "registry", "g", "docker.io", "registry URL (use special value \"internal\" to discover internally deployed registry)")
	installCommand.Flags().BoolVarP(&wait, "wait", "w", false, "wait for installation to succeed")
}

var installCommand = &cobra.Command{
	Use:   "install",
	Short: "Install Turandot",
	Run: func(cmd *cobra.Command, args []string) {
		err := NewClient().Turandot().Install(site, registry, wait)
		util.FailOnError(err)
	},
}
