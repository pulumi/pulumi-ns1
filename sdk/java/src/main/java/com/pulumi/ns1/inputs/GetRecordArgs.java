// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetRecordArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRecordArgs Empty = new GetRecordArgs();

    /**
     * The records&#39; domain.
     * 
     */
    @Import(name="domain", required=true)
    private Output<String> domain;

    /**
     * @return The records&#39; domain.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }

    /**
     * The records&#39; RR type.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The records&#39; RR type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * The zone the record belongs to.
     * 
     */
    @Import(name="zone", required=true)
    private Output<String> zone;

    /**
     * @return The zone the record belongs to.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    private GetRecordArgs() {}

    private GetRecordArgs(GetRecordArgs $) {
        this.domain = $.domain;
        this.type = $.type;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRecordArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRecordArgs $;

        public Builder() {
            $ = new GetRecordArgs();
        }

        public Builder(GetRecordArgs defaults) {
            $ = new GetRecordArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain The records&#39; domain.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The records&#39; domain.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param type The records&#39; RR type.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The records&#39; RR type.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param zone The zone the record belongs to.
         * 
         * @return builder
         * 
         */
        public Builder zone(Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone The zone the record belongs to.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public GetRecordArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            $.zone = Objects.requireNonNull($.zone, "expected parameter 'zone' to be non-null");
            return $;
        }
    }

}
