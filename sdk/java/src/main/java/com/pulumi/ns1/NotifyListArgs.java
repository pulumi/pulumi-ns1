// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ns1.inputs.NotifyListNotificationArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NotifyListArgs extends com.pulumi.resources.ResourceArgs {

    public static final NotifyListArgs Empty = new NotifyListArgs();

    /**
     * The free-form display name for the notify list.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The free-form display name for the notify list.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
     * 
     */
    @Import(name="notifications")
    private @Nullable Output<List<NotifyListNotificationArgs>> notifications;

    /**
     * @return A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
     * 
     */
    public Optional<Output<List<NotifyListNotificationArgs>>> notifications() {
        return Optional.ofNullable(this.notifications);
    }

    private NotifyListArgs() {}

    private NotifyListArgs(NotifyListArgs $) {
        this.name = $.name;
        this.notifications = $.notifications;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NotifyListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NotifyListArgs $;

        public Builder() {
            $ = new NotifyListArgs();
        }

        public Builder(NotifyListArgs defaults) {
            $ = new NotifyListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The free-form display name for the notify list.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The free-form display name for the notify list.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param notifications A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
         * 
         * @return builder
         * 
         */
        public Builder notifications(@Nullable Output<List<NotifyListNotificationArgs>> notifications) {
            $.notifications = notifications;
            return this;
        }

        /**
         * @param notifications A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
         * 
         * @return builder
         * 
         */
        public Builder notifications(List<NotifyListNotificationArgs> notifications) {
            return notifications(Output.of(notifications));
        }

        /**
         * @param notifications A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
         * 
         * @return builder
         * 
         */
        public Builder notifications(NotifyListNotificationArgs... notifications) {
            return notifications(List.of(notifications));
        }

        public NotifyListArgs build() {
            return $;
        }
    }

}
