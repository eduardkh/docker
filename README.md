# docker build and config files

>run a stack from yaml file (has to bee in swarm mode)

`docker stack deploy --compose-file docker-compose.yml jenkins`

>build an image (from Dockerfile)

`docker build -t eduardkh/go_web_test:001 .`

>run the newly built container

`docker run -d -p 8080:8080 eduardkh/go_web_test:001`