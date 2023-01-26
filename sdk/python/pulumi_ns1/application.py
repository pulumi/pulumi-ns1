# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 browser_wait_millis: Optional[pulumi.Input[int]] = None,
                 default_config: Optional[pulumi.Input['ApplicationDefaultConfigArgs']] = None,
                 jobs_per_transaction: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Application resource.
        :param pulumi.Input[bool] active: Indicates whether or not this application is currently active and usable for traffic
               steering.
        :param pulumi.Input[int] browser_wait_millis: The amount of time (in milliseconds) the browser should wait before running
               measurements.
        :param pulumi.Input['ApplicationDefaultConfigArgs'] default_config: Default job configuration. If a field is present here and not on a specific job
               associated with this application, the default value specified here is used..
        :param pulumi.Input[int] jobs_per_transaction: Number of jobs to measure per user impression.
        :param pulumi.Input[str] name: Descriptive name for this Pulsar app.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if browser_wait_millis is not None:
            pulumi.set(__self__, "browser_wait_millis", browser_wait_millis)
        if default_config is not None:
            pulumi.set(__self__, "default_config", default_config)
        if jobs_per_transaction is not None:
            pulumi.set(__self__, "jobs_per_transaction", jobs_per_transaction)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether or not this application is currently active and usable for traffic
        steering.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter(name="browserWaitMillis")
    def browser_wait_millis(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of time (in milliseconds) the browser should wait before running
        measurements.
        """
        return pulumi.get(self, "browser_wait_millis")

    @browser_wait_millis.setter
    def browser_wait_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "browser_wait_millis", value)

    @property
    @pulumi.getter(name="defaultConfig")
    def default_config(self) -> Optional[pulumi.Input['ApplicationDefaultConfigArgs']]:
        """
        Default job configuration. If a field is present here and not on a specific job
        associated with this application, the default value specified here is used..
        """
        return pulumi.get(self, "default_config")

    @default_config.setter
    def default_config(self, value: Optional[pulumi.Input['ApplicationDefaultConfigArgs']]):
        pulumi.set(self, "default_config", value)

    @property
    @pulumi.getter(name="jobsPerTransaction")
    def jobs_per_transaction(self) -> Optional[pulumi.Input[int]]:
        """
        Number of jobs to measure per user impression.
        """
        return pulumi.get(self, "jobs_per_transaction")

    @jobs_per_transaction.setter
    def jobs_per_transaction(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "jobs_per_transaction", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Descriptive name for this Pulsar app.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ApplicationState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 browser_wait_millis: Optional[pulumi.Input[int]] = None,
                 default_config: Optional[pulumi.Input['ApplicationDefaultConfigArgs']] = None,
                 jobs_per_transaction: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Application resources.
        :param pulumi.Input[bool] active: Indicates whether or not this application is currently active and usable for traffic
               steering.
        :param pulumi.Input[int] browser_wait_millis: The amount of time (in milliseconds) the browser should wait before running
               measurements.
        :param pulumi.Input['ApplicationDefaultConfigArgs'] default_config: Default job configuration. If a field is present here and not on a specific job
               associated with this application, the default value specified here is used..
        :param pulumi.Input[int] jobs_per_transaction: Number of jobs to measure per user impression.
        :param pulumi.Input[str] name: Descriptive name for this Pulsar app.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if browser_wait_millis is not None:
            pulumi.set(__self__, "browser_wait_millis", browser_wait_millis)
        if default_config is not None:
            pulumi.set(__self__, "default_config", default_config)
        if jobs_per_transaction is not None:
            pulumi.set(__self__, "jobs_per_transaction", jobs_per_transaction)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether or not this application is currently active and usable for traffic
        steering.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter(name="browserWaitMillis")
    def browser_wait_millis(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of time (in milliseconds) the browser should wait before running
        measurements.
        """
        return pulumi.get(self, "browser_wait_millis")

    @browser_wait_millis.setter
    def browser_wait_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "browser_wait_millis", value)

    @property
    @pulumi.getter(name="defaultConfig")
    def default_config(self) -> Optional[pulumi.Input['ApplicationDefaultConfigArgs']]:
        """
        Default job configuration. If a field is present here and not on a specific job
        associated with this application, the default value specified here is used..
        """
        return pulumi.get(self, "default_config")

    @default_config.setter
    def default_config(self, value: Optional[pulumi.Input['ApplicationDefaultConfigArgs']]):
        pulumi.set(self, "default_config", value)

    @property
    @pulumi.getter(name="jobsPerTransaction")
    def jobs_per_transaction(self) -> Optional[pulumi.Input[int]]:
        """
        Number of jobs to measure per user impression.
        """
        return pulumi.get(self, "jobs_per_transaction")

    @jobs_per_transaction.setter
    def jobs_per_transaction(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "jobs_per_transaction", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Descriptive name for this Pulsar app.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Application(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 browser_wait_millis: Optional[pulumi.Input[int]] = None,
                 default_config: Optional[pulumi.Input[pulumi.InputType['ApplicationDefaultConfigArgs']]] = None,
                 jobs_per_transaction: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a NS1 Pulsar application resource. This can be used to create, modify, and delete applications.

        ## NS1 Documentation

        [Application Api Docs](https://ns1.com/api#get-list-pulsar-applications)

        ## Import

        ```sh
         $ pulumi import ns1:index/application:Application `ns1_application`
        ```

         So for the example above

        ```sh
         $ pulumi import ns1:index/application:Application example terraform.example.io`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Indicates whether or not this application is currently active and usable for traffic
               steering.
        :param pulumi.Input[int] browser_wait_millis: The amount of time (in milliseconds) the browser should wait before running
               measurements.
        :param pulumi.Input[pulumi.InputType['ApplicationDefaultConfigArgs']] default_config: Default job configuration. If a field is present here and not on a specific job
               associated with this application, the default value specified here is used..
        :param pulumi.Input[int] jobs_per_transaction: Number of jobs to measure per user impression.
        :param pulumi.Input[str] name: Descriptive name for this Pulsar app.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ApplicationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NS1 Pulsar application resource. This can be used to create, modify, and delete applications.

        ## NS1 Documentation

        [Application Api Docs](https://ns1.com/api#get-list-pulsar-applications)

        ## Import

        ```sh
         $ pulumi import ns1:index/application:Application `ns1_application`
        ```

         So for the example above

        ```sh
         $ pulumi import ns1:index/application:Application example terraform.example.io`
        ```

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 browser_wait_millis: Optional[pulumi.Input[int]] = None,
                 default_config: Optional[pulumi.Input[pulumi.InputType['ApplicationDefaultConfigArgs']]] = None,
                 jobs_per_transaction: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            __props__.__dict__["active"] = active
            __props__.__dict__["browser_wait_millis"] = browser_wait_millis
            __props__.__dict__["default_config"] = default_config
            __props__.__dict__["jobs_per_transaction"] = jobs_per_transaction
            __props__.__dict__["name"] = name
        super(Application, __self__).__init__(
            'ns1:index/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            browser_wait_millis: Optional[pulumi.Input[int]] = None,
            default_config: Optional[pulumi.Input[pulumi.InputType['ApplicationDefaultConfigArgs']]] = None,
            jobs_per_transaction: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Indicates whether or not this application is currently active and usable for traffic
               steering.
        :param pulumi.Input[int] browser_wait_millis: The amount of time (in milliseconds) the browser should wait before running
               measurements.
        :param pulumi.Input[pulumi.InputType['ApplicationDefaultConfigArgs']] default_config: Default job configuration. If a field is present here and not on a specific job
               associated with this application, the default value specified here is used..
        :param pulumi.Input[int] jobs_per_transaction: Number of jobs to measure per user impression.
        :param pulumi.Input[str] name: Descriptive name for this Pulsar app.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationState.__new__(_ApplicationState)

        __props__.__dict__["active"] = active
        __props__.__dict__["browser_wait_millis"] = browser_wait_millis
        __props__.__dict__["default_config"] = default_config
        __props__.__dict__["jobs_per_transaction"] = jobs_per_transaction
        __props__.__dict__["name"] = name
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether or not this application is currently active and usable for traffic
        steering.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter(name="browserWaitMillis")
    def browser_wait_millis(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of time (in milliseconds) the browser should wait before running
        measurements.
        """
        return pulumi.get(self, "browser_wait_millis")

    @property
    @pulumi.getter(name="defaultConfig")
    def default_config(self) -> pulumi.Output[Optional['outputs.ApplicationDefaultConfig']]:
        """
        Default job configuration. If a field is present here and not on a specific job
        associated with this application, the default value specified here is used..
        """
        return pulumi.get(self, "default_config")

    @property
    @pulumi.getter(name="jobsPerTransaction")
    def jobs_per_transaction(self) -> pulumi.Output[Optional[int]]:
        """
        Number of jobs to measure per user impression.
        """
        return pulumi.get(self, "jobs_per_transaction")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Descriptive name for this Pulsar app.
        """
        return pulumi.get(self, "name")

