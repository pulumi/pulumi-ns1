// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.GetDNSSecArgs;
import com.pulumi.ns1.inputs.GetDNSSecPlainArgs;
import com.pulumi.ns1.inputs.GetRecordArgs;
import com.pulumi.ns1.inputs.GetRecordPlainArgs;
import com.pulumi.ns1.inputs.GetZoneArgs;
import com.pulumi.ns1.inputs.GetZonePlainArgs;
import com.pulumi.ns1.outputs.GetDNSSecResult;
import com.pulumi.ns1.outputs.GetRecordResult;
import com.pulumi.ns1.outputs.GetZoneResult;
import java.util.concurrent.CompletableFuture;

public final class Ns1Functions {
    /**
     * Provides DNSSEC details about a NS1 Zone.
     * 
     * ## Example Usage
     * 
     */
    public static Output<GetDNSSecResult> getDNSSec(GetDNSSecArgs args) {
        return getDNSSec(args, InvokeOptions.Empty);
    }
    /**
     * Provides DNSSEC details about a NS1 Zone.
     * 
     * ## Example Usage
     * 
     */
    public static CompletableFuture<GetDNSSecResult> getDNSSecPlain(GetDNSSecPlainArgs args) {
        return getDNSSecPlain(args, InvokeOptions.Empty);
    }
    /**
     * Provides DNSSEC details about a NS1 Zone.
     * 
     * ## Example Usage
     * 
     */
    public static Output<GetDNSSecResult> getDNSSec(GetDNSSecArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ns1:index/getDNSSec:getDNSSec", TypeShape.of(GetDNSSecResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Provides DNSSEC details about a NS1 Zone.
     * 
     * ## Example Usage
     * 
     */
    public static CompletableFuture<GetDNSSecResult> getDNSSecPlain(GetDNSSecPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ns1:index/getDNSSec:getDNSSec", TypeShape.of(GetDNSSecResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Provides details about a NS1 Record. Use this if you would simply like to read
     * information from NS1 into your configurations. For read/write operations, you
     * should use a resource.
     * 
     * ## Example Usage
     * 
     */
    public static Output<GetRecordResult> getRecord(GetRecordArgs args) {
        return getRecord(args, InvokeOptions.Empty);
    }
    /**
     * Provides details about a NS1 Record. Use this if you would simply like to read
     * information from NS1 into your configurations. For read/write operations, you
     * should use a resource.
     * 
     * ## Example Usage
     * 
     */
    public static CompletableFuture<GetRecordResult> getRecordPlain(GetRecordPlainArgs args) {
        return getRecordPlain(args, InvokeOptions.Empty);
    }
    /**
     * Provides details about a NS1 Record. Use this if you would simply like to read
     * information from NS1 into your configurations. For read/write operations, you
     * should use a resource.
     * 
     * ## Example Usage
     * 
     */
    public static Output<GetRecordResult> getRecord(GetRecordArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ns1:index/getRecord:getRecord", TypeShape.of(GetRecordResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Provides details about a NS1 Record. Use this if you would simply like to read
     * information from NS1 into your configurations. For read/write operations, you
     * should use a resource.
     * 
     * ## Example Usage
     * 
     */
    public static CompletableFuture<GetRecordResult> getRecordPlain(GetRecordPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ns1:index/getRecord:getRecord", TypeShape.of(GetRecordResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Provides details about a NS1 Zone. Use this if you would simply like to read
     * information from NS1 into your configurations. For read/write operations, you
     * should use a resource.
     * 
     * ## Example Usage
     * 
     */
    public static Output<GetZoneResult> getZone(GetZoneArgs args) {
        return getZone(args, InvokeOptions.Empty);
    }
    /**
     * Provides details about a NS1 Zone. Use this if you would simply like to read
     * information from NS1 into your configurations. For read/write operations, you
     * should use a resource.
     * 
     * ## Example Usage
     * 
     */
    public static CompletableFuture<GetZoneResult> getZonePlain(GetZonePlainArgs args) {
        return getZonePlain(args, InvokeOptions.Empty);
    }
    /**
     * Provides details about a NS1 Zone. Use this if you would simply like to read
     * information from NS1 into your configurations. For read/write operations, you
     * should use a resource.
     * 
     * ## Example Usage
     * 
     */
    public static Output<GetZoneResult> getZone(GetZoneArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ns1:index/getZone:getZone", TypeShape.of(GetZoneResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Provides details about a NS1 Zone. Use this if you would simply like to read
     * information from NS1 into your configurations. For read/write operations, you
     * should use a resource.
     * 
     * ## Example Usage
     * 
     */
    public static CompletableFuture<GetZoneResult> getZonePlain(GetZonePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ns1:index/getZone:getZone", TypeShape.of(GetZoneResult.class), args, Utilities.withVersion(options));
    }
}
