# Docker example: Build Zork game with Go

```bash
docker build -t zork .
docker images | grep "zork"
docker run -it --name zork zork
```