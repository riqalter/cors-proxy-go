# CORS Proxy
an attempt to make a simple CORS proxy written in Go

### Local Development
```bash
go run main.go
```
or if you're using [air](https://github.com/cosmtrek/air)
```bash
air
```

### Docker
```bash
# Using docker-compose
docker-compose up -d

# Or with Docker directly
docker build -t cors-proxy .
docker run -p 8080:8080 cors-proxy
```

Server runs on `http://localhost:8080`