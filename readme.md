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

**Run development server:**

```sh
make dev
```

## Docker:

### Running on your local machine:

**1. Create Network**

```sh
# Create a network, which allows containers to communicate
# with each other, by using their container name as a hostname
docker network create app_network
```

**2. Build postgres container**

```sh
docker-compose up --build postgres
```

**3. Migrate database**

```sh
migrate -path db/migrations -database "postgres://postgres:postgres@0.0.0.0:5430/eniqlo-store?sslmode=disable" up
```

**4. Running API:**

```sh
docker-compose up --build api
```

Open http://localhost:8080

### Publishing Docker images:

**1. Migrate database (for production, if needed)**

```sh
docker-compose --env-file .env up --build migrate
```

**2. Push to the registry**

```sh
export DOCKER_PASSWORD=<PASSWORD??> && bash ./docker-deploy.sh
```

### API:

- [x] http://localhost:8080/v1/staff/register
- [x] http://localhost:8080/v1/staff/login
- [ ] http://localhost:8080/v1/product (POST)
- [ ] http://localhost:8080/v1/product (GET)
- [ ] http://localhost:8080/v1/product/{id} (PUT)
- [ ] http://localhost:8080/v1/product/{id} (DELETE)
- [x] http://localhost:8080/v1/product/customer
- [x] http://localhost:8080/v1/customer/register
- [x] http://localhost:8080/v1/customer
- [ ] http://localhost:8080/v1/product/checkout
- [x] http://localhost:8080/v1/product/checkout/history
