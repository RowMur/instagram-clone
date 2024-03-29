# insta-go

## Development

### GraphQL

All GraphQL related code lives in `/internal/graph/`. To make changes, modify the schema file and then run the generate script.

```bash
go run github.com/99designs/gqlgen 
```

### SQL

#### Migrations

Load the local `.env` file into the terminal session.

```bash
set -o allexport; source .env; set +o allexport
```

Navigate into the `schema` directory.

```bash
cd sql/schema
```

Perform the migrations.

```bash
goose postgres $DB_CONNECTION_STRING up 
```

#### Generate Go for Queries

In the project root...

```bash
sqlc generate
```
