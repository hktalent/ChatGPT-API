# ChatGPT-API

# How use

```
curl -XPOST -v -H "user-agent: Mozilla/5.0 (Windows NT 6.1; rv:45.0) Gecko/20100101 Firefox/45.0" 'https://51pwn.com/chatGPT?q=介绍支持matrix协议、开源客户端聊天软件'
```

#### download from release
https://github.com/hktalent/ChatGPT-API/releases

```
wget -c https://github.com/hktalent/ChatGPT-API/releases/download/0.0.3/ChatGPT_0.0.3_macOS_amd64.zip
unzip -x ChatGPT_0.0.3_macOS_amd64.zip
./ChatGPT -i -k [your key]

```
<img width="700" alt="Screenshot 2022-12-18 at 18 10 56" src="https://user-images.githubusercontent.com/18223385/208293150-c7a18250-6ce5-41aa-99de-cd105e95eaf1.png">

<img src=https://user-images.githubusercontent.com/18223385/208293119-45384470-56ec-4e53-ab0a-67bf524a81bf.gif width=500>

# Example
```go
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


```
