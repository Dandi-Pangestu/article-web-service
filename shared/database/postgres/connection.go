package postgres

import (
	"fmt"
	"strings"

	log "article-web-service/shared/log/app"
	"github.com/spf13/viper"
)

func ConnectionDsn(database string) string {
	instance := viper.GetString(fmt.Sprintf("application.resources.database.%s.instance", database))
	port := viper.GetInt(fmt.Sprintf("application.resources.database.%s.port", database))
	dbname := viper.GetString(fmt.Sprintf("application.resources.database.%s.dbname", database))
	username := viper.GetString(fmt.Sprintf("application.resources.database.%s.username", database))
	password := viper.GetString(fmt.Sprintf("application.resources.database.%s.password", database))
	options := viper.GetString(fmt.Sprintf("application.resources.database.%s.options", database))

	if len(instance) == 0 {
		log.Panic(nil, "Database instance is required")
	}

	if port == 0 {
		log.Panic(nil, "Database port is required")
	}

	if len(dbname) == 0 {
		log.Panic(nil, "Database name is required")
	}

	if len(username) == 0 {
		log.Panic(nil, "Database username is required")
	}

	if len(password) == 0 {
		log.Panic(nil, "Database password is required")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", instance, username, password, dbname, port)

	if len(options) > 0 {
		opts := strings.Split(options, "&")
		for _, opt := range opts {
			dsn += " " + opt
		}
	}

	return dsn
}
