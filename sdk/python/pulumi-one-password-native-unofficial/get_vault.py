# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetVaultResult',
    'AwaitableGetVaultResult',
    'get_vault',
    'get_vault_output',
]

@pulumi.output_type
class GetVaultResult:
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The UUID of the vault to retrieve. This field will be populated with the UUID of the vault if the vault it looked up by its name.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the vault to retrieve. This field will be populated with the name of the vault if the vault it looked up by its UUID.
        """
        return pulumi.get(self, "name")


class AwaitableGetVaultResult(GetVaultResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVaultResult(
            id=self.id,
            name=self.name)


def get_vault(vault: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVaultResult:
    """
    Use this data source to get details of a vault by either its name or uuid.


    :param str vault: The vault to get information of.  Can be either the name or the UUID.
    """
    __args__ = dict()
    __args__['vault'] = vault
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('one-password-native-unofficial:index:GetVault', __args__, opts=opts, typ=GetVaultResult).value

    return AwaitableGetVaultResult(
        id=__ret__.id,
        name=__ret__.name)


@_utilities.lift_output_func(get_vault)
def get_vault_output(vault: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVaultResult]:
    """
    Use this data source to get details of a vault by either its name or uuid.


    :param str vault: The vault to get information of.  Can be either the name or the UUID.
    """
    ...
