# Docker-rootless

> add git user on the host

```bash
sudo adduser --system --shell /bin/bash --gecos 'Git Version Control' --group --disabled-password --home /home/git git
```

> SSH Container Passthrough

```bash
sudo tee -a /usr/local/bin/gitea-shell >/dev/null <<'EOF'
#!/bin/sh
/usr/bin/docker exec -i --env SSH_ORIGINAL_COMMAND="$SSH_ORIGINAL_COMMAND" gitea sh "$@"
EOF

sudo chmod +x /usr/local/bin/gitea-shell
sudo usermod -s /usr/local/bin/gitea-shell git
```

> SSH authentication on the host

```bash
sudo tee -a /etc/ssh/sshd_config >/dev/null <<'EOF'
Match User git
  AuthorizedKeysCommandUser git
  AuthorizedKeysCommand /usr/bin/docker exec -i gitea /usr/local/bin/gitea keys -c /etc/gitea/app.ini -e git -u %u -t %t -k %k
EOF
cat /etc/ssh/sshd_config
sudo systemctl restart sshd
sudo systemctl status sshd
```

> deploy docker stack as git

```bash
sudo chown -R git:git *
sudo usermod -aG docker git
sudo -u git docker stack deploy -c docker-compose.yml gitea
```

> windows ssh problem fix (ssh config file)

[issues 18528]([https://](https://github.com/go-gitea/gitea/issues/18528))
[issues 17798]([https://](https://github.com/go-gitea/gitea/issues/17798))
