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

**2. Start a Postgres instance**

```sh
docker-compose up --build postgres
```

**3. Migrate database**

```sh
docker-compose up --build migrate
```

**4. Running API**

```sh
export DB_HOST=host.docker.internal && docker-compose up --build api
```

Open http://localhost:8080

### Publishing Docker images:

**1. Prepare database (if not exists)**

```sh
docker exec -it enigmanations_postgres_container psql -h host.docker.internal -p 5430 -U postgres -c 'create database "eniqlo-store"'
```

**2. Run migration**

```sh
docker-compose up --build migrate
```

**3. Push to the registry**

```sh
# The password needs to be requested from the author, @natserract.
export DOCKER_PASSWORD="" && bash ./docker-deploy.sh
```

### Run from registry:

**1. Pull from registry** https://hub.docker.com/repository/docker/natserract/enigmanations-inventory

```sh
docker pull natserract/enigmanations-inventory:latest
```

**2. Run**

```sh
# Create a network, which allows containers to communicate
# with each other, by using their container name as a hostname
docker network create app_network

# Using env file
docker run -it --rm --network app_network \
    --env-file .env natserract/enigmanations-inventory

# Using env variable exported
docker run -it --rm --network app_network -p 8080:8080 \
    -e ENV=production \
    -e DB_HOST=host.docker.internal \
    -e DB_USERNAME=postgres \
    -e DB_PASSWORD=postgres \
    -e DB_NAME=eniqlo-store natserract/enigmanations-inventory
```

### API:

- [x] http://localhost:8080/v1/staff/register
- [x] http://localhost:8080/v1/staff/login
- [x] http://localhost:8080/v1/product (POST)
- [x] http://localhost:8080/v1/product (GET)
- [x] http://localhost:8080/v1/product/{id} (PUT)
- [x] http://localhost:8080/v1/product/{id} (DELETE)
- [x] http://localhost:8080/v1/product/customer
- [x] http://localhost:8080/v1/customer/register
- [x] http://localhost:8080/v1/customer
- [x] http://localhost:8080/v1/product/checkout
- [x] http://localhost:8080/v1/product/checkout/history
