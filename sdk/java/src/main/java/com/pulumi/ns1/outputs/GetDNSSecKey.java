// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ns1.outputs.GetDNSSecKeyDnskey;
import java.lang.Integer;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDNSSecKey {
    /**
     * @return (Computed) List of Keys. Key is documented below.
     * 
     */
    private List<GetDNSSecKeyDnskey> dnskeys;
    /**
     * @return (Computed) TTL for the Keys (int).
     * 
     */
    private Integer ttl;

    private GetDNSSecKey() {}
    /**
     * @return (Computed) List of Keys. Key is documented below.
     * 
     */
    public List<GetDNSSecKeyDnskey> dnskeys() {
        return this.dnskeys;
    }
    /**
     * @return (Computed) TTL for the Keys (int).
     * 
     */
    public Integer ttl() {
        return this.ttl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDNSSecKey defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetDNSSecKeyDnskey> dnskeys;
        private Integer ttl;
        public Builder() {}
        public Builder(GetDNSSecKey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dnskeys = defaults.dnskeys;
    	      this.ttl = defaults.ttl;
        }

        @CustomType.Setter
        public Builder dnskeys(List<GetDNSSecKeyDnskey> dnskeys) {
            if (dnskeys == null) {
              throw new MissingRequiredPropertyException("GetDNSSecKey", "dnskeys");
            }
            this.dnskeys = dnskeys;
            return this;
        }
        public Builder dnskeys(GetDNSSecKeyDnskey... dnskeys) {
            return dnskeys(List.of(dnskeys));
        }
        @CustomType.Setter
        public Builder ttl(Integer ttl) {
            if (ttl == null) {
              throw new MissingRequiredPropertyException("GetDNSSecKey", "ttl");
            }
            this.ttl = ttl;
            return this;
        }
        public GetDNSSecKey build() {
            final var _resultValue = new GetDNSSecKey();
            _resultValue.dnskeys = dnskeys;
            _resultValue.ttl = ttl;
            return _resultValue;
        }
    }
}
