package main

import (
    "reflect"
	"io/ioutil"
	"flag"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
    "encoding/json"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))
var verifytmp = template.Must(template.New("verify").Parse(templateStrVerify))

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

type Message struct {
    Slots []map[string]interface{} `slotEntities`
}

type TData struct {
    StandardValue string
    OriginalValue string
}

func QR(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	r , _ := httputil.DumpRequest(req,true)
	log.Println(string(r))
    var f map[string]interface{}
    var s string
    var o string
    body,_ := ioutil.ReadAll(req.Body)
    err := json.Unmarshal(body, &f)
    log.Println(err)
    for k := range f{
        if k=="slotEntities" {
            log.Println(reflect.TypeOf(f[k]))
            b, ok := f[k].([]interface{})
            if ok {
                b2, ok := b[0].(map[string]interface{})
                if ok {
                    log.Println(reflect.TypeOf(b2["intentParameterName"]))
                    log.Println(b2["intentParameterName"])
                    s = b2["standardValue"].(string)
                    o = b2["originalValue"].(string)
                    log.Println(s)
                    log.Println(o)
                }
            }
        }
    }
    //log.Println(f.Slots[0]["intentParameterName"])
	//templ.Execute(w, req.FormValue("s"))
    var td TData
    td.StandardValue = s
    td.OriginalValue = o
    t, err := template.New("qr").Parse(templateStr)
    err = t.Execute(w, &td)
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
        "reply": "欢迎测试设备控制技能,你是要{{.OriginalValue}}台灯",
        "resultType": "RESULT",
        "properties": { "actions": "[{\"name\":\"dataResult\",\"nluReplyText\":\"你是要{{.OriginalValue}}台灯\",\"parameters\":{\"bizInfo\":\"{ \\\"operate\\\": \\\"{{.StandardValue}}\\\", \\\"oterh\\\": \\\"yes\\\" }\"}}]" }
    }
}
`
