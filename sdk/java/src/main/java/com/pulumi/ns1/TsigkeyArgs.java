// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TsigkeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final TsigkeyArgs Empty = new TsigkeyArgs();

    /**
     * The algorithm used to hash the TSIG key&#39;s secret.
     * 
     */
    @Import(name="algorithm", required=true)
    private Output<String> algorithm;

    /**
     * @return The algorithm used to hash the TSIG key&#39;s secret.
     * 
     */
    public Output<String> algorithm() {
        return this.algorithm;
    }

    /**
     * The free form name of the tsigkey.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The free form name of the tsigkey.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The key&#39;s secret to be hashed.
     * 
     */
    @Import(name="secret", required=true)
    private Output<String> secret;

    /**
     * @return The key&#39;s secret to be hashed.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }

    private TsigkeyArgs() {}

    private TsigkeyArgs(TsigkeyArgs $) {
        this.algorithm = $.algorithm;
        this.name = $.name;
        this.secret = $.secret;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TsigkeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TsigkeyArgs $;

        public Builder() {
            $ = new TsigkeyArgs();
        }

        public Builder(TsigkeyArgs defaults) {
            $ = new TsigkeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm The algorithm used to hash the TSIG key&#39;s secret.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm The algorithm used to hash the TSIG key&#39;s secret.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param name The free form name of the tsigkey.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The free form name of the tsigkey.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param secret The key&#39;s secret to be hashed.
         * 
         * @return builder
         * 
         */
        public Builder secret(Output<String> secret) {
            $.secret = secret;
            return this;
        }

        /**
         * @param secret The key&#39;s secret to be hashed.
         * 
         * @return builder
         * 
         */
        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        public TsigkeyArgs build() {
            if ($.algorithm == null) {
                throw new MissingRequiredPropertyException("TsigkeyArgs", "algorithm");
            }
            if ($.secret == null) {
                throw new MissingRequiredPropertyException("TsigkeyArgs", "secret");
            }
            return $;
        }
    }

}
