# plaintweet

* Provides a plaintext representation of a tweet
* Idea taken from [nerab/plaintweet](https://github.com/nerab/plaintweet) and rewritten in Go.

# Configuration

The environment variables `$TWITTER_CONSUMER_KEY` and `$TWITTER_CONSUMER_SECRET` are required. Get your key at [developer.twitter.com](https://developer.twitter.com/).

# Print

`plaintweet print` prints a plain-text representation of a single tweet. Provide a tweet's URL or its number as an argument.

# Serve

`plaintweet serve` serves a plain-text representation of a single tweet via HTTP. It will start an HTTP server listening on $PORT (defaults to `8080`) and provide the same functionality as `print` (see above).

# Run

* Grab the [latest release](https://github.com/suhlig/plaintweet/releases/latest), unpack it and run `plaintweet` for more information.
* There is also a Docker image, run it with:

  ```command
  $ docker run --env TWITTER_CONSUMER_KEY --env TWITTER_CONSUMER_SECRET -it suhlig/plaintweet
  ```
