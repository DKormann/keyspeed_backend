package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"keyspeed/storage"
	"keyspeed/util"
	"net/http"
	"strings"
)

const (
	OK          = 200
	NotFound    = 400
	ServerError = 500
	Denied      = 409
)

func setup() {

	handle("hello", func(arg []string) (data any, err error) {
		util.GLogln("hello")
		data = "world"
		return
	})

	handle("register/<username>/<pwd_hash>", func(arg []string) (token any, err error) {
		username := arg[0]
		passhash := arg[1]

		fmt.Println("Register ", username, " with password ", passhash)

		token = "1234567890"
		return
	})

	handle("getdata/<id>", func(arg []string) (data any, err error) {
		id := arg[0]
		util.GLogln("getdata ", id)
		data, err = storage.Get(id)
		if err != nil {
			util.RLogln(err)
		}
		return
	})

	handle("list", func(arg []string) (data any, err error) {
		util.GLogln("list")
		data, err = storage.ListAll()
		if err != nil {
			util.RLogln(err)
		}
		return
	})

	http.HandleFunc("/keyspeed_api/setdata/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Print("got post request: ")
		util.BLogln(r.URL.Path)

		id := strings.Split(r.URL.Path, "/")[3]
		util.GLogln("id: ", id)

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		util.BLogln(string(b))

		err = storage.Append(id, b)

		if err != nil {
			util.RLogln(err)
			return
		}

	})

}

func handle(path string, handler func([]string) (any, error)) {

	path = strings.Trim(path, "/")

	prefix := "/keyspeed_api/"
	path = prefix + strings.Trim(path, "/") + "/"

	items := strings.SplitN(path, "<", 2)
	path = items[0]

	params := ""

	if len(items) > 1 {
		params = "<" + items[1]
	}

	pathlen := len(strings.Split(strings.Trim(path, "/"), "/"))
	paramslen := 0
	if params != "" {
		paramslen = len(strings.Split(strings.Trim(params, "/"), "/"))
	}

	wrapper := func(w http.ResponseWriter, r *http.Request) {
		util.ColorLog("cyan", " >>> serving ", r.URL.Path+"\n")

		params := strings.Split(strings.Trim(r.URL.Path, "/"), "/")[pathlen:]

		util.ColorLog("blue", " ", strings.Join(params, "/"), "\n")

		if len(params) != paramslen {
			msg := fmt.Sprintf("params not matching. expected %v but got %v ", paramslen, len(params))
			util.RLogln(msg)
			w.Write([]byte(msg))
			w.WriteHeader(NotFound)
			return
		}

		(w).Header().Set("Access-Control-Allow-Origin", "*")
		resp, err := handler(params)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			util.RLogln(err.Error())
			return
		}
		data, err := json.Marshal(resp)
		if err != nil {
			util.RLog(err)
			w.WriteHeader(http.StatusInternalServerError)
			util.RLogln(err.Error())
			return
		}

		util.GLogln("OK")
		w.Write(data)
	}
	http.HandleFunc(path, wrapper)
}
