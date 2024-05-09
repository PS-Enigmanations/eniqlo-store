# Eniqlo Store

https://openidea-projectsprint.notion.site/EniQilo-Store-93d69f62951c4c8aaf91e6c090127886?pvs=4

**Environment:**

```sh
cp .env.example .env
```

**Database:**

```sh
DB_HOST=localhost
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=eniqlo-store
DB_PORT=5432
DB_PARAMS="sslmode=disabled"
JWT_SECRET=a-very-secretive-secret-key
BCRYPT_SALT=8
```

**Run migration:**

1. Install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)

2. Run scripts

```sh
make migrateup
```

**Setup:**

```sh
go mod download
```

**Run development mode:**

1. Install [go air](https://github.com/cosmtrek/air):

```sh
go install github.com/cosmtrek/air@latest
```

2. Start development server:

```sh
air
```
