// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ns1.outputs.GetZoneSecondary;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetZoneResult {
    private @Nullable List<Integer> additionalPorts;
    /**
     * @return List of additional IPv4 addresses for the primary
     * zone.
     * 
     */
    private @Nullable List<String> additionalPrimaries;
    /**
     * @return Authoritative Name Servers.
     * 
     */
    private String dnsServers;
    /**
     * @return Whether or not DNSSEC is enabled for the zone.
     * 
     */
    private Boolean dnssec;
    /**
     * @return The SOA Expiry.
     * 
     */
    private Integer expiry;
    /**
     * @return The SOA Hostmaster.
     * 
     */
    private String hostmaster;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The linked target zone.
     * 
     */
    private String link;
    /**
     * @return List of network IDs (`int`) for which the zone should be made
     * available. Default is network 0, the primary NSONE Global Network.
     * 
     */
    private List<Integer> networks;
    /**
     * @return The SOA NX TTL.
     * 
     */
    private Integer nxTtl;
    /**
     * @return The primary zones&#39; IPv4 address.
     * 
     */
    private String primary;
    private @Nullable Integer primaryPort;
    /**
     * @return The SOA Refresh.
     * 
     */
    private Integer refresh;
    /**
     * @return The SOA Retry.
     * 
     */
    private Integer retry;
    /**
     * @return List of secondary servers. Secondaries is
     * documented below.
     * 
     */
    private List<GetZoneSecondary> secondaries;
    /**
     * @return The SOA TTL.
     * 
     */
    private Integer ttl;
    private String zone;

    private GetZoneResult() {}
    public List<Integer> additionalPorts() {
        return this.additionalPorts == null ? List.of() : this.additionalPorts;
    }
    /**
     * @return List of additional IPv4 addresses for the primary
     * zone.
     * 
     */
    public List<String> additionalPrimaries() {
        return this.additionalPrimaries == null ? List.of() : this.additionalPrimaries;
    }
    /**
     * @return Authoritative Name Servers.
     * 
     */
    public String dnsServers() {
        return this.dnsServers;
    }
    /**
     * @return Whether or not DNSSEC is enabled for the zone.
     * 
     */
    public Boolean dnssec() {
        return this.dnssec;
    }
    /**
     * @return The SOA Expiry.
     * 
     */
    public Integer expiry() {
        return this.expiry;
    }
    /**
     * @return The SOA Hostmaster.
     * 
     */
    public String hostmaster() {
        return this.hostmaster;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The linked target zone.
     * 
     */
    public String link() {
        return this.link;
    }
    /**
     * @return List of network IDs (`int`) for which the zone should be made
     * available. Default is network 0, the primary NSONE Global Network.
     * 
     */
    public List<Integer> networks() {
        return this.networks;
    }
    /**
     * @return The SOA NX TTL.
     * 
     */
    public Integer nxTtl() {
        return this.nxTtl;
    }
    /**
     * @return The primary zones&#39; IPv4 address.
     * 
     */
    public String primary() {
        return this.primary;
    }
    public Optional<Integer> primaryPort() {
        return Optional.ofNullable(this.primaryPort);
    }
    /**
     * @return The SOA Refresh.
     * 
     */
    public Integer refresh() {
        return this.refresh;
    }
    /**
     * @return The SOA Retry.
     * 
     */
    public Integer retry() {
        return this.retry;
    }
    /**
     * @return List of secondary servers. Secondaries is
     * documented below.
     * 
     */
    public List<GetZoneSecondary> secondaries() {
        return this.secondaries;
    }
    /**
     * @return The SOA TTL.
     * 
     */
    public Integer ttl() {
        return this.ttl;
    }
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetZoneResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<Integer> additionalPorts;
        private @Nullable List<String> additionalPrimaries;
        private String dnsServers;
        private Boolean dnssec;
        private Integer expiry;
        private String hostmaster;
        private String id;
        private String link;
        private List<Integer> networks;
        private Integer nxTtl;
        private String primary;
        private @Nullable Integer primaryPort;
        private Integer refresh;
        private Integer retry;
        private List<GetZoneSecondary> secondaries;
        private Integer ttl;
        private String zone;
        public Builder() {}
        public Builder(GetZoneResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.additionalPorts = defaults.additionalPorts;
    	      this.additionalPrimaries = defaults.additionalPrimaries;
    	      this.dnsServers = defaults.dnsServers;
    	      this.dnssec = defaults.dnssec;
    	      this.expiry = defaults.expiry;
    	      this.hostmaster = defaults.hostmaster;
    	      this.id = defaults.id;
    	      this.link = defaults.link;
    	      this.networks = defaults.networks;
    	      this.nxTtl = defaults.nxTtl;
    	      this.primary = defaults.primary;
    	      this.primaryPort = defaults.primaryPort;
    	      this.refresh = defaults.refresh;
    	      this.retry = defaults.retry;
    	      this.secondaries = defaults.secondaries;
    	      this.ttl = defaults.ttl;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder additionalPorts(@Nullable List<Integer> additionalPorts) {
            this.additionalPorts = additionalPorts;
            return this;
        }
        public Builder additionalPorts(Integer... additionalPorts) {
            return additionalPorts(List.of(additionalPorts));
        }
        @CustomType.Setter
        public Builder additionalPrimaries(@Nullable List<String> additionalPrimaries) {
            this.additionalPrimaries = additionalPrimaries;
            return this;
        }
        public Builder additionalPrimaries(String... additionalPrimaries) {
            return additionalPrimaries(List.of(additionalPrimaries));
        }
        @CustomType.Setter
        public Builder dnsServers(String dnsServers) {
            this.dnsServers = Objects.requireNonNull(dnsServers);
            return this;
        }
        @CustomType.Setter
        public Builder dnssec(Boolean dnssec) {
            this.dnssec = Objects.requireNonNull(dnssec);
            return this;
        }
        @CustomType.Setter
        public Builder expiry(Integer expiry) {
            this.expiry = Objects.requireNonNull(expiry);
            return this;
        }
        @CustomType.Setter
        public Builder hostmaster(String hostmaster) {
            this.hostmaster = Objects.requireNonNull(hostmaster);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder link(String link) {
            this.link = Objects.requireNonNull(link);
            return this;
        }
        @CustomType.Setter
        public Builder networks(List<Integer> networks) {
            this.networks = Objects.requireNonNull(networks);
            return this;
        }
        public Builder networks(Integer... networks) {
            return networks(List.of(networks));
        }
        @CustomType.Setter
        public Builder nxTtl(Integer nxTtl) {
            this.nxTtl = Objects.requireNonNull(nxTtl);
            return this;
        }
        @CustomType.Setter
        public Builder primary(String primary) {
            this.primary = Objects.requireNonNull(primary);
            return this;
        }
        @CustomType.Setter
        public Builder primaryPort(@Nullable Integer primaryPort) {
            this.primaryPort = primaryPort;
            return this;
        }
        @CustomType.Setter
        public Builder refresh(Integer refresh) {
            this.refresh = Objects.requireNonNull(refresh);
            return this;
        }
        @CustomType.Setter
        public Builder retry(Integer retry) {
            this.retry = Objects.requireNonNull(retry);
            return this;
        }
        @CustomType.Setter
        public Builder secondaries(List<GetZoneSecondary> secondaries) {
            this.secondaries = Objects.requireNonNull(secondaries);
            return this;
        }
        public Builder secondaries(GetZoneSecondary... secondaries) {
            return secondaries(List.of(secondaries));
        }
        @CustomType.Setter
        public Builder ttl(Integer ttl) {
            this.ttl = Objects.requireNonNull(ttl);
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            this.zone = Objects.requireNonNull(zone);
            return this;
        }
        public GetZoneResult build() {
            final var o = new GetZoneResult();
            o.additionalPorts = additionalPorts;
            o.additionalPrimaries = additionalPrimaries;
            o.dnsServers = dnsServers;
            o.dnssec = dnssec;
            o.expiry = expiry;
            o.hostmaster = hostmaster;
            o.id = id;
            o.link = link;
            o.networks = networks;
            o.nxTtl = nxTtl;
            o.primary = primary;
            o.primaryPort = primaryPort;
            o.refresh = refresh;
            o.retry = retry;
            o.secondaries = secondaries;
            o.ttl = ttl;
            o.zone = zone;
            return o;
        }
    }
}