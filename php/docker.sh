docker run -it --rm -v "$(pwd):/usr/src/myapp" -w /usr/src/myapp -p 8080:8080 php:8.2-cli php -S 0.0.0.0:8080