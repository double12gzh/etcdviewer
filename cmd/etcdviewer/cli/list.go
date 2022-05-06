// Package cli
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */
package cli

import (
	"github.com/spf13/cobra"

	"github.com/double12gzh/etcdviewer/pkg/etcdviewer/v3"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list data",
	Long:  `list the k8s data stored in etcd.`,
	Run: func(cmd *cobra.Command, args []string) {
		v3.List(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("list", "l", "/registry", "list all data stored in etcd by kubernetes")
}
