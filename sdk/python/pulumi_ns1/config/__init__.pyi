# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

apikey: Optional[str]
"""
The ns1 API key, this is required
"""

enableDdi: Optional[bool]

endpoint: Optional[str]

ignoreSsl: Optional[bool]

rateLimitParallelism: Optional[int]

