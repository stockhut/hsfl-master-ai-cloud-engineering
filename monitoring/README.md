# Monitoring

# Add to minikube

Apply the config to minikube via the following command.
```shell
minikube apply -f . --recursive
```

# Access Grafana

Run the following command in a terminal to open Grafana in your browser.
```shell
minikube service -n monitoring grafana
```

# Import Dashboard

After logging in, you can import `dashboard.json` and `kubernetes-dashboard.json` to have access to our pre-built dashboards.