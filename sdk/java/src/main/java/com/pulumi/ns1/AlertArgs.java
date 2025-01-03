// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlertArgs extends com.pulumi.resources.ResourceArgs {

    public static final AlertArgs Empty = new AlertArgs();

    /**
     * The free-form display name for the alert.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The free-form display name for the alert.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A list of id&#39;s for notification lists whose notifiers will be triggered by the alert.
     * 
     */
    @Import(name="notificationLists")
    private @Nullable Output<List<String>> notificationLists;

    /**
     * @return A list of id&#39;s for notification lists whose notifiers will be triggered by the alert.
     * 
     */
    public Optional<Output<List<String>>> notificationLists() {
        return Optional.ofNullable(this.notificationLists);
    }

    /**
     * A list of record id&#39;s this alert applies to.
     * 
     */
    @Import(name="recordIds")
    private @Nullable Output<List<String>> recordIds;

    /**
     * @return A list of record id&#39;s this alert applies to.
     * 
     */
    public Optional<Output<List<String>>> recordIds() {
        return Optional.ofNullable(this.recordIds);
    }

    /**
     * The type of the alert.
     * 
     */
    @Import(name="subtype", required=true)
    private Output<String> subtype;

    /**
     * @return The type of the alert.
     * 
     */
    public Output<String> subtype() {
        return this.subtype;
    }

    /**
     * The type of the alert.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of the alert.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * A list of zones this alert applies to.
     * 
     */
    @Import(name="zoneNames")
    private @Nullable Output<List<String>> zoneNames;

    /**
     * @return A list of zones this alert applies to.
     * 
     */
    public Optional<Output<List<String>>> zoneNames() {
        return Optional.ofNullable(this.zoneNames);
    }

    private AlertArgs() {}

    private AlertArgs(AlertArgs $) {
        this.name = $.name;
        this.notificationLists = $.notificationLists;
        this.recordIds = $.recordIds;
        this.subtype = $.subtype;
        this.type = $.type;
        this.zoneNames = $.zoneNames;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlertArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlertArgs $;

        public Builder() {
            $ = new AlertArgs();
        }

        public Builder(AlertArgs defaults) {
            $ = new AlertArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The free-form display name for the alert.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The free-form display name for the alert.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param notificationLists A list of id&#39;s for notification lists whose notifiers will be triggered by the alert.
         * 
         * @return builder
         * 
         */
        public Builder notificationLists(@Nullable Output<List<String>> notificationLists) {
            $.notificationLists = notificationLists;
            return this;
        }

        /**
         * @param notificationLists A list of id&#39;s for notification lists whose notifiers will be triggered by the alert.
         * 
         * @return builder
         * 
         */
        public Builder notificationLists(List<String> notificationLists) {
            return notificationLists(Output.of(notificationLists));
        }

        /**
         * @param notificationLists A list of id&#39;s for notification lists whose notifiers will be triggered by the alert.
         * 
         * @return builder
         * 
         */
        public Builder notificationLists(String... notificationLists) {
            return notificationLists(List.of(notificationLists));
        }

        /**
         * @param recordIds A list of record id&#39;s this alert applies to.
         * 
         * @return builder
         * 
         */
        public Builder recordIds(@Nullable Output<List<String>> recordIds) {
            $.recordIds = recordIds;
            return this;
        }

        /**
         * @param recordIds A list of record id&#39;s this alert applies to.
         * 
         * @return builder
         * 
         */
        public Builder recordIds(List<String> recordIds) {
            return recordIds(Output.of(recordIds));
        }

        /**
         * @param recordIds A list of record id&#39;s this alert applies to.
         * 
         * @return builder
         * 
         */
        public Builder recordIds(String... recordIds) {
            return recordIds(List.of(recordIds));
        }

        /**
         * @param subtype The type of the alert.
         * 
         * @return builder
         * 
         */
        public Builder subtype(Output<String> subtype) {
            $.subtype = subtype;
            return this;
        }

        /**
         * @param subtype The type of the alert.
         * 
         * @return builder
         * 
         */
        public Builder subtype(String subtype) {
            return subtype(Output.of(subtype));
        }

        /**
         * @param type The type of the alert.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of the alert.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param zoneNames A list of zones this alert applies to.
         * 
         * @return builder
         * 
         */
        public Builder zoneNames(@Nullable Output<List<String>> zoneNames) {
            $.zoneNames = zoneNames;
            return this;
        }

        /**
         * @param zoneNames A list of zones this alert applies to.
         * 
         * @return builder
         * 
         */
        public Builder zoneNames(List<String> zoneNames) {
            return zoneNames(Output.of(zoneNames));
        }

        /**
         * @param zoneNames A list of zones this alert applies to.
         * 
         * @return builder
         * 
         */
        public Builder zoneNames(String... zoneNames) {
            return zoneNames(List.of(zoneNames));
        }

        public AlertArgs build() {
            if ($.subtype == null) {
                throw new MissingRequiredPropertyException("AlertArgs", "subtype");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("AlertArgs", "type");
            }
            return $;
        }
    }

}