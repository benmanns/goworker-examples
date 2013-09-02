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

## Running on Heroku

These examples will run on Heroku.

First, clone the examples.

```sh
git clone https://github.com/benmanns/goworker-examples.git myapp
cd myapp
```

Then, create an app with the go buildpack.

```sh
heroku create -b https://github.com/kr/heroku-buildpack-go
```

Next, install the `redistogo:nano` and `heroku-postgresql:dev` addons (both free).

```sh
heroku addons:add redistogo:nano
heroku addons:add heroku-postgresql:dev
```

Promote your Redis and Postgres databases. Be sure to change `[COLOR]` to whatever was provided by the `heroku-postgresql:dev` addon.

```sh
heroku config:set REDIS_PROVIDER=REDISTOGO_URL
heroku pg:promote HEROKU_POSTGRESQL_[COLOR]_URL
```

Configure the goworker process with environment variables. See the [Procfile](https://github.com/benmanns/goworker-examples/blob/master/Procfile) for more information.

```sh
heroku config:set \
  QUEUES=hello,insert,multiply,sleep \
  INTERVAL=1.0 \
  CONCURRENCY=50 \
  CONNECTIONS=2
```

Finally, push and start your workers.

```sh
git push heroku master
heroku ps:scale worker=1
```
