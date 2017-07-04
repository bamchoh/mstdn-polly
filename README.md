# mstdn-polly
mstdn-polly is to speech text via go-mastodon stream json file

# go-mastodon

https://github.com/mattn/go-mastodon

# Usage

After start mstdn-polly, you need to configure access key / secret key in mstdn-poll.yml to access Amazon polly.

For configuration, please refer to https://github.com/bamchoh/pollydent#configuration-yaml-file

This example is using mstdn which is CLI that is including in go-mastodon.
```
$ mstdn stream --simplejson | mstdn-polly
```

# Install

```
$ go get github.com/bamchoh/mstdn-polly
```

# Licence

MIT

# Author

Yoshihiko Yamanaka (a.k.a. bamchoh)
