// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ns1.inputs.PulsarJobBlendMetricWeightsArgs;
import com.pulumi.ns1.inputs.PulsarJobConfigArgs;
import com.pulumi.ns1.inputs.PulsarJobWeightArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PulsarJobArgs extends com.pulumi.resources.ResourceArgs {

    public static final PulsarJobArgs Empty = new PulsarJobArgs();

    @Import(name="active")
    private @Nullable Output<Boolean> active;

    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    @Import(name="appId", required=true)
    private Output<String> appId;

    public Output<String> appId() {
        return this.appId;
    }

    @Import(name="blendMetricWeights")
    private @Nullable Output<PulsarJobBlendMetricWeightsArgs> blendMetricWeights;

    public Optional<Output<PulsarJobBlendMetricWeightsArgs>> blendMetricWeights() {
        return Optional.ofNullable(this.blendMetricWeights);
    }

    @Import(name="config")
    private @Nullable Output<PulsarJobConfigArgs> config;

    public Optional<Output<PulsarJobConfigArgs>> config() {
        return Optional.ofNullable(this.config);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    @Import(name="typeId", required=true)
    private Output<String> typeId;

    public Output<String> typeId() {
        return this.typeId;
    }

    @Import(name="weights")
    private @Nullable Output<List<PulsarJobWeightArgs>> weights;

    public Optional<Output<List<PulsarJobWeightArgs>>> weights() {
        return Optional.ofNullable(this.weights);
    }

    private PulsarJobArgs() {}

    private PulsarJobArgs(PulsarJobArgs $) {
        this.active = $.active;
        this.appId = $.appId;
        this.blendMetricWeights = $.blendMetricWeights;
        this.config = $.config;
        this.name = $.name;
        this.shared = $.shared;
        this.typeId = $.typeId;
        this.weights = $.weights;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PulsarJobArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PulsarJobArgs $;

        public Builder() {
            $ = new PulsarJobArgs();
        }

        public Builder(PulsarJobArgs defaults) {
            $ = new PulsarJobArgs(Objects.requireNonNull(defaults));
        }

        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        public Builder appId(Output<String> appId) {
            $.appId = appId;
            return this;
        }

        public Builder appId(String appId) {
            return appId(Output.of(appId));
        }

        public Builder blendMetricWeights(@Nullable Output<PulsarJobBlendMetricWeightsArgs> blendMetricWeights) {
            $.blendMetricWeights = blendMetricWeights;
            return this;
        }

        public Builder blendMetricWeights(PulsarJobBlendMetricWeightsArgs blendMetricWeights) {
            return blendMetricWeights(Output.of(blendMetricWeights));
        }

        public Builder config(@Nullable Output<PulsarJobConfigArgs> config) {
            $.config = config;
            return this;
        }

        public Builder config(PulsarJobConfigArgs config) {
            return config(Output.of(config));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder shared(@Nullable Output<Boolean> shared) {
            $.shared = shared;
            return this;
        }

        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        public Builder typeId(Output<String> typeId) {
            $.typeId = typeId;
            return this;
        }

        public Builder typeId(String typeId) {
            return typeId(Output.of(typeId));
        }

        public Builder weights(@Nullable Output<List<PulsarJobWeightArgs>> weights) {
            $.weights = weights;
            return this;
        }

        public Builder weights(List<PulsarJobWeightArgs> weights) {
            return weights(Output.of(weights));
        }

        public Builder weights(PulsarJobWeightArgs... weights) {
            return weights(List.of(weights));
        }

        public PulsarJobArgs build() {
            if ($.appId == null) {
                throw new MissingRequiredPropertyException("PulsarJobArgs", "appId");
            }
            if ($.typeId == null) {
                throw new MissingRequiredPropertyException("PulsarJobArgs", "typeId");
            }
            return $;
        }
    }

}