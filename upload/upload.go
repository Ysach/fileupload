package upload

import (
	"net/http"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
	"html/template"
	//"encoding/json"
	//"log"
	//"io/ioutil"
	"encoding/json"
)

type test_struct struct {
	Test string
}

type ReturnRes struct {
	Status string `json:"status"`
}

type ReturnSlice struct {
	ResStatus []ReturnRes `json:"resStatus"`
}

// 处理/upload 逻辑
func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		fmt.Println("时间：", crutime)
		h := md5.New()
		fmt.Println("md5数值:", h)
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("Token的数值：", token)

		t, _ := template.ParseFiles("templates/upload.html")
		t.Execute(w, nil)
	} else {
		r.ParseMultipartForm(32 << 20)
		/*
		decoder := json.NewDecoder(r.Body)
		var t test_struct
		err := decoder.Decode(&t)
		defer r.Body.Close()
		log.Println(t.Test)
		*/


		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		file, handler, err := r.FormFile("filename")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./filedir/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		//res := `{"sucess": "sucess"}`
		var res ReturnSlice
		res.ResStatus = append(res.ResStatus, ReturnRes{Status:"Success"})
		b, _ := json.Marshal(res)

		fmt.Fprintf(w, string(b))
	}
}
