// get in json form post request, posts same json to nano node, and responds by post
package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	//"reflect"
)

var jsonMap interface{}


func main() {
	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {

		//json.NewDecoder(r.Body).Decode(&jsonMap)
		//fmt.Println(jsonMap)
		//jsonBytes, _ := json.Marshal(jsonMap)
		//fmt.Println(jsonBytes)

		jsonBytes, _ := ioutil.ReadAll(r.Body)
		respJson := PassRequestToNanoNode(jsonBytes)

		w.Header().Set("Content-Type", "application/json")
		w.Write(respJson)

		//json.NewEncoder(w).Encode(resp) // sending json type
	})

	fmt.Println("Listening on :9077/rpc for json")
	http.ListenAndServe(":9077", nil)
}


func PassRequestToNanoNode(b []uint8) []uint8 {

	url := "http://[::1]:7076" // localhost

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { panic(err) }

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println("resp body: ", string(body))

	return body
}

// example use:
// curl -s -XPOST -d'{ "action": "block_count" }' CONTAINER.PUBLIC.IP.ADDRESS:9077/rpc
// curl -g -d '{ "accounts": ["xrb_ADDRESS"], "action": "accounts_balances"  }' CONTAINER.PUBLIC.IP.ADDRESS:9077/rpc
