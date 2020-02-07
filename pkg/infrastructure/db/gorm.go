package db

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// CreateDBConnection returns a handle to the DB object
func CreateDBConnection(url *url.URL, timeout time.Duration) *gorm.DB {

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

	t := time.Now()
	log.Infof("Database connection timeout: %f seconds", timeout.Seconds())
	for {
		db, err := gorm.Open(url.Scheme, connString)
		if err == nil {
			return db
		}
		if time.Since(t).Seconds() < timeout.Seconds() {
			continue
		} else {
			d := time.Since(t)
			log.Fatalf("Database connection is not open for %f seconds. %s", d.Seconds(), err)
		}
	}
}
