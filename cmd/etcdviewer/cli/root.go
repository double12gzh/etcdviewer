// Package cli
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package cli

import (
	"flag"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"

	"github.com/double12gzh/etcdviewer/options"
)

var etcdViewerOptions = options.NewCmdOptions()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "etcdviewer",
	Short:         "k8s etcd data viewer",
	Long:          `make it easy to get or watch etcd data stored by kubernetes based on etcd API V3`,
	SilenceErrors: true,
	SilenceUsage:  true,
	PreRun: func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(cmd.Flags())
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
	Example: helper(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func helper() string {
	result := `
	Get all data
	$ etcdviewer list /

	Get data from given path
	$ etcdviewer get /resitry/xxx/xxx

	Watch data from given path
	$ etcdviewer watch /resitry/xxx/xxx
`

	return result
}

func init() {
	klog.InitFlags(flag.CommandLine)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	rootCmd.Flags().StringSliceVar(&etcdViewerOptions.Inclusions, options.FlagInclude, []string{"Deployments", "Pods"}, "Filter by resource type (plural form or short name).")
	rootCmd.Flags().StringSliceVar(&etcdViewerOptions.Exclusions, options.FlagExclude, []string{}, "Filter by resource type (plural form or short name).")

	etcdViewerOptions.GenericCliOptConfigFlags.AddFlags(rootCmd.Flags())
	etcdViewerOptions.PrintFlags.AddFlags(rootCmd)

	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		klog.Errorf("failed bind flags: %s", err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// read in environment variables that match
	viper.SetEnvPrefix("etcdviewer")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
}
