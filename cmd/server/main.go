package main

import (
	"fmt"

	"github.com/marcos-nsantos/comments-rest-api/internal/comment"
	"github.com/marcos-nsantos/comments-rest-api/internal/database"
	transportHttp "github.com/marcos-nsantos/comments-rest-api/internal/transport/http"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application.
func Run() error {
	fmt.Println("starting up our application")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err = db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	httpHandler := transportHttp.NewHandler(cmtService)
	if err = httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
