# docker build and config files

>run a stack from yaml file (has to bee in swarm mode)

`docker stack deploy --compose-file docker-compose.yml jenkins`

>build an image (from Dockerfile)

`docker build -t eduardkh/go_web_test:001 .`

>run the newly built container (multi-stage)

`docker build --target Builder -t eduardkh/go_web_test_builder:001 .`

`docker build -t eduardkh/go_web_test_final:001 .`