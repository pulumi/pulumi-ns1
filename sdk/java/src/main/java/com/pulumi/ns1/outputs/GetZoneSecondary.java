// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetZoneSecondary {
    /**
     * @return IPv4 address of the secondary server.
     * 
     */
    private String ip;
    /**
     * @return List of network IDs (`int`) for which the zone should be made
     * available. Default is network 0, the primary NSONE Global Network.
     * 
     */
    private List<Integer> networks;
    /**
     * @return Whether we send `NOTIFY` messages to the secondary host
     * when the zone changes. Default `false`.
     * 
     */
    private Boolean notify;
    /**
     * @return Port of the the secondary server. Default `53`.
     * 
     */
    private Integer port;

    private GetZoneSecondary() {}
    /**
     * @return IPv4 address of the secondary server.
     * 
     */
    public String ip() {
        return this.ip;
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
     * @return Whether we send `NOTIFY` messages to the secondary host
     * when the zone changes. Default `false`.
     * 
     */
    public Boolean notify_() {
        return this.notify;
    }
    /**
     * @return Port of the the secondary server. Default `53`.
     * 
     */
    public Integer port() {
        return this.port;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetZoneSecondary defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ip;
        private List<Integer> networks;
        private Boolean notify;
        private Integer port;
        public Builder() {}
        public Builder(GetZoneSecondary defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ip = defaults.ip;
    	      this.networks = defaults.networks;
    	      this.notify = defaults.notify;
    	      this.port = defaults.port;
        }

        @CustomType.Setter
        public Builder ip(String ip) {
            if (ip == null) {
              throw new MissingRequiredPropertyException("GetZoneSecondary", "ip");
            }
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder networks(List<Integer> networks) {
            if (networks == null) {
              throw new MissingRequiredPropertyException("GetZoneSecondary", "networks");
            }
            this.networks = networks;
            return this;
        }
        public Builder networks(Integer... networks) {
            return networks(List.of(networks));
        }
        @CustomType.Setter("notify")
        public Builder notify_(Boolean notify) {
            if (notify == null) {
              throw new MissingRequiredPropertyException("GetZoneSecondary", "notify");
            }
            this.notify = notify;
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            if (port == null) {
              throw new MissingRequiredPropertyException("GetZoneSecondary", "port");
            }
            this.port = port;
            return this;
        }
        public GetZoneSecondary build() {
            final var _resultValue = new GetZoneSecondary();
            _resultValue.ip = ip;
            _resultValue.networks = networks;
            _resultValue.notify = notify;
            _resultValue.port = port;
            return _resultValue;
        }
    }
}
