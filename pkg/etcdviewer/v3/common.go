// Package v3
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package v3

import (
	"fmt"
	"time"

	"go.etcd.io/etcd/client/pkg/v3/transport"
	clientv3 "go.etcd.io/etcd/client/v3"

	logger2 "github.com/double12gzh/etcdviewer/pkg/logger"
)

type v3 struct {
	client *clientv3.Client
}

var logger = logger2.NewLogger()

// NewClientV3 init etcd api v3 client
func NewClientV3(caFile, certFile, keyFile string, endpoints []string) *v3 {
	if caFile == "" || certFile == "" || keyFile == "" {
		logger.Error(fmt.Errorf("caFile, certFile and keyFile are needed"))
		return nil
	}

	tlsInfo := transport.TLSInfo{
		TrustedCAFile: caFile,
		CertFile:      certFile,
		KeyFile:       keyFile,
	}

	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		logger.Error(fmt.Errorf("failed to get tls info, error: %s", err.Error()))
		return nil
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
		TLS:         tlsConfig,
	})
	if err != nil {
		logger.Error(fmt.Errorf("failed to create etcd client v3, err: %s", err.Error()))
		return nil
	}

	return &v3{client: cli}
}
