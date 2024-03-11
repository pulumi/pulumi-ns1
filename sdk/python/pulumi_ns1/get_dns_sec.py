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

__all__ = [
    'GetDNSSecResult',
    'AwaitableGetDNSSecResult',
    'get_dns_sec',
    'get_dns_sec_output',
]

@pulumi.output_type
class GetDNSSecResult:
    """
    A collection of values returned by getDNSSec.
    """
    def __init__(__self__, delegations=None, id=None, keys=None, zone=None):
        if delegations and not isinstance(delegations, list):
            raise TypeError("Expected argument 'delegations' to be a list")
        pulumi.set(__self__, "delegations", delegations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def delegations(self) -> Sequence['outputs.GetDNSSecDelegationResult']:
        """
        (Computed) - Delegation field is documented
        below.
        """
        return pulumi.get(self, "delegations")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keys(self) -> Sequence['outputs.GetDNSSecKeyResult']:
        """
        (Computed) - Keys field is documented below.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetDNSSecResult(GetDNSSecResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDNSSecResult(
            delegations=self.delegations,
            id=self.id,
            keys=self.keys,
            zone=self.zone)


def get_dns_sec(zone: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDNSSecResult:
    """
    Provides DNSSEC details about a NS1 Zone.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ns1 as ns1

    # Get DNSSEC details about a NS1 Zone.
    example_zone = ns1.Zone("exampleZone",
        dnssec=True,
        zone="terraform.example.io")
    example_dns_sec = ns1.get_dns_sec_output(zone=example_zone.zone)
    ```
    <!--End PulumiCodeChooser -->


    :param str zone: The name of the zone to get DNSSEC details for.
    """
    __args__ = dict()
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ns1:index/getDNSSec:getDNSSec', __args__, opts=opts, typ=GetDNSSecResult).value

    return AwaitableGetDNSSecResult(
        delegations=pulumi.get(__ret__, 'delegations'),
        id=pulumi.get(__ret__, 'id'),
        keys=pulumi.get(__ret__, 'keys'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_dns_sec)
def get_dns_sec_output(zone: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDNSSecResult]:
    """
    Provides DNSSEC details about a NS1 Zone.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ns1 as ns1

    # Get DNSSEC details about a NS1 Zone.
    example_zone = ns1.Zone("exampleZone",
        dnssec=True,
        zone="terraform.example.io")
    example_dns_sec = ns1.get_dns_sec_output(zone=example_zone.zone)
    ```
    <!--End PulumiCodeChooser -->


    :param str zone: The name of the zone to get DNSSEC details for.
    """
    ...
