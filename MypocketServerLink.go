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
var MyPocket_Email = "hoshina@trust.ocn.ne.jp"
var MyPocket_Password = "Kaei8349aA8k1"
var MyPocket_AccessKey = "Mypocket58316053905d254dd29c6ed97ba53caa471f3ce977db9a1fbc3af9cd2955614d6ad9cf4499fee5c694c0a54e91fd25ae7f17b442631605270001"

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
func FileCopy() {
	req, _ := http.NewRequest("POST", "https://cocoa.ntt.com/rest/storage/v1/files/copy", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Content-length", "68")
	req.Header.Set("Authorization", "")

}

/*=====================================================================
#     (main関数)
#=====================================================================*/
func main() {

	var X_WSSE_DATA string
	X_WSSE_DATA += "UsernameToken Username=\""
	X_WSSE_DATA += MyPocket_Email
	X_WSSE_DATA += "\",PasswordText=\""
	X_WSSE_DATA += MyPocket_Password
	X_WSSE_DATA += "\",AccessKey=\""
	X_WSSE_DATA += MyPocket_AccessKey
	X_WSSE_DATA += "\",UsernameType=\"1\""
	fmt.Println(X_WSSE_DATA)
	fmt.Printf(GetToken(X_WSSE_DATA))
}
