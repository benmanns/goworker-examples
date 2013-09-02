# goworker-examples

These are examples of goworker workers.

To benchmark them yourself, run

```sh
redis-cli -r 100 RPUSH resque:queue:hello '{"class":"Hello","args":[]}'
redis-cli -r 100 RPUSH resque:queue:insert '{"class":"Insert","args":["John Doe","jdoe@example.com","(555) 555-5555"]}'
redis-cli -r 100 RPUSH resque:queue:multiply '{"class":"Multiply","args":[]}'
redis-cli -r 100 RPUSH resque:queue:sleep '{"class":"Sleep","args":[1]}'

export DATABASE_URL=postgres://goworker:goworker@localhost:5432/goworker_development

go get github.com/benmanns/goworker-examples
time $GOPATH/bin/goworker-examples -queues=hello,insert,multiply,sleep -exit-on-complete=true
```

This app requires a local Postgres database. See the above `DATABASE_URL` configuration string for more information.
