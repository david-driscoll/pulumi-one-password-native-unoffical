# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import bankaccount as _bankaccount
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['BankAccountItemArgs', 'BankAccountItem']

@pulumi.input_type
class BankAccountItemArgs:
    def __init__(__self__, *,
                 vault: pulumi.Input[str],
                 account_number: Optional[pulumi.Input[str]] = None,
                 bank_name: Optional[pulumi.Input[str]] = None,
                 branch_information: Optional[pulumi.Input['_bankaccount.BranchInformationSectionArgs']] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]] = None,
                 iban: Optional[pulumi.Input[str]] = None,
                 input_attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 name_on_account: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 routing_number: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]] = None,
                 swift: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BankAccountItem resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        :param pulumi.Input[str] category: The category of the vault the item is in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        """
        pulumi.set(__self__, "vault", vault)
        if account_number is not None:
            pulumi.set(__self__, "account_number", account_number)
        if bank_name is not None:
            pulumi.set(__self__, "bank_name", bank_name)
        if branch_information is not None:
            pulumi.set(__self__, "branch_information", branch_information)
        if category is not None:
            pulumi.set(__self__, "category", 'Bank Account')
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if iban is not None:
            pulumi.set(__self__, "iban", iban)
        if input_attachments is not None:
            pulumi.set(__self__, "input_attachments", input_attachments)
        if name_on_account is not None:
            pulumi.set(__self__, "name_on_account", name_on_account)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if pin is not None:
            pulumi.set(__self__, "pin", pin)
        if routing_number is not None:
            pulumi.set(__self__, "routing_number", routing_number)
        if sections is not None:
            pulumi.set(__self__, "sections", sections)
        if swift is not None:
            pulumi.set(__self__, "swift", swift)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)

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
    @pulumi.getter(name="accountNumber")
    def account_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "account_number")

    @account_number.setter
    def account_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_number", value)

    @property
    @pulumi.getter(name="bankName")
    def bank_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bank_name")

    @bank_name.setter
    def bank_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bank_name", value)

    @property
    @pulumi.getter(name="branchInformation")
    def branch_information(self) -> Optional[pulumi.Input['_bankaccount.BranchInformationSectionArgs']]:
        return pulumi.get(self, "branch_information")

    @branch_information.setter
    def branch_information(self, value: Optional[pulumi.Input['_bankaccount.BranchInformationSectionArgs']]):
        pulumi.set(self, "branch_information", value)

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
    def fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def iban(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iban")

    @iban.setter
    def iban(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iban", value)

    @property
    @pulumi.getter(name="inputAttachments")
    def input_attachments(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]]:
        return pulumi.get(self, "input_attachments")

    @input_attachments.setter
    def input_attachments(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]]):
        pulumi.set(self, "input_attachments", value)

    @property
    @pulumi.getter(name="nameOnAccount")
    def name_on_account(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name_on_account")

    @name_on_account.setter
    def name_on_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_on_account", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter
    def pin(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pin")

    @pin.setter
    def pin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pin", value)

    @property
    @pulumi.getter(name="routingNumber")
    def routing_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "routing_number")

    @routing_number.setter
    def routing_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "routing_number", value)

    @property
    @pulumi.getter
    def sections(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]:
        return pulumi.get(self, "sections")

    @sections.setter
    def sections(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]):
        pulumi.set(self, "sections", value)

    @property
    @pulumi.getter
    def swift(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "swift")

    @swift.setter
    def swift(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "swift", value)

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


@pulumi.input_type
class _BankAccountItemState:
    def __init__(__self__, *,
                 vault: pulumi.Input[str]):
        """
        Input properties used for looking up and filtering BankAccountItem resources.
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


class BankAccountItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_number: Optional[pulumi.Input[str]] = None,
                 bank_name: Optional[pulumi.Input[str]] = None,
                 branch_information: Optional[pulumi.Input[pulumi.InputType['_bankaccount.BranchInformationSectionArgs']]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 iban: Optional[pulumi.Input[str]] = None,
                 input_attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 name_on_account: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 routing_number: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 swift: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a BankAccountItem resource with the given unique name, props, and options.
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
                 args: BankAccountItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BankAccountItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BankAccountItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BankAccountItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_number: Optional[pulumi.Input[str]] = None,
                 bank_name: Optional[pulumi.Input[str]] = None,
                 branch_information: Optional[pulumi.Input[pulumi.InputType['_bankaccount.BranchInformationSectionArgs']]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 iban: Optional[pulumi.Input[str]] = None,
                 input_attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 name_on_account: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 pin: Optional[pulumi.Input[str]] = None,
                 routing_number: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 swift: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = BankAccountItemArgs.__new__(BankAccountItemArgs)

            __props__.__dict__["account_number"] = account_number
            __props__.__dict__["bank_name"] = bank_name
            __props__.__dict__["branch_information"] = branch_information
            __props__.__dict__["category"] = 'Bank Account'
            __props__.__dict__["fields"] = fields
            __props__.__dict__["iban"] = iban
            __props__.__dict__["input_attachments"] = input_attachments
            __props__.__dict__["name_on_account"] = name_on_account
            __props__.__dict__["notes"] = notes
            __props__.__dict__["pin"] = None if pin is None else pulumi.Output.secret(pin)
            __props__.__dict__["routing_number"] = routing_number
            __props__.__dict__["sections"] = sections
            __props__.__dict__["swift"] = swift
            __props__.__dict__["tags"] = tags
            __props__.__dict__["title"] = title
            __props__.__dict__["type"] = type
            if vault is None and not opts.urn:
                raise TypeError("Missing required property 'vault'")
            __props__.__dict__["vault"] = vault
            __props__.__dict__["attachments"] = None
            __props__.__dict__["references"] = None
            __props__.__dict__["uuid"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["attachments", "fields", "pin", "references", "sections"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(BankAccountItem, __self__).__init__(
            'one-password-native-unoffical:index:BankAccountItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            vault: Optional[pulumi.Input[str]] = None) -> 'BankAccountItem':
        """
        Get an existing BankAccountItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BankAccountItemState.__new__(_BankAccountItemState)

        __props__.__dict__["vault"] = vault
        __props__.__dict__["account_number"] = None
        __props__.__dict__["attachments"] = None
        __props__.__dict__["bank_name"] = None
        __props__.__dict__["branch_information"] = None
        __props__.__dict__["category"] = None
        __props__.__dict__["fields"] = None
        __props__.__dict__["iban"] = None
        __props__.__dict__["name_on_account"] = None
        __props__.__dict__["notes"] = None
        __props__.__dict__["pin"] = None
        __props__.__dict__["references"] = None
        __props__.__dict__["routing_number"] = None
        __props__.__dict__["sections"] = None
        __props__.__dict__["swift"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["title"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["uuid"] = None
        return BankAccountItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountNumber")
    def account_number(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "account_number")

    @property
    @pulumi.getter
    def attachments(self) -> pulumi.Output[Mapping[str, 'outputs.OutAttachment']]:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter(name="bankName")
    def bank_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "bank_name")

    @property
    @pulumi.getter(name="branchInformation")
    def branch_information(self) -> pulumi.Output[Optional['_bankaccount.outputs.BranchInformationSection']]:
        return pulumi.get(self, "branch_information")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Mapping[str, 'outputs.OutField']]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def iban(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "iban")

    @property
    @pulumi.getter(name="nameOnAccount")
    def name_on_account(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name_on_account")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def pin(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "pin")

    @property
    @pulumi.getter
    def references(self) -> pulumi.Output[Mapping[str, 'outputs.OutField']]:
        return pulumi.get(self, "references")

    @property
    @pulumi.getter(name="routingNumber")
    def routing_number(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "routing_number")

    @property
    @pulumi.getter
    def sections(self) -> pulumi.Output[Mapping[str, 'outputs.OutSection']]:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def swift(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "swift")

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
                       name: pulumi.Input[str]) -> pulumi.Output['BankAccountItem.GetAttachmentResult']:
        """

        :param pulumi.Input[str] name: The name or uuid of the attachment to get
        """
        __args__ = dict()
        __args__['__self__'] = __self__
        __args__['name'] = name
        return pulumi.runtime.call('one-password-native-unoffical:index:BankAccountItem/attachment', __args__, res=__self__, typ=BankAccountItem.GetAttachmentResult)

