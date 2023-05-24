# Gitea Rocky Linux 8 - rootless

[howtoforge](https://www.howtoforge.com/how-to-install-gitea-code-hosting-using-docker-on-rocky-linux-8/)

> Prerequisites

[disable selinux](https://linuxize.com/post/how-to-disable-selinux-on-centos-7/)

```bash
# disable selinux
sudo sed -i s/SELINUX=enforcing/SELINUX=disabled/ /etc/selinux/config
sestatus
sudo setenforce 0
sudo reboot now

# update and install tools
sudo dnf update
sudo dnf install yum-utils nano curl -y

```

> Check and Configure Firewall

```bash
sudo firewall-cmd --state
sudo firewall-cmd --permanent --list-services

sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --permanent --add-port=2222/tcp

sudo firewall-cmd --permanent --list-all
sudo firewall-cmd --reload
```

> Install Docker and Configure Docker

```bash
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

sudo dnf install docker-ce docker-ce-cli containerd.io -y
sudo systemctl enable docker --now
sudo systemctl status docker
sudo usermod -aG docker $(whoami) # reload session
```

> Install Docker Compose

```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```

> Configure and Install Gitea

```bash
timedatectl
sudo timedatectl set-timezone Asia/Jerusalem

mkdir ~/gitea-docker
cd ~/gitea-docker
mkdir {gitea,postgres}

# check user and group id for docker-compose file
echo $(id -u)
echo $(id -g)

nano docker-compose.yml
docker-compose up -d
```

> Create a Certificate (self-signed)

```bash
openssl req -newkey rsa:2048 -nodes -keyout key.pem -x509 -days 365 -out cert.pem -config - <<EOF
[req]
default_bits = 2048
default_md = sha256
prompt = no
encrypt_key = no
distinguished_name = req_distinguished_name
req_extensions = req_ext

[req_distinguished_name]
CN = gitea.example.local
O = gitea example LTD.
L = Tel Aviv-Jaffa
ST = Israel
C = IL
emailAddress = info@gitea.example.local

[req_ext]
subjectAltName = DNS:gitea.example.local, DNS:blog.gitea.example.local, DNS:www.gitea.example.local
EOF

```

> Install and Configure Nginx

```bash
sudo dnf install nginx -y
sudo systemctl enable nginx --now

# as root
cat > /etc/nginx/nginx.conf <<EOF
user nginx;
worker_processes auto;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;

events {
    worker_connections 1024;
}

http {
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                    '$status $body_bytes_sent "$http_referer" '
                    '"$http_user_agent" "$http_x_forwarded_for"';

    access_log /var/log/nginx/access.log main;

    sendfile on;
    tcp_nopush on;
    tcp_nodelay on;
    keepalive_timeout 65;
    types_hash_max_size 2048;

    include /etc/nginx/mime.types;
    default_type application/octet-stream;

    server {
        listen 80 default_server;
        server_name _;
        return 301 https://\$host\$request_uri;
    }

    server {
        listen 443 ssl http2;
        server_name _;

        ssl_certificate /home/vagrant/gitea-docker/cert.pem;
        ssl_certificate_key /home/vagrant/gitea-docker/key.pem;
        # ssl_trusted_certificate /etc/letsencrypt/live/gitea.example.local/chain.pem;
        ssl_session_timeout 1d;
        ssl_session_cache shared:MozSSL:10m;
        ssl_session_tickets off;
        # ssl_stapling on;
        # ssl_stapling_verify on;
        # ssl_dhparam /etc/ssl/certs/dhparam.pem;

        ssl_protocols TLSv1.2 TLSv1.3;
        ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384;

        access_log /var/log/nginx/gitea.example.local.access.log main;
        error_log /var/log/nginx/gitea.example.local.error.log;

        location / {
            client_max_body_size 100M;
            proxy_pass http://localhost:3000;
            # proxy_set_header Host ;
            # proxy_set_header X-Real-IP ;
            # proxy_set_header X-Forwarded-For ;
            # proxy_set_header X-Forwarded-Proto ;
        }
    }
}
EOF

systemctl restart nginx.service
```
