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
from ._inputs import *

__all__ = ['CreditCardItemArgs', 'CreditCardItem']

@pulumi.input_type
class CreditCardItemArgs:
    def __init__(__self__, *,
                 vault: pulumi.Input[str],
                 additional_details: Optional[pulumi.Input['_creditcard.AdditionalDetailsSectionArgs']] = None,
                 cardholder_name: Optional[pulumi.Input[str]] = None,
                 contact_information: Optional[pulumi.Input['_creditcard.ContactInformationSectionArgs']] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input['FieldArgs']]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Sequence[pulumi.Input['SectionArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 verification_number: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CreditCardItem resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        """
        pulumi.set(__self__, "vault", vault)
        if additional_details is not None:
            pulumi.set(__self__, "additional_details", additional_details)
        if cardholder_name is not None:
            pulumi.set(__self__, "cardholder_name", cardholder_name)
        if contact_information is not None:
            pulumi.set(__self__, "contact_information", contact_information)
        if expiry_date is not None:
            pulumi.set(__self__, "expiry_date", expiry_date)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if number is not None:
            pulumi.set(__self__, "number", number)
        if sections is not None:
            pulumi.set(__self__, "sections", sections)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if valid_from is not None:
            pulumi.set(__self__, "valid_from", valid_from)
        if verification_number is not None:
            pulumi.set(__self__, "verification_number", verification_number)

    @property
    @pulumi.getter
    def vault(self) -> pulumi.Input[str]:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")

    @vault.setter
    def vault(self, value: pulumi.Input[str]):
        pulumi.set(self, "vault", value)

    @property
    @pulumi.getter(name="additionalDetails")
    def additional_details(self) -> Optional[pulumi.Input['_creditcard.AdditionalDetailsSectionArgs']]:
        return pulumi.get(self, "additional_details")

    @additional_details.setter
    def additional_details(self, value: Optional[pulumi.Input['_creditcard.AdditionalDetailsSectionArgs']]):
        pulumi.set(self, "additional_details", value)

    @property
    @pulumi.getter(name="cardholderName")
    def cardholder_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cardholder_name")

    @cardholder_name.setter
    def cardholder_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cardholder_name", value)

    @property
    @pulumi.getter(name="contactInformation")
    def contact_information(self) -> Optional[pulumi.Input['_creditcard.ContactInformationSectionArgs']]:
        return pulumi.get(self, "contact_information")

    @contact_information.setter
    def contact_information(self, value: Optional[pulumi.Input['_creditcard.ContactInformationSectionArgs']]):
        pulumi.set(self, "contact_information", value)

    @property
    @pulumi.getter(name="expiryDate")
    def expiry_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiry_date")

    @expiry_date.setter
    def expiry_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry_date", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FieldArgs']]]]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "number", value)

    @property
    @pulumi.getter
    def sections(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SectionArgs']]]]:
        return pulumi.get(self, "sections")

    @sections.setter
    def sections(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SectionArgs']]]]):
        pulumi.set(self, "sections", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of strings of the tags assigned to the item.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "valid_from")

    @valid_from.setter
    def valid_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_from", value)

    @property
    @pulumi.getter(name="verificationNumber")
    def verification_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "verification_number")

    @verification_number.setter
    def verification_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verification_number", value)


class CreditCardItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_details: Optional[pulumi.Input[pulumi.InputType['_creditcard.AdditionalDetailsSectionArgs']]] = None,
                 cardholder_name: Optional[pulumi.Input[str]] = None,
                 contact_information: Optional[pulumi.Input[pulumi.InputType['_creditcard.ContactInformationSectionArgs']]] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 verification_number: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a CreditCardItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CreditCardItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CreditCardItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CreditCardItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CreditCardItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_details: Optional[pulumi.Input[pulumi.InputType['_creditcard.AdditionalDetailsSectionArgs']]] = None,
                 cardholder_name: Optional[pulumi.Input[str]] = None,
                 contact_information: Optional[pulumi.Input[pulumi.InputType['_creditcard.ContactInformationSectionArgs']]] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 verification_number: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CreditCardItemArgs.__new__(CreditCardItemArgs)

            __props__.__dict__["additional_details"] = additional_details
            __props__.__dict__["cardholder_name"] = cardholder_name
            __props__.__dict__["contact_information"] = contact_information
            __props__.__dict__["expiry_date"] = expiry_date
            __props__.__dict__["fields"] = fields
            __props__.__dict__["notes"] = notes
            __props__.__dict__["number"] = number
            __props__.__dict__["sections"] = sections
            __props__.__dict__["tags"] = tags
            __props__.__dict__["title"] = title
            __props__.__dict__["type"] = type
            __props__.__dict__["valid_from"] = valid_from
            if vault is None and not opts.urn:
                raise TypeError("Missing required property 'vault'")
            __props__.__dict__["vault"] = vault
            __props__.__dict__["verification_number"] = verification_number
            __props__.__dict__["category"] = None
            __props__.__dict__["id"] = None
            __props__.__dict__["uuid"] = None
        super(CreditCardItem, __self__).__init__(
            'onepassword:index:CreditCardItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CreditCardItem':
        """
        Get an existing CreditCardItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CreditCardItemArgs.__new__(CreditCardItemArgs)

        __props__.__dict__["additional_details"] = None
        __props__.__dict__["cardholder_name"] = None
        __props__.__dict__["category"] = None
        __props__.__dict__["contact_information"] = None
        __props__.__dict__["expiry_date"] = None
        __props__.__dict__["fields"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["notes"] = None
        __props__.__dict__["number"] = None
        __props__.__dict__["sections"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["title"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["uuid"] = None
        __props__.__dict__["valid_from"] = None
        __props__.__dict__["vault"] = None
        __props__.__dict__["verification_number"] = None
        return CreditCardItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalDetails")
    def additional_details(self) -> pulumi.Output[Optional['_creditcard.outputs.AdditionalDetailsSection']]:
        return pulumi.get(self, "additional_details")

    @property
    @pulumi.getter(name="cardholderName")
    def cardholder_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "cardholder_name")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="contactInformation")
    def contact_information(self) -> pulumi.Output[Optional['_creditcard.outputs.ContactInformationSection']]:
        return pulumi.get(self, "contact_information")

    @property
    @pulumi.getter(name="expiryDate")
    def expiry_date(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expiry_date")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Optional[Sequence['outputs.GetField']]]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def number(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "number")

    @property
    @pulumi.getter
    def sections(self) -> pulumi.Output[Optional[Sequence['outputs.GetSection']]]:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of strings of the tags assigned to the item.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        The title of the item.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "valid_from")

    @property
    @pulumi.getter
    def vault(self) -> pulumi.Output[str]:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")

    @property
    @pulumi.getter(name="verificationNumber")
    def verification_number(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "verification_number")

