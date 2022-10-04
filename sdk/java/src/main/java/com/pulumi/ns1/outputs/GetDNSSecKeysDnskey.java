// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDNSSecKeysDnskey {
    /**
     * @return (Computed) Algorithm of the key.
     * 
     */
    private String algorithm;
    /**
     * @return (Computed) Flags for the key.
     * 
     */
    private String flags;
    /**
     * @return (Computed) Protocol of the key.
     * 
     */
    private String protocol;
    /**
     * @return (Computed) Public key for the key.
     * 
     */
    private String publicKey;

    private GetDNSSecKeysDnskey() {}
    /**
     * @return (Computed) Algorithm of the key.
     * 
     */
    public String algorithm() {
        return this.algorithm;
    }
    /**
     * @return (Computed) Flags for the key.
     * 
     */
    public String flags() {
        return this.flags;
    }
    /**
     * @return (Computed) Protocol of the key.
     * 
     */
    public String protocol() {
        return this.protocol;
    }
    /**
     * @return (Computed) Public key for the key.
     * 
     */
    public String publicKey() {
        return this.publicKey;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDNSSecKeysDnskey defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String algorithm;
        private String flags;
        private String protocol;
        private String publicKey;
        public Builder() {}
        public Builder(GetDNSSecKeysDnskey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithm = defaults.algorithm;
    	      this.flags = defaults.flags;
    	      this.protocol = defaults.protocol;
    	      this.publicKey = defaults.publicKey;
        }

        @CustomType.Setter
        public Builder algorithm(String algorithm) {
            this.algorithm = Objects.requireNonNull(algorithm);
            return this;
        }
        @CustomType.Setter
        public Builder flags(String flags) {
            this.flags = Objects.requireNonNull(flags);
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        @CustomType.Setter
        public Builder publicKey(String publicKey) {
            this.publicKey = Objects.requireNonNull(publicKey);
            return this;
        }
        public GetDNSSecKeysDnskey build() {
            final var o = new GetDNSSecKeysDnskey();
            o.algorithm = algorithm;
            o.flags = flags;
            o.protocol = protocol;
            o.publicKey = publicKey;
            return o;
        }
    }
}
