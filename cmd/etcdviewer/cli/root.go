// Package cli
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package cli

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/double12gzh/etcdviewer/options"
	logger2 "github.com/double12gzh/etcdviewer/pkg/logger"
)

var logger = logger2.NewLogger()
var etcdViewerOptions = options.NewEtcdViewerOptions()

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
	Example: helper(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// global config for all commands
	rootCmd.PersistentFlags().StringSliceVar(&etcdViewerOptions.Endpoints, options.ENDPOINTS, []string{}, "etcd endpoints.")
	rootCmd.PersistentFlags().StringVar(&etcdViewerOptions.CAFile, options.CACERT, "", "etcd cert file.")
	rootCmd.PersistentFlags().StringVar(&etcdViewerOptions.CertFile, options.CERT, "", "etcd ca file.")
	rootCmd.PersistentFlags().StringVar(&etcdViewerOptions.KeyFile, options.KEY, "", "etcd key file.")

	_ = rootCmd.MarkFlagRequired(options.CACERT)
	_ = rootCmd.MarkFlagRequired(options.KEY)
	_ = rootCmd.MarkFlagRequired(options.CERT)
}

func helper() string {
	result := `
	Get all data
	$ etcdviewer list /

	Get data from given path
	$ etcdviewer get /resitry/xxx/xxx

	Watch data from given path
	$ etcdviewer watch /resitry/xxx/xxx

	Dump data
	$ etcdviewer dump
`
	return result
}

func setDefault(etcdViewerOptions *options.EtcdViewerOptions) {
	if etcdViewerOptions.KeyFile == "" {
		// get from environment variable: ETCD_VIEWER_KEY
		etcdViewerOptions.KeyFile = viper.GetString(options.EtcdViewerENVKey)
		if etcdViewerOptions.KeyFile == "" {
			etcdViewerOptions.KeyFile = options.DefaultKeyFile
		}
	}

	if etcdViewerOptions.CertFile == "" {
		// get from environment variable: ETCD_VIEWER_CERT
		etcdViewerOptions.CertFile = viper.GetString(options.EtcdViewerENVCert)
		if etcdViewerOptions.CertFile == "" {
			etcdViewerOptions.CertFile = options.DefaultCertFile
		}
	}

	if etcdViewerOptions.CAFile == "" {
		// get from environment variable: ETCD_VIEWER_CA_CERT
		etcdViewerOptions.CAFile = viper.GetString(options.EtcdViewerENVCaCert)
		if etcdViewerOptions.CAFile == "" {
			etcdViewerOptions.CAFile = options.DefaultCaCertFile
		}
	}

	if len(etcdViewerOptions.Endpoints) == 0 {
		// get from environment variable: ETCD_VIEWER_ENDPOINTS
		etcdViewerOptions.Endpoints = viper.GetStringSlice(options.EtcdViewerENVEndpoints)
		if len(etcdViewerOptions.Endpoints) == 0 {
			etcdViewerOptions.Endpoints = strings.Split(options.DefaultEndpoints, ",")
		}
	}
}

func initConfig() {
	// setup viper to get value from environment variables
	viper.SetEnvPrefix(options.EtcdViewerENVPrefix)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
}
