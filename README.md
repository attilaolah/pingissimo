# Pingissimo

A dead-simple Google App Engine application scaffold that can be used
out-of-the-box for periodically pinging websites. Just create a Google App
Engine application, update the App ID in `app.yaml`, and deploy it.

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

Then run `appcfg.py update_cron` to deploy the crons.

## The API

The API is as simple as it can get. Send a GET request to `/ping/{method}` with
the `url` parameter being the website that you want to ping. The `url`
parameter can be repeated. `/ping` is an alias for `/ping/head`.

## Future Features

So far I'm not planning on adding any. However, it would be intersting to add
some more features, like:

* Custom header support
* `POST` payload support
* Using the Datastore to persist all jobs, and have a single cron that executes them
* ICMP message support (using outbound sockets)
* User accounts with a dashboard and ping/latency history
* Email alerts on failure

I might implement these when/if I need them. Feel free to fork and help out.
