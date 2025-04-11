# microservice-auth

## Preparation

1. Copy .env.example file to .env and set environment variables

```bash
cp .env.example .env.{ENV}
```

## Production

1. Read .env file to environments

```bash
export ENV=prod
```

2. Make sure docker network service-net is in place for microservices communication. If none exists, then create network:

```bash
task docker:prod:network
```

3. Build image

```bash
task docker:prod:build
```

4. To deploy Auth Service:

```bash
task docker:prod:up
```

5. To stop Auth Service:

```bash
task docker:prod:stop
```


## Development
1. Read .env file to environments

```bash
export ENV=local
```

2. Make sure docker network service-net is in place for microservices communication. If none exists, then create network:

```bash
task docker:dev:network
```

3. Build image

```bash
task docker:dev:build
```

4. To deploy Auth Service:

```bash
task docker:dev:up
```

5. To stop Auth Service:

```bash
task docker:dev:stop
```

