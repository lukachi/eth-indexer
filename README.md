# ETH indexer svc

## Getting started

### Install tools

```bash
  npm install -g @quobix/vacuum @typespec/compiler
```

### Generate openapi.yaml

```bash
  cd docs/
```

```bash
  npm install
```

Add spec to [components](./docs/components)

import them at [main](./docs/main.tsp)

update version at [package.json](./docs/package.json)

Build openapi file:

```bash
  npm run build
```

Lint generated file:

```bash
npm run lint
```

Generate html file (`redoc-static.html`):

```bash
  npm run build-docs
```

### Generate Resources

```bash
  cd ../
```

```bash
  sh ./generate-resources.sh
```

Check generated resources [here](./resources/generated.go)

## Work with DB

### Install tools

```bash
  brew install golang-migrate
```

```bash
    brew install xo/xo/xo
```

### Run migrations

```bash
    export $(grep -v '^#' .env | xargs) && migrate -path internal/assets/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up
```

### Generate models

```bash
    sh ./generate-xo-models.sh
```
