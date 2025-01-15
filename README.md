# microservice-auth

1. Read .env file to environments

```bash
export $(xargs < .env.local)
```

_ENV is used then as a config name. Possible ENV values are now stage and prod as these configs are now in the repository._

2. Make sure docker network service-net is in place for microservices communication. If none exists, then create network:

```bash
task docker:network
```

3. Build image

```bash
task docker:build
```

4. To deploy Auth Service:

```bash
task docker:deploy
```

5. To stop Auth Service:

```bash
task docker:stop
```
