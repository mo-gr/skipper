// skipper program main
//
// for a summary about skipper, please see the readme file.
package main

import (
	"flag"
	"github.com/zalando/skipper/run"
	"log"
	"strings"
)

const (
	defaultAddress     = ":9090"
	defaultEtcdUrls    = "http://127.0.0.1:2379,http://127.0.0.1:4001"
	defaultStorageRoot = "/skipper"
	addressUsage       = "address where skipper should listen on"
	etcdUrlsUsage      = "urls where etcd can be found"
	insecureUsage      = "set this flag to allow invalid certificates for tls connections"
	storageRootUsage   = "prefix for skipper related data in the provided etcd storage"
	routesFileUsage    = "routes file to use instead of etcd"
)

var (
	address     string
	etcdUrls    string
	insecure    bool
	storageRoot string
	routesFile  string
)

func init() {
	flag.StringVar(&address, "address", defaultAddress, addressUsage)
	flag.StringVar(&etcdUrls, "etcd-urls", defaultEtcdUrls, etcdUrlsUsage)
	flag.BoolVar(&insecure, "insecure", false, insecureUsage)
	flag.StringVar(&storageRoot, "storage-root", defaultStorageRoot, storageRootUsage)
	flag.StringVar(&routesFile, "routes-file", "", routesFileUsage)
	flag.Parse()
}

func main() {
	log.Fatal(run.Run(address, strings.Split(etcdUrls, ","), storageRoot, insecure, routesFile))
}
