# ebpf 

Testing out some ebpf goodness, code taken initially from [here](https://github.com/lizrice/libbpfgo-beginners)

## Launching

Use the Docker Dev Environments feature, create a dev env with this repo, start to code.

## Running

First build the project

```shell
$ make
```

And then you can run it

```shell
$ sudo ./hello
```

Stop it (ctrl+c) after a while and you will be greeted with the number of syscalls each process made

```shell
$ sudo ./hello
^Ckube-apiserver: 956
etcd: 554
ip6tables: 73
containerd: 298
kube-controller: 784
docker-init: 2
logwrite: 31
storage-provisi: 27
iptables: 73
dockerd: 1553
containerd-shim: 66
hello: 16499
node: 91
kube-scheduler: 77
vpnkit-bridge: 342
coredns: 295
memlogd: 133
kubelet: 15180
```