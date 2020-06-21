import * as pulumi from "@pulumi/pulumi";
import * as ns1 from "@pulumi/ns1";

const myZone = new ns1.Zone("demo-zone", {
    zone: "demo-ts.pulumi.io"
});

export const zoneName = myZone.zone;
