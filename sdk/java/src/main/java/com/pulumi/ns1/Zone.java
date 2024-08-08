// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.ZoneArgs;
import com.pulumi.ns1.inputs.ZoneState;
import com.pulumi.ns1.outputs.ZoneSecondary;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * ```sh
 * $ pulumi import ns1:index/zone:Zone &lt;name&gt; &lt;zone&gt;`
 * ```
 * 
 * So for the example above:
 * 
 * ```sh
 * $ pulumi import ns1:index/zone:Zone example terraform.example.io`
 * ```
 * 
 */
@ResourceType(type="ns1:index/zone:Zone")
public class Zone extends com.pulumi.resources.CustomResource {
    @Export(name="additionalPorts", refs={List.class,Integer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Integer>> additionalPorts;

    public Output<Optional<List<Integer>>> additionalPorts() {
        return Codegen.optional(this.additionalPorts);
    }
    /**
     * List of additional IPv4 addresses for the primary
     * zone. Conflicts with `secondaries`.
     * 
     */
    @Export(name="additionalPrimaries", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> additionalPrimaries;

    /**
     * @return List of additional IPv4 addresses for the primary
     * zone. Conflicts with `secondaries`.
     * 
     */
    public Output<Optional<List<String>>> additionalPrimaries() {
        return Codegen.optional(this.additionalPrimaries);
    }
    @Export(name="autogenerateNsRecord", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autogenerateNsRecord;

    public Output<Optional<Boolean>> autogenerateNsRecord() {
        return Codegen.optional(this.autogenerateNsRecord);
    }
    /**
     * (Computed) Authoritative Name Servers.
     * 
     */
    @Export(name="dnsServers", refs={String.class}, tree="[0]")
    private Output<String> dnsServers;

    /**
     * @return (Computed) Authoritative Name Servers.
     * 
     */
    public Output<String> dnsServers() {
        return this.dnsServers;
    }
    /**
     * Whether or not DNSSEC is enabled for the zone.
     * Note that DNSSEC must be enabled on the account by support for this to be set
     * to `true`.
     * 
     */
    @Export(name="dnssec", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> dnssec;

    /**
     * @return Whether or not DNSSEC is enabled for the zone.
     * Note that DNSSEC must be enabled on the account by support for this to be set
     * to `true`.
     * 
     */
    public Output<Boolean> dnssec() {
        return this.dnssec;
    }
    /**
     * The SOA Expiry. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    @Export(name="expiry", refs={Integer.class}, tree="[0]")
    private Output<Integer> expiry;

    /**
     * @return The SOA Expiry. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    public Output<Integer> expiry() {
        return this.expiry;
    }
    /**
     * (Computed) The SOA Hostmaster.
     * 
     */
    @Export(name="hostmaster", refs={String.class}, tree="[0]")
    private Output<String> hostmaster;

    /**
     * @return (Computed) The SOA Hostmaster.
     * 
     */
    public Output<String> hostmaster() {
        return this.hostmaster;
    }
    /**
     * The target zone(domain name) to link to.
     * 
     */
    @Export(name="link", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> link;

    /**
     * @return The target zone(domain name) to link to.
     * 
     */
    public Output<Optional<String>> link() {
        return Codegen.optional(this.link);
    }
    /**
     * List of network IDs for which the zone is
     * available. If no network is provided, the zone will be created in network 0,
     * the primary NS1 Global Network.
     * 
     */
    @Export(name="networks", refs={List.class,Integer.class}, tree="[0,1]")
    private Output<List<Integer>> networks;

    /**
     * @return List of network IDs for which the zone is
     * available. If no network is provided, the zone will be created in network 0,
     * the primary NS1 Global Network.
     * 
     */
    public Output<List<Integer>> networks() {
        return this.networks;
    }
    /**
     * The SOA NX TTL. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    @Export(name="nxTtl", refs={Integer.class}, tree="[0]")
    private Output<Integer> nxTtl;

    /**
     * @return The SOA NX TTL. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    public Output<Integer> nxTtl() {
        return this.nxTtl;
    }
    /**
     * The primary zones&#39; IPv4 address. This makes the zone a
     * secondary. Conflicts with `secondaries`.
     * 
     */
    @Export(name="primary", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> primary;

    /**
     * @return The primary zones&#39; IPv4 address. This makes the zone a
     * secondary. Conflicts with `secondaries`.
     * 
     */
    public Output<Optional<String>> primary() {
        return Codegen.optional(this.primary);
    }
    @Export(name="primaryPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> primaryPort;

    public Output<Integer> primaryPort() {
        return this.primaryPort;
    }
    /**
     * The SOA Refresh. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    @Export(name="refresh", refs={Integer.class}, tree="[0]")
    private Output<Integer> refresh;

    /**
     * @return The SOA Refresh. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    public Output<Integer> refresh() {
        return this.refresh;
    }
    /**
     * The SOA Retry. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    @Export(name="retry", refs={Integer.class}, tree="[0]")
    private Output<Integer> retry;

    /**
     * @return The SOA Retry. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    public Output<Integer> retry() {
        return this.retry;
    }
    /**
     * List of secondary servers. This makes the zone a
     * primary. Conflicts with `primary` and `additional_primaries`.
     * Secondaries is documented below.
     * 
     */
    @Export(name="secondaries", refs={List.class,ZoneSecondary.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ZoneSecondary>> secondaries;

    /**
     * @return List of secondary servers. This makes the zone a
     * primary. Conflicts with `primary` and `additional_primaries`.
     * Secondaries is documented below.
     * 
     */
    public Output<Optional<List<ZoneSecondary>>> secondaries() {
        return Codegen.optional(this.secondaries);
    }
    /**
     * map of tags in the form of `&#34;key&#34; = &#34;value&#34;` where both key and value are strings
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return map of tags in the form of `&#34;key&#34; = &#34;value&#34;` where both key and value are strings
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * TSIG is documented below
     * 
     */
    @Export(name="tsig", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tsig;

    /**
     * @return TSIG is documented below
     * 
     */
    public Output<Optional<Map<String,String>>> tsig() {
        return Codegen.optional(this.tsig);
    }
    /**
     * The SOA TTL.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output<Integer> ttl;

    /**
     * @return The SOA TTL.
     * 
     */
    public Output<Integer> ttl() {
        return this.ttl;
    }
    /**
     * The domain name of the zone.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return The domain name of the zone.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Zone(java.lang.String name) {
        this(name, ZoneArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Zone(java.lang.String name, ZoneArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Zone(java.lang.String name, ZoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/zone:Zone", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Zone(java.lang.String name, Output<java.lang.String> id, @Nullable ZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/zone:Zone", name, state, makeResourceOptions(options, id), false);
    }

    private static ZoneArgs makeArgs(ZoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ZoneArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Zone get(java.lang.String name, Output<java.lang.String> id, @Nullable ZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Zone(name, id, state, options);
    }
}
