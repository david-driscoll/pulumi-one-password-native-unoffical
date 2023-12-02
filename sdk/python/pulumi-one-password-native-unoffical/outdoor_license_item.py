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

__all__ = ['OutdoorLicenseItemArgs', 'OutdoorLicenseItem']

@pulumi.input_type
class OutdoorLicenseItemArgs:
    def __init__(__self__, *,
                 vault: pulumi.Input[str],
                 approved_wildlife: Optional[pulumi.Input[str]] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 maximum_quota: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OutdoorLicenseItem resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        :param pulumi.Input[str] category: The category of the vault the item is in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        """
        pulumi.set(__self__, "vault", vault)
        if approved_wildlife is not None:
            pulumi.set(__self__, "approved_wildlife", approved_wildlife)
        if attachments is not None:
            pulumi.set(__self__, "attachments", attachments)
        if category is not None:
            pulumi.set(__self__, "category", 'Outdoor License')
        if country is not None:
            pulumi.set(__self__, "country", country)
        if expires is not None:
            pulumi.set(__self__, "expires", expires)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if full_name is not None:
            pulumi.set(__self__, "full_name", full_name)
        if maximum_quota is not None:
            pulumi.set(__self__, "maximum_quota", maximum_quota)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if sections is not None:
            pulumi.set(__self__, "sections", sections)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if valid_from is not None:
            pulumi.set(__self__, "valid_from", valid_from)

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
    @pulumi.getter(name="approvedWildlife")
    def approved_wildlife(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "approved_wildlife")

    @approved_wildlife.setter
    def approved_wildlife(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "approved_wildlife", value)

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
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter
    def expires(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expires")

    @expires.setter
    def expires(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "full_name")

    @full_name.setter
    def full_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_name", value)

    @property
    @pulumi.getter(name="maximumQuota")
    def maximum_quota(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "maximum_quota")

    @maximum_quota.setter
    def maximum_quota(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maximum_quota", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def sections(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]:
        return pulumi.get(self, "sections")

    @sections.setter
    def sections(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]):
        pulumi.set(self, "sections", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

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
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "valid_from")

    @valid_from.setter
    def valid_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_from", value)


@pulumi.input_type
class _OutdoorLicenseItemState:
    def __init__(__self__, *,
                 vault: pulumi.Input[str]):
        """
        Input properties used for looking up and filtering OutdoorLicenseItem resources.
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


class OutdoorLicenseItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approved_wildlife: Optional[pulumi.Input[str]] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 maximum_quota: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a OutdoorLicenseItem resource with the given unique name, props, and options.
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
                 args: OutdoorLicenseItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a OutdoorLicenseItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param OutdoorLicenseItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OutdoorLicenseItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 approved_wildlife: Optional[pulumi.Input[str]] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 expires: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 maximum_quota: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
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
            __props__ = OutdoorLicenseItemArgs.__new__(OutdoorLicenseItemArgs)

            __props__.__dict__["approved_wildlife"] = approved_wildlife
            __props__.__dict__["attachments"] = attachments
            __props__.__dict__["category"] = 'Outdoor License'
            __props__.__dict__["country"] = country
            __props__.__dict__["expires"] = expires
            __props__.__dict__["fields"] = fields
            __props__.__dict__["full_name"] = full_name
            __props__.__dict__["maximum_quota"] = maximum_quota
            __props__.__dict__["notes"] = notes
            __props__.__dict__["sections"] = sections
            __props__.__dict__["state"] = state
            __props__.__dict__["tags"] = tags
            __props__.__dict__["title"] = title
            __props__.__dict__["valid_from"] = valid_from
            if vault is None and not opts.urn:
                raise TypeError("Missing required property 'vault'")
            __props__.__dict__["vault"] = vault
            __props__.__dict__["references"] = None
            __props__.__dict__["uuid"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["attachments", "fields", "references", "sections"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(OutdoorLicenseItem, __self__).__init__(
            'one-password-native-unoffical:index:OutdoorLicenseItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            vault: Optional[pulumi.Input[str]] = None) -> 'OutdoorLicenseItem':
        """
        Get an existing OutdoorLicenseItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OutdoorLicenseItemState.__new__(_OutdoorLicenseItemState)

        __props__.__dict__["vault"] = vault
        __props__.__dict__["approved_wildlife"] = None
        __props__.__dict__["attachments"] = None
        __props__.__dict__["category"] = None
        __props__.__dict__["country"] = None
        __props__.__dict__["expires"] = None
        __props__.__dict__["fields"] = None
        __props__.__dict__["full_name"] = None
        __props__.__dict__["maximum_quota"] = None
        __props__.__dict__["notes"] = None
        __props__.__dict__["references"] = None
        __props__.__dict__["sections"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["title"] = None
        __props__.__dict__["uuid"] = None
        __props__.__dict__["valid_from"] = None
        return OutdoorLicenseItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="approvedWildlife")
    def approved_wildlife(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "approved_wildlife")

    @property
    @pulumi.getter
    def attachments(self) -> pulumi.Output[Mapping[str, 'outputs.OutField']]:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def country(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "country")

    @property
    @pulumi.getter
    def expires(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Mapping[str, 'outputs.OutField']]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter(name="maximumQuota")
    def maximum_quota(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "maximum_quota")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def references(self) -> pulumi.Output[Mapping[str, 'outputs.OutField']]:
        return pulumi.get(self, "references")

    @property
    @pulumi.getter
    def sections(self) -> pulumi.Output[Mapping[str, 'outputs.OutSection']]:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "state")

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
        def value(self) -> str:
            """
            the value of the attachment
            """
            return pulumi.get(self, "value")

    def get_attachment(__self__, *,
                       name: pulumi.Input[str]) -> pulumi.Output['OutdoorLicenseItem.GetAttachmentResult']:
        """

        :param pulumi.Input[str] name: The name or uuid of the attachment to get
        """
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['name'] = name
        return pulumi.runtime.call('one-password-native-unoffical:index:OutdoorLicenseItem/attachment', __args__, res=__self__, typ=OutdoorLicenseItem.GetAttachmentResult)

