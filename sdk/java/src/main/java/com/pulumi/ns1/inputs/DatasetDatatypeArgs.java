// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;


public final class DatasetDatatypeArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatasetDatatypeArgs Empty = new DatasetDatatypeArgs();

    @Import(name="data", required=true)
    private Output<Map<String,Object>> data;

    public Output<Map<String,Object>> data() {
        return this.data;
    }

    @Import(name="scope", required=true)
    private Output<String> scope;

    public Output<String> scope() {
        return this.scope;
    }

    @Import(name="type", required=true)
    private Output<String> type;

    public Output<String> type() {
        return this.type;
    }

    private DatasetDatatypeArgs() {}

    private DatasetDatatypeArgs(DatasetDatatypeArgs $) {
        this.data = $.data;
        this.scope = $.scope;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatasetDatatypeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatasetDatatypeArgs $;

        public Builder() {
            $ = new DatasetDatatypeArgs();
        }

        public Builder(DatasetDatatypeArgs defaults) {
            $ = new DatasetDatatypeArgs(Objects.requireNonNull(defaults));
        }

        public Builder data(Output<Map<String,Object>> data) {
            $.data = data;
            return this;
        }

        public Builder data(Map<String,Object> data) {
            return data(Output.of(data));
        }

        public Builder scope(Output<String> scope) {
            $.scope = scope;
            return this;
        }

        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public DatasetDatatypeArgs build() {
            if ($.data == null) {
                throw new MissingRequiredPropertyException("DatasetDatatypeArgs", "data");
            }
            if ($.scope == null) {
                throw new MissingRequiredPropertyException("DatasetDatatypeArgs", "scope");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("DatasetDatatypeArgs", "type");
            }
            return $;
        }
    }

}