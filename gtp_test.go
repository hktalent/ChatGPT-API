package gtp

import (
	"log"
	"testing"
)

func TestCompletions(t *testing.T) {
	for _, k := range Keys {
		if got, err := Completions("从百度查询xx并保存为csv的python代码", k); nil == err {

			if "" != got {
				println(k)
				log.Println(got)
				break
			}
		}

	}

}
