### Dockerfile

```
docker build -t ma-sso-auth . --no-cache

docker run -it --rm -p 4435:4435 --name running-ma-sso-auth ma-sso-auth:latest
```
