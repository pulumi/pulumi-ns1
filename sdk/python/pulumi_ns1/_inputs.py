# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'MonitoringJobRuleArgs',
    'NotifyListNotificationArgs',
    'RecordAnswerArgs',
    'RecordFilterArgs',
    'RecordRegionArgs',
    'TeamIpWhitelistArgs',
    'ZoneSecondaryArgs',
]

@pulumi.input_type
class MonitoringJobRuleArgs:
    def __init__(__self__, *,
                 comparison: pulumi.Input[str],
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "comparison", comparison)
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def comparison(self) -> pulumi.Input[str]:
        return pulumi.get(self, "comparison")

    @comparison.setter
    def comparison(self, value: pulumi.Input[str]):
        pulumi.set(self, "comparison", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class NotifyListNotificationArgs:
    def __init__(__self__, *,
                 config: pulumi.Input[Mapping[str, Any]],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[Mapping[str, Any]] config: Configuration details for the given notifier type.
        :param pulumi.Input[str] type: The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        Configuration details for the given notifier type.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class RecordAnswerArgs:
    def __init__(__self__, *,
                 answer: Optional[pulumi.Input[str]] = None,
                 meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] answer: Space delimited string of RDATA fields dependent on the record type.
        :param pulumi.Input[str] region: The region (Answer Group really) that this answer
               belongs to. This should be one of the names specified in `regions`. Only a
               single `region` per answer is currently supported. If you want an answer in
               multiple regions, duplicating the answer (including metadata) is the correct
               approach.
               * ` meta` - (Optional) meta is supported at the `answer` level. Meta
               is documented below.
        """
        if answer is not None:
            pulumi.set(__self__, "answer", answer)
        if meta is not None:
            pulumi.set(__self__, "meta", meta)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def answer(self) -> Optional[pulumi.Input[str]]:
        """
        Space delimited string of RDATA fields dependent on the record type.
        """
        return pulumi.get(self, "answer")

    @answer.setter
    def answer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "answer", value)

    @property
    @pulumi.getter
    def meta(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "meta")

    @meta.setter
    def meta(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "meta", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region (Answer Group really) that this answer
        belongs to. This should be one of the names specified in `regions`. Only a
        single `region` per answer is currently supported. If you want an answer in
        multiple regions, duplicating the answer (including metadata) is the correct
        approach.
        * ` meta` - (Optional) meta is supported at the `answer` level. Meta
        is documented below.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class RecordFilterArgs:
    def __init__(__self__, *,
                 filter: pulumi.Input[str],
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] filter: The type of filter.
        :param pulumi.Input[Mapping[str, Any]] config: The filters' configuration. Simple key/value pairs
               determined by the filter type.
        :param pulumi.Input[bool] disabled: Determines whether the filter is applied in the
               filter chain.
        """
        pulumi.set(__self__, "filter", filter)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Input[str]:
        """
        The type of filter.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: pulumi.Input[str]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The filters' configuration. Simple key/value pairs
        determined by the filter type.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether the filter is applied in the
        filter chain.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)


@pulumi.input_type
class RecordRegionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 meta: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        :param pulumi.Input[str] name: Name of the region (or Answer Group).
        """
        pulumi.set(__self__, "name", name)
        if meta is not None:
            pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the region (or Answer Group).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def meta(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "meta")

    @meta.setter
    def meta(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "meta", value)


@pulumi.input_type
class TeamIpWhitelistArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 values: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[str] name: The free form name of the team.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The free form name of the team.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class ZoneSecondaryArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 notify: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] ip: IPv4 address of the secondary server.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] networks: - List of network IDs (`int`) for which the zone
               should be made available. Default is network 0, the primary NSONE Global
               Network. Normally, you should not have to worry about this.
        :param pulumi.Input[bool] notify: Whether we send `NOTIFY` messages to the secondary host
               when the zone changes. Default `false`.
        :param pulumi.Input[int] port: Port of the the secondary server. Default `53`.
        """
        pulumi.set(__self__, "ip", ip)
        if networks is not None:
            pulumi.set(__self__, "networks", networks)
        if notify is not None:
            pulumi.set(__self__, "notify", notify)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IPv4 address of the secondary server.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        - List of network IDs (`int`) for which the zone
        should be made available. Default is network 0, the primary NSONE Global
        Network. Normally, you should not have to worry about this.
        """
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter
    def notify(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether we send `NOTIFY` messages to the secondary host
        when the zone changes. Default `false`.
        """
        return pulumi.get(self, "notify")

    @notify.setter
    def notify(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "notify", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port of the the secondary server. Default `53`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


