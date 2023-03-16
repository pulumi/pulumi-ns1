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
    'GetNetworksResult',
    'AwaitableGetNetworksResult',
    'get_networks',
]

@pulumi.output_type
class GetNetworksResult:
    """
    A collection of values returned by getNetworks.
    """
    def __init__(__self__, id=None, networks=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetNetworksNetworkResult']:
        """
        A set of the available networks. Networks is
        documented below.
        """
        return pulumi.get(self, "networks")


class AwaitableGetNetworksResult(GetNetworksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworksResult(
            id=self.id,
            networks=self.networks)


def get_networks(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworksResult:
    """
    Provides details about NS1 Networks. Use this if you would simply like to read
    information from NS1 into your configurations. For read/write operations, you
    should use a resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ns1 as ns1

    example = ns1.get_networks()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ns1:index/getNetworks:getNetworks', __args__, opts=opts, typ=GetNetworksResult).value

    return AwaitableGetNetworksResult(
        id=__ret__.id,
        networks=__ret__.networks)
