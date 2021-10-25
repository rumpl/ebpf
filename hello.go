package main

import "C"

import (
	"fmt"
	"os"
	"os/signal"

	bpf "github.com/aquasecurity/tracee/libbpfgo"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	bpfModule, err := bpf.NewModuleFromFile("hello.bpf.o")
	must(err)

	err = bpfModule.BPFLoadObject()
	must(err)

	prog, err := bpfModule.GetProgram("hello_bpftrace")
	must(err)
	_, err = prog.AttachRawTracepoint("sys_enter")
	must(err)

	e := make(chan []byte, 300)
	p, err := bpfModule.InitPerfBuf("events", e, nil, 1024)
	must(err)

	p.Start()

	counter := make(map[string]int, 350)
	go func() {
		for data := range e {
			comm := string(data)
			counter[comm]++
		}
	}()

	<-sig

	for comm, n := range counter {
		fmt.Printf("%s: %d\n", comm, n)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
