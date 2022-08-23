package main

import (
	"context"
	"fmt"

	"github.com/marcos-nsantos/comments-rest-api/internal/database"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application.
func Run() error {
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err = db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfully connected and pinged the database")

	fmt.Println("starting up our application")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
