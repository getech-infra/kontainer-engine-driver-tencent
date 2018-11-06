package main

import (
	"github.com/sirupsen/logrus"
	"sync"

	"github.com/rancher/kontainer-engine/types"
	"github.com/rancher/rancher-tke-driver/driver"
)

var wg = &sync.WaitGroup{}

func main() {
	addr := make(chan string)
	go types.NewServer(driver.NewDriver(), addr).Serve()
	logrus.Infof("tke driver up and running on at %v", <-addr)

	wg.Add(1)
	wg.Wait() // wait forever, we only exit if killed by parent process
}
