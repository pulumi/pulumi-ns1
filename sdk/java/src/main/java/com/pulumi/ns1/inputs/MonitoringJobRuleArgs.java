// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class MonitoringJobRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final MonitoringJobRuleArgs Empty = new MonitoringJobRuleArgs();

    @Import(name="comparison", required=true)
    private Output<String> comparison;

    public Output<String> comparison() {
        return this.comparison;
    }

    @Import(name="key", required=true)
    private Output<String> key;

    public Output<String> key() {
        return this.key;
    }

    @Import(name="value", required=true)
    private Output<String> value;

    public Output<String> value() {
        return this.value;
    }

    private MonitoringJobRuleArgs() {}

    private MonitoringJobRuleArgs(MonitoringJobRuleArgs $) {
        this.comparison = $.comparison;
        this.key = $.key;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MonitoringJobRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MonitoringJobRuleArgs $;

        public Builder() {
            $ = new MonitoringJobRuleArgs();
        }

        public Builder(MonitoringJobRuleArgs defaults) {
            $ = new MonitoringJobRuleArgs(Objects.requireNonNull(defaults));
        }

        public Builder comparison(Output<String> comparison) {
            $.comparison = comparison;
            return this;
        }

        public Builder comparison(String comparison) {
            return comparison(Output.of(comparison));
        }

        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        public Builder value(String value) {
            return value(Output.of(value));
        }

        public MonitoringJobRuleArgs build() {
            $.comparison = Objects.requireNonNull($.comparison, "expected parameter 'comparison' to be non-null");
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}
