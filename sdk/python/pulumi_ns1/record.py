# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Record(pulumi.CustomResource):
    answers: pulumi.Output[list]
    """
    One or more NS1 answers for the records' specified type.
    Answers are documented below.

      * `answer` (`str`) - Space delimited string of RDATA fields dependent on the record type.
      * `meta` (`dict`)
      * `region` (`str`) - The region (Answer Group really) that this answer
        belongs to. This should be one of the names specified in `regions`. Only a
        single `region` per answer is currently supported. If you want an answer in
        multiple regions, duplicating the answer (including metadata) is the correct
        approach.
        * ` meta` - (Optional) meta is supported at the `answer` level. Meta
        is documented below.
    """
    domain: pulumi.Output[str]
    """
    The records' domain. Cannot have leading or trailing
    dots - see the example above and `FQDN formatting` below.
    """
    filters: pulumi.Output[list]
    """
    One or more NS1 filters for the record(order matters).
    Filters are documented below.

      * `config` (`dict`) - The filters' configuration. Simple key/value pairs
        determined by the filter type.
      * `disabled` (`bool`) - Determines whether the filter is applied in the
        filter chain.
      * `filter` (`str`) - The type of filter.
    """
    link: pulumi.Output[str]
    """
    The target record to link to. This means this record is a
    'linked' record, and it inherits all properties from its target.
    """
    meta: pulumi.Output[dict]
    regions: pulumi.Output[list]
    """
    One or more "regions" for the record. These are really
    just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
    but remain `regions` here for legacy reasons. Regions are
    documented below. Please note the ordering requirement!

      * `meta` (`dict`)
      * `name` (`str`) - Name of the region (or Answer Group).
    """
    short_answers: pulumi.Output[list]
    ttl: pulumi.Output[float]
    """
    The records' time to live.
    """
    type: pulumi.Output[str]
    """
    The records' RR type.
    """
    use_client_subnet: pulumi.Output[bool]
    """
    Whether to use EDNS client subnet data when
    available(in filter chain).
    * ` meta` - (Optional) meta is supported at the `record` level. Meta
    is documented below.
    """
    zone: pulumi.Output[str]
    """
    The zone the record belongs to. Cannot have leading or
    trailing dots (".") - see the example above and `FQDN formatting` below.
    """
    def __init__(__self__, resource_name, opts=None, answers=None, domain=None, filters=None, link=None, meta=None, regions=None, short_answers=None, ttl=None, type=None, use_client_subnet=None, zone=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a NS1 Record resource. This can be used to create, modify, and delete records.

        ## NS1 Documentation

        [Record Api Doc](https://ns1.com/api#records)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] answers: One or more NS1 answers for the records' specified type.
               Answers are documented below.
        :param pulumi.Input[str] domain: The records' domain. Cannot have leading or trailing
               dots - see the example above and `FQDN formatting` below.
        :param pulumi.Input[list] filters: One or more NS1 filters for the record(order matters).
               Filters are documented below.
        :param pulumi.Input[str] link: The target record to link to. This means this record is a
               'linked' record, and it inherits all properties from its target.
        :param pulumi.Input[list] regions: One or more "regions" for the record. These are really
               just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
               but remain `regions` here for legacy reasons. Regions are
               documented below. Please note the ordering requirement!
        :param pulumi.Input[float] ttl: The records' time to live.
        :param pulumi.Input[str] type: The records' RR type.
        :param pulumi.Input[bool] use_client_subnet: Whether to use EDNS client subnet data when
               available(in filter chain).
               * ` meta` - (Optional) meta is supported at the `record` level. Meta
               is documented below.
        :param pulumi.Input[str] zone: The zone the record belongs to. Cannot have leading or
               trailing dots (".") - see the example above and `FQDN formatting` below.

        The **answers** object supports the following:

          * `answer` (`pulumi.Input[str]`) - Space delimited string of RDATA fields dependent on the record type.
          * `meta` (`pulumi.Input[dict]`)
          * `region` (`pulumi.Input[str]`) - The region (Answer Group really) that this answer
            belongs to. This should be one of the names specified in `regions`. Only a
            single `region` per answer is currently supported. If you want an answer in
            multiple regions, duplicating the answer (including metadata) is the correct
            approach.
            * ` meta` - (Optional) meta is supported at the `answer` level. Meta
            is documented below.

        The **filters** object supports the following:

          * `config` (`pulumi.Input[dict]`) - The filters' configuration. Simple key/value pairs
            determined by the filter type.
          * `disabled` (`pulumi.Input[bool]`) - Determines whether the filter is applied in the
            filter chain.
          * `filter` (`pulumi.Input[str]`) - The type of filter.

        The **regions** object supports the following:

          * `meta` (`pulumi.Input[dict]`)
          * `name` (`pulumi.Input[str]`) - Name of the region (or Answer Group).
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['answers'] = answers
            if domain is None:
                raise TypeError("Missing required property 'domain'")
            __props__['domain'] = domain
            __props__['filters'] = filters
            __props__['link'] = link
            __props__['meta'] = meta
            __props__['regions'] = regions
            if short_answers is not None:
                warnings.warn("short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.", DeprecationWarning)
                pulumi.log.warn("short_answers is deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.")
            __props__['short_answers'] = short_answers
            __props__['ttl'] = ttl
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['use_client_subnet'] = use_client_subnet
            if zone is None:
                raise TypeError("Missing required property 'zone'")
            __props__['zone'] = zone
        super(Record, __self__).__init__(
            'ns1:index/record:Record',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, answers=None, domain=None, filters=None, link=None, meta=None, regions=None, short_answers=None, ttl=None, type=None, use_client_subnet=None, zone=None):
        """
        Get an existing Record resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] answers: One or more NS1 answers for the records' specified type.
               Answers are documented below.
        :param pulumi.Input[str] domain: The records' domain. Cannot have leading or trailing
               dots - see the example above and `FQDN formatting` below.
        :param pulumi.Input[list] filters: One or more NS1 filters for the record(order matters).
               Filters are documented below.
        :param pulumi.Input[str] link: The target record to link to. This means this record is a
               'linked' record, and it inherits all properties from its target.
        :param pulumi.Input[list] regions: One or more "regions" for the record. These are really
               just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
               but remain `regions` here for legacy reasons. Regions are
               documented below. Please note the ordering requirement!
        :param pulumi.Input[float] ttl: The records' time to live.
        :param pulumi.Input[str] type: The records' RR type.
        :param pulumi.Input[bool] use_client_subnet: Whether to use EDNS client subnet data when
               available(in filter chain).
               * ` meta` - (Optional) meta is supported at the `record` level. Meta
               is documented below.
        :param pulumi.Input[str] zone: The zone the record belongs to. Cannot have leading or
               trailing dots (".") - see the example above and `FQDN formatting` below.

        The **answers** object supports the following:

          * `answer` (`pulumi.Input[str]`) - Space delimited string of RDATA fields dependent on the record type.
          * `meta` (`pulumi.Input[dict]`)
          * `region` (`pulumi.Input[str]`) - The region (Answer Group really) that this answer
            belongs to. This should be one of the names specified in `regions`. Only a
            single `region` per answer is currently supported. If you want an answer in
            multiple regions, duplicating the answer (including metadata) is the correct
            approach.
            * ` meta` - (Optional) meta is supported at the `answer` level. Meta
            is documented below.

        The **filters** object supports the following:

          * `config` (`pulumi.Input[dict]`) - The filters' configuration. Simple key/value pairs
            determined by the filter type.
          * `disabled` (`pulumi.Input[bool]`) - Determines whether the filter is applied in the
            filter chain.
          * `filter` (`pulumi.Input[str]`) - The type of filter.

        The **regions** object supports the following:

          * `meta` (`pulumi.Input[dict]`)
          * `name` (`pulumi.Input[str]`) - Name of the region (or Answer Group).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["answers"] = answers
        __props__["domain"] = domain
        __props__["filters"] = filters
        __props__["link"] = link
        __props__["meta"] = meta
        __props__["regions"] = regions
        __props__["short_answers"] = short_answers
        __props__["ttl"] = ttl
        __props__["type"] = type
        __props__["use_client_subnet"] = use_client_subnet
        __props__["zone"] = zone
        return Record(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
