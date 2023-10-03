package main

import (
	"chubby/chubby/api"
	"chubby/chubby/client"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var leader_election_id1		string		// ID of this client.

func init() {
	flag.StringVar(&leader_election_id1, "leader_clientID1", "leader_election_id1", "ID of this client1")
}

func main() {
	// Parse flags from command line.
	flag.Parse()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sess, err := client.InitSession(api.ClientID(leader_election_id1))
	if err != nil {
		log.Fatal(err)
	}
	errOpenLock := sess.OpenLock("Lock/Lock1")
	if errOpenLock != nil {
		log.Fatal(errOpenLock)
	}
	isSuccessful, err := sess.TryAcquireLock("Lock/Lock1", api.EXCLUSIVE)
	if !isSuccessful {
		fmt.Printf("Lock Acquire Unexpected Failure")
	}
	if err != nil {
		log.Fatal(err)
	}
	isSuccessful, err = sess.WriteContent("Lock/Lock1", leader_election_id1)
	if !isSuccessful {
		fmt.Println("Unexpected Error Writing to Lock")
	}
	if err != nil {
		log.Fatal(err)
	}
	content, err := sess.ReadContent("Lock/Lock1")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read Content is %s\n",content)
	}
	for {
		// Exit on signal.
		select {
		case <- quitCh:
			break
		default:
			continue
		}
	}
}

