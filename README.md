# libdetectcloud

Originally forked from https://github.com/perlogix/libdetectcloud, adding VMware detection and potentially other vendors

http.Client timeout is set to `300ms`. Sometimes hitting the metadata service to fast will return empty instead of the cloud provider detected.

```go
    package main

    import (
    	"fmt"
    	"gitlab.com/taskfitio/lib/detectcloud"
    )

    func main() {

        // detectcloud.Detect() will return an empty string or
        // Amazon Web Services, Microsoft Azure, Digital Ocean
        // Google Compute Engine, OpenStack, SoftLayer, Vultr
        // K8S Container, Container

    	fmt.Println(detectcloud.Detect())

    }
```
