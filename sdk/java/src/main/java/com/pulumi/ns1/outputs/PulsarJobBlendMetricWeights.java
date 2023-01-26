// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class PulsarJobBlendMetricWeights {
    private Integer timestamp;

    private PulsarJobBlendMetricWeights() {}
    public Integer timestamp() {
        return this.timestamp;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PulsarJobBlendMetricWeights defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer timestamp;
        public Builder() {}
        public Builder(PulsarJobBlendMetricWeights defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.timestamp = defaults.timestamp;
        }

        @CustomType.Setter
        public Builder timestamp(Integer timestamp) {
            this.timestamp = Objects.requireNonNull(timestamp);
            return this;
        }
        public PulsarJobBlendMetricWeights build() {
            final var o = new PulsarJobBlendMetricWeights();
            o.timestamp = timestamp;
            return o;
        }
    }
}