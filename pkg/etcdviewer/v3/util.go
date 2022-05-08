// Package v3
/* Copyright © 2022 JeffreyGuan <double12gzh@gmail.com> */

package v3

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	k8sSerializerJson "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/kubectl/pkg/scheme"
)

// ListKeys gets all keys in etcd
func (cli *v3) ListKeys(key string, isPretty bool) {
	defer func() {
		_ = cli.client.Close()
	}()

	resp, err := clientv3.NewKV(cli.client).Get(context.TODO(), key, clientv3.WithSerializable(), clientv3.WithFromKey(), clientv3.WithKeysOnly())
	if err != nil {
		logger.Error(err)
		return
	}

	for _, value := range resp.Kvs {
		logger.Info(string(value.Key))
	}
}

// Get retrieve value by key
func (cli *v3) Get(key string, isPretty bool) {
	defer func() {
		_ = cli.client.Close()
	}()

	resp, err := clientv3.NewKV(cli.client).Get(context.TODO(), key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		logger.Error(err)
		return
	}

	universalDeserializer := scheme.Codecs.UniversalDeserializer()
	strictJSONSerializer := k8sSerializerJson.NewSerializerWithOptions(
		k8sSerializerJson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
		k8sSerializerJson.SerializerOptions{
			Yaml:   false,
			Pretty: true,
			Strict: false,
		},
	)

	for _, value := range resp.Kvs {
		obj, _, err := universalDeserializer.Decode(value.Value, nil, nil)
		if err != nil {
			logger.Error(err)
			continue
		}

		ioWriter := &bytes.Buffer{}
		if err = strictJSONSerializer.Encode(obj, ioWriter); err != nil {
			logger.Error(err)
			continue
		}

		i := map[string]interface{}{}
		if err = json.Unmarshal(ioWriter.Bytes(), &i); err != nil {
			logger.Error(err)
			continue
		}

		data := make([]byte, 0)
		if isPretty {
			data, err = json.MarshalIndent(i, "", "    ")
		} else {
			data, err = json.Marshal(i)
		}

		if err != nil {
			logger.Error(err)
			continue
		}

		logger.Info(string(data))
	}
}

// Dump dump data from etcd
func (cli *v3) Dump(key string, isPretty bool) {
	defer func() {
		_ = cli.client.Close()
	}()

	response, err := clientv3.NewKV(cli.client).Get(context.TODO(), key, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		logger.Error(err)
		return
	}

	type dumpedData struct {
		Key            string `json:"key,omitempty"`
		Value          string `json:"value,omitempty"`
		Version        int64  `json:"version,omitempty"`
		Lease          int64  `json:"lease,omitempty"`
		CreateRevision int64  `json:"create_revision,omitempty"`
		ModRevision    int64  `json:"mod_revision,omitempty"`
	}

	result := make([]dumpedData, 0)
	universalDeserializer := scheme.Codecs.UniversalDeserializer()
	strictJSONSerializer := k8sSerializerJson.NewSerializerWithOptions(
		k8sSerializerJson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
		k8sSerializerJson.SerializerOptions{
			Yaml:   false,
			Pretty: true,
			Strict: false,
		},
	)

	for _, value := range response.Kvs {
		obj, _, err := universalDeserializer.Decode(value.Value, nil, nil)
		if err != nil {
			logger.Error(err)
			continue
		}
		ioWriter := &bytes.Buffer{}
		if err := strictJSONSerializer.Encode(obj, ioWriter); err != nil {
			logger.Error(err)
			continue
		}
		result = append(result, dumpedData{
			Key:            string(value.Key),
			Value:          string(ioWriter.Bytes()),
			Lease:          value.Lease,
			Version:        value.Version,
			CreateRevision: value.CreateRevision,
			ModRevision:    value.ModRevision,
		})
	}

	var data []byte
	if isPretty {
		data, err = json.MarshalIndent(result, "", "    ")
	} else {
		data, err = json.Marshal(result)
	}

	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info(string(data))
}

// Watch monitor changes
func (cli *v3) Watch(key string) {
	// TODO
}

// GetMembers get endpoints
func (cli *v3) GetMembers(ctx context.Context, timeout time.Duration) {
	cx, cancel := context.WithTimeout(ctx, timeout)
	resp, err := cli.client.MemberList(cx)

	defer func() {
		if err := cli.client.Close(); err != nil {
			logger.Error(err)
		}
		cancel()
	}()

	if err != nil {
		logger.Error(err)
	}

	for _, v := range resp.Members {
		logger.Info(v.String())
	}
}
