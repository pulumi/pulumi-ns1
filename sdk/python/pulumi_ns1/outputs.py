# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs

__all__ = [
    'MonitoringJobRule',
    'NotifyListNotification',
    'RecordAnswer',
    'RecordFilter',
    'RecordRegion',
    'TeamIpWhitelist',
    'ZoneSecondary',
    'GetDNSSecDelegationResult',
    'GetDNSSecDelegationDResult',
    'GetDNSSecDelegationDnskeyResult',
    'GetDNSSecKeysResult',
    'GetDNSSecKeysDnskeyResult',
    'GetZoneSecondaryResult',
]

@pulumi.output_type
class MonitoringJobRule(dict):
    def __init__(__self__, *,
                 comparison: str,
                 key: str,
                 value: str):
        """
        :param str comparison: The comparison to perform on the the output.
        :param str key: The output key.
        :param str value: The value to compare to.
        """
        pulumi.set(__self__, "comparison", comparison)
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def comparison(self) -> str:
        """
        The comparison to perform on the the output.
        """
        return pulumi.get(self, "comparison")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The output key.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value to compare to.
        """
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class NotifyListNotification(dict):
    def __init__(__self__, *,
                 config: Mapping[str, Any],
                 type: str):
        """
        :param Mapping[str, Any] config: Configuration details for the given notifier type.
        :param str type: The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
        """
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def config(self) -> Mapping[str, Any]:
        """
        Configuration details for the given notifier type.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RecordAnswer(dict):
    def __init__(__self__, *,
                 answer: Optional[str] = None,
                 meta: Optional[Mapping[str, Any]] = None,
                 region: Optional[str] = None):
        """
        :param str answer: Space delimited string of RDATA fields dependent on the record type.
        :param str region: The region (Answer Group really) that this answer
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
    def answer(self) -> Optional[str]:
        """
        Space delimited string of RDATA fields dependent on the record type.
        """
        return pulumi.get(self, "answer")

    @property
    @pulumi.getter
    def meta(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
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

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RecordFilter(dict):
    def __init__(__self__, *,
                 filter: str,
                 config: Optional[Mapping[str, Any]] = None,
                 disabled: Optional[bool] = None):
        """
        :param str filter: The type of filter.
        :param Mapping[str, Any] config: The filters' configuration. Simple key/value pairs
               determined by the filter type.
        :param bool disabled: Determines whether the filter is applied in the
               filter chain.
        """
        pulumi.set(__self__, "filter", filter)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)

    @property
    @pulumi.getter
    def filter(self) -> str:
        """
        The type of filter.
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def config(self) -> Optional[Mapping[str, Any]]:
        """
        The filters' configuration. Simple key/value pairs
        determined by the filter type.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def disabled(self) -> Optional[bool]:
        """
        Determines whether the filter is applied in the
        filter chain.
        """
        return pulumi.get(self, "disabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class RecordRegion(dict):
    def __init__(__self__, *,
                 name: str,
                 meta: Optional[Mapping[str, Any]] = None):
        """
        :param str name: Name of the region (or Answer Group).
        """
        pulumi.set(__self__, "name", name)
        if meta is not None:
            pulumi.set(__self__, "meta", meta)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the region (or Answer Group).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def meta(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "meta")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TeamIpWhitelist(dict):
    def __init__(__self__, *,
                 name: str,
                 values: List[str]):
        """
        :param str name: The free form name of the team.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The free form name of the team.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> List[str]:
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ZoneSecondary(dict):
    def __init__(__self__, *,
                 ip: str,
                 networks: Optional[List[float]] = None,
                 notify: Optional[bool] = None,
                 port: Optional[float] = None):
        """
        :param str ip: IPv4 address of the secondary server.
        :param List[float] networks: - List of network IDs (`int`) for which the zone
               should be made available. Default is network 0, the primary NSONE Global
               Network. Normally, you should not have to worry about this.
        :param bool notify: Whether we send `NOTIFY` messages to the secondary host
               when the zone changes. Default `false`.
        :param float port: Port of the the secondary server. Default `53`.
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
    def ip(self) -> str:
        """
        IPv4 address of the secondary server.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def networks(self) -> Optional[List[float]]:
        """
        - List of network IDs (`int`) for which the zone
        should be made available. Default is network 0, the primary NSONE Global
        Network. Normally, you should not have to worry about this.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def notify(self) -> Optional[bool]:
        """
        Whether we send `NOTIFY` messages to the secondary host
        when the zone changes. Default `false`.
        """
        return pulumi.get(self, "notify")

    @property
    @pulumi.getter
    def port(self) -> Optional[float]:
        """
        Port of the the secondary server. Default `53`.
        """
        return pulumi.get(self, "port")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetDNSSecDelegationResult(dict):
    def __init__(__self__, *,
                 dnskeys: List['outputs.GetDNSSecDelegationDnskeyResult'],
                 ds: List['outputs.GetDNSSecDelegationDResult'],
                 ttl: float):
        """
        :param List['GetDNSSecDelegationDnskeyArgs'] dnskeys: (Computed) List of Keys. Key is documented below.
        :param List['GetDNSSecDelegationDArgs'] ds: (Computed) List of Keys. Key is documented below.
        :param float ttl: (Computed) TTL for the Keys (int).
        """
        pulumi.set(__self__, "dnskeys", dnskeys)
        pulumi.set(__self__, "ds", ds)
        pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def dnskeys(self) -> List['outputs.GetDNSSecDelegationDnskeyResult']:
        """
        (Computed) List of Keys. Key is documented below.
        """
        return pulumi.get(self, "dnskeys")

    @property
    @pulumi.getter
    def ds(self) -> List['outputs.GetDNSSecDelegationDResult']:
        """
        (Computed) List of Keys. Key is documented below.
        """
        return pulumi.get(self, "ds")

    @property
    @pulumi.getter
    def ttl(self) -> float:
        """
        (Computed) TTL for the Keys (int).
        """
        return pulumi.get(self, "ttl")


@pulumi.output_type
class GetDNSSecDelegationDResult(dict):
    def __init__(__self__, *,
                 algorithm: str,
                 flags: str,
                 protocol: str,
                 public_key: str):
        """
        :param str algorithm: (Computed) Algorithm of the key.
        :param str flags: (Computed) Flags for the key.
        :param str protocol: (Computed) Protocol of the key.
        :param str public_key: (Computed) Public key for the key.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "flags", flags)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        (Computed) Algorithm of the key.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def flags(self) -> str:
        """
        (Computed) Flags for the key.
        """
        return pulumi.get(self, "flags")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        (Computed) Protocol of the key.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> str:
        """
        (Computed) Public key for the key.
        """
        return pulumi.get(self, "public_key")


