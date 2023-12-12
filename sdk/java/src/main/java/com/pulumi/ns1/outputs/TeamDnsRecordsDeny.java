// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class TeamDnsRecordsDeny {
    private String domain;
    private Boolean includeSubdomains;
    private String type;
    private String zone;

    private TeamDnsRecordsDeny() {}
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

    public static Builder builder(TeamDnsRecordsDeny defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String domain;
        private Boolean includeSubdomains;
        private String type;
        private String zone;
        public Builder() {}
        public Builder(TeamDnsRecordsDeny defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domain = defaults.domain;
    	      this.includeSubdomains = defaults.includeSubdomains;
    	      this.type = defaults.type;
    	      this.zone = defaults.zone;
        }

        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder includeSubdomains(Boolean includeSubdomains) {
            this.includeSubdomains = Objects.requireNonNull(includeSubdomains);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder zone(String zone) {
            this.zone = Objects.requireNonNull(zone);
            return this;
        }
        public TeamDnsRecordsDeny build() {
            final var _resultValue = new TeamDnsRecordsDeny();
            _resultValue.domain = domain;
            _resultValue.includeSubdomains = includeSubdomains;
            _resultValue.type = type;
            _resultValue.zone = zone;
            return _resultValue;
        }
    }
}
