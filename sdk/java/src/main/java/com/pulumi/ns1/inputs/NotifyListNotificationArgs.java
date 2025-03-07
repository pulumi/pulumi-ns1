// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;


public final class NotifyListNotificationArgs extends com.pulumi.resources.ResourceArgs {

    public static final NotifyListNotificationArgs Empty = new NotifyListNotificationArgs();

    /**
     * Configuration details for the given notifier type.
     * 
     */
    @Import(name="config", required=true)
    private Output<Map<String,String>> config;

    /**
     * @return Configuration details for the given notifier type.
     * 
     */
    public Output<Map<String,String>> config() {
        return this.config;
    }

    /**
     * The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private NotifyListNotificationArgs() {}

    private NotifyListNotificationArgs(NotifyListNotificationArgs $) {
        this.config = $.config;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NotifyListNotificationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NotifyListNotificationArgs $;

        public Builder() {
            $ = new NotifyListNotificationArgs();
        }

        public Builder(NotifyListNotificationArgs defaults) {
            $ = new NotifyListNotificationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param config Configuration details for the given notifier type.
         * 
         * @return builder
         * 
         */
        public Builder config(Output<Map<String,String>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config Configuration details for the given notifier type.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,String> config) {
            return config(Output.of(config));
        }

        /**
         * @param type The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public NotifyListNotificationArgs build() {
            if ($.config == null) {
                throw new MissingRequiredPropertyException("NotifyListNotificationArgs", "config");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("NotifyListNotificationArgs", "type");
            }
            return $;
        }
    }

}
