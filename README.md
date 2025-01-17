# Youless Prometheus Exporter

You can run this to export metrics from a YouLess device to Prometheus and then create graphs in Grafana.

Tested with:
- YouLess LS120, Firmware 1.6.3-EL

## Usage

```
Usage of ./youless-go:
  -listen string
    	listen address (default ":8002")
  -url string
    	URL base for YouLess (default "http://192.168.1.20")
```

## Metrics

These are the metrics, numbers all set to 0 in the example for privacy reasons.

```
$ curl -s http://127.0.0.1:8002/metrics
youless_time 0
youless_power 0
youless_netto 0.0
youless_times0 0
youless_total 0.0
youless_powers0 0
youless_p1 0.0
youless_p2 0.0
youless_n1 0.0
youless_n2 0.0
youless_gas 0.0
youless_gas_timestamp 0
youless_water 0.0
youless_water_timestamp 0
youless_tarif 0
youless_current1 0.0
youless_current2 0.0
youless_current3 0.0
youless_voltage1 0.0
youless_voltage2 0.0
youless_voltage3 0.0
youless_power1 0
youless_power2 0
youless_power3 0
```

## Resources

- https://youless.nl
- https://wiki.td-er.nl/index.php?title=YouLess
