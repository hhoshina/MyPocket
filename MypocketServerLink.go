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
import "os"

/*=====================================================================
#     (Struct)
#=====================================================================*/
type Token struct {
	Issued string `json:"issued"`
	Token  string `json:"token"`
}

/*=====================================================================
#     (Global定数)
#=====================================================================*/
var MyPocket_Email = os.Getenv("MyPocket_Email")
var MyPocket_Password = os.Getenv("MyPocket_Password")
var MyPocket_AccessKey = os.Getenv("MyPocket_AccessKey")

/*=====================================================================
#     (GetToken関数)
#=====================================================================*/
func GetToken(X_WSSE_DATA string) string {
	req, _ := http.NewRequest("GET", "https://cocoa.ntt.com/rest/users/v1/token?mode=0", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-WSSE", X_WSSE_DATA)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)

	var output Token

	err = json.Unmarshal(byteArray, &output)

	if err != nil {
		fmt.Println(err.Error())
	}

	return string(output.Token)
}

/*=====================================================================
#     (FileCopy関数)
#=====================================================================*/
func FileCopy(token string) {
	var AUTHORI string
	AUTHORI += "Bearer "
	AUTHORI += token
	req, _ := http.NewRequest("POST", "https://cocoa.ntt.com/rest/storage/v1/files/copy", nil)
	req.Header.Set("Authorization", AUTHORI)
	req.Header.Set("X-Authorization", MyPocket_AccessKey)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Content-length", "68")

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token)
	fmt.Println(string(byteArray))

}

/*=====================================================================
#     (main関数)
#=====================================================================*/
func main() {
	var token string
	var X_WSSE_DATA string
	X_WSSE_DATA += "UsernameToken Username=\""
	X_WSSE_DATA += MyPocket_Email
	X_WSSE_DATA += "\",PasswordText=\""
	X_WSSE_DATA += MyPocket_Password
	X_WSSE_DATA += "\",AccessKey=\""
	X_WSSE_DATA += MyPocket_AccessKey
	X_WSSE_DATA += "\",UsernameType=\"1\""

	token = GetToken(X_WSSE_DATA)
	fmt.Println(token)
}
