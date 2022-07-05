package persistence

import (
	"dislinkt/job_offer_service/ent"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func GetClient(host, port string) (*ent.Client, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=postgres password=ftn dbname=dislinkt sslmode=disable", host, port)
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}
