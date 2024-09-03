// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DatasetTimeframe {
    private String aggregation;
    private @Nullable Integer cycles;
    private @Nullable Integer from;
    private @Nullable Integer to;

    private DatasetTimeframe() {}
    public String aggregation() {
        return this.aggregation;
    }
    public Optional<Integer> cycles() {
        return Optional.ofNullable(this.cycles);
    }
    public Optional<Integer> from() {
        return Optional.ofNullable(this.from);
    }
    public Optional<Integer> to() {
        return Optional.ofNullable(this.to);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DatasetTimeframe defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String aggregation;
        private @Nullable Integer cycles;
        private @Nullable Integer from;
        private @Nullable Integer to;
        public Builder() {}
        public Builder(DatasetTimeframe defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aggregation = defaults.aggregation;
    	      this.cycles = defaults.cycles;
    	      this.from = defaults.from;
    	      this.to = defaults.to;
        }

        @CustomType.Setter
        public Builder aggregation(String aggregation) {
            if (aggregation == null) {
              throw new MissingRequiredPropertyException("DatasetTimeframe", "aggregation");
            }
            this.aggregation = aggregation;
            return this;
        }
        @CustomType.Setter
        public Builder cycles(@Nullable Integer cycles) {

            this.cycles = cycles;
            return this;
        }
        @CustomType.Setter
        public Builder from(@Nullable Integer from) {

            this.from = from;
            return this;
        }
        @CustomType.Setter
        public Builder to(@Nullable Integer to) {

            this.to = to;
            return this;
        }
        public DatasetTimeframe build() {
            final var _resultValue = new DatasetTimeframe();
            _resultValue.aggregation = aggregation;
            _resultValue.cycles = cycles;
            _resultValue.from = from;
            _resultValue.to = to;
            return _resultValue;
        }
    }
}