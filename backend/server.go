package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/smallnest/rpcx/client"
)

var (
	hashServiceAddr = "127.0.0.1:8972"
	xclient         client.XClient
)

func main() {
	initXClient()

	http.HandleFunc("/api/hash", hashHandler)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("listen on port 9981")
	http.ListenAndServe(":9981", nil)

	xclient.Close()
}

type hashRequest struct {
	Target string `json:"target"`
}

func initXClient() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+hashServiceAddr, "")
	opt := client.DefaultOption

	xclient = client.NewXClient("Hash", client.Failtry, client.RandomSelect, d, opt)
}

func hashHandler(w http.ResponseWriter, r *http.Request) {
	var decoded hashRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)
	checkError(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hard, err := strconv.Atoi(decoded.Target)
	checkError(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var reply string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = xclient.Call(ctx, "Calc", hard, &reply)
	cancel()

	checkError(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, `{ "hash": "%s" }`, reply)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
