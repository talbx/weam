# weam
weam (<b>we</b>bex-<b>a</b>lert<b>m</b>anager) is a simple middleware for transportation of prometheus alertmanager alerts into cisco webex.


## Get started
First, provide a `config.yaml` with webex apiKey and the receiver

`config.yaml`:
```yaml
---
apiKey: <apiKey>
receiver: <receiverId> # webex room or person id
```

and then point your alertmanager to the HTTP endpoint. When using prometheus-operator CRDs it can look like this:

`alertmanagerConfig.yaml`:
```yaml
---
apiVersion: monitoring.coreos.com/v1alpha1
kind: AlertmanagerConfig
metadata:
  name: my-alertmanager-config
spec:
  receivers:
    - name: "webex"
      webhookConfigs:
        - sendResolved: true
          url: weam.example.com/notify
  route:
    groupBy: [namespace, job]
    groupWait: 5m
    groupInterval: 5m
    repeatInterval: 5m
    receiver: "null"
    routes:
      - matchers:
        - name: namespace
          matchType: =
          value: some-special-namespace
        receiver: webex
```

the middleware exposes the HTTP-POST endpoint on `:2000/notify`