@pulumi.output_type
class GetDNSSecDelegationDnskeyResult(dict):
    def __init__(__self__, *,
                 algorithm: str,
                 flags: str,
                 protocol: str,
                 public_key: str):
        """
        :param str algorithm: (Computed) Algorithm of the key.
        :param str flags: (Computed) Flags for the key.
        :param str protocol: (Computed) Protocol of the key.
        :param str public_key: (Computed) Public key for the key.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "flags", flags)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        (Computed) Algorithm of the key.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def flags(self) -> str:
        """
        (Computed) Flags for the key.
        """
        return pulumi.get(self, "flags")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        (Computed) Protocol of the key.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> str:
        """
        (Computed) Public key for the key.
        """
        return pulumi.get(self, "public_key")


@pulumi.output_type
class GetDNSSecKeysResult(dict):
    def __init__(__self__, *,
                 dnskeys: List['outputs.GetDNSSecKeysDnskeyResult'],
                 ttl: float):
        """
        :param List['GetDNSSecKeysDnskeyArgs'] dnskeys: (Computed) List of Keys. Key is documented below.
        :param float ttl: (Computed) TTL for the Keys (int).
        """
        pulumi.set(__self__, "dnskeys", dnskeys)
        pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def dnskeys(self) -> List['outputs.GetDNSSecKeysDnskeyResult']:
        """
        (Computed) List of Keys. Key is documented below.
        """
        return pulumi.get(self, "dnskeys")

    @property
    @pulumi.getter
    def ttl(self) -> float:
        """
        (Computed) TTL for the Keys (int).
        """
        return pulumi.get(self, "ttl")


@pulumi.output_type
class GetDNSSecKeysDnskeyResult(dict):
    def __init__(__self__, *,
                 algorithm: str,
                 flags: str,
                 protocol: str,
                 public_key: str):
        """
        :param str algorithm: (Computed) Algorithm of the key.
        :param str flags: (Computed) Flags for the key.
        :param str protocol: (Computed) Protocol of the key.
        :param str public_key: (Computed) Public key for the key.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "flags", flags)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        (Computed) Algorithm of the key.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def flags(self) -> str:
        """
        (Computed) Flags for the key.
        """
        return pulumi.get(self, "flags")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        (Computed) Protocol of the key.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> str:
        """
        (Computed) Public key for the key.
        """
        return pulumi.get(self, "public_key")


@pulumi.output_type
class GetZoneSecondaryResult(dict):
    def __init__(__self__, *,
                 ip: str,
                 networks: List[float],
                 notify: bool,
                 port: float):
        """
        :param str ip: IPv4 address of the secondary server.
        :param List[float] networks: List of network IDs (`int`) for which the zone should be made
               available. Default is network 0, the primary NSONE Global Network.
        :param bool notify: Whether we send `NOTIFY` messages to the secondary host
               when the zone changes. Default `false`.
        :param float port: Port of the the secondary server. Default `53`.
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "networks", networks)
        pulumi.set(__self__, "notify", notify)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        IPv4 address of the secondary server.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def networks(self) -> List[float]:
        """
        List of network IDs (`int`) for which the zone should be made
        available. Default is network 0, the primary NSONE Global Network.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def notify(self) -> bool:
        """
        Whether we send `NOTIFY` messages to the secondary host
        when the zone changes. Default `false`.
        """
        return pulumi.get(self, "notify")

    @property
    @pulumi.getter
    def port(self) -> float:
        """
        Port of the the secondary server. Default `53`.
        """
        return pulumi.get(self, "port")


