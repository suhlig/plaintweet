# plaintweet

* Provides a plaintext representation of a tweet
* Idea taken from [nerab/plaintweet](https://github.com/nerab/plaintweet) and rewritten in Go.

# Configuration

The environment variables `$TWITTER_CONSUMER_KEY` and `$TWITTER_CONSUMER_SECRET` are required. Get your key at [developer.twitter.com](https://developer.twitter.com/).

# Print

`plaintweet print` prints a plain-text representation of a single tweet. Provide a tweet's URL or its number as an argument.

# Serve

`plaintweet serve` serves a plain-text representation of a single tweet via HTTP. It will start an HTTP server listening on `$PORT` (defaults to `8080`) and provide the same functionality as `print` (see above).

# Run

* Grab the [latest release](https://github.com/suhlig/plaintweet/releases/latest), unpack it and run `plaintweet` for more information.
* There is also a Docker image, run it with:

  ```command
  $ docker run --env TWITTER_CONSUMER_KEY --env TWITTER_CONSUMER_SECRET -it suhlig/plaintweet
  ```

# Bonus

This application is intended for use in my course "[Web Services](https://ws.uhlig.it/)" at [DHBW](https://www.ravensburg.dhbw.de/studienangebot/bachelor-studiengaenge/informatik). For this purpose, it has some endpoints that showcase Kubernetes' [liveness and readiness probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/):

* `/liveness` will generally return `200`, unless the environment variable `MAX_UPTIME` was set to a value [`time.ParseDuration`](https://pkg.go.dev/time#ParseDuration) accepts and the given time since server start has elapsed. Other paths will still work, but `/liveness` will return `500` thereafter.
* `/readiness` will return `200` if the authentication with Twitter is successful, otherwise `500` will be returned.
