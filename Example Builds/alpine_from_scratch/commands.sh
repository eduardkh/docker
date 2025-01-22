# https://hub.docker.com/layers/library/alpine/latest/images/sha256-f3240395711384fc3c07daa46cbc8d73aa5ba25ad1deb97424992760f8cb2b94
# https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/x86/

# get the file
wget https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/x86/alpine-minirootfs-3.21.2-x86.tar.gz

# build the container
docker build -t myalpine3212custom:v01 .

# run the container
docker run -ti --rm myalpine3212custom:v