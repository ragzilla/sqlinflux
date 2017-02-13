package main

import (
	"database/sql"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	_ "github.com/minus5/gofreetds"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	Dsn          string
	Query        string
	Translations map[string]string
	Tags         map[string]string
}

func main() {
	filename := os.Args[1]

	c := Config{}

	s, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(s, &c)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mssql", c.Dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(c.Query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fields := make(map[string]interface{})
	for rows.Next() {
		var Variable string
		var Value uint
		err = rows.Scan(&Variable, &Value)
		if err != nil {
			panic(err)
		}
		if val, ok := c.Translations[Variable]; ok {
			Variable = val
		}
		fields[Variable] = Value
	}
	if len(fields) == 0 {
		return
	}
	Pt, err := client.NewPoint("sqlValues", c.Tags, fields, time.Now().UTC().Round(time.Minute))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", Pt.String())
}
