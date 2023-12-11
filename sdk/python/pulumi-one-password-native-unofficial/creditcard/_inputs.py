# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AdditionalDetailsSectionArgs',
    'ContactInformationSectionArgs',
]

@pulumi.input_type
class AdditionalDetailsSectionArgs:
    def __init__(__self__, *,
                 cash_withdrawal_limit: Optional[pulumi.Input[str]] = None,
                 credit_limit: Optional[pulumi.Input[str]] = None,
                 interest_rate: Optional[pulumi.Input[str]] = None,
                 issue_number: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None):
        if cash_withdrawal_limit is not None:
            pulumi.set(__self__, "cash_withdrawal_limit", cash_withdrawal_limit)
        if credit_limit is not None:
            pulumi.set(__self__, "credit_limit", credit_limit)
        if interest_rate is not None:
            pulumi.set(__self__, "interest_rate", interest_rate)
        if issue_number is not None:
            pulumi.set(__self__, "issue_number", issue_number)
        if pin is not None:
            pulumi.set(__self__, "pin", pin)

    @property
    @pulumi.getter(name="cashWithdrawalLimit")
    def cash_withdrawal_limit(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cash_withdrawal_limit")

    @cash_withdrawal_limit.setter
    def cash_withdrawal_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cash_withdrawal_limit", value)

    @property
    @pulumi.getter(name="creditLimit")
    def credit_limit(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "credit_limit")

    @credit_limit.setter
    def credit_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credit_limit", value)

    @property
    @pulumi.getter(name="interestRate")
    def interest_rate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interest_rate")

    @interest_rate.setter
    def interest_rate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interest_rate", value)

    @property
    @pulumi.getter(name="issueNumber")
    def issue_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issue_number")

    @issue_number.setter
    def issue_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issue_number", value)

    @property
    @pulumi.getter
    def pin(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pin")

    @pin.setter
    def pin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pin", value)


@pulumi.input_type
class ContactInformationSectionArgs:
    def __init__(__self__, *,
                 issuing_bank: Optional[pulumi.Input[str]] = None,
                 phone_intl: Optional[pulumi.Input[str]] = None,
                 phone_local: Optional[pulumi.Input[str]] = None,
                 phone_toll_free: Optional[pulumi.Input[str]] = None,
                 website: Optional[pulumi.Input[str]] = None):
        if issuing_bank is not None:
            pulumi.set(__self__, "issuing_bank", issuing_bank)
        if phone_intl is not None:
            pulumi.set(__self__, "phone_intl", phone_intl)
        if phone_local is not None:
            pulumi.set(__self__, "phone_local", phone_local)
        if phone_toll_free is not None:
            pulumi.set(__self__, "phone_toll_free", phone_toll_free)
        if website is not None:
            pulumi.set(__self__, "website", website)

    @property
    @pulumi.getter(name="issuingBank")
    def issuing_bank(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issuing_bank")

    @issuing_bank.setter
    def issuing_bank(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuing_bank", value)

    @property
    @pulumi.getter(name="phoneIntl")
    def phone_intl(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "phone_intl")

    @phone_intl.setter
    def phone_intl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_intl", value)

    @property
    @pulumi.getter(name="phoneLocal")
    def phone_local(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "phone_local")

    @phone_local.setter
    def phone_local(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_local", value)

    @property
    @pulumi.getter(name="phoneTollFree")
    def phone_toll_free(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "phone_toll_free")

    @phone_toll_free.setter
    def phone_toll_free(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone_toll_free", value)

    @property
    @pulumi.getter
    def website(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "website")

    @website.setter
    def website(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "website", value)


