package main

import (
	"k8s.io/client-go/rest"
)

func main() {

	rest.InClusterConfig()
}
