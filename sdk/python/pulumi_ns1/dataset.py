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

__all__ = ['DatasetArgs', 'Dataset']

@pulumi.input_type
class DatasetArgs:
    def __init__(__self__, *,
                 datatype: pulumi.Input['DatasetDatatypeArgs'],
                 export_type: pulumi.Input[str],
                 timeframe: pulumi.Input['DatasetTimeframeArgs'],
                 name: Optional[pulumi.Input[str]] = None,
                 recipient_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repeat: Optional[pulumi.Input['DatasetRepeatArgs']] = None):
        """
        The set of arguments for constructing a Dataset resource.
        """
        pulumi.set(__self__, "datatype", datatype)
        pulumi.set(__self__, "export_type", export_type)
        pulumi.set(__self__, "timeframe", timeframe)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if recipient_emails is not None:
            pulumi.set(__self__, "recipient_emails", recipient_emails)
        if repeat is not None:
            pulumi.set(__self__, "repeat", repeat)

    @property
    @pulumi.getter
    def datatype(self) -> pulumi.Input['DatasetDatatypeArgs']:
        return pulumi.get(self, "datatype")

    @datatype.setter
    def datatype(self, value: pulumi.Input['DatasetDatatypeArgs']):
        pulumi.set(self, "datatype", value)

    @property
    @pulumi.getter(name="exportType")
    def export_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "export_type")

    @export_type.setter
    def export_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "export_type", value)

    @property
    @pulumi.getter
    def timeframe(self) -> pulumi.Input['DatasetTimeframeArgs']:
        return pulumi.get(self, "timeframe")

    @timeframe.setter
    def timeframe(self, value: pulumi.Input['DatasetTimeframeArgs']):
        pulumi.set(self, "timeframe", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="recipientEmails")
    def recipient_emails(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "recipient_emails")

    @recipient_emails.setter
    def recipient_emails(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "recipient_emails", value)

    @property
    @pulumi.getter
    def repeat(self) -> Optional[pulumi.Input['DatasetRepeatArgs']]:
        return pulumi.get(self, "repeat")

    @repeat.setter
    def repeat(self, value: Optional[pulumi.Input['DatasetRepeatArgs']]):
        pulumi.set(self, "repeat", value)


@pulumi.input_type
class _DatasetState:
    def __init__(__self__, *,
                 datatype: Optional[pulumi.Input['DatasetDatatypeArgs']] = None,
                 export_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recipient_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repeat: Optional[pulumi.Input['DatasetRepeatArgs']] = None,
                 reports: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetReportArgs']]]] = None,
                 timeframe: Optional[pulumi.Input['DatasetTimeframeArgs']] = None):
        """
        Input properties used for looking up and filtering Dataset resources.
        """
        if datatype is not None:
            pulumi.set(__self__, "datatype", datatype)
        if export_type is not None:
            pulumi.set(__self__, "export_type", export_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if recipient_emails is not None:
            pulumi.set(__self__, "recipient_emails", recipient_emails)
        if repeat is not None:
            pulumi.set(__self__, "repeat", repeat)
        if reports is not None:
            pulumi.set(__self__, "reports", reports)
        if timeframe is not None:
            pulumi.set(__self__, "timeframe", timeframe)

    @property
    @pulumi.getter
    def datatype(self) -> Optional[pulumi.Input['DatasetDatatypeArgs']]:
        return pulumi.get(self, "datatype")

    @datatype.setter
    def datatype(self, value: Optional[pulumi.Input['DatasetDatatypeArgs']]):
        pulumi.set(self, "datatype", value)

    @property
    @pulumi.getter(name="exportType")
    def export_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "export_type")

    @export_type.setter
    def export_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "export_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="recipientEmails")
    def recipient_emails(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "recipient_emails")

    @recipient_emails.setter
    def recipient_emails(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "recipient_emails", value)

    @property
    @pulumi.getter
    def repeat(self) -> Optional[pulumi.Input['DatasetRepeatArgs']]:
        return pulumi.get(self, "repeat")

    @repeat.setter
    def repeat(self, value: Optional[pulumi.Input['DatasetRepeatArgs']]):
        pulumi.set(self, "repeat", value)

    @property
    @pulumi.getter
    def reports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DatasetReportArgs']]]]:
        return pulumi.get(self, "reports")

    @reports.setter
    def reports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DatasetReportArgs']]]]):
        pulumi.set(self, "reports", value)

    @property
    @pulumi.getter
    def timeframe(self) -> Optional[pulumi.Input['DatasetTimeframeArgs']]:
        return pulumi.get(self, "timeframe")

    @timeframe.setter
    def timeframe(self, value: Optional[pulumi.Input['DatasetTimeframeArgs']]):
        pulumi.set(self, "timeframe", value)


