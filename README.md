# Container updater

As the name says, this is a (currently only for docker) basic container updater.

**This is currently a very basic MVP/proof of concept that I created for personal use!**

## How to
This (currently) only supports starting/replacing container images.

You start/replace images by sending a `POST` request with a json payload that looks like this (this is probably going to change):
```json
{
    "Image": "some/image:tag",
	"PortBindings": [
		{
			"HostIP": "127.0.0.1",
			"HostPort": "1234",
			"ContainerPort": "3456"
		}
	]
}
```

This is equivalent to running:
```sh
docker run -d -p "127.0.0.1:1234:3456" some/image:tag
```

