// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class APIKeyDnsRecordsAllow {
    private String domain;
    private Boolean includeSubdomains;
    private String type;
    private String zone;

    private APIKeyDnsRecordsAllow() {}
    public String domain() {
        return this.domain;
    }
    public Boolean includeSubdomains() {
        return this.includeSubdomains;
    }
    public String type() {
        return this.type;
    }
    public String zone() {
        return this.zone;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(APIKeyDnsRecordsAllow defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String domain;
        private Boolean includeSubdomains;
        private String type;
        private String zone;
        public Builder() {}
        public Builder(APIKeyDnsRecordsAllow defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domain = defaults.domain;
    	      this.includeSubdomains = defaults.includeSubdomains;
    	      this.type = defaults.type;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder domain(String domain) {
            if (domain == null) {
              throw new MissingRequiredPropertyException("APIKeyDnsRecordsAllow", "domain");
            }
            this.domain = domain;
            return this;
        }
        @CustomType.Setter
        public Builder includeSubdomains(Boolean includeSubdomains) {
            if (includeSubdomains == null) {
              throw new MissingRequiredPropertyException("APIKeyDnsRecordsAllow", "includeSubdomains");
            }
            this.includeSubdomains = includeSubdomains;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("APIKeyDnsRecordsAllow", "type");
            }
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            if (zone == null) {
              throw new MissingRequiredPropertyException("APIKeyDnsRecordsAllow", "zone");
            }
            this.zone = zone;
            return this;
        }
        public APIKeyDnsRecordsAllow build() {
            final var _resultValue = new APIKeyDnsRecordsAllow();
            _resultValue.domain = domain;
            _resultValue.includeSubdomains = includeSubdomains;
            _resultValue.type = type;
            _resultValue.zone = zone;
            return _resultValue;
        }
    }
}
