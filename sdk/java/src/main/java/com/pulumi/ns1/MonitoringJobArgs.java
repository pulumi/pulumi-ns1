// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ns1.inputs.MonitoringJobRuleArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MonitoringJobArgs extends com.pulumi.resources.ResourceArgs {

    public static final MonitoringJobArgs Empty = new MonitoringJobArgs();

    /**
     * Indicates if the job is active or temporarily disabled.
     * 
     */
    @Import(name="active")
    private @Nullable Output<Boolean> active;

    /**
     * @return Indicates if the job is active or temporarily disabled.
     * 
     */
    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    /**
     * A configuration dictionary with keys and values depending on the job_type. Configuration details for each job_type are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     * 
     */
    @Import(name="config", required=true)
    private Output<Map<String,Object>> config;

    /**
     * @return A configuration dictionary with keys and values depending on the job_type. Configuration details for each job_type are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     * 
     */
    public Output<Map<String,Object>> config() {
        return this.config;
    }

    /**
     * The frequency, in seconds, at which to run the monitoring job in each region.
     * 
     */
    @Import(name="frequency", required=true)
    private Output<Integer> frequency;

    /**
     * @return The frequency, in seconds, at which to run the monitoring job in each region.
     * 
     */
    public Output<Integer> frequency() {
        return this.frequency;
    }

    /**
     * The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
     * 
     */
    @Import(name="jobType", required=true)
    private Output<String> jobType;

    /**
     * @return The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
     * 
     */
    public Output<String> jobType() {
        return this.jobType;
    }

    /**
     * turn off the notifications for the monitoring job.
     * 
     */
    @Import(name="mute")
    private @Nullable Output<Boolean> mute;

    /**
     * @return turn off the notifications for the monitoring job.
     * 
     */
    public Optional<Output<Boolean>> mute() {
        return Optional.ofNullable(this.mute);
    }

    /**
     * The free-form display name for the monitoring job.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The free-form display name for the monitoring job.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Freeform notes to be included in any notifications about this job.
     * 
     */
    @Import(name="notes")
    private @Nullable Output<String> notes;

    /**
     * @return Freeform notes to be included in any notifications about this job.
     * 
     */
    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    /**
     * The time in seconds after a failure to wait before sending a notification.
     * 
     */
    @Import(name="notifyDelay")
    private @Nullable Output<Integer> notifyDelay;

    /**
     * @return The time in seconds after a failure to wait before sending a notification.
     * 
     */
    public Optional<Output<Integer>> notifyDelay() {
        return Optional.ofNullable(this.notifyDelay);
    }

    /**
     * If true, a notification is sent when a job returns to an &#34;up&#34; state.
     * 
     */
    @Import(name="notifyFailback")
    private @Nullable Output<Boolean> notifyFailback;

    /**
     * @return If true, a notification is sent when a job returns to an &#34;up&#34; state.
     * 
     */
    public Optional<Output<Boolean>> notifyFailback() {
        return Optional.ofNullable(this.notifyFailback);
    }

    @Import(name="notifyList")
    private @Nullable Output<String> notifyList;

    public Optional<Output<String>> notifyList() {
        return Optional.ofNullable(this.notifyList);
    }

    /**
     * If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
     * 
     */
    @Import(name="notifyRegional")
    private @Nullable Output<Boolean> notifyRegional;

    /**
     * @return If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
     * 
     */
    public Optional<Output<Boolean>> notifyRegional() {
        return Optional.ofNullable(this.notifyRegional);
    }

    /**
     * The time in seconds between repeat notifications of a failed job.
     * 
     */
    @Import(name="notifyRepeat")
    private @Nullable Output<Integer> notifyRepeat;

    /**
     * @return The time in seconds between repeat notifications of a failed job.
     * 
     */
    public Optional<Output<Integer>> notifyRepeat() {
        return Optional.ofNullable(this.notifyRepeat);
    }

    /**
     * The policy for determining the monitor&#39;s global status
     * based on the status of the job in all regions. See NS1 API docs for supported values.
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return The policy for determining the monitor&#39;s global status
     * based on the status of the job in all regions. See NS1 API docs for supported values.
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    /**
     * If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
     * 
     */
    @Import(name="rapidRecheck")
    private @Nullable Output<Boolean> rapidRecheck;

    /**
     * @return If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
     * 
     */
    public Optional<Output<Boolean>> rapidRecheck() {
        return Optional.ofNullable(this.rapidRecheck);
    }

    /**
     * The list of region codes in which to run the monitoring
     * job. See NS1 API docs for supported values.
     * 
     */
    @Import(name="regions", required=true)
    private Output<List<String>> regions;

    /**
     * @return The list of region codes in which to run the monitoring
     * job. See NS1 API docs for supported values.
     * 
     */
    public Output<List<String>> regions() {
        return this.regions;
    }

    /**
     * A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {&#34;key&#34;:&#34;rtt&#34;, &#34;comparison&#34;:&#34;&lt;&#34;, &#34;value&#34;:100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<MonitoringJobRuleArgs>> rules;

    /**
     * @return A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {&#34;key&#34;:&#34;rtt&#34;, &#34;comparison&#34;:&#34;&lt;&#34;, &#34;value&#34;:100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     * 
     */
    public Optional<Output<List<MonitoringJobRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    private MonitoringJobArgs() {}

    private MonitoringJobArgs(MonitoringJobArgs $) {
        this.active = $.active;
        this.config = $.config;
        this.frequency = $.frequency;
        this.jobType = $.jobType;
        this.mute = $.mute;
        this.name = $.name;
        this.notes = $.notes;
        this.notifyDelay = $.notifyDelay;
        this.notifyFailback = $.notifyFailback;
        this.notifyList = $.notifyList;
        this.notifyRegional = $.notifyRegional;
        this.notifyRepeat = $.notifyRepeat;
        this.policy = $.policy;
        this.rapidRecheck = $.rapidRecheck;
        this.regions = $.regions;
        this.rules = $.rules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MonitoringJobArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MonitoringJobArgs $;

        public Builder() {
            $ = new MonitoringJobArgs();
        }

        public Builder(MonitoringJobArgs defaults) {
            $ = new MonitoringJobArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param active Indicates if the job is active or temporarily disabled.
         * 
         * @return builder
         * 
         */
        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        /**
         * @param active Indicates if the job is active or temporarily disabled.
         * 
         * @return builder
         * 
         */
        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        /**
         * @param config A configuration dictionary with keys and values depending on the job_type. Configuration details for each job_type are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
         * 
         * @return builder
         * 
         */
        public Builder config(Output<Map<String,Object>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config A configuration dictionary with keys and values depending on the job_type. Configuration details for each job_type are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,Object> config) {
            return config(Output.of(config));
        }

        /**
         * @param frequency The frequency, in seconds, at which to run the monitoring job in each region.
         * 
         * @return builder
         * 
         */
        public Builder frequency(Output<Integer> frequency) {
            $.frequency = frequency;
            return this;
        }

        /**
         * @param frequency The frequency, in seconds, at which to run the monitoring job in each region.
         * 
         * @return builder
         * 
         */
        public Builder frequency(Integer frequency) {
            return frequency(Output.of(frequency));
        }

        /**
         * @param jobType The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
         * 
         * @return builder
         * 
         */
        public Builder jobType(Output<String> jobType) {
            $.jobType = jobType;
            return this;
        }

        /**
         * @param jobType The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
         * 
         * @return builder
         * 
         */
        public Builder jobType(String jobType) {
            return jobType(Output.of(jobType));
        }

        /**
         * @param mute turn off the notifications for the monitoring job.
         * 
         * @return builder
         * 
         */
        public Builder mute(@Nullable Output<Boolean> mute) {
            $.mute = mute;
            return this;
        }

        /**
         * @param mute turn off the notifications for the monitoring job.
         * 
         * @return builder
         * 
         */
        public Builder mute(Boolean mute) {
            return mute(Output.of(mute));
        }

        /**
         * @param name The free-form display name for the monitoring job.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The free-form display name for the monitoring job.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param notes Freeform notes to be included in any notifications about this job.
         * 
         * @return builder
         * 
         */
        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        /**
         * @param notes Freeform notes to be included in any notifications about this job.
         * 
         * @return builder
         * 
         */
        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        /**
         * @param notifyDelay The time in seconds after a failure to wait before sending a notification.
         * 
         * @return builder
         * 
         */
        public Builder notifyDelay(@Nullable Output<Integer> notifyDelay) {
            $.notifyDelay = notifyDelay;
            return this;
        }

        /**
         * @param notifyDelay The time in seconds after a failure to wait before sending a notification.
         * 
         * @return builder
         * 
         */
        public Builder notifyDelay(Integer notifyDelay) {
            return notifyDelay(Output.of(notifyDelay));
        }

        /**
         * @param notifyFailback If true, a notification is sent when a job returns to an &#34;up&#34; state.
         * 
         * @return builder
         * 
         */
        public Builder notifyFailback(@Nullable Output<Boolean> notifyFailback) {
            $.notifyFailback = notifyFailback;
            return this;
        }

        /**
         * @param notifyFailback If true, a notification is sent when a job returns to an &#34;up&#34; state.
         * 
         * @return builder
         * 
         */
        public Builder notifyFailback(Boolean notifyFailback) {
            return notifyFailback(Output.of(notifyFailback));
        }

        public Builder notifyList(@Nullable Output<String> notifyList) {
            $.notifyList = notifyList;
            return this;
        }

        public Builder notifyList(String notifyList) {
            return notifyList(Output.of(notifyList));
        }

        /**
         * @param notifyRegional If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
         * 
         * @return builder
         * 
         */
        public Builder notifyRegional(@Nullable Output<Boolean> notifyRegional) {
            $.notifyRegional = notifyRegional;
            return this;
        }

        /**
         * @param notifyRegional If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
         * 
         * @return builder
         * 
         */
        public Builder notifyRegional(Boolean notifyRegional) {
            return notifyRegional(Output.of(notifyRegional));
        }

        /**
         * @param notifyRepeat The time in seconds between repeat notifications of a failed job.
         * 
         * @return builder
         * 
         */
        public Builder notifyRepeat(@Nullable Output<Integer> notifyRepeat) {
            $.notifyRepeat = notifyRepeat;
            return this;
        }

        /**
         * @param notifyRepeat The time in seconds between repeat notifications of a failed job.
         * 
         * @return builder
         * 
         */
        public Builder notifyRepeat(Integer notifyRepeat) {
            return notifyRepeat(Output.of(notifyRepeat));
        }

        /**
         * @param policy The policy for determining the monitor&#39;s global status
         * based on the status of the job in all regions. See NS1 API docs for supported values.
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy The policy for determining the monitor&#39;s global status
         * based on the status of the job in all regions. See NS1 API docs for supported values.
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        /**
         * @param rapidRecheck If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
         * 
         * @return builder
         * 
         */
        public Builder rapidRecheck(@Nullable Output<Boolean> rapidRecheck) {
            $.rapidRecheck = rapidRecheck;
            return this;
        }

        /**
         * @param rapidRecheck If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
         * 
         * @return builder
         * 
         */
        public Builder rapidRecheck(Boolean rapidRecheck) {
            return rapidRecheck(Output.of(rapidRecheck));
        }

        /**
         * @param regions The list of region codes in which to run the monitoring
         * job. See NS1 API docs for supported values.
         * 
         * @return builder
         * 
         */
        public Builder regions(Output<List<String>> regions) {
            $.regions = regions;
            return this;
        }

        /**
         * @param regions The list of region codes in which to run the monitoring
         * job. See NS1 API docs for supported values.
         * 
         * @return builder
         * 
         */
        public Builder regions(List<String> regions) {
            return regions(Output.of(regions));
        }

        /**
         * @param regions The list of region codes in which to run the monitoring
         * job. See NS1 API docs for supported values.
         * 
         * @return builder
         * 
         */
        public Builder regions(String... regions) {
            return regions(List.of(regions));
        }

        /**
         * @param rules A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {&#34;key&#34;:&#34;rtt&#34;, &#34;comparison&#34;:&#34;&lt;&#34;, &#34;value&#34;:100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<MonitoringJobRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {&#34;key&#34;:&#34;rtt&#34;, &#34;comparison&#34;:&#34;&lt;&#34;, &#34;value&#34;:100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<MonitoringJobRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {&#34;key&#34;:&#34;rtt&#34;, &#34;comparison&#34;:&#34;&lt;&#34;, &#34;value&#34;:100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
         * 
         * @return builder
         * 
         */
        public Builder rules(MonitoringJobRuleArgs... rules) {
            return rules(List.of(rules));
        }

        public MonitoringJobArgs build() {
            if ($.config == null) {
                throw new MissingRequiredPropertyException("MonitoringJobArgs", "config");
            }
            if ($.frequency == null) {
                throw new MissingRequiredPropertyException("MonitoringJobArgs", "frequency");
            }
            if ($.jobType == null) {
                throw new MissingRequiredPropertyException("MonitoringJobArgs", "jobType");
            }
            if ($.regions == null) {
                throw new MissingRequiredPropertyException("MonitoringJobArgs", "regions");
            }
            return $;
        }
    }

}
