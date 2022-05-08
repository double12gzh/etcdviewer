// Package cli
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package cli

import (
	"github.com/spf13/cobra"

	"github.com/double12gzh/etcdviewer/options"
	v3 "github.com/double12gzh/etcdviewer/pkg/etcdviewer/v3"
)

// dumpCmd represents the list command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump data",
	Long:  `dump the k8s data stored in etcd`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		isPretty, _ := cmd.Flags().GetBool(options.PRETTY)

		setDefault(etcdViewerOptions)

		v3.NewClientV3(etcdViewerOptions.CAFile, etcdViewerOptions.CertFile, etcdViewerOptions.KeyFile, etcdViewerOptions.Endpoints).Dump("/", isPretty)
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)

	// define local flags for command list
	dumpCmd.Flags().BoolVarP(&pretty, options.PRETTY, "p", true, "print the formatted value.")
}
