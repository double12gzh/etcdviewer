// Package options
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package options

type EtcdViewerOptions struct {
	Endpoints []string
	CAFile    string
	CertFile  string
	KeyFile   string
}

func NewEtcdViewerOptions() *EtcdViewerOptions {
	return &EtcdViewerOptions{
		Endpoints: []string{},
		CAFile:    "",
		CertFile:  "",
		KeyFile:   "",
	}
}
