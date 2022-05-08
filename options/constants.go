// Package options
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package options

const (
	ENDPOINTS = "endpoints"
	CACERT    = "ca-cert"
	KEY       = "key"
	CERT      = "cert"
	PRETTY    = "pretty"
)

const (
	EtcdViewerENVPrefix    = "ETCD_VIEWER"
	EtcdViewerENVKey       = "KEY"       // environment variable: ETCD_VIEWER_KEY
	EtcdViewerENVCert      = "CERT"      // environment variable: ETCD_VIEWER_CERT
	EtcdViewerENVCaCert    = "CA-CERT"   // environment variable: ETCD_VIEWER_CA_CERT
	EtcdViewerENVEndpoints = "ENDPOINTS" // environment variable: ETCD_VIEWER_ENDPOINTS
)

const (
	DefaultEndpoints  = "https://127.0.0.1:2379"
	DefaultKeyFile    = "/etc/kubernetes/pki/apiserver-etcd-client.key"
	DefaultCertFile   = "/etc/kubernetes/pki/apiserver-etcd-client.crt"
	DefaultCaCertFile = "/etc/kubernetes/pki/etcd/ca.crt"
)
