https://github.com/openfaas/faasd
https://blog.alexellis.io/faasd-for-lightweight-serverless/

Grafana
https://willschenk.com/articles/2021/installing_faasd/
ssh -L 3000:127.0.0.1:3000 pi@raspberrypi
http://localhost:3000/

InfluxDB
https://github.com/alexellis/growlab/tree/master/bme280-logger
ssh -L 8086:127.0.0.1:8086 pi@raspberrypi
http://localhost:8086

```cli
curl -G 'http://localhost:8086/query?pretty=true' -u admin:PASSWORD --data-urlencode "db=defaultdb" --data-urlencode "q=SELECT \"value\" FROM \"cpu_load_short\" WHERE \"region\"='us-west'"
{
    "results": [
        {
            "statement_id": 0
        }
    ]
}
```

# Build send-tweet func

```cli
faas-cli publish -f send-tweet.yml --platforms linux/arm/v7
```

# Publish

```cli
faas-cli deploy -f send-tweet.yml --gateway http://192.168.1.42:8080
```

# Debug

```cli
faas-cli logs send-tweet --gateway http://192.168.1.42:8080
```
