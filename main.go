// go-head 显示文件或标准输入的开头若干行（-n）或若干字节（-b）。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// headLines 取前 n 行，headBytes 取前 n 字节。
func headLines(r io.Reader, n int) []string {
	sc := bufio.NewScanner(r)
	var out []string
	count := 0
	for sc.Scan() && count < n {
		out = append(out, sc.Text())
		count++
	}
	return out
}

func headBytes(r io.Reader, n int) ([]byte, error) {
	buf := make([]byte, n)
	m, err := io.ReadFull(r, buf)
	if err != nil && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	return buf[:m], nil
}

func main() {
	n := flag.Int("n", 10, "显示前 N 行")
	b := flag.Int("b", 0, "显示前 N 字节（>0 时优先于 -n）")
	flag.Parse()

	var r io.Reader = os.Stdin
	args := flag.Args()
	if len(args) > 0 {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "打开失败:", err)
			os.Exit(1)
		}
		defer f.Close()
		r = f
	}

	if *b > 0 {
		buf, err := headBytes(r, *b)
		if err != nil {
			fmt.Fprintln(os.Stderr, "读取失败:", err)
			os.Exit(1)
		}
		fmt.Print(string(buf))
		return
	}

	for _, line := range headLines(r, *n) {
		fmt.Println(line)
	}
}
