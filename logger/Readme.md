### Dockerfile

```
docker build -t ma-logger . --no-cache

docker run -it --rm -p 4432:4432 --name running-ma-logger ma-logger:latest
```
