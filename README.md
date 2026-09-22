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
        // Amazon Web Services, Microsoft Azure, Digital Ocean,
        // Google Compute Engine, SoftLayer, Vultr, Oracle Cloud,
        // Alibaba Cloud, Tencent Cloud, Hetzner, K8S Container,
        // Container, VMware, Nutanix, oVirt, CloudStack, KubeVirt,
        // OpenStack, Hyper-V, Xen, VirtualBox, OpenNebula, Linode,
        // UpCloud, Huawei Cloud, KVM

    	fmt.Println(detectcloud.Detect())

    }
```
