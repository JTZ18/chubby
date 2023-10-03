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
	"time"
)

var acquireLock_clientID		string		// ID of this client.

func init() {
	flag.StringVar(&acquireLock_clientID, "acquireLock_clientID1", "acquireLock_clientID", "ID of this client2")
}

func main() {
	// Parse flags from command line.
	flag.Parse()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	time.Sleep(5*time.Second)
	sess, err := client.InitSession(api.ClientID(acquireLock_clientID))
	if err != nil {
		log.Fatal(err)
	}
	errOpenLock := sess.OpenLock("Lock/Lock1")
	if errOpenLock != nil {
		log.Fatal(errOpenLock)
	}
	startTime := time.Now()
	for {
		isSuccessful, err := sess.TryAcquireLock("Lock/Lock1", api.EXCLUSIVE)
		if err != nil {
			log.Println(err)
		}
		if isSuccessful && err == nil {
			isSuccessful, err = sess.WriteContent("Lock/Lock1", acquireLock_clientID)
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
        		if content == acquireLock_clientID {
                		elapsed := time.Since(startTime)
                		log.Printf("Successfully acquired lock after %s\n",elapsed)
        		}
			return
		}
	}

	content, err := sess.ReadContent("Lock/Lock1")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read Content is %s\n",content)
	}
	if content == acquireLock_clientID {
		elapsed := time.Since(startTime)
		log.Printf("Successfully acquired lock after %s\n",elapsed)
	}

	// Exit on signal.
	<-quitCh
}
