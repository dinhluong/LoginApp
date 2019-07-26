> docker build -t cloud9-image:latest  c9/

> docker run --privileged -e "IP=127.0.0.1" -e "PORT=8080" --name c9 --mount type=bind,source=%GOPATH%,target=/home/ubuntu/workspace/ -d -t -p 5050:5050 -p 8080-8082:8080-8082 cloud9-image:latest
```
Can custom mount volume, instead of GOPATH, you can set your workspace on local.
```