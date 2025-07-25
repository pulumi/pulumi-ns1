# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['DataFeedArgs', 'DataFeed']

@pulumi.input_type
class DataFeedArgs:
    def __init__(__self__, *,
                 source_id: pulumi.Input[_builtins.str],
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a DataFeed resource.
        :param pulumi.Input[_builtins.str] source_id: The data source id that this feed is connected to.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] config: The feeds configuration matching the specification in
               `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
        :param pulumi.Input[_builtins.str] name: The free form name of the data feed.
        """
        pulumi.set(__self__, "source_id", source_id)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @_builtins.property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> pulumi.Input[_builtins.str]:
        """
        The data source id that this feed is connected to.
        """
        return pulumi.get(self, "source_id")

    @source_id.setter
    def source_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "source_id", value)

    @_builtins.property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        The feeds configuration matching the specification in
        `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "config", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The free form name of the data feed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DataFeedState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 source_id: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering DataFeed resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] config: The feeds configuration matching the specification in
               `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
        :param pulumi.Input[_builtins.str] name: The free form name of the data feed.
        :param pulumi.Input[_builtins.str] source_id: The data source id that this feed is connected to.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if source_id is not None:
            pulumi.set(__self__, "source_id", source_id)

    @_builtins.property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        The feeds configuration matching the specification in
        `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "config", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The free form name of the data feed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The data source id that this feed is connected to.
        """
        return pulumi.get(self, "source_id")

    @source_id.setter
    def source_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "source_id", value)


@pulumi.type_token("ns1:index/dataFeed:DataFeed")
class DataFeed(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 source_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Provides a NS1 Data Feed resource. This can be used to create, modify, and delete data feeds.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.DataSource("example",
            name="example",
            sourcetype="nsone_v1")
        example_monitoring = ns1.DataSource("example_monitoring",
            name="example_monitoring",
            sourcetype="nsone_monitoring")
        uswest_feed = ns1.DataFeed("uswest_feed",
            name="uswest_feed",
            source_id=example.id,
            config={
                "label": "uswest",
            })
        useast_feed = ns1.DataFeed("useast_feed",
            name="useast_feed",
            source_id=example.id,
            config={
                "label": "useast",
            })
        useast_monitor_feed = ns1.DataFeed("useast_monitor_feed",
            name="useast_monitor_feed",
            source_id=example_monitoring.id,
            config={
                "jobid": example_job["id"],
            })
        ```

        ## NS1 Documentation

        [Datafeed Api Doc](https://ns1.com/api#data-feeds)

        ## Import

        ```sh
        $ pulumi import ns1:index/dataFeed:DataFeed <name> <datasource_id>/<datafeed_id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] config: The feeds configuration matching the specification in
               `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
        :param pulumi.Input[_builtins.str] name: The free form name of the data feed.
        :param pulumi.Input[_builtins.str] source_id: The data source id that this feed is connected to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataFeedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NS1 Data Feed resource. This can be used to create, modify, and delete data feeds.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.DataSource("example",
            name="example",
            sourcetype="nsone_v1")
        example_monitoring = ns1.DataSource("example_monitoring",
            name="example_monitoring",
            sourcetype="nsone_monitoring")
        uswest_feed = ns1.DataFeed("uswest_feed",
            name="uswest_feed",
            source_id=example.id,
            config={
                "label": "uswest",
            })
        useast_feed = ns1.DataFeed("useast_feed",
            name="useast_feed",
            source_id=example.id,
            config={
                "label": "useast",
            })
        useast_monitor_feed = ns1.DataFeed("useast_monitor_feed",
            name="useast_monitor_feed",
            source_id=example_monitoring.id,
            config={
                "jobid": example_job["id"],
            })
        ```

        ## NS1 Documentation

        [Datafeed Api Doc](https://ns1.com/api#data-feeds)

        ## Import

        ```sh
        $ pulumi import ns1:index/dataFeed:DataFeed <name> <datasource_id>/<datafeed_id>`
        ```

        :param str resource_name: The name of the resource.
        :param DataFeedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataFeedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 source_id: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataFeedArgs.__new__(DataFeedArgs)

            __props__.__dict__["config"] = config
            __props__.__dict__["name"] = name
            if source_id is None and not opts.urn:
                raise TypeError("Missing required property 'source_id'")
            __props__.__dict__["source_id"] = source_id
        super(DataFeed, __self__).__init__(
            'ns1:index/dataFeed:DataFeed',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            source_id: Optional[pulumi.Input[_builtins.str]] = None) -> 'DataFeed':
        """
        Get an existing DataFeed resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] config: The feeds configuration matching the specification in
               `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
        :param pulumi.Input[_builtins.str] name: The free form name of the data feed.
        :param pulumi.Input[_builtins.str] source_id: The data source id that this feed is connected to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataFeedState.__new__(_DataFeedState)

        __props__.__dict__["config"] = config
        __props__.__dict__["name"] = name
        __props__.__dict__["source_id"] = source_id
        return DataFeed(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def config(self) -> pulumi.Output[Optional[Mapping[str, _builtins.str]]]:
        """
        The feeds configuration matching the specification in
        `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
        """
        return pulumi.get(self, "config")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The free form name of the data feed.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="sourceId")
    def source_id(self) -> pulumi.Output[_builtins.str]:
        """
        The data source id that this feed is connected to.
        """
        return pulumi.get(self, "source_id")

