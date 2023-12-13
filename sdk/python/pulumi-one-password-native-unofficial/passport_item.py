# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['PassportItemArgs', 'PassportItem']

@pulumi.input_type
class PassportItemArgs:
    def __init__(__self__, *,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 date_of_birth: Optional[pulumi.Input[str]] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['FieldArgs', str]]]]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 issued_on: Optional[pulumi.Input[str]] = None,
                 issuing_authority: Optional[pulumi.Input[str]] = None,
                 issuing_country: Optional[pulumi.Input[str]] = None,
                 nationality: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 place_of_birth: Optional[pulumi.Input[str]] = None,
                 references: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UrlArgs', str]]]]] = None,
                 vault: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PassportItem resource.
        :param pulumi.Input[str] category: The category of the vault the item is in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        if attachments is not None:
            pulumi.set(__self__, "attachments", attachments)
        if category is not None:
            pulumi.set(__self__, "category", 'Passport')
        if date_of_birth is not None:
            pulumi.set(__self__, "date_of_birth", date_of_birth)
        if expiry_date is not None:
            pulumi.set(__self__, "expiry_date", expiry_date)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if full_name is not None:
            pulumi.set(__self__, "full_name", full_name)
        if gender is not None:
            pulumi.set(__self__, "gender", gender)
        if issued_on is not None:
            pulumi.set(__self__, "issued_on", issued_on)
        if issuing_authority is not None:
            pulumi.set(__self__, "issuing_authority", issuing_authority)
        if issuing_country is not None:
            pulumi.set(__self__, "issuing_country", issuing_country)
        if nationality is not None:
            pulumi.set(__self__, "nationality", nationality)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if number is not None:
            pulumi.set(__self__, "number", number)
        if place_of_birth is not None:
            pulumi.set(__self__, "place_of_birth", place_of_birth)
        if references is not None:
            pulumi.set(__self__, "references", references)
        if sections is not None:
            pulumi.set(__self__, "sections", sections)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if urls is not None:
            pulumi.set(__self__, "urls", urls)
        if vault is not None:
            pulumi.set(__self__, "vault", vault)

    @property
    @pulumi.getter
    def attachments(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]]:
        return pulumi.get(self, "attachments")

    @attachments.setter
    def attachments(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]]):
        pulumi.set(self, "attachments", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        """
        The category of the vault the item is in.
        """
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

    @property
    @pulumi.getter(name="dateOfBirth")
    def date_of_birth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "date_of_birth")

    @date_of_birth.setter
    def date_of_birth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date_of_birth", value)

    @property
    @pulumi.getter(name="expiryDate")
    def expiry_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiry_date")

    @expiry_date.setter
    def expiry_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry_date", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['FieldArgs', str]]]]]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union['FieldArgs', str]]]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "full_name")

    @full_name.setter
    def full_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_name", value)

    @property
    @pulumi.getter
    def gender(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gender")

    @gender.setter
    def gender(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gender", value)

    @property
    @pulumi.getter(name="issuedOn")
    def issued_on(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issued_on")

    @issued_on.setter
    def issued_on(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issued_on", value)

    @property
    @pulumi.getter(name="issuingAuthority")
    def issuing_authority(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issuing_authority")

    @issuing_authority.setter
    def issuing_authority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuing_authority", value)

    @property
    @pulumi.getter(name="issuingCountry")
    def issuing_country(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "issuing_country")

    @issuing_country.setter
    def issuing_country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuing_country", value)

    @property
    @pulumi.getter
    def nationality(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "nationality")

    @nationality.setter
    def nationality(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nationality", value)

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
    @pulumi.getter(name="placeOfBirth")
    def place_of_birth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "place_of_birth")

    @place_of_birth.setter
    def place_of_birth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "place_of_birth", value)

    @property
    @pulumi.getter
    def references(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "references")

    @references.setter
    def references(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "references", value)

    @property
    @pulumi.getter
    def sections(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]:
        return pulumi.get(self, "sections")

    @sections.setter
    def sections(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]):
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
    @pulumi.getter
    def urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Union['UrlArgs', str]]]]]:
        return pulumi.get(self, "urls")

    @urls.setter
    def urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UrlArgs', str]]]]]):
        pulumi.set(self, "urls", value)

    @property
    @pulumi.getter
    def vault(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")

    @vault.setter
    def vault(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vault", value)


@pulumi.input_type
class _PassportItemState:
    def __init__(__self__, *,
                 vault: pulumi.Input[str]):
        """
        Input properties used for looking up and filtering PassportItem resources.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        pulumi.set(__self__, "vault", vault)

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


class PassportItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 date_of_birth: Optional[pulumi.Input[str]] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.InputType['FieldArgs'], str]]]]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 issued_on: Optional[pulumi.Input[str]] = None,
                 issuing_authority: Optional[pulumi.Input[str]] = None,
                 issuing_country: Optional[pulumi.Input[str]] = None,
                 nationality: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 place_of_birth: Optional[pulumi.Input[str]] = None,
                 references: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[Union[pulumi.InputType['UrlArgs'], str]]]]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PassportItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category: The category of the vault the item is in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PassportItemArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PassportItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PassportItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PassportItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 date_of_birth: Optional[pulumi.Input[str]] = None,
                 expiry_date: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.InputType['FieldArgs'], str]]]]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 issued_on: Optional[pulumi.Input[str]] = None,
                 issuing_authority: Optional[pulumi.Input[str]] = None,
                 issuing_country: Optional[pulumi.Input[str]] = None,
                 nationality: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 place_of_birth: Optional[pulumi.Input[str]] = None,
                 references: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[Union[pulumi.InputType['UrlArgs'], str]]]]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PassportItemArgs.__new__(PassportItemArgs)

            __props__.__dict__["attachments"] = attachments
            __props__.__dict__["category"] = 'Passport'
            __props__.__dict__["date_of_birth"] = date_of_birth
            __props__.__dict__["expiry_date"] = expiry_date
            __props__.__dict__["fields"] = fields
            __props__.__dict__["full_name"] = full_name
            __props__.__dict__["gender"] = gender
            __props__.__dict__["issued_on"] = issued_on
            __props__.__dict__["issuing_authority"] = issuing_authority
            __props__.__dict__["issuing_country"] = issuing_country
            __props__.__dict__["nationality"] = nationality
            __props__.__dict__["notes"] = notes
            __props__.__dict__["number"] = number
            __props__.__dict__["place_of_birth"] = place_of_birth
            __props__.__dict__["references"] = references
            __props__.__dict__["sections"] = sections
            __props__.__dict__["tags"] = tags
            __props__.__dict__["title"] = title
            __props__.__dict__["type"] = type
            __props__.__dict__["urls"] = urls
            __props__.__dict__["vault"] = vault
            __props__.__dict__["id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["attachments", "fields", "sections"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(PassportItem, __self__).__init__(
            'one-password-native-unofficial:index:PassportItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            vault: Optional[pulumi.Input[str]] = None) -> 'PassportItem':
        """
        Get an existing PassportItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PassportItemState.__new__(_PassportItemState)

        __props__.__dict__["vault"] = vault
        __props__.__dict__["attachments"] = None
        __props__.__dict__["category"] = None
        __props__.__dict__["date_of_birth"] = None
        __props__.__dict__["expiry_date"] = None
        __props__.__dict__["fields"] = None
        __props__.__dict__["full_name"] = None
        __props__.__dict__["gender"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["issued_on"] = None
        __props__.__dict__["issuing_authority"] = None
        __props__.__dict__["issuing_country"] = None
        __props__.__dict__["nationality"] = None
        __props__.__dict__["notes"] = None
        __props__.__dict__["number"] = None
        __props__.__dict__["place_of_birth"] = None
        __props__.__dict__["references"] = None
        __props__.__dict__["sections"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["title"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["urls"] = None
        return PassportItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attachments(self) -> pulumi.Output[Mapping[str, 'outputs.OutputAttachment']]:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="dateOfBirth")
    def date_of_birth(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "date_of_birth")

    @property
    @pulumi.getter(name="expiryDate")
    def expiry_date(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expiry_date")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Mapping[str, 'outputs.OutputField']]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter
    def gender(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "gender")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="issuedOn")
    def issued_on(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "issued_on")

    @property
    @pulumi.getter(name="issuingAuthority")
    def issuing_authority(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "issuing_authority")

    @property
    @pulumi.getter(name="issuingCountry")
    def issuing_country(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "issuing_country")

    @property
    @pulumi.getter
    def nationality(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "nationality")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def number(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "number")

    @property
    @pulumi.getter(name="placeOfBirth")
    def place_of_birth(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "place_of_birth")

    @property
    @pulumi.getter
    def references(self) -> pulumi.Output[Sequence['outputs.OutputReference']]:
        return pulumi.get(self, "references")

    @property
    @pulumi.getter
    def sections(self) -> pulumi.Output[Mapping[str, 'outputs.OutputSection']]:
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
    def urls(self) -> pulumi.Output[Optional[Sequence['outputs.OutputUrl']]]:
        return pulumi.get(self, "urls")

    @property
    @pulumi.getter
    def vault(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "vault")

