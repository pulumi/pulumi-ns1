// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class DatasetRepeat {
    private Integer endAfterN;
    private String repeatsEvery;
    private Integer start;

    private DatasetRepeat() {}
    public Integer endAfterN() {
        return this.endAfterN;
    }
    public String repeatsEvery() {
        return this.repeatsEvery;
    }
    public Integer start() {
        return this.start;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DatasetRepeat defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer endAfterN;
        private String repeatsEvery;
        private Integer start;
        public Builder() {}
        public Builder(DatasetRepeat defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.endAfterN = defaults.endAfterN;
    	      this.repeatsEvery = defaults.repeatsEvery;
    	      this.start = defaults.start;
        }

        @CustomType.Setter
        public Builder endAfterN(Integer endAfterN) {
            if (endAfterN == null) {
              throw new MissingRequiredPropertyException("DatasetRepeat", "endAfterN");
            }
            this.endAfterN = endAfterN;
            return this;
        }
        @CustomType.Setter
        public Builder repeatsEvery(String repeatsEvery) {
            if (repeatsEvery == null) {
              throw new MissingRequiredPropertyException("DatasetRepeat", "repeatsEvery");
            }
            this.repeatsEvery = repeatsEvery;
            return this;
        }
        @CustomType.Setter
        public Builder start(Integer start) {
            if (start == null) {
              throw new MissingRequiredPropertyException("DatasetRepeat", "start");
            }
            this.start = start;
            return this;
        }
        public DatasetRepeat build() {
            final var _resultValue = new DatasetRepeat();
            _resultValue.endAfterN = endAfterN;
            _resultValue.repeatsEvery = repeatsEvery;
            _resultValue.start = start;
            return _resultValue;
        }
    }
}
