package main

import (
	//"io/ioutil"
	"flag"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))
var verifytmp = template.Must(template.New("qr").Parse(templateStrVerify))

func main() {

	var i uint
	arr := make([]byte, 32)
	for i = 0; i < ((uint)(len(arr))); i++ {
		arr[i] = (byte)(i)
		log.Println(arr[i])
	}

	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	http.Handle("/aligenie/4b94f01be364f47c025eecf9b80d5bfe.txt", http.HandlerFunc(verify))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	r , _ := httputil.DumpRequest(req,true)
	log.Println(string(r))
	templ.Execute(w, req.FormValue("s"))
}
func verify(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	r , _ := httputil.DumpRequest(req,true)
	log.Println(string(r))
	verifytmp.Execute(w, req.FormValue("s"))

}
const templateStrVerify = `Jfc4Z4Ur15JwUBuvUQD5wg7Nu8+l+HscqYlfofbyJdb07RwiSRq7c8M8Z8Z6w7c+`

const templateStr = `
{
    "returnCode": "0",
    "returnErrorSolution": "",
    "returnMessage": "",
    "returnValue": {
        "reply": "欢迎测试设备控制技能",
        "resultType": "RESULT",
        "properties": { "actions": "[{\"name\":\"dataResult\",\"nluReplyText\":\"你是要打开台灯不\",\"parameters\":{\"bizInfo\":\"{ \\\"operate\\\": \\\"true\\\", \\\"oterh\\\": \\\"yes\\\" }\"}}]" }
    }
}
`

func QQQ() string {
	return "hello world"
}
