package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/benmanns/goworker"
	"github.com/coocood/qbs"
	"os"
)

type Contact struct {
	Id    int64
	Name  string
	Email string
	Phone string
}

var (
	errorTableCreationFailed = errors.New("Table creation failed.")
)

func newInsertWorker(uri string, connections int) (func(string, ...interface{}) error, error) {
	dsn, err := postgresDSNFromUri(uri)
	if err != nil {
		return nil, err
	}

	qbs.RegisterWithDataSourceName(dsn)

	qbs.SetConnectionLimit(connections, true)

	migration, err := qbs.GetMigration()
	if err != nil {
		return nil, err
	}
	defer migration.Close()

	err = func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = errorTableCreationFailed
			}
		}()
		err = migration.CreateTableIfNotExists(new(Contact))
		return
	}()
	if err != nil {
		return nil, err
	}

	return func(queue string, args ...interface{}) error {
		name, ok := args[0].(string)
		if !ok {
			return fmt.Errorf("Invalid parameters %v to insert worker. Expected string, was %T.", args, args[0])
		}
		email, ok := args[1].(string)
		if !ok {
			return fmt.Errorf("Invalid parameters %v to insert worker. Expected string, was %T.", args, args[1])
		}
		phone, ok := args[2].(string)
		if !ok {
			return fmt.Errorf("Invalid parameters %v to insert worker. Expected string, was %T.", args, args[2])
		}
		return qbs.WithQbs(func(db *qbs.Qbs) error {
			contact := &Contact{Name: name, Email: email, Phone: phone}
			_, err := db.Save(contact)
			return err
		})
	}, nil
}

func init() {
	qbs.StructNameToTableName = toSnakePlural
	qbs.TableNameToStructName = snakePluralToUpperCamel

	var connections int
	flag.IntVar(&connections, "insert-connections", 5, "maximum DB connections for insert worker")

	insertWorker, err := newInsertWorker(os.Getenv("DATABASE_URL"), connections)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	goworker.Register("Insert", insertWorker)
}
