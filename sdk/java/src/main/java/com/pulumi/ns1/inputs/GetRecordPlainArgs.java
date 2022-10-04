// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetRecordPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRecordPlainArgs Empty = new GetRecordPlainArgs();

    /**
     * The records&#39; domain.
     * 
     */
    @Import(name="domain", required=true)
    private String domain;

    /**
     * @return The records&#39; domain.
     * 
     */
    public String domain() {
        return this.domain;
    }

    /**
     * The records&#39; RR type.
     * 
     */
    @Import(name="type", required=true)
    private String type;

    /**
     * @return The records&#39; RR type.
     * 
     */
    public String type() {
        return this.type;
    }

    /**
     * The zone the record belongs to.
     * 
     */
    @Import(name="zone", required=true)
    private String zone;

    /**
     * @return The zone the record belongs to.
     * 
     */
    public String zone() {
        return this.zone;
    }

    private GetRecordPlainArgs() {}

    private GetRecordPlainArgs(GetRecordPlainArgs $) {
        this.domain = $.domain;
        this.type = $.type;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRecordPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRecordPlainArgs $;

        public Builder() {
            $ = new GetRecordPlainArgs();
        }

        public Builder(GetRecordPlainArgs defaults) {
            $ = new GetRecordPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param domain The records&#39; domain.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param type The records&#39; RR type.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            $.type = type;
            return this;
        }

        /**
         * @param zone The zone the record belongs to.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            $.zone = zone;
            return this;
        }

        public GetRecordPlainArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            $.zone = Objects.requireNonNull($.zone, "expected parameter 'zone' to be non-null");
            return $;
        }
    }

}
