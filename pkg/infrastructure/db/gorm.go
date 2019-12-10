package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"net/url"
	"strings"
)

//returns a handle to the DB object
func CreateDBConnection(url *url.URL) *gorm.DB {

	if url.Scheme != "postgres" {
		logrus.Fatalf("Unsupported database scheme [%s]", url.Scheme)
	}

	connFmt := "host=%s port=%s user=%s password=%s dbname=%s %s"
	host := url.Hostname()
	port := url.Port()
	user := url.User.Username()
	password, _ := url.User.Password()
	database := strings.TrimLeft(url.Path, "/")
	query := url.RawQuery // Here will be passed all options of connections, like a `sslmode=disable`
	connString := fmt.Sprintf(connFmt, host, port, user, password, database, query)

	db, err := gorm.Open(url.Scheme, connString)
	if err != nil {
		logrus.Fatalf("Database connection is not open. %s", err)
	}

	return db
}