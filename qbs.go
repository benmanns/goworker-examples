package main

import (
	"bytes"
	"errors"
	"github.com/coocood/qbs"
	_ "github.com/lib/pq"
	"net/url"
	"strings"
)

var (
	errorInvalidScheme   = errors.New("Invalid Postgres database URI scheme.")
	errorMissingDatabase = errors.New("You must specify a Postgres database.")
)

// Converts struct name FooBar to table name
// foo_bars.
func toSnakePlural(s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				buf.WriteByte('_')
			}
			buf.WriteByte(c + 32)
		} else {
			buf.WriteByte(c)
		}
	}
	buf.WriteByte('s')
	return buf.String()
}

// Converts table name foo_bars to struct name
// FooBar.
func snakePluralToUpperCamel(s string) string {
	buf := new(bytes.Buffer)
	first := true
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' && first {
			buf.WriteByte(c - 32)
			first = false
		} else if c == '_' {
			first = true
			continue
		} else {
			if i != len(s)-1 || c != 's' {
				buf.WriteByte(c)
			}
		}
	}
	return buf.String()
}

// Creates a DSN from a database URI. Expects
// format postgres://user:pass@host:port/db.
func postgresDSNFromUri(uriString string) (*qbs.DataSourceName, error) {
	uri, err := url.Parse(uriString)
	if err != nil {
		return nil, err
	}

	if uri.Scheme != "postgres" {
		return nil, errorInvalidScheme
	}

	if len(uri.Path) <= 1 {
		return nil, errorMissingDatabase
	}

	dsn := qbs.DefaultPostgresDataSourceName(uri.Path[1:])

	if uri.User != nil {
		dsn.Username = uri.User.Username()
		dsn.Password, _ = uri.User.Password()
	}

	host := strings.Split(uri.Host, ":")
	dsn.Host = host[0]
	if len(host) > 1 {
		dsn.Port = host[1]
	}

	return dsn, nil
}
