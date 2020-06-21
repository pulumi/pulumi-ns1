import pulumi
import pulumi_ns1 as ns1

zone = ns1.Zone("my-demo-zone",
                zone="demo-py.pulumi.io")

pulumi.export("zone_name", zone.zone)
