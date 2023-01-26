// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ns1.inputs.ZoneSecondaryArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneArgs extends com.pulumi.resources.ResourceArgs {

    public static final ZoneArgs Empty = new ZoneArgs();

    @Import(name="additionalPorts")
    private @Nullable Output<List<Integer>> additionalPorts;

    public Optional<Output<List<Integer>>> additionalPorts() {
        return Optional.ofNullable(this.additionalPorts);
    }

    /**
     * List of additional IPv4 addresses for the primary
     * zone. Conflicts with `secondaries`.
     * 
     */
    @Import(name="additionalPrimaries")
    private @Nullable Output<List<String>> additionalPrimaries;

    /**
     * @return List of additional IPv4 addresses for the primary
     * zone. Conflicts with `secondaries`.
     * 
     */
    public Optional<Output<List<String>>> additionalPrimaries() {
        return Optional.ofNullable(this.additionalPrimaries);
    }

    @Import(name="autogenerateNsRecord")
    private @Nullable Output<Boolean> autogenerateNsRecord;

    public Optional<Output<Boolean>> autogenerateNsRecord() {
        return Optional.ofNullable(this.autogenerateNsRecord);
    }

    /**
     * Whether or not DNSSEC is enabled for the zone.
     * Note that DNSSEC must be enabled on the account by support for this to be set
     * to `true`.
     * 
     */
    @Import(name="dnssec")
    private @Nullable Output<Boolean> dnssec;

    /**
     * @return Whether or not DNSSEC is enabled for the zone.
     * Note that DNSSEC must be enabled on the account by support for this to be set
     * to `true`.
     * 
     */
    public Optional<Output<Boolean>> dnssec() {
        return Optional.ofNullable(this.dnssec);
    }

    /**
     * The SOA Expiry. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    @Import(name="expiry")
    private @Nullable Output<Integer> expiry;

    /**
     * @return The SOA Expiry. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    public Optional<Output<Integer>> expiry() {
        return Optional.ofNullable(this.expiry);
    }

    /**
     * (Computed) The SOA Hostmaster.
     * 
     */
    @Import(name="hostmaster")
    private @Nullable Output<String> hostmaster;

    /**
     * @return (Computed) The SOA Hostmaster.
     * 
     */
    public Optional<Output<String>> hostmaster() {
        return Optional.ofNullable(this.hostmaster);
    }

    /**
     * The target zone(domain name) to link to.
     * 
     */
    @Import(name="link")
    private @Nullable Output<String> link;

    /**
     * @return The target zone(domain name) to link to.
     * 
     */
    public Optional<Output<String>> link() {
        return Optional.ofNullable(this.link);
    }

    /**
     * List of network IDs for which the zone is
     * available. If no network is provided, the zone will be created in network 0,
     * the primary NS1 Global Network.
     * 
     */
    @Import(name="networks")
    private @Nullable Output<List<Integer>> networks;

    /**
     * @return List of network IDs for which the zone is
     * available. If no network is provided, the zone will be created in network 0,
     * the primary NS1 Global Network.
     * 
     */
    public Optional<Output<List<Integer>>> networks() {
        return Optional.ofNullable(this.networks);
    }

    /**
     * The SOA NX TTL. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    @Import(name="nxTtl")
    private @Nullable Output<Integer> nxTtl;

    /**
     * @return The SOA NX TTL. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    public Optional<Output<Integer>> nxTtl() {
        return Optional.ofNullable(this.nxTtl);
    }

    /**
     * The primary zones&#39; IPv4 address. This makes the zone a
     * secondary. Conflicts with `secondaries`.
     * 
     */
    @Import(name="primary")
    private @Nullable Output<String> primary;

    /**
     * @return The primary zones&#39; IPv4 address. This makes the zone a
     * secondary. Conflicts with `secondaries`.
     * 
     */
    public Optional<Output<String>> primary() {
        return Optional.ofNullable(this.primary);
    }

    @Import(name="primaryPort")
    private @Nullable Output<Integer> primaryPort;

    public Optional<Output<Integer>> primaryPort() {
        return Optional.ofNullable(this.primaryPort);
    }

    /**
     * The SOA Refresh. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    @Import(name="refresh")
    private @Nullable Output<Integer> refresh;

    /**
     * @return The SOA Refresh. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    public Optional<Output<Integer>> refresh() {
        return Optional.ofNullable(this.refresh);
    }

    /**
     * The SOA Retry. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    @Import(name="retry")
    private @Nullable Output<Integer> retry;

    /**
     * @return The SOA Retry. Conflicts with `primary` and
     * `additional_primaries` (default must be accepted).
     * 
     */
    public Optional<Output<Integer>> retry() {
        return Optional.ofNullable(this.retry);
    }

    /**
     * List of secondary servers. This makes the zone a
     * primary. Conflicts with `primary` and `additional_primaries`.
     * Secondaries is documented below.
     * 
     */
    @Import(name="secondaries")
    private @Nullable Output<List<ZoneSecondaryArgs>> secondaries;

    /**
     * @return List of secondary servers. This makes the zone a
     * primary. Conflicts with `primary` and `additional_primaries`.
     * Secondaries is documented below.
     * 
     */
    public Optional<Output<List<ZoneSecondaryArgs>>> secondaries() {
        return Optional.ofNullable(this.secondaries);
    }

    /**
     * TSIG is documented below
     * 
     */
    @Import(name="tsig")
    private @Nullable Output<Map<String,String>> tsig;

    /**
     * @return TSIG is documented below
     * 
     */
    public Optional<Output<Map<String,String>>> tsig() {
        return Optional.ofNullable(this.tsig);
    }

    /**
     * The SOA TTL.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return The SOA TTL.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * The domain name of the zone.
     * 
     */
    @Import(name="zone", required=true)
    private Output<String> zone;

    /**
     * @return The domain name of the zone.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    private ZoneArgs() {}

    private ZoneArgs(ZoneArgs $) {
        this.additionalPorts = $.additionalPorts;
        this.additionalPrimaries = $.additionalPrimaries;
        this.autogenerateNsRecord = $.autogenerateNsRecord;
        this.dnssec = $.dnssec;
        this.expiry = $.expiry;
        this.hostmaster = $.hostmaster;
        this.link = $.link;
        this.networks = $.networks;
        this.nxTtl = $.nxTtl;
        this.primary = $.primary;
        this.primaryPort = $.primaryPort;
        this.refresh = $.refresh;
        this.retry = $.retry;
        this.secondaries = $.secondaries;
        this.tsig = $.tsig;
        this.ttl = $.ttl;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneArgs $;

        public Builder() {
            $ = new ZoneArgs();
        }

        public Builder(ZoneArgs defaults) {
            $ = new ZoneArgs(Objects.requireNonNull(defaults));
        }

        public Builder additionalPorts(@Nullable Output<List<Integer>> additionalPorts) {
            $.additionalPorts = additionalPorts;
            return this;
        }

        public Builder additionalPorts(List<Integer> additionalPorts) {
            return additionalPorts(Output.of(additionalPorts));
        }

        public Builder additionalPorts(Integer... additionalPorts) {
            return additionalPorts(List.of(additionalPorts));
        }

        /**
         * @param additionalPrimaries List of additional IPv4 addresses for the primary
         * zone. Conflicts with `secondaries`.
         * 
         * @return builder
         * 
         */
        public Builder additionalPrimaries(@Nullable Output<List<String>> additionalPrimaries) {
            $.additionalPrimaries = additionalPrimaries;
            return this;
        }

        /**
         * @param additionalPrimaries List of additional IPv4 addresses for the primary
         * zone. Conflicts with `secondaries`.
         * 
         * @return builder
         * 
         */
        public Builder additionalPrimaries(List<String> additionalPrimaries) {
            return additionalPrimaries(Output.of(additionalPrimaries));
        }

        /**
         * @param additionalPrimaries List of additional IPv4 addresses for the primary
         * zone. Conflicts with `secondaries`.
         * 
         * @return builder
         * 
         */
        public Builder additionalPrimaries(String... additionalPrimaries) {
            return additionalPrimaries(List.of(additionalPrimaries));
        }

        public Builder autogenerateNsRecord(@Nullable Output<Boolean> autogenerateNsRecord) {
            $.autogenerateNsRecord = autogenerateNsRecord;
            return this;
        }

        public Builder autogenerateNsRecord(Boolean autogenerateNsRecord) {
            return autogenerateNsRecord(Output.of(autogenerateNsRecord));
        }

        /**
         * @param dnssec Whether or not DNSSEC is enabled for the zone.
         * Note that DNSSEC must be enabled on the account by support for this to be set
         * to `true`.
         * 
         * @return builder
         * 
         */
        public Builder dnssec(@Nullable Output<Boolean> dnssec) {
            $.dnssec = dnssec;
            return this;
        }

        /**
         * @param dnssec Whether or not DNSSEC is enabled for the zone.
         * Note that DNSSEC must be enabled on the account by support for this to be set
         * to `true`.
         * 
         * @return builder
         * 
         */
        public Builder dnssec(Boolean dnssec) {
            return dnssec(Output.of(dnssec));
        }

        /**
         * @param expiry The SOA Expiry. Conflicts with `primary` and
         * `additional_primaries` (default must be accepted).
         * 
         * @return builder
         * 
         */
        public Builder expiry(@Nullable Output<Integer> expiry) {
            $.expiry = expiry;
            return this;
        }

        /**
         * @param expiry The SOA Expiry. Conflicts with `primary` and
         * `additional_primaries` (default must be accepted).
         * 
         * @return builder
         * 
         */
        public Builder expiry(Integer expiry) {
            return expiry(Output.of(expiry));
        }

        /**
         * @param hostmaster (Computed) The SOA Hostmaster.
         * 
         * @return builder
         * 
         */
        public Builder hostmaster(@Nullable Output<String> hostmaster) {
            $.hostmaster = hostmaster;
            return this;
        }

        /**
         * @param hostmaster (Computed) The SOA Hostmaster.
         * 
         * @return builder
         * 
         */
        public Builder hostmaster(String hostmaster) {
            return hostmaster(Output.of(hostmaster));
        }

        /**
         * @param link The target zone(domain name) to link to.
         * 
         * @return builder
         * 
         */
        public Builder link(@Nullable Output<String> link) {
            $.link = link;
            return this;
        }

        /**
         * @param link The target zone(domain name) to link to.
         * 
         * @return builder
         * 
         */
        public Builder link(String link) {
            return link(Output.of(link));
        }

        /**
         * @param networks List of network IDs for which the zone is
         * available. If no network is provided, the zone will be created in network 0,
         * the primary NS1 Global Network.
         * 
         * @return builder
         * 
         */
        public Builder networks(@Nullable Output<List<Integer>> networks) {
            $.networks = networks;
            return this;
        }

        /**
         * @param networks List of network IDs for which the zone is
         * available. If no network is provided, the zone will be created in network 0,
         * the primary NS1 Global Network.
         * 
         * @return builder
         * 
         */
        public Builder networks(List<Integer> networks) {
            return networks(Output.of(networks));
        }

        /**
         * @param networks List of network IDs for which the zone is
         * available. If no network is provided, the zone will be created in network 0,
         * the primary NS1 Global Network.
         * 
         * @return builder
         * 
         */
        public Builder networks(Integer... networks) {
            return networks(List.of(networks));
        }

        /**
         * @param nxTtl The SOA NX TTL. Conflicts with `primary` and
         * `additional_primaries` (default must be accepted).
         * 
         * @return builder
         * 
         */
        public Builder nxTtl(@Nullable Output<Integer> nxTtl) {
            $.nxTtl = nxTtl;
            return this;
        }

        /**
         * @param nxTtl The SOA NX TTL. Conflicts with `primary` and
         * `additional_primaries` (default must be accepted).
         * 
         * @return builder
         * 
         */
        public Builder nxTtl(Integer nxTtl) {
            return nxTtl(Output.of(nxTtl));
        }

        /**
         * @param primary The primary zones&#39; IPv4 address. This makes the zone a
         * secondary. Conflicts with `secondaries`.
         * 
         * @return builder
         * 
         */
        public Builder primary(@Nullable Output<String> primary) {
            $.primary = primary;
            return this;
        }

        /**
         * @param primary The primary zones&#39; IPv4 address. This makes the zone a
         * secondary. Conflicts with `secondaries`.
         * 
         * @return builder
         * 
         */
        public Builder primary(String primary) {
            return primary(Output.of(primary));
        }

        public Builder primaryPort(@Nullable Output<Integer> primaryPort) {
            $.primaryPort = primaryPort;
            return this;
        }

        public Builder primaryPort(Integer primaryPort) {
            return primaryPort(Output.of(primaryPort));
        }

        /**
         * @param refresh The SOA Refresh. Conflicts with `primary` and
         * `additional_primaries` (default must be accepted).
         * 
         * @return builder
         * 
         */
        public Builder refresh(@Nullable Output<Integer> refresh) {
            $.refresh = refresh;
            return this;
        }

        /**
         * @param refresh The SOA Refresh. Conflicts with `primary` and
         * `additional_primaries` (default must be accepted).
         * 
         * @return builder
         * 
         */
        public Builder refresh(Integer refresh) {
            return refresh(Output.of(refresh));
        }

        /**
         * @param retry The SOA Retry. Conflicts with `primary` and
         * `additional_primaries` (default must be accepted).
         * 
         * @return builder
         * 
         */
        public Builder retry(@Nullable Output<Integer> retry) {
            $.retry = retry;
            return this;
        }

        /**
         * @param retry The SOA Retry. Conflicts with `primary` and
         * `additional_primaries` (default must be accepted).
         * 
         * @return builder
         * 
         */
        public Builder retry(Integer retry) {
            return retry(Output.of(retry));
        }

        /**
         * @param secondaries List of secondary servers. This makes the zone a
         * primary. Conflicts with `primary` and `additional_primaries`.
         * Secondaries is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secondaries(@Nullable Output<List<ZoneSecondaryArgs>> secondaries) {
            $.secondaries = secondaries;
            return this;
        }

        /**
         * @param secondaries List of secondary servers. This makes the zone a
         * primary. Conflicts with `primary` and `additional_primaries`.
         * Secondaries is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secondaries(List<ZoneSecondaryArgs> secondaries) {
            return secondaries(Output.of(secondaries));
        }

        /**
         * @param secondaries List of secondary servers. This makes the zone a
         * primary. Conflicts with `primary` and `additional_primaries`.
         * Secondaries is documented below.
         * 
         * @return builder
         * 
         */
        public Builder secondaries(ZoneSecondaryArgs... secondaries) {
            return secondaries(List.of(secondaries));
        }

        /**
         * @param tsig TSIG is documented below
         * 
         * @return builder
         * 
         */
        public Builder tsig(@Nullable Output<Map<String,String>> tsig) {
            $.tsig = tsig;
            return this;
        }

        /**
         * @param tsig TSIG is documented below
         * 
         * @return builder
         * 
         */
        public Builder tsig(Map<String,String> tsig) {
            return tsig(Output.of(tsig));
        }

        /**
         * @param ttl The SOA TTL.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The SOA TTL.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param zone The domain name of the zone.
         * 
         * @return builder
         * 
         */
        public Builder zone(Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone The domain name of the zone.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public ZoneArgs build() {
            $.zone = Objects.requireNonNull($.zone, "expected parameter 'zone' to be non-null");
            return $;
        }
    }

}