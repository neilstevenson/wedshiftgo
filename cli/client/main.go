package main

import (
	"context"
	"fmt"
	"log"
	"os"
        "time"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/logger"
)

const map_name = "neil"
var loggingLevel = logger.DebugLevel

func getClient(ctx context.Context) *hazelcast.Client {
        host_ip := os.Getenv("HOST_IP")
	fmt.Printf("--------------------------------------\n")
	fmt.Printf("HOST_IP '%s'\n", host_ip)
	fmt.Printf("--------------------------------------\n")

	config := hazelcast.Config{}
	config.ClientName = "neil-go"
	config.Cluster.Name = "dev"
	config.Cluster.Network.SetAddresses(host_ip)
        config.Cluster.Unisocket = true
	config.Logger.Level = loggingLevel
        config.Stats.Enabled = true
	client, err := hazelcast.StartNewClientWithConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
        ctx := context.TODO()
	hazelcastClient := getClient(ctx)

	fmt.Printf("======================================\n")
	neil_map, _ := hazelcastClient.GetMap(ctx, map_name)
	size, _ := neil_map.Size(ctx)
	fmt.Printf("Map '%s' has size %d\n", map_name, size)
	fmt.Printf("======================================\n")

	time.Sleep(30 * time.Minute)
	fmt.Printf("Disconnecting\n")
	hazelcastClient.Shutdown(ctx)
}
