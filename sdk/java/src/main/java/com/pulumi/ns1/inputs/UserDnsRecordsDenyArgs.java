// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class UserDnsRecordsDenyArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserDnsRecordsDenyArgs Empty = new UserDnsRecordsDenyArgs();

    @Import(name="domain", required=true)
    private Output<String> domain;

    public Output<String> domain() {
        return this.domain;
    }

    @Import(name="includeSubdomains", required=true)
    private Output<Boolean> includeSubdomains;

    public Output<Boolean> includeSubdomains() {
        return this.includeSubdomains;
    }

    @Import(name="type", required=true)
    private Output<String> type;

    public Output<String> type() {
        return this.type;
    }

    @Import(name="zone", required=true)
    private Output<String> zone;

    public Output<String> zone() {
        return this.zone;
    }

    private UserDnsRecordsDenyArgs() {}

    private UserDnsRecordsDenyArgs(UserDnsRecordsDenyArgs $) {
        this.domain = $.domain;
        this.includeSubdomains = $.includeSubdomains;
        this.type = $.type;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserDnsRecordsDenyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserDnsRecordsDenyArgs $;

        public Builder() {
            $ = new UserDnsRecordsDenyArgs();
        }

        public Builder(UserDnsRecordsDenyArgs defaults) {
            $ = new UserDnsRecordsDenyArgs(Objects.requireNonNull(defaults));
        }

        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        public Builder includeSubdomains(Output<Boolean> includeSubdomains) {
            $.includeSubdomains = includeSubdomains;
            return this;
        }

        public Builder includeSubdomains(Boolean includeSubdomains) {
            return includeSubdomains(Output.of(includeSubdomains));
        }

        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public Builder zone(Output<String> zone) {
            $.zone = zone;
            return this;
        }

        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public UserDnsRecordsDenyArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            $.includeSubdomains = Objects.requireNonNull($.includeSubdomains, "expected parameter 'includeSubdomains' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            $.zone = Objects.requireNonNull($.zone, "expected parameter 'zone' to be non-null");
            return $;
        }
    }

}
