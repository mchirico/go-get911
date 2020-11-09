package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"time"
)

type Entity struct {
	Value string
}

func E() {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "mchirico")
	if err != nil {
		fmt.Println(err)
	}


	query := datastore.NewQuery("Garbo2")

	it := client.Run(ctx, query)
	for {
		var task Entity
		_, err := it.Next(&task)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next task: %v", err)
		}
		fmt.Printf("Task %q, \n", task.Value)
	}


}


func D() {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "mchirico")
	if err != nil {
		fmt.Println(err)
	}


	datastore.NewQuery("Garbo2")

	k := datastore.NameKey("Garbo2", "Garbo2", nil)
	e := new(Entity)
	if err := client.Get(ctx, k, e); err != nil {
		fmt.Println(err)
	}

	old := e.Value
	e.Value = "Hello World!"

	if _, err := client.Put(ctx, k, e); err != nil {
		// Handle error.
		fmt.Println(err)
	}

	fmt.Printf("Updated value from %q to %q\n", old, e.Value)
}


func main() {

	D()
	E()
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "mchirico")
	if err != nil {
		fmt.Println(err)
	}

	type Weather struct {
		TimeStamp       time.Time `datastore:timeStamp`
		Description string `datastore:desc`
	}
	key := datastore.NameKey("Weather", "weather", nil)
	weather := &Weather{}
	if err := client.Get(ctx, key, weather); err != nil {
		fmt.Println(err)
	}
}