# Docker Build & Push

## Build Image

```bash
# Build
docker build -t agent-harness:latest .

# Or with buildx
docker buildx build --platform linux/amd64,linux/arm64 -t your-registry/agent-harness:latest .
```

## Push to Registry

```bash
# Docker Hub
docker push your-registry/agent-harness:latest

# GHCR (GitHub Container Registry)
docker push ghcr.io/youruser/agent-harness:latest
```

## GitHub Actions

```yaml
# .github/workflows/docker.yml
name: Docker

on:
  push:
    branches: [main]
  release:
    types: [published]

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
        
      - name: Login to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}
          
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: youruser/agent-harness:latest
```

## Run

```bash
# Local
docker run -d -p 3000:3000 \
  -e OPENAI_API_KEY=xxx \
  -e GITHUB_TOKEN=xxx \
  agent-harness:latest

# Compose
docker-compose up -d
```