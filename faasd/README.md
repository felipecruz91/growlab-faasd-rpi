Update your `/var/lib/faasd/docker-compose.yml` with the value of the `docker-compose.yml` file and then restart `faasd`.

```cli
sudo systemctl daemon-reload && sudo systemctl restart faasd
```