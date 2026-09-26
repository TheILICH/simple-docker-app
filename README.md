# Simple Go App in Docker

A small Go web server, containerized with Docker.

## Build

```bash
docker build -t simple .
```

### Run
```bash
docker run -p 8080:8080 -e APP_ENV=in-docker --name in-docker simple:latest
```


App will be running at `http://localhost:8080`.

## Env Variable

`APP_ENV` — defaults to `development` if not set.
When run in Docker (see command above), it's set to `in-docker`.

## Routes

- `/` — returns JSON with status, current `APP_ENV`, and a message
- `/health` — returns `OK`, used for health checks
