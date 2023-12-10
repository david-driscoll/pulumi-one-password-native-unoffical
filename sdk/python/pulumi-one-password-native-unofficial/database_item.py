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

__all__ = ['DatabaseItemArgs', 'DatabaseItem']

@pulumi.input_type
class DatabaseItemArgs:
    def __init__(__self__, *,
                 vault: pulumi.Input[str],
                 alias: Optional[pulumi.Input[str]] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 connection_options: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input['UrlArgs']]]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabaseItem resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        :param pulumi.Input[str] category: The category of the vault the item is in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: An array of strings of the tags assigned to the item.
        :param pulumi.Input[str] title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
        """
        pulumi.set(__self__, "vault", vault)
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if attachments is not None:
            pulumi.set(__self__, "attachments", attachments)
        if category is not None:
            pulumi.set(__self__, "category", 'Database')
        if connection_options is not None:
            pulumi.set(__self__, "connection_options", connection_options)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if sections is not None:
            pulumi.set(__self__, "sections", sections)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if sid is not None:
            pulumi.set(__self__, "sid", sid)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if urls is not None:
            pulumi.set(__self__, "urls", urls)
        if username is not None:
            pulumi.set(__self__, "username", username)

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
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

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
    @pulumi.getter(name="connectionOptions")
    def connection_options(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "connection_options")

    @connection_options.setter
    def connection_options(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_options", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]:
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['FieldArgs']]]]):
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
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def sections(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]:
        return pulumi.get(self, "sections")

    @sections.setter
    def sections(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input['SectionArgs']]]]):
        pulumi.set(self, "sections", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def sid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sid")

    @sid.setter
    def sid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sid", value)

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
    def urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UrlArgs']]]]:
        return pulumi.get(self, "urls")

    @urls.setter
    def urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UrlArgs']]]]):
        pulumi.set(self, "urls", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _DatabaseItemState:
    def __init__(__self__, *,
                 vault: pulumi.Input[str]):
        """
        Input properties used for looking up and filtering DatabaseItem resources.
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


class DatabaseItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 connection_options: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UrlArgs']]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vault: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DatabaseItem resource with the given unique name, props, and options.
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
                 args: DatabaseItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DatabaseItem resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DatabaseItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 attachments: Optional[pulumi.Input[Mapping[str, pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 connection_options: Optional[pulumi.Input[str]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['FieldArgs']]]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 sections: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['SectionArgs']]]]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 sid: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UrlArgs']]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
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
            __props__ = DatabaseItemArgs.__new__(DatabaseItemArgs)

            __props__.__dict__["alias"] = alias
            __props__.__dict__["attachments"] = attachments
            __props__.__dict__["category"] = 'Database'
            __props__.__dict__["connection_options"] = connection_options
            __props__.__dict__["database"] = database
            __props__.__dict__["fields"] = fields
            __props__.__dict__["notes"] = notes
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["port"] = port
            __props__.__dict__["sections"] = sections
            __props__.__dict__["server"] = server
            __props__.__dict__["sid"] = sid
            __props__.__dict__["tags"] = tags
            __props__.__dict__["title"] = title
            __props__.__dict__["type"] = type
            __props__.__dict__["urls"] = urls
            __props__.__dict__["username"] = username
            if vault is None and not opts.urn:
                raise TypeError("Missing required property 'vault'")
            __props__.__dict__["vault"] = vault
            __props__.__dict__["id"] = None
            __props__.__dict__["references"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["attachments", "fields", "password", "sections"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(DatabaseItem, __self__).__init__(
            'one-password-native-unofficial:index:DatabaseItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            vault: Optional[pulumi.Input[str]] = None) -> 'DatabaseItem':
        """
        Get an existing DatabaseItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vault: The UUID of the vault the item is in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseItemState.__new__(_DatabaseItemState)

        __props__.__dict__["vault"] = vault
        __props__.__dict__["alias"] = None
        __props__.__dict__["attachments"] = None
        __props__.__dict__["category"] = None
        __props__.__dict__["connection_options"] = None
        __props__.__dict__["database"] = None
        __props__.__dict__["fields"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["notes"] = None
        __props__.__dict__["password"] = None
        __props__.__dict__["port"] = None
        __props__.__dict__["references"] = None
        __props__.__dict__["sections"] = None
        __props__.__dict__["server"] = None
        __props__.__dict__["sid"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["title"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["urls"] = None
        __props__.__dict__["username"] = None
        return DatabaseItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def attachments(self) -> pulumi.Output[Mapping[str, 'outputs.OutputAttachment']]:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="connectionOptions")
    def connection_options(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "connection_options")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Mapping[str, 'outputs.OutputField']]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "port")

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
    def server(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def sid(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "sid")

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
    def username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vault(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "vault")

