package app

import (
	"log"
)

type AppConfig struct {
	Logger *log.Logger /* An example of an 'injection', we could also add things like a DB connection, etc. */
}
