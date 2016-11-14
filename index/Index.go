package index

import (
	"net/http"
	"fmt"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "---》"))
	}
	//写入到w，可以输出到客户端
	fmt.Fprintf(w, "欢迎来到我的网站!")
}
