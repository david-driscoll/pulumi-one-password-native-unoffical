# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import creditcard as _creditcard
from . import outputs
from ._enums import *

__all__ = [
    'GetCreditCardResult',
    'AwaitableGetCreditCardResult',
    'get_credit_card',
    'get_credit_card_output',
]

@pulumi.output_type
class GetCreditCardResult:
    def __init__(__self__, additional_details=None, attachments=None, cardholder_name=None, category=None, contact_information=None, expiry_date=None, fields=None, notes=None, number=None, references=None, sections=None, tags=None, title=None, type=None, uuid=None, valid_from=None, vault=None, verification_number=None):
        if additional_details and not isinstance(additional_details, dict):
            raise TypeError("Expected argument 'additional_details' to be a dict")
        pulumi.set(__self__, "additional_details", additional_details)
        if attachments and not isinstance(attachments, dict):
            raise TypeError("Expected argument 'attachments' to be a dict")
        pulumi.set(__self__, "attachments", attachments)
        if cardholder_name and not isinstance(cardholder_name, str):
            raise TypeError("Expected argument 'cardholder_name' to be a str")
        pulumi.set(__self__, "cardholder_name", cardholder_name)
        if category and not isinstance(category, dict):
            raise TypeError("Expected argument 'category' to be a dict")
        pulumi.set(__self__, "category", category)
        if contact_information and not isinstance(contact_information, dict):
            raise TypeError("Expected argument 'contact_information' to be a dict")
        pulumi.set(__self__, "contact_information", contact_information)
        if expiry_date and not isinstance(expiry_date, str):
            raise TypeError("Expected argument 'expiry_date' to be a str")
        pulumi.set(__self__, "expiry_date", expiry_date)
        if fields and not isinstance(fields, dict):
            raise TypeError("Expected argument 'fields' to be a dict")
        pulumi.set(__self__, "fields", fields)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if number and not isinstance(number, str):
            raise TypeError("Expected argument 'number' to be a str")
        pulumi.set(__self__, "number", number)
        if references and not isinstance(references, dict):
            raise TypeError("Expected argument 'references' to be a dict")
        pulumi.set(__self__, "references", references)
        if sections and not isinstance(sections, dict):
            raise TypeError("Expected argument 'sections' to be a dict")
        pulumi.set(__self__, "sections", sections)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if valid_from and not isinstance(valid_from, str):
            raise TypeError("Expected argument 'valid_from' to be a str")
        pulumi.set(__self__, "valid_from", valid_from)
        if vault and not isinstance(vault, str):
            raise TypeError("Expected argument 'vault' to be a str")
        pulumi.set(__self__, "vault", vault)
        if verification_number and not isinstance(verification_number, str):
            raise TypeError("Expected argument 'verification_number' to be a str")
        pulumi.set(__self__, "verification_number", verification_number)

    @property
    @pulumi.getter(name="additionalDetails")
    def additional_details(self) -> Optional['_creditcard.outputs.AdditionalDetailsSection']:
        return pulumi.get(self, "additional_details")

    @property
    @pulumi.getter
    def attachments(self) -> Mapping[str, 'outputs.OutField']:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter(name="cardholderName")
    def cardholder_name(self) -> Optional[str]:
        return pulumi.get(self, "cardholder_name")

    @property
    @pulumi.getter
    def category(self) -> str:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="contactInformation")
    def contact_information(self) -> Optional['_creditcard.outputs.ContactInformationSection']:
        return pulumi.get(self, "contact_information")

    @property
    @pulumi.getter(name="expiryDate")
    def expiry_date(self) -> Optional[str]:
        return pulumi.get(self, "expiry_date")

    @property
    @pulumi.getter
    def fields(self) -> Mapping[str, 'outputs.OutField']:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def number(self) -> Optional[str]:
        return pulumi.get(self, "number")

    @property
    @pulumi.getter
    def references(self) -> Mapping[str, 'outputs.OutField']:
        return pulumi.get(self, "references")

    @property
    @pulumi.getter
    def sections(self) -> Mapping[str, 'outputs.OutSection']:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        An array of strings of the tags assigned to the item.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        The title of the item.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> Optional[str]:
        return pulumi.get(self, "valid_from")

    @property
    @pulumi.getter
    def vault(self) -> str:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")

    @property
    @pulumi.getter(name="verificationNumber")
    def verification_number(self) -> Optional[str]:
        return pulumi.get(self, "verification_number")


class AwaitableGetCreditCardResult(GetCreditCardResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCreditCardResult(
            additional_details=self.additional_details,
            attachments=self.attachments,
            cardholder_name=self.cardholder_name,
            category=self.category,
            contact_information=self.contact_information,
            expiry_date=self.expiry_date,
            fields=self.fields,
            notes=self.notes,
            number=self.number,
            references=self.references,
            sections=self.sections,
            tags=self.tags,
            title=self.title,
            type=self.type,
            uuid=self.uuid,
            valid_from=self.valid_from,
            vault=self.vault,
            verification_number=self.verification_number)


def get_credit_card(title: Optional[str] = None,
                    uuid: Optional[str] = None,
                    vault: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCreditCardResult:
    """
    Use this data source to access information about an existing resource.

    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str uuid: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str vault: The UUID of the vault the item is in.
    """
    __args__ = dict()
    __args__['title'] = title
    __args__['uuid'] = uuid
    __args__['vault'] = vault
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('one-password-native-unoffical:index:GetCreditCard', __args__, opts=opts, typ=GetCreditCardResult).value

    return AwaitableGetCreditCardResult(
        additional_details=__ret__.additional_details,
        attachments=__ret__.attachments,
        cardholder_name=__ret__.cardholder_name,
        category=__ret__.category,
        contact_information=__ret__.contact_information,
        expiry_date=__ret__.expiry_date,
        fields=__ret__.fields,
        notes=__ret__.notes,
        number=__ret__.number,
        references=__ret__.references,
        sections=__ret__.sections,
        tags=__ret__.tags,
        title=__ret__.title,
        type=__ret__.type,
        uuid=__ret__.uuid,
        valid_from=__ret__.valid_from,
        vault=__ret__.vault,
        verification_number=__ret__.verification_number)


@_utilities.lift_output_func(get_credit_card)
def get_credit_card_output(title: Optional[pulumi.Input[Optional[str]]] = None,
                           uuid: Optional[pulumi.Input[Optional[str]]] = None,
                           vault: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCreditCardResult]:
    """
    Use this data source to access information about an existing resource.

    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str uuid: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str vault: The UUID of the vault the item is in.
    """
    ...
