package main

import (
	"strings"
	"testing"
)

func TestHeadLines(t *testing.T) {
	in := "a\nb\nc\nd\n"
	got := headLines(strings.NewReader(in), 2)
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("取前两行出错: %v", got)
	}
}

func TestHeadBytes(t *testing.T) {
	buf, err := headBytes(strings.NewReader("hello world"), 5)
	if err != nil {
		t.Fatal(err)
	}
	if string(buf) != "hello" {
		t.Fatalf("取前 5 字节出错: %q", buf)
	}
}
