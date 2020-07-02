# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetZoneResult:
    """
    A collection of values returned by getZone.
    """
    def __init__(__self__, additional_primaries=None, dns_servers=None, dnssec=None, expiry=None, hostmaster=None, id=None, link=None, networks=None, nx_ttl=None, primary=None, refresh=None, retry=None, secondaries=None, ttl=None, zone=None):
        if additional_primaries and not isinstance(additional_primaries, list):
            raise TypeError("Expected argument 'additional_primaries' to be a list")
        __self__.additional_primaries = additional_primaries
        """
        List of additional IPv4 addresses for the primary
        zone.
        """
        if dns_servers and not isinstance(dns_servers, str):
            raise TypeError("Expected argument 'dns_servers' to be a str")
        __self__.dns_servers = dns_servers
        """
        Authoritative Name Servers.
        """
        if dnssec and not isinstance(dnssec, bool):
            raise TypeError("Expected argument 'dnssec' to be a bool")
        __self__.dnssec = dnssec
        """
        Whether or not DNSSEC is enabled for the zone.
        """
        if expiry and not isinstance(expiry, float):
            raise TypeError("Expected argument 'expiry' to be a float")
        __self__.expiry = expiry
        """
        The SOA Expiry.
        """
        if hostmaster and not isinstance(hostmaster, str):
            raise TypeError("Expected argument 'hostmaster' to be a str")
        __self__.hostmaster = hostmaster
        """
        The SOA Hostmaster.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if link and not isinstance(link, str):
            raise TypeError("Expected argument 'link' to be a str")
        __self__.link = link
        """
        The linked target zone.
        """
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        __self__.networks = networks
        """
        List of network IDs (`int`) for which the zone should be made
        available. Default is network 0, the primary NSONE Global Network.
        """
        if nx_ttl and not isinstance(nx_ttl, float):
            raise TypeError("Expected argument 'nx_ttl' to be a float")
        __self__.nx_ttl = nx_ttl
        """
        The SOA NX TTL.
        """
        if primary and not isinstance(primary, str):
            raise TypeError("Expected argument 'primary' to be a str")
        __self__.primary = primary
        """
        The primary zones' IPv4 address.
        """
        if refresh and not isinstance(refresh, float):
            raise TypeError("Expected argument 'refresh' to be a float")
        __self__.refresh = refresh
        """
        The SOA Refresh.
        """
        if retry and not isinstance(retry, float):
            raise TypeError("Expected argument 'retry' to be a float")
        __self__.retry = retry
        """
        The SOA Retry.
        """
        if secondaries and not isinstance(secondaries, list):
            raise TypeError("Expected argument 'secondaries' to be a list")
        __self__.secondaries = secondaries
        """
        List of secondary servers. Secondaries is
        documented below.
        """
        if ttl and not isinstance(ttl, float):
            raise TypeError("Expected argument 'ttl' to be a float")
        __self__.ttl = ttl
        """
        The SOA TTL.
        """
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        __self__.zone = zone
class AwaitableGetZoneResult(GetZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneResult(
            additional_primaries=self.additional_primaries,
            dns_servers=self.dns_servers,
            dnssec=self.dnssec,
            expiry=self.expiry,
            hostmaster=self.hostmaster,
            id=self.id,
            link=self.link,
            networks=self.networks,
            nx_ttl=self.nx_ttl,
            primary=self.primary,
            refresh=self.refresh,
            retry=self.retry,
            secondaries=self.secondaries,
            ttl=self.ttl,
            zone=self.zone)

def get_zone(additional_primaries=None,zone=None,opts=None):
    """
    Provides details about a NS1 Zone. Use this if you would simply like to read
    information from NS1 into your configurations. For read/write operations, you
    should use a resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ns1 as ns1

    example = ns1.get_zone(zone="terraform.example.io")
    ```


    :param list additional_primaries: List of additional IPv4 addresses for the primary
           zone.
    :param str zone: The domain name of the zone.
    """
    __args__ = dict()


    __args__['additionalPrimaries'] = additional_primaries
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('ns1:index/getZone:getZone', __args__, opts=opts).value

    return AwaitableGetZoneResult(
        additional_primaries=__ret__.get('additionalPrimaries'),
        dns_servers=__ret__.get('dnsServers'),
        dnssec=__ret__.get('dnssec'),
        expiry=__ret__.get('expiry'),
        hostmaster=__ret__.get('hostmaster'),
        id=__ret__.get('id'),
        link=__ret__.get('link'),
        networks=__ret__.get('networks'),
        nx_ttl=__ret__.get('nxTtl'),
        primary=__ret__.get('primary'),
        refresh=__ret__.get('refresh'),
        retry=__ret__.get('retry'),
        secondaries=__ret__.get('secondaries'),
        ttl=__ret__.get('ttl'),
        zone=__ret__.get('zone'))
