### Dockerfile

```
docker build -t ma-moodle-proxy . --no-cache

docker run -it --rm -p 10443:10443 --name running-ma-moodle-proxy ma-moodle-proxy:latest
```
