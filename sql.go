/*
* @Author: scottxiong
* @Date:   2019-08-02 16:46:32
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-08-02 17:39:40
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://micky.life/sql/temp.sql"

var name string

func main() {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error Status Code: %d", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./temp.sql")
	if err != nil {
		panic(err)
	}
	f.Write(all)
}
