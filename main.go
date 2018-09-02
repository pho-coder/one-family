package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func getRandInt() int {
	seed := rand.NewSource(time.Now().Unix())
	r := rand.New(seed)
	return r.Intn(1000000000)
}

func getExternalIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	ipStr, _ := ioutil.ReadAll(resp.Body)
	return string(ipStr)
}

func getNewName() (name string, err error) {
	myname := ""
	host, err := os.Hostname()
	if err != nil {
		log.Printf("%s", err)
		return myname, err
	}
	myname += host
	myname += "-"
	myname += time.Now().Format("20060102150405")
	myname += "-"
	myname += fmt.Sprintf("%d", getRandInt())
	return myname, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func getName() (name string, err error) {
	nameFile := "my-name"
	// name file exists
	if fileExists(nameFile) {
		myname, err1 := ioutil.ReadFile(nameFile)
		if err != nil {
			log.Printf("%s", err1)
			return "", err1
		}
		return string(myname), nil
	}
	// name file not exists, get hostname and name myself
	myname, err2 := getNewName()
	if err != nil {
		log.Printf("%s", err2)
		return myname, err2
	}
	ioutil.WriteFile("my-name", []byte(myname), 0777)
	return myname, nil
}

func getBornTime() (bornTime time.Time) {
	nameFile := "my-name"
	f, err := os.Stat(nameFile)
	if err != nil {
		log.Fatalln(err)
	}
	UTCLoc, err := time.LoadLocation("")
	if err != nil {
		log.Fatalln(err)
	}
	return f.ModTime().In(UTCLoc)
}

func index(writer http.ResponseWriter, request *http.Request) {
	myname, err := getName()
	if err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
	info := ""
	info += "My name is "
	info += myname
	info += "\n"
	info += "I was born in "
	info += getBornTime().String()
	info += "\n"
	info += "My address is "
	info += getExternalIP()
	info += "\n"
	writer.Write([]byte(info))
}

func main() {
	pid := os.Getpid()
	ioutil.WriteFile(".pid", []byte(fmt.Sprintf("%d", pid)), 0777)
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr:    "0.0.0.0:7474",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
