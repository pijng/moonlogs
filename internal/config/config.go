package config

import "moonlogs/ent"

var client *ent.Client

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}
