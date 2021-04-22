# Intro

This setup is to test alerts based on loki rules. We want to be notified when we get an error in logs.

The demo system consists of the following services:

* grafana
* loki
* alertmanager
* go-logs - demo application that write logs with incrementing number every 5 seconds
* ngrok - or some other tool that can get webhooks

# Setup

## Loki log driver for Docker
We do not configure promtail to collect logs and send them to Loki. Instead, we use Loki dirver for Docker. Please, install it so logs from `go-logs` will be deliverd to Loki:

```
docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions
```

## Ngrok to receive webhook calls

We need a way to see what and if we get alert. To do this we can use `ngrok`. Install it and run:

```
ngrok http 80
```

then copy received Forwarding url like https://7ba12c5b699b.ngrok.io and paste it into [./grafana/alertmanager/config.yml](./grafana/alertmanager/config.yml) in `url`

# Run 

## Launch all services
To run everything and see logs from `go-logs`:

* cd ./grafana
* docker-compose up


## Configure Loki as a source in Grafana
To see the logs from `go-logs` in Grafana:

* open http://localhost:3000 and login under `admin` / `admin`.
* go to http://localhost:3000/datasources (Configuration)
* Click "Add data source" and add Loki using "http://loki:3100"

Now, hopefully, you can view all logs: http://localhost:3000/explore?orgId=1&left=%5B%22now-1h%22,%22now%22,%22Loki%22,%7B%22expr%22:%22%7Bcontainer_name%3D%5C%22grafana_go-logs_1%5C%22%7D%22%7D%5D

# Test

My idea is to get notification when we receive `number="6"` in logs.

Unfortunately, I was not able to even see if Loki handles the rule...
