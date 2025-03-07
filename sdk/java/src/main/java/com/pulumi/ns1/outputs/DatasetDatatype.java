// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class DatasetDatatype {
    private Map<String,String> data;
    private String scope;
    private String type;

    private DatasetDatatype() {}
    public Map<String,String> data() {
        return this.data;
    }
    public String scope() {
        return this.scope;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DatasetDatatype defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> data;
        private String scope;
        private String type;
        public Builder() {}
        public Builder(DatasetDatatype defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.data = defaults.data;
    	      this.scope = defaults.scope;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder data(Map<String,String> data) {
            if (data == null) {
              throw new MissingRequiredPropertyException("DatasetDatatype", "data");
            }
            this.data = data;
            return this;
        }
        @CustomType.Setter
        public Builder scope(String scope) {
            if (scope == null) {
              throw new MissingRequiredPropertyException("DatasetDatatype", "scope");
            }
            this.scope = scope;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("DatasetDatatype", "type");
            }
            this.type = type;
            return this;
        }
        public DatasetDatatype build() {
            final var _resultValue = new DatasetDatatype();
            _resultValue.data = data;
            _resultValue.scope = scope;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
