// Package cli
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package cli

import (
	"github.com/spf13/cobra"

	"github.com/double12gzh/etcdviewer/options"
	v3 "github.com/double12gzh/etcdviewer/pkg/etcdviewer/v3"
)

// getCmd represents the list command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get data",
	Long:  `get the k8s data stored in etcd`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		isPretty, _ := cmd.Flags().GetBool(options.PRETTY)
		key := args[0]

		setDefault(etcdViewerOptions)

		v3.NewClientV3(etcdViewerOptions.CAFile, etcdViewerOptions.CertFile, etcdViewerOptions.KeyFile, etcdViewerOptions.Endpoints).Get(key, isPretty)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// define local flags for command list
	getCmd.Flags().BoolVarP(&pretty, options.PRETTY, "p", true, "print the formatted value.")
}
