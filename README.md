# Clean Architecture training by Golang Web API

- I referred to the following repository.Thanks.
  - https://github.com/hirotakan/go-cleanarchitecture-sample

# Start

```
git clone git@github.com:stobita/golang-ca-sample.git
cd go-ca-sample
dep ensure
docker-compose up -d
docker-compose run --rm app go run db/migrate.go up
```

endpoint is localhost:8080
