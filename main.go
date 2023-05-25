package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func pingdb() error {
	db, err := sql.Open("postgres", os.Args[1])
	defer db.Close()

	if err != nil {
		log.Print("Could not connect to database: ", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Print("Could not ping database: ", err)
		return err
	}

	return nil
}

func main() {
	var exitCode int

	defer func() {
		log.Print("Exiting with code: ", exitCode)
		os.Exit(exitCode)
	}()

	if len(os.Args) < 2 {
		log.Print("Provide a connection string as the first argument")
		exitCode = 2
		return
	}

	const MAX_ATTEMPTS int = 3

	for i := 0; i < MAX_ATTEMPTS; i++ {
		err := pingdb()
		if err == nil {
			log.Print("No errors")
			exitCode = 0
			return
		}

		log.Print("Sleeping for 1s")
		time.Sleep(time.Second)
	}

	log.Print("Too many attempts")
	exitCode = 3
}
