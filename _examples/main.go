package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/briandowns/jail"
)

func main() {
	o := &jail.Opts{
		Path:     "/zroot/jails/build", //Make sure this directory exists
		Name:     "jailname",
		Hostname: "hostname",
		IP4:      "192.168.0.200",
		Chdir:    true,
	}
	jid, err := jail.Jail(o)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("JID: %d - / director listing in jail", jid)
	time.Sleep(30 * time.Second)
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
