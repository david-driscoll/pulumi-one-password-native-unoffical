# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAttachmentResult',
    'AwaitableGetAttachmentResult',
    'get_attachment',
    'get_attachment_output',
]

@pulumi.output_type
class GetAttachmentResult:
    """
    The resolved reference value
    """
    def __init__(__self__, value=None):
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


class AwaitableGetAttachmentResult(GetAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAttachmentResult(
            value=self.value)


def get_attachment(reference: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAttachmentResult:
    """
    Use this data source to access information about an existing resource.

    :param str reference: The 1Password secret reference path to the item.  eg: op://vault/item/[section]/file 
    """
    __args__ = dict()
    __args__['reference'] = reference
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('one-password-native-unoffical:index:GetAttachment', __args__, opts=opts, typ=GetAttachmentResult).value

    return AwaitableGetAttachmentResult(
        value=__ret__.value)


@_utilities.lift_output_func(get_attachment)
def get_attachment_output(reference: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAttachmentResult]:
    """
    Use this data source to access information about an existing resource.

    :param str reference: The 1Password secret reference path to the item.  eg: op://vault/item/[section]/file 
    """
    ...
