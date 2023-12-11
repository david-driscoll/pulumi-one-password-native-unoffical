# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'BranchInformationSectionArgs',
]

@pulumi.input_type
class BranchInformationSectionArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None):
        if address is not None:
            pulumi.set(__self__, "address", address)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)


