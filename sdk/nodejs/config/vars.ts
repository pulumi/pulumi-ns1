// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("ns1");

/**
 * The ns1 API key, this is required
 */
export let apikey: string | undefined = __config.get("apikey");
export let enableDdi: boolean | undefined = __config.getObject<boolean>("enableDdi");
export let endpoint: string | undefined = __config.get("endpoint");
export let ignoreSsl: boolean | undefined = __config.getObject<boolean>("ignoreSsl");
export let rateLimitParallelism: number | undefined = __config.getObject<number>("rateLimitParallelism");
