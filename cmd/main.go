package main

import "L0/pkg"

func main() {
	var a pkg.App
	a.Run("test-cluster", "service", "nats://0.0.0.0:4222")
}


