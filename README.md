![](https://img.shields.io/github/issues/double12gzh/etcdviewer) ![](https://img.shields.io/github/forks/double12gzh/etcdviewer) ![](	https://img.shields.io/github/stars/double12gzh/etcdviewer) ![](https://img.shields.io/github/license/double12gzh/etcdviewer)
# etcdviewer kubectl

A `kubectl` `etcdviewer` plugin that get or watch Kubernetes data stored in etcd.

## How it works


## Quick start

### Prerequisites
Note: You will need `git` to install the krew plugin.

The `etcdviwer` plugin is installed using the krew plugin manager for Kubernetes CLI. 
Installation instructions for `krew` can be found [here](https://krew.sigs.k8s.io/docs/user-guide/setup/install/).

### Installation
```bash
kubectl krew install etcdviewer
```

### Usage

```
kubectl etcdviewer \
    --endpoint=https://127.0.0.1:2379 \
    --key=/etc/kubernetes/pki/apiserver-etcd-client.key \
    --cert=/etc/kubernetes/pki/apiserver-etcd-client.crt \
    --cacert=/etc/kubernetes/pki/etcd/ca.crt \
    get --pretty /registry/pods/kube-system/kube-scheduler-mast
```

