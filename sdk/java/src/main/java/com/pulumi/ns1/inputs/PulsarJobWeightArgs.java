// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PulsarJobWeightArgs extends com.pulumi.resources.ResourceArgs {

    public static final PulsarJobWeightArgs Empty = new PulsarJobWeightArgs();

    @Import(name="defaultValue", required=true)
    private Output<Double> defaultValue;

    public Output<Double> defaultValue() {
        return this.defaultValue;
    }

    @Import(name="maximize")
    private @Nullable Output<Boolean> maximize;

    public Optional<Output<Boolean>> maximize() {
        return Optional.ofNullable(this.maximize);
    }

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    @Import(name="weight", required=true)
    private Output<Integer> weight;

    public Output<Integer> weight() {
        return this.weight;
    }

    private PulsarJobWeightArgs() {}

    private PulsarJobWeightArgs(PulsarJobWeightArgs $) {
        this.defaultValue = $.defaultValue;
        this.maximize = $.maximize;
        this.name = $.name;
        this.weight = $.weight;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PulsarJobWeightArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PulsarJobWeightArgs $;

        public Builder() {
            $ = new PulsarJobWeightArgs();
        }

        public Builder(PulsarJobWeightArgs defaults) {
            $ = new PulsarJobWeightArgs(Objects.requireNonNull(defaults));
        }

        public Builder defaultValue(Output<Double> defaultValue) {
            $.defaultValue = defaultValue;
            return this;
        }

        public Builder defaultValue(Double defaultValue) {
            return defaultValue(Output.of(defaultValue));
        }

        public Builder maximize(@Nullable Output<Boolean> maximize) {
            $.maximize = maximize;
            return this;
        }

        public Builder maximize(Boolean maximize) {
            return maximize(Output.of(maximize));
        }

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder weight(Output<Integer> weight) {
            $.weight = weight;
            return this;
        }

        public Builder weight(Integer weight) {
            return weight(Output.of(weight));
        }

        public PulsarJobWeightArgs build() {
            if ($.defaultValue == null) {
                throw new MissingRequiredPropertyException("PulsarJobWeightArgs", "defaultValue");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("PulsarJobWeightArgs", "name");
            }
            if ($.weight == null) {
                throw new MissingRequiredPropertyException("PulsarJobWeightArgs", "weight");
            }
            return $;
        }
    }

}
