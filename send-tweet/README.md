### Using private repos

```cli
kubectl create secret docker-registry $NAME \
    --dry-run \
    --docker-server=$SERVER \
    --docker-username=$USERNAME \
    --docker-password=$PASSWORD \
    -o yaml > config.json
```

To use private image repos, `config.json` needs to be copied to `/var/lib/faasd/.docker/config.json`.

Restart `faasd`

```cli
sudo systemctl daemon-reload && sudo systemctl restart faasd
```

===

```cli
faas-cli secret create secret-twitter-api-key \
  --from-literal=<YOUR_TWITTER_API_KEY>
```

```cli
faas-cli secret create secret-twitter-api-secret \
  --from-literal=<YOUR_TWITTER_API_SECRET>
```

```cli
faas-cli secret create secret-twitter-access-token \
  --from-literal=<YOUR_TWITTER_ACCESS_TOKEN>
```

```cli
faas-cli secret create secret-twitter-access-secret \
  --from-literal=<YOUR_TWITTER_ACCESS_SECRET>
```