class Dataset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datatype: Optional[pulumi.Input[Union['DatasetDatatypeArgs', 'DatasetDatatypeArgsDict']]] = None,
                 export_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recipient_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repeat: Optional[pulumi.Input[Union['DatasetRepeatArgs', 'DatasetRepeatArgsDict']]] = None,
                 timeframe: Optional[pulumi.Input[Union['DatasetTimeframeArgs', 'DatasetTimeframeArgsDict']]] = None,
                 __props__=None):
        """
        Create a Dataset resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatasetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Dataset resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DatasetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datatype: Optional[pulumi.Input[Union['DatasetDatatypeArgs', 'DatasetDatatypeArgsDict']]] = None,
                 export_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recipient_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repeat: Optional[pulumi.Input[Union['DatasetRepeatArgs', 'DatasetRepeatArgsDict']]] = None,
                 timeframe: Optional[pulumi.Input[Union['DatasetTimeframeArgs', 'DatasetTimeframeArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatasetArgs.__new__(DatasetArgs)

            if datatype is None and not opts.urn:
                raise TypeError("Missing required property 'datatype'")
            __props__.__dict__["datatype"] = datatype
            if export_type is None and not opts.urn:
                raise TypeError("Missing required property 'export_type'")
            __props__.__dict__["export_type"] = export_type
            __props__.__dict__["name"] = name
            __props__.__dict__["recipient_emails"] = recipient_emails
            __props__.__dict__["repeat"] = repeat
            if timeframe is None and not opts.urn:
                raise TypeError("Missing required property 'timeframe'")
            __props__.__dict__["timeframe"] = timeframe
            __props__.__dict__["reports"] = None
        super(Dataset, __self__).__init__(
            'ns1:index/dataset:Dataset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            datatype: Optional[pulumi.Input[Union['DatasetDatatypeArgs', 'DatasetDatatypeArgsDict']]] = None,
            export_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            recipient_emails: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            repeat: Optional[pulumi.Input[Union['DatasetRepeatArgs', 'DatasetRepeatArgsDict']]] = None,
            reports: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DatasetReportArgs', 'DatasetReportArgsDict']]]]] = None,
            timeframe: Optional[pulumi.Input[Union['DatasetTimeframeArgs', 'DatasetTimeframeArgsDict']]] = None) -> 'Dataset':
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatasetState.__new__(_DatasetState)

        __props__.__dict__["datatype"] = datatype
        __props__.__dict__["export_type"] = export_type
        __props__.__dict__["name"] = name
        __props__.__dict__["recipient_emails"] = recipient_emails
        __props__.__dict__["repeat"] = repeat
        __props__.__dict__["reports"] = reports
        __props__.__dict__["timeframe"] = timeframe
        return Dataset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def datatype(self) -> pulumi.Output['outputs.DatasetDatatype']:
        return pulumi.get(self, "datatype")

    @property
    @pulumi.getter(name="exportType")
    def export_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "export_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="recipientEmails")
    def recipient_emails(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "recipient_emails")

    @property
    @pulumi.getter
    def repeat(self) -> pulumi.Output[Optional['outputs.DatasetRepeat']]:
        return pulumi.get(self, "repeat")

    @property
    @pulumi.getter
    def reports(self) -> pulumi.Output[Sequence['outputs.DatasetReport']]:
        return pulumi.get(self, "reports")

    @property
    @pulumi.getter
    def timeframe(self) -> pulumi.Output['outputs.DatasetTimeframe']:
        return pulumi.get(self, "timeframe")

