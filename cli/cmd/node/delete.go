package node

import (
	"github.com/gravitl/netmaker/cli/functions"
	"github.com/spf13/cobra"
)

var force bool

var nodeDeleteCmd = &cobra.Command{
	Use:   "delete [NETWORK NAME] [NODE ID]",
	Args:  cobra.ExactArgs(2),
	Short: "Delete a Node",
	Long:  `Delete a Node`,
	Run: func(cmd *cobra.Command, args []string) {
		functions.PrettyPrint(functions.DeleteNode(args[0], args[1], force))
	},
}

func init() {
	rootCmd.AddCommand(nodeDeleteCmd)
	nodeDeleteCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force delete a node")
}
