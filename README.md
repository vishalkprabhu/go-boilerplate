# GO Lang Boilerplate Dockerised   [![Build Status](https://app.travis-ci.com/vishalkprabhu/go-boilerplate.svg?branch=main)](https://app.travis-ci.com/vishalkprabhu/go-boilerplate)
###  [ Echo+MySql+GORM+Adminer ]

Golang development made easy with docker container, just build and deploy.

## Installation

Clone the repository and hit blow command.

```bash
docker-compose up
```

## URL

```python
http://localhost:1323/students
```
A json with some students results should be visible.
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Migrations

You should install go migrate cli locally

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

For linux, to install you can try this command

```
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate.linux-amd64 /bin/migrate
```

How to create a migration file

```
migrate create -ext sql -dir migrations -seq create_users_table

```

To Run Migration (up)

```
make migrate-up-local

```

To execute (down)

```
make migrate-down-local

```