// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.util.Objects;


public final class PulsarJobBlendMetricWeightsArgs extends com.pulumi.resources.ResourceArgs {

    public static final PulsarJobBlendMetricWeightsArgs Empty = new PulsarJobBlendMetricWeightsArgs();

    @Import(name="timestamp", required=true)
    private Output<Integer> timestamp;

    public Output<Integer> timestamp() {
        return this.timestamp;
    }

    private PulsarJobBlendMetricWeightsArgs() {}

    private PulsarJobBlendMetricWeightsArgs(PulsarJobBlendMetricWeightsArgs $) {
        this.timestamp = $.timestamp;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PulsarJobBlendMetricWeightsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PulsarJobBlendMetricWeightsArgs $;

        public Builder() {
            $ = new PulsarJobBlendMetricWeightsArgs();
        }

        public Builder(PulsarJobBlendMetricWeightsArgs defaults) {
            $ = new PulsarJobBlendMetricWeightsArgs(Objects.requireNonNull(defaults));
        }

        public Builder timestamp(Output<Integer> timestamp) {
            $.timestamp = timestamp;
            return this;
        }

        public Builder timestamp(Integer timestamp) {
            return timestamp(Output.of(timestamp));
        }

        public PulsarJobBlendMetricWeightsArgs build() {
            if ($.timestamp == null) {
                throw new MissingRequiredPropertyException("PulsarJobBlendMetricWeightsArgs", "timestamp");
            }
            return $;
        }
    }

}
