package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/go-gorp/gorp"
	"github.com/scanbadge/api/utility"
	"io/ioutil"
)

// Config contains the current configuration settings.
var Config Configuration

// JwtKey contains the key used for signing and verifying a JWT with HS256.
var JwtKey []byte

// Dbmap contains a pointer to the gorp.DpMap
var Dbmap *gorp.DbMap

// Configuration stores the main configuration for the application.
type Configuration struct {
	ServerHost string // The hostname on which the HTTP server will run, e.g. 'localhost'.
	ServerPort int    // The port on which the HTTP server will run, e.g. '8080'.
	Key        string // The relative path of the base64-encoded key used for signing JWT. Recommended size of key: 256 bits.
	Database   MySQLConfiguration
}

// MySQLConfiguration stores the specific MySQL configuration for this application.
type MySQLConfiguration struct {
	Username     string // The username of the MySQL user.
	Password     string // The password of the MySQL user.
	DatabaseName string // The MySQL database to use.
	Protocol     string // The protocol to use to connect to the MySQL server, either 'tcp' or 'udp'.
	Host         string // The host to use to connect to the MySQL server, e.g. 'localhost'.
	Port         string // The port to use to connect to the MySQL server, e.g. '3306'.
	Charset      string // The charset to use for connection string, e.g. 'utf8mb4,utf8'.
	Engine       string // The database engine of the MySQL server, e.g. 'InnoDB'.
	Encoding     string // The encoding used on the MySQL server, e.g. 'UTF8'.
}

// Read reads the configuration file and parses it. If any errors are found, a panic will occur.
func Read() {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err.Error())
	}
}

// ReadKey reads the key used for JWT authenticating.
func ReadKey() {
	cfile := Config.Key

	if cfile != "" {
		f, err := utility.ReadData(Config.Key)

		if err != nil {
			panic(err.Error())
		}

		db, err := utility.DecodeBase64(f)

		if err != nil {
			panic(err.Error())
		}

		JwtKey = db
	} else {
		panic(fmt.Errorf("cannot read key, because its value is not set in configuration file"))
	}
}
