// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class GetBillingUsageByNetworkDaily {
    /**
     * @return Clean queries for this day.
     * 
     */
    private Integer cleanQueries;
    /**
     * @return DDoS queries for this day.
     * 
     */
    private Integer ddosQueries;
    /**
     * @return NXD responses for this day.
     * 
     */
    private Integer nxdResponses;
    /**
     * @return The timestamp for the day.
     * 
     */
    private Integer timestamp;

    private GetBillingUsageByNetworkDaily() {}
    /**
     * @return Clean queries for this day.
     * 
     */
    public Integer cleanQueries() {
        return this.cleanQueries;
    }
    /**
     * @return DDoS queries for this day.
     * 
     */
    public Integer ddosQueries() {
        return this.ddosQueries;
    }
    /**
     * @return NXD responses for this day.
     * 
     */
    public Integer nxdResponses() {
        return this.nxdResponses;
    }
    /**
     * @return The timestamp for the day.
     * 
     */
    public Integer timestamp() {
        return this.timestamp;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBillingUsageByNetworkDaily defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer cleanQueries;
        private Integer ddosQueries;
        private Integer nxdResponses;
        private Integer timestamp;
        public Builder() {}
        public Builder(GetBillingUsageByNetworkDaily defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cleanQueries = defaults.cleanQueries;
    	      this.ddosQueries = defaults.ddosQueries;
    	      this.nxdResponses = defaults.nxdResponses;
    	      this.timestamp = defaults.timestamp;
        }

        @CustomType.Setter
        public Builder cleanQueries(Integer cleanQueries) {
            if (cleanQueries == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageByNetworkDaily", "cleanQueries");
            }
            this.cleanQueries = cleanQueries;
            return this;
        }
        @CustomType.Setter
        public Builder ddosQueries(Integer ddosQueries) {
            if (ddosQueries == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageByNetworkDaily", "ddosQueries");
            }
            this.ddosQueries = ddosQueries;
            return this;
        }
        @CustomType.Setter
        public Builder nxdResponses(Integer nxdResponses) {
            if (nxdResponses == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageByNetworkDaily", "nxdResponses");
            }
            this.nxdResponses = nxdResponses;
            return this;
        }
        @CustomType.Setter
        public Builder timestamp(Integer timestamp) {
            if (timestamp == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageByNetworkDaily", "timestamp");
            }
            this.timestamp = timestamp;
            return this;
        }
        public GetBillingUsageByNetworkDaily build() {
            final var _resultValue = new GetBillingUsageByNetworkDaily();
            _resultValue.cleanQueries = cleanQueries;
            _resultValue.ddosQueries = ddosQueries;
            _resultValue.nxdResponses = nxdResponses;
            _resultValue.timestamp = timestamp;
            return _resultValue;
        }
    }
}
