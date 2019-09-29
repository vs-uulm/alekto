### Dockerfile

```
docker build -t ma-policy-engine . --no-cache

docker run -it --rm -p 4438:4438 --name running-ma-policy-engine ma-policy-engine:latest
```
