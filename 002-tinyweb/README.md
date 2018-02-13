# Docker example: Build a tiny web server

```bash
docker build –t tinyweb .
docker run –d –p 8080:80 tinyweb
```

Navigate to [http://localhost:8080](http://localhost:8080)