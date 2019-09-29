### Dockerfile

```
docker build -t ma-trust-engine . --no-cache

docker run -it --rm -p 4439:4439 --name running-ma-trust-engine ma-trust-engine:latest
```
