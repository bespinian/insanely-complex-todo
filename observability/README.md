# Monitoring & Observability

We are running a Grafana/Prometheus stack to get some observability into our setup. This stack also allows us to
monitor our services.

See the [main Readme](../README.md) on how to get it all set up.

* **Prometheus**: Collects metrics from different sources and keeps them for a couple of days.
* **Grafana**: Connects to Prometheus and displays its data in nice dashboards and charts.
* **Node Exporter**: Collects metrics from our host. That's either our local machine or the Kubernetes host, depending on the setup.

## Docker Compose Setup

After everything is running, visit <http://grafana.localhost:8000> and log in using `admin` as username and as password.

You can also directly access Prometheus' UI at <http://prometheus.localhost:8000>.
