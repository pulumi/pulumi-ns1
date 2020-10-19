# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['MonitoringJob']


class MonitoringJob(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 frequency: Optional[pulumi.Input[int]] = None,
                 job_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 notify_delay: Optional[pulumi.Input[int]] = None,
                 notify_failback: Optional[pulumi.Input[bool]] = None,
                 notify_list: Optional[pulumi.Input[str]] = None,
                 notify_regional: Optional[pulumi.Input[bool]] = None,
                 notify_repeat: Optional[pulumi.Input[int]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 rapid_recheck: Optional[pulumi.Input[bool]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MonitoringJobRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a NS1 Monitoring Job resource. This can be used to create, modify, and delete monitoring jobs.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        uswest_monitor = ns1.MonitoringJob("uswestMonitor",
            active=True,
            config={
                "host": "example-elb-uswest.aws.amazon.com",
                "port": 443,
                "send": \"\"\"HEAD / HTTP/1.0
        

        \"\"\",
                "ssl": 1,
            },
            frequency=60,
            job_type="tcp",
            policy="quorum",
            rapid_recheck=True,
            regions=[
                "sjc",
                "sin",
                "lga",
            ],
            rules=[ns1.MonitoringJobRuleArgs(
                comparison="contains",
                key="output",
                value="200 OK",
            )])
        ```
        ## NS1 Documentation

        [MonitoringJob Api Doc](https://ns1.com/api#monitoring-jobs)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Indicates if the job is active or temporarily disabled.
        :param pulumi.Input[Mapping[str, Any]] config: A configuration dictionary with keys and values depending on the job_type. Configuration details for each job_type are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
        :param pulumi.Input[int] frequency: The frequency, in seconds, at which to run the monitoring job in each region.
        :param pulumi.Input[str] job_type: The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
        :param pulumi.Input[str] name: The free-form display name for the monitoring job.
        :param pulumi.Input[str] notes: Freeform notes to be included in any notifications about this job.
        :param pulumi.Input[int] notify_delay: The time in seconds after a failure to wait before sending a notification.
        :param pulumi.Input[bool] notify_failback: If true, a notification is sent when a job returns to an "up" state.
        :param pulumi.Input[bool] notify_regional: If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
        :param pulumi.Input[int] notify_repeat: The time in seconds between repeat notifications of a failed job.
        :param pulumi.Input[str] policy: The policy for determining the monitor's global status
               based on the status of the job in all regions. See NS1 API docs for supported values.
        :param pulumi.Input[bool] rapid_recheck: If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: The list of region codes in which to run the monitoring
               job. See NS1 API docs for supported values.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MonitoringJobRuleArgs']]]] rules: A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['active'] = active
            if config is None:
                raise TypeError("Missing required property 'config'")
            __props__['config'] = config
            if frequency is None:
                raise TypeError("Missing required property 'frequency'")
            __props__['frequency'] = frequency
            if job_type is None:
                raise TypeError("Missing required property 'job_type'")
            __props__['job_type'] = job_type
            __props__['name'] = name
            __props__['notes'] = notes
            __props__['notify_delay'] = notify_delay
            __props__['notify_failback'] = notify_failback
            __props__['notify_list'] = notify_list
            __props__['notify_regional'] = notify_regional
            __props__['notify_repeat'] = notify_repeat
            __props__['policy'] = policy
            __props__['rapid_recheck'] = rapid_recheck
            if regions is None:
                raise TypeError("Missing required property 'regions'")
            __props__['regions'] = regions
            __props__['rules'] = rules
        super(MonitoringJob, __self__).__init__(
            'ns1:index/monitoringJob:MonitoringJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            frequency: Optional[pulumi.Input[int]] = None,
            job_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            notify_delay: Optional[pulumi.Input[int]] = None,
            notify_failback: Optional[pulumi.Input[bool]] = None,
            notify_list: Optional[pulumi.Input[str]] = None,
            notify_regional: Optional[pulumi.Input[bool]] = None,
            notify_repeat: Optional[pulumi.Input[int]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            rapid_recheck: Optional[pulumi.Input[bool]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MonitoringJobRuleArgs']]]]] = None) -> 'MonitoringJob':
        """
        Get an existing MonitoringJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Indicates if the job is active or temporarily disabled.
        :param pulumi.Input[Mapping[str, Any]] config: A configuration dictionary with keys and values depending on the job_type. Configuration details for each job_type are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
        :param pulumi.Input[int] frequency: The frequency, in seconds, at which to run the monitoring job in each region.
        :param pulumi.Input[str] job_type: The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
        :param pulumi.Input[str] name: The free-form display name for the monitoring job.
        :param pulumi.Input[str] notes: Freeform notes to be included in any notifications about this job.
        :param pulumi.Input[int] notify_delay: The time in seconds after a failure to wait before sending a notification.
        :param pulumi.Input[bool] notify_failback: If true, a notification is sent when a job returns to an "up" state.
        :param pulumi.Input[bool] notify_regional: If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
        :param pulumi.Input[int] notify_repeat: The time in seconds between repeat notifications of a failed job.
        :param pulumi.Input[str] policy: The policy for determining the monitor's global status
               based on the status of the job in all regions. See NS1 API docs for supported values.
        :param pulumi.Input[bool] rapid_recheck: If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] regions: The list of region codes in which to run the monitoring
               job. See NS1 API docs for supported values.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MonitoringJobRuleArgs']]]] rules: A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["active"] = active
        __props__["config"] = config
        __props__["frequency"] = frequency
        __props__["job_type"] = job_type
        __props__["name"] = name
        __props__["notes"] = notes
        __props__["notify_delay"] = notify_delay
        __props__["notify_failback"] = notify_failback
        __props__["notify_list"] = notify_list
        __props__["notify_regional"] = notify_regional
        __props__["notify_repeat"] = notify_repeat
        __props__["policy"] = policy
        __props__["rapid_recheck"] = rapid_recheck
        __props__["regions"] = regions
        __props__["rules"] = rules
        return MonitoringJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the job is active or temporarily disabled.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        A configuration dictionary with keys and values depending on the job_type. Configuration details for each job_type are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def frequency(self) -> pulumi.Output[int]:
        """
        The frequency, in seconds, at which to run the monitoring job in each region.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> pulumi.Output[str]:
        """
        The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
        """
        return pulumi.get(self, "job_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The free-form display name for the monitoring job.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        """
        Freeform notes to be included in any notifications about this job.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="notifyDelay")
    def notify_delay(self) -> pulumi.Output[Optional[int]]:
        """
        The time in seconds after a failure to wait before sending a notification.
        """
        return pulumi.get(self, "notify_delay")

    @property
    @pulumi.getter(name="notifyFailback")
    def notify_failback(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, a notification is sent when a job returns to an "up" state.
        """
        return pulumi.get(self, "notify_failback")

    @property
    @pulumi.getter(name="notifyList")
    def notify_list(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notify_list")

    @property
    @pulumi.getter(name="notifyRegional")
    def notify_regional(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
        """
        return pulumi.get(self, "notify_regional")

    @property
    @pulumi.getter(name="notifyRepeat")
    def notify_repeat(self) -> pulumi.Output[Optional[int]]:
        """
        The time in seconds between repeat notifications of a failed job.
        """
        return pulumi.get(self, "notify_repeat")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[Optional[str]]:
        """
        The policy for determining the monitor's global status
        based on the status of the job in all regions. See NS1 API docs for supported values.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="rapidRecheck")
    def rapid_recheck(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
        """
        return pulumi.get(self, "rapid_recheck")

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of region codes in which to run the monitoring
        job. See NS1 API docs for supported values.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.MonitoringJobRule']]]:
        """
        A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
        """
        return pulumi.get(self, "rules")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

