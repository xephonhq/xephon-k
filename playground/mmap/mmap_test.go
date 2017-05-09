package mmap

import (
	"testing"
	"os"
	"syscall"
)

// TODO:
// - write
// - write & read struct
// - write in golang, read in C/C++

func TestMmap_Read(t *testing.T) {
	// C source files not allowed when not using cgo or SWIG
	f, err := os.Open("demo.txt")
	if err != nil {
		t.Fatal(err)
		return
	}
	stat, err := f.Stat()
	size := stat.Size()
	// NOTE: from https://github.com/google/codesearch/blob/master/index/mmap_linux.go#L19
	// we can't handle single file larger than 4GB, due to the len function does not work with slice longer than 2^32
	// TODO: why not use int64(int(size)) == size
	if int64(int(size+4095)) != size+4095 {
		t.Fatalf("%s: too large for mmap", f.Name())
		return
	}
	n := int(size)
	if n == 0 {
		t.Fatal("file is empty!")
		return
	}
	// TODO: why codesearch is using (n+4095)&^4095? align to 4KB? I think so
	t.Log((n+4095)&^4095)
	data, err := syscall.Mmap(int(f.Fd()),0, n, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		t.Fatalf("mmap %s: %v", f.Name(), err)
	}
	t.Log(len(data))
	t.Log(string(data[0:10]))
}
