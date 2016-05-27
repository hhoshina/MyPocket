package main

/*=========================================================================
#  Copyright (C) 2016, H.Hoshina
#-------------------------------------------------------------------------
#  My Pocket Server Link
#
#  MypocketServerLink.go
#
#  2016/5/29  H.Hoshina
#-------------------------------------------------------------------------
# Attension:
#=========================================================================*/

/*=========================================================================
#-------------------------------------------------------------------------
# Update:
#=========================================================================*/

import "net/http"
import "io/ioutil"
import "encoding/json"
import "fmt"
import "log"

/*=====================================================================
#     (failOnError)
#=====================================================================*/
func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

/*=====================================================================
#     (HTTP Request関数)
#=====================================================================*/
func HttpRequest() {
	type Token struct {
		issued string
		token  string
	}

	req, _ := http.NewRequest("GET", "https://cocoa.ntt.com/rest/users/v1/token?mode=0", nil)
	req.Header.Set("X-WSSE", "UsernameToken Username=\"hoshina@trust.ocn.ne.jp\",PasswordText=\"Kaei8349aA8k1\",AccessKey=\"Mypocket58316053905d254dd29c6ed97ba53caa471f3ce977db9a1fbc3af9cd2955614d6ad9cf4499fee5c694c0a54e91fd25ae7f17b442631605270001\",UsernameType=\"1\"")

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	var js Token
	json.Unmarshal([]byte(string(byteArray)), &js)
	fmt.Println(js)
}

/*=====================================================================
#     (main関数)
#=====================================================================*/
func main() {
	HttpRequest()
}
