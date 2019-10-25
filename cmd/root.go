package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "kube-rabbit",
		Short: "kube-rabbit a simple rabbitmq cluster within kubernetes",
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("kube-rabbit root command prerun func")
		},
	}
	c.AddCommand(startCmd())
	return c
}
