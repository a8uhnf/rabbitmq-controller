package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func startCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start kubernetes controller for kube-rabbit",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("starting rabbitmq controller...")
		},
	}
}
