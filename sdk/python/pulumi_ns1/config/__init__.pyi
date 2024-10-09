# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

apikey: Optional[str]
"""
The ns1 API key (required)
"""

endpoint: Optional[str]
"""
URL prefix (including version) for API calls
"""

ignoreSsl: Optional[bool]
"""
Don't validate server SSL/TLS certificate
"""

rateLimitParallelism: Optional[int]
"""
Tune response to rate limits, see docs
"""

retryMax: Optional[int]
"""
Maximum retries for 50x errors (-1 to disable)
"""

userAgent: Optional[str]
"""
User-Agent string to use in NS1 API requests
"""

