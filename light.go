package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"io/ioutil"
)

func runCmd(cmd string) {
	a := exec.Command(
		"screen",
		"-S",
		"skegcraft",
		"-X",
		"stuff",
		cmd + "\n",
	)

	b, err := a.Output()
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("%s", string(b))
}

func index(w http.ResponseWriter, req *http.Request) {
	html, err := ioutil.ReadFile("index.html")
	if err != nil {
		return
	}

	fmt.Fprintf(w, string(html));

	if req.FormValue("dan-on") != "" {
		fmt.Fprintf(w, "lights turned on")
		runCmd("setblock -81 78 247 lever[face=floor,powered=true]")
	} else if req.FormValue("dan-off") != "" {
		fmt.Fprintf(w, "lights turned off")
		runCmd("setblock -81 78 247 lever[face=floor,powered=false]")
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8060", nil)
}
