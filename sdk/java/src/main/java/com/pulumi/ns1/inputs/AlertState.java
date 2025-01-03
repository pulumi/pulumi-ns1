// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AlertState extends com.pulumi.resources.ResourceArgs {

    public static final AlertState Empty = new AlertState();

    /**
     * (Read Only) The Unix timestamp representing when the alert configuration was created.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<Integer> createdAt;

    /**
     * @return (Read Only) The Unix timestamp representing when the alert configuration was created.
     * 
     */
    public Optional<Output<Integer>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * (Read Only) The user or apikey that created this alert.
     * 
     */
    @Import(name="createdBy")
    private @Nullable Output<String> createdBy;

    /**
     * @return (Read Only) The user or apikey that created this alert.
     * 
     */
    public Optional<Output<String>> createdBy() {
        return Optional.ofNullable(this.createdBy);
    }

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
    @Import(name="subtype")
    private @Nullable Output<String> subtype;

    /**
     * @return The type of the alert.
     * 
     */
    public Optional<Output<String>> subtype() {
        return Optional.ofNullable(this.subtype);
    }

    /**
     * The type of the alert.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of the alert.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * (Read Only) The Unix timestamp representing when the alert configuration was last modified.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<Integer> updatedAt;

    /**
     * @return (Read Only) The Unix timestamp representing when the alert configuration was last modified.
     * 
     */
    public Optional<Output<Integer>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    /**
     * (Read Only) The user or apikey that last modified this alert.
     * 
     */
    @Import(name="updatedBy")
    private @Nullable Output<String> updatedBy;

    /**
     * @return (Read Only) The user or apikey that last modified this alert.
     * 
     */
    public Optional<Output<String>> updatedBy() {
        return Optional.ofNullable(this.updatedBy);
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

    private AlertState() {}

    private AlertState(AlertState $) {
        this.createdAt = $.createdAt;
        this.createdBy = $.createdBy;
        this.name = $.name;
        this.notificationLists = $.notificationLists;
        this.recordIds = $.recordIds;
        this.subtype = $.subtype;
        this.type = $.type;
        this.updatedAt = $.updatedAt;
        this.updatedBy = $.updatedBy;
        this.zoneNames = $.zoneNames;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AlertState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AlertState $;

        public Builder() {
            $ = new AlertState();
        }

        public Builder(AlertState defaults) {
            $ = new AlertState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdAt (Read Only) The Unix timestamp representing when the alert configuration was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<Integer> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt (Read Only) The Unix timestamp representing when the alert configuration was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(Integer createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param createdBy (Read Only) The user or apikey that created this alert.
         * 
         * @return builder
         * 
         */
        public Builder createdBy(@Nullable Output<String> createdBy) {
            $.createdBy = createdBy;
            return this;
        }

        /**
         * @param createdBy (Read Only) The user or apikey that created this alert.
         * 
         * @return builder
         * 
         */
        public Builder createdBy(String createdBy) {
            return createdBy(Output.of(createdBy));
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
        public Builder subtype(@Nullable Output<String> subtype) {
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
        public Builder type(@Nullable Output<String> type) {
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
         * @param updatedAt (Read Only) The Unix timestamp representing when the alert configuration was last modified.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<Integer> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt (Read Only) The Unix timestamp representing when the alert configuration was last modified.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(Integer updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        /**
         * @param updatedBy (Read Only) The user or apikey that last modified this alert.
         * 
         * @return builder
         * 
         */
        public Builder updatedBy(@Nullable Output<String> updatedBy) {
            $.updatedBy = updatedBy;
            return this;
        }

        /**
         * @param updatedBy (Read Only) The user or apikey that last modified this alert.
         * 
         * @return builder
         * 
         */
        public Builder updatedBy(String updatedBy) {
            return updatedBy(Output.of(updatedBy));
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

        public AlertState build() {
            return $;
        }
    }

}
