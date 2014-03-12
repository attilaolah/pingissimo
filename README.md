# Pingissimo

A dead-simple Google App Engine application scaffold that can be used
out-of-the-box for periodically pinging websites.

Just create a Google App Engine application, update the App ID in `app.yaml`,
and deploy it.

To ping websites, simply add cron entries to `cron.yaml`, like this:

```yaml
cron:
- description: HEAD Example every minute
  url: /ping?url=http://www.example.com
  schedule: every 1 minutes
- description: POST MyApp and Example every two hours
  url: /ping/post?url=https://www.myapp.com&url=http://www.example.com
  schedule: every 2 hours
```

The `url` parameter can be repeated. `/ping` is an alias for `/ping/head`.
