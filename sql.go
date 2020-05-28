/*
* @Author: scottxiong
* @Date:   2019-08-02 16:46:32
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-08-07 11:43:08
 */
package main

import (
	"flag"
	"fmt"
	"github.com/scott-x/gutils/fs"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	domain           = "https://micky.life"
)

var templates_folder = fs.HOME+"/go/src/github.com/scott-x/templates/sql"

var local, remote bool

func main() {
	flag.BoolVar(&local, "l", false, "get the templates via local")
	flag.BoolVar(&remote, "r", false, "get the templates via remote server")
	flag.Parse()
	switch len(os.Args) {
	case 1:
		flag.PrintDefaults()
	case 2:
		if os.Args[1] != "-l" && os.Args[1] != "-r" {
			flag.PrintDefaults()
		}
		switch os.Args[1] {
		case "-l":
			fs.CopyFolder(templates_folder, "./")
		case "-r":
			source := []string{"temp.sql", "example/user1.sql", "example/user2.sql", "example/blog1.sql", "example/blog2.sql"}
			createDirIfNotExist("example")
			for _, s := range source {
				getContents(s, "./"+s)
			}
			fmt.Println("import data: mysql [-h host] -uuser -p > xxx.sql")
		}
	}

}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func getContents(remote, local string) {
	resp, err := http.Get(domain + "/sql/" + remote)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error Status Code: %d", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(local)
	if err != nil {
		panic(err)
	}
	f.Write(all)
}
