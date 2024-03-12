# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'APIKeyDnsRecordsAllowArgs',
    'APIKeyDnsRecordsDenyArgs',
    'ApplicationDefaultConfigArgs',
    'DatasetDatatypeArgs',
    'DatasetRepeatArgs',
    'DatasetReportArgs',
    'DatasetTimeframeArgs',
    'MonitoringJobRuleArgs',
    'NotifyListNotificationArgs',
    'PulsarJobBlendMetricWeightsArgs',
    'PulsarJobConfigArgs',
    'PulsarJobWeightArgs',
    'RecordAnswerArgs',
    'RecordFilterArgs',
    'RecordRegionArgs',
    'TeamDnsRecordsAllowArgs',
    'TeamDnsRecordsDenyArgs',
    'TeamIpWhitelistArgs',
    'UserDnsRecordsAllowArgs',
    'UserDnsRecordsDenyArgs',
    'ZoneSecondaryArgs',
    'GetMonitoringRegionsRegionArgs',
]

@pulumi.input_type
class APIKeyDnsRecordsAllowArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 include_subdomains: pulumi.Input[bool],
                 type: pulumi.Input[str],
                 zone: pulumi.Input[str]):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "include_subdomains", include_subdomains)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="includeSubdomains")
    def include_subdomains(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "include_subdomains")

    @include_subdomains.setter
    def include_subdomains(self, value: pulumi.Input[bool]):
        pulumi.set(self, "include_subdomains", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class APIKeyDnsRecordsDenyArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 include_subdomains: pulumi.Input[bool],
                 type: pulumi.Input[str],
                 zone: pulumi.Input[str]):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "include_subdomains", include_subdomains)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="includeSubdomains")
    def include_subdomains(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "include_subdomains")

    @include_subdomains.setter
    def include_subdomains(self, value: pulumi.Input[bool]):
        pulumi.set(self, "include_subdomains", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class ApplicationDefaultConfigArgs:
    def __init__(__self__, *,
                 http: pulumi.Input[bool],
                 https: Optional[pulumi.Input[bool]] = None,
                 job_timeout_millis: Optional[pulumi.Input[int]] = None,
                 request_timeout_millis: Optional[pulumi.Input[int]] = None,
                 static_values: Optional[pulumi.Input[bool]] = None,
                 use_xhr: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] http: Indicates whether or not to use HTTP in measurements.
        :param pulumi.Input[bool] https: Indicates whether or not to use HTTPS in measurements.
        :param pulumi.Input[int] job_timeout_millis: Maximum timeout per job
               0, the primary NSONE Global Network. Normally, you should not have to worry about this.
        :param pulumi.Input[int] request_timeout_millis: Maximum timeout per request.
        :param pulumi.Input[bool] static_values: Indicates whether or not to skip aggregation for this job's measurements
        :param pulumi.Input[bool] use_xhr: Whether to use XMLHttpRequest (XHR) when taking measurements.
        """
        pulumi.set(__self__, "http", http)
        if https is not None:
            pulumi.set(__self__, "https", https)
        if job_timeout_millis is not None:
            pulumi.set(__self__, "job_timeout_millis", job_timeout_millis)
        if request_timeout_millis is not None:
            pulumi.set(__self__, "request_timeout_millis", request_timeout_millis)
        if static_values is not None:
            pulumi.set(__self__, "static_values", static_values)
        if use_xhr is not None:
            pulumi.set(__self__, "use_xhr", use_xhr)

    @property
    @pulumi.getter
    def http(self) -> pulumi.Input[bool]:
        """
        Indicates whether or not to use HTTP in measurements.
        """
        return pulumi.get(self, "http")

    @http.setter
    def http(self, value: pulumi.Input[bool]):
        pulumi.set(self, "http", value)

    @property
    @pulumi.getter
    def https(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether or not to use HTTPS in measurements.
        """
        return pulumi.get(self, "https")

    @https.setter
    def https(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https", value)

    @property
    @pulumi.getter(name="jobTimeoutMillis")
    def job_timeout_millis(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum timeout per job
        0, the primary NSONE Global Network. Normally, you should not have to worry about this.
        """
        return pulumi.get(self, "job_timeout_millis")

    @job_timeout_millis.setter
    def job_timeout_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "job_timeout_millis", value)

    @property
    @pulumi.getter(name="requestTimeoutMillis")
    def request_timeout_millis(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum timeout per request.
        """
        return pulumi.get(self, "request_timeout_millis")

    @request_timeout_millis.setter
    def request_timeout_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "request_timeout_millis", value)

    @property
    @pulumi.getter(name="staticValues")
    def static_values(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether or not to skip aggregation for this job's measurements
        """
        return pulumi.get(self, "static_values")

    @static_values.setter
    def static_values(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "static_values", value)

    @property
    @pulumi.getter(name="useXhr")
    def use_xhr(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use XMLHttpRequest (XHR) when taking measurements.
        """
        return pulumi.get(self, "use_xhr")

    @use_xhr.setter
    def use_xhr(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_xhr", value)


@pulumi.input_type
class DatasetDatatypeArgs:
    def __init__(__self__, *,
                 data: pulumi.Input[Mapping[str, Any]],
                 scope: pulumi.Input[str],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "data", data)
        pulumi.set(__self__, "scope", scope)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Input[Mapping[str, Any]]:
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input[str]:
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class DatasetRepeatArgs:
    def __init__(__self__, *,
                 end_after_n: pulumi.Input[int],
                 repeats_every: pulumi.Input[str],
                 start: pulumi.Input[int]):
        pulumi.set(__self__, "end_after_n", end_after_n)
        pulumi.set(__self__, "repeats_every", repeats_every)
        pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter(name="endAfterN")
    def end_after_n(self) -> pulumi.Input[int]:
        return pulumi.get(self, "end_after_n")

    @end_after_n.setter
    def end_after_n(self, value: pulumi.Input[int]):
        pulumi.set(self, "end_after_n", value)

    @property
    @pulumi.getter(name="repeatsEvery")
    def repeats_every(self) -> pulumi.Input[str]:
        return pulumi.get(self, "repeats_every")

    @repeats_every.setter
    def repeats_every(self, value: pulumi.Input[str]):
        pulumi.set(self, "repeats_every", value)

    @property
    @pulumi.getter
    def start(self) -> pulumi.Input[int]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: pulumi.Input[int]):
        pulumi.set(self, "start", value)


@pulumi.input_type
class DatasetReportArgs:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[int]] = None,
                 end: Optional[pulumi.Input[int]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if end is not None:
            pulumi.set(__self__, "end", end)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if start is not None:
            pulumi.set(__self__, "start", start)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class DatasetTimeframeArgs:
    def __init__(__self__, *,
                 aggregation: pulumi.Input[str],
                 cycles: Optional[pulumi.Input[int]] = None,
                 from_: Optional[pulumi.Input[int]] = None,
                 to: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "aggregation", aggregation)
        if cycles is not None:
            pulumi.set(__self__, "cycles", cycles)
        if from_ is not None:
            pulumi.set(__self__, "from_", from_)
        if to is not None:
            pulumi.set(__self__, "to", to)

    @property
    @pulumi.getter
    def aggregation(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aggregation")

    @aggregation.setter
    def aggregation(self, value: pulumi.Input[str]):
        pulumi.set(self, "aggregation", value)

    @property
    @pulumi.getter
    def cycles(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "cycles")

    @cycles.setter
    def cycles(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cycles", value)

    @property
    @pulumi.getter(name="from")
    def from_(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "from_")

    @from_.setter
    def from_(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "from_", value)

    @property
    @pulumi.getter
    def to(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "to")

    @to.setter
    def to(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "to", value)


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
class PulsarJobBlendMetricWeightsArgs:
    def __init__(__self__, *,
                 timestamp: pulumi.Input[int]):
        pulumi.set(__self__, "timestamp", timestamp)

    @property
    @pulumi.getter
    def timestamp(self) -> pulumi.Input[int]:
        return pulumi.get(self, "timestamp")

    @timestamp.setter
    def timestamp(self, value: pulumi.Input[int]):
        pulumi.set(self, "timestamp", value)


@pulumi.input_type
class PulsarJobConfigArgs:
    def __init__(__self__, *,
                 host: Optional[pulumi.Input[str]] = None,
                 http: Optional[pulumi.Input[bool]] = None,
                 https: Optional[pulumi.Input[bool]] = None,
                 job_timeout_millis: Optional[pulumi.Input[int]] = None,
                 request_timeout_millis: Optional[pulumi.Input[int]] = None,
                 static_values: Optional[pulumi.Input[bool]] = None,
                 url_path: Optional[pulumi.Input[str]] = None,
                 use_xhr: Optional[pulumi.Input[bool]] = None):
        if host is not None:
            pulumi.set(__self__, "host", host)
        if http is not None:
            pulumi.set(__self__, "http", http)
        if https is not None:
            pulumi.set(__self__, "https", https)
        if job_timeout_millis is not None:
            pulumi.set(__self__, "job_timeout_millis", job_timeout_millis)
        if request_timeout_millis is not None:
            pulumi.set(__self__, "request_timeout_millis", request_timeout_millis)
        if static_values is not None:
            pulumi.set(__self__, "static_values", static_values)
        if url_path is not None:
            pulumi.set(__self__, "url_path", url_path)
        if use_xhr is not None:
            pulumi.set(__self__, "use_xhr", use_xhr)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def http(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "http")

    @http.setter
    def http(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "http", value)

    @property
    @pulumi.getter
    def https(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "https")

    @https.setter
    def https(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https", value)

    @property
    @pulumi.getter(name="jobTimeoutMillis")
    def job_timeout_millis(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "job_timeout_millis")

    @job_timeout_millis.setter
    def job_timeout_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "job_timeout_millis", value)

    @property
    @pulumi.getter(name="requestTimeoutMillis")
    def request_timeout_millis(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "request_timeout_millis")

    @request_timeout_millis.setter
    def request_timeout_millis(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "request_timeout_millis", value)

    @property
    @pulumi.getter(name="staticValues")
    def static_values(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "static_values")

    @static_values.setter
    def static_values(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "static_values", value)

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_path", value)

    @property
    @pulumi.getter(name="useXhr")
    def use_xhr(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "use_xhr")

    @use_xhr.setter
    def use_xhr(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_xhr", value)


@pulumi.input_type
class PulsarJobWeightArgs:
    def __init__(__self__, *,
                 default_value: pulumi.Input[float],
                 name: pulumi.Input[str],
                 weight: pulumi.Input[int],
                 maximize: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "default_value", default_value)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "weight", weight)
        if maximize is not None:
            pulumi.set(__self__, "maximize", maximize)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> pulumi.Input[float]:
        return pulumi.get(self, "default_value")

    @default_value.setter
    def default_value(self, value: pulumi.Input[float]):
        pulumi.set(self, "default_value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Input[int]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: pulumi.Input[int]):
        pulumi.set(self, "weight", value)

    @property
    @pulumi.getter
    def maximize(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "maximize")

    @maximize.setter
    def maximize(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "maximize", value)


@pulumi.input_type
class RecordAnswerArgs:
    def __init__(__self__, *,
                 answer: Optional[pulumi.Input[str]] = None,
                 meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] answer: Space delimited string of RDATA fields dependent on the record type.
               
               A:
               
               answer = "1.2.3.4"
               
               CNAME:
               
               answer = "www.example.com"
               
               MX:
               
               answer = "5 mail.example.com"
               
               SRV:
               
               answer = "10 0 2380 node-1.example.com"
               
               SPF:
               
               answer = "v=DKIM1; k=rsa; p=XXXXXXXX"
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

        A:

        answer = "1.2.3.4"

        CNAME:

        answer = "www.example.com"

        MX:

        answer = "5 mail.example.com"

        SRV:

        answer = "10 0 2380 node-1.example.com"

        SPF:

        answer = "v=DKIM1; k=rsa; p=XXXXXXXX"
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
class TeamDnsRecordsAllowArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 include_subdomains: pulumi.Input[bool],
                 type: pulumi.Input[str],
                 zone: pulumi.Input[str]):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "include_subdomains", include_subdomains)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="includeSubdomains")
    def include_subdomains(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "include_subdomains")

    @include_subdomains.setter
    def include_subdomains(self, value: pulumi.Input[bool]):
        pulumi.set(self, "include_subdomains", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class TeamDnsRecordsDenyArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 include_subdomains: pulumi.Input[bool],
                 type: pulumi.Input[str],
                 zone: pulumi.Input[str]):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "include_subdomains", include_subdomains)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="includeSubdomains")
    def include_subdomains(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "include_subdomains")

    @include_subdomains.setter
    def include_subdomains(self, value: pulumi.Input[bool]):
        pulumi.set(self, "include_subdomains", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)


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
class UserDnsRecordsAllowArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 include_subdomains: pulumi.Input[bool],
                 type: pulumi.Input[str],
                 zone: pulumi.Input[str]):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "include_subdomains", include_subdomains)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="includeSubdomains")
    def include_subdomains(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "include_subdomains")

    @include_subdomains.setter
    def include_subdomains(self, value: pulumi.Input[bool]):
        pulumi.set(self, "include_subdomains", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class UserDnsRecordsDenyArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 include_subdomains: pulumi.Input[bool],
                 type: pulumi.Input[str],
                 zone: pulumi.Input[str]):
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "include_subdomains", include_subdomains)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="includeSubdomains")
    def include_subdomains(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "include_subdomains")

    @include_subdomains.setter
    def include_subdomains(self, value: pulumi.Input[bool]):
        pulumi.set(self, "include_subdomains", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class ZoneSecondaryArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 notify: Optional[pulumi.Input[bool]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] ip: IPv4 address of the secondary server.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] networks: List of network IDs (`int`) for which the zone
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
        List of network IDs (`int`) for which the zone
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


@pulumi.input_type
class GetMonitoringRegionsRegionArgs:
    def __init__(__self__, *,
                 code: Optional[str] = None,
                 name: Optional[str] = None,
                 subnets: Optional[Sequence[str]] = None):
        """
        :param str code: 3-letter city code identifying the location of the monitor.
        :param str name: City name identifying the location of the monitor.
        :param Sequence[str] subnets: A list of IPv4 and IPv6 subnets the monitor sources requests from.
        """
        if code is not None:
            pulumi.set(__self__, "code", code)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if subnets is not None:
            pulumi.set(__self__, "subnets", subnets)

    @property
    @pulumi.getter
    def code(self) -> Optional[str]:
        """
        3-letter city code identifying the location of the monitor.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[str]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        City name identifying the location of the monitor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def subnets(self) -> Optional[Sequence[str]]:
        """
        A list of IPv4 and IPv6 subnets the monitor sources requests from.
        """
        return pulumi.get(self, "subnets")

    @subnets.setter
    def subnets(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "subnets", value)


