package main

import (
	"bufio"
	"flag"
	"fmt"
	gtp "github.com/hktalent/ChatGPT-API"
	"os"
	"strings"
)

func doOne(q, k string) {
	if got, err := gtp.Completions(q, k); nil == err {
		fmt.Println(got)
	} else if nil != err {
		fmt.Println("gtp.Completions is err:", err)
	}
}

func main() {
	key := flag.String("k", "", "your ChatGPT-API token key")
	interact := flag.Bool("i", false, "interact")
	q := flag.String("q", "", "your question")
	flag.Parse()
	if *key != "" {
		if *interact {
			buf := bufio.NewScanner(os.Stdin)
			for buf.Scan() {
				s := buf.Text()
				x1 := strings.TrimSpace(strings.ToLower(s))
				if x1 == "exit" || x1 == "quit" || x1 == "q" || x1 == "x" {
					break
				}
				if "" != s {
					doOne(s, *key)
				}
				if nil != buf.Err() {
					break
				}
			}
		} else if "" != *q {
			doOne(*q, *key)
		}
	}
}
