// Package cli
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package cli

import (
	"github.com/spf13/cobra"

	"github.com/double12gzh/etcdviewer/options"
	v3 "github.com/double12gzh/etcdviewer/pkg/etcdviewer/v3"
)

var pretty bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list data",
	Long:  `list the k8s data stored in etcd`,
	Run: func(cmd *cobra.Command, args []string) {
		isPretty, _ := cmd.Flags().GetBool(options.PRETTY)
		key := "/"
		if len(args) >= 0 {
			key = args[0]
		}

		setDefault(etcdViewerOptions)

		v3.NewClientV3(etcdViewerOptions.CAFile, etcdViewerOptions.CertFile, etcdViewerOptions.KeyFile, etcdViewerOptions.Endpoints).ListKeys(key, isPretty)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// define local flags for command list
	listCmd.Flags().BoolVarP(&pretty, options.PRETTY, "p", true, "print the formatted value.")
}
