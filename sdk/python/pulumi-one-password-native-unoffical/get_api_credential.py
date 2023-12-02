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

__all__ = [
    'GetAPICredentialResult',
    'AwaitableGetAPICredentialResult',
    'get_api_credential',
    'get_api_credential_output',
]

@pulumi.output_type
class GetAPICredentialResult:
    def __init__(__self__, attachments=None, category=None, credential=None, expires=None, fields=None, filename=None, hostname=None, notes=None, references=None, sections=None, tags=None, title=None, type=None, username=None, uuid=None, valid_from=None, vault=None):
        if attachments and not isinstance(attachments, dict):
            raise TypeError("Expected argument 'attachments' to be a dict")
        pulumi.set(__self__, "attachments", attachments)
        if category and not isinstance(category, dict):
            raise TypeError("Expected argument 'category' to be a dict")
        pulumi.set(__self__, "category", category)
        if credential and not isinstance(credential, str):
            raise TypeError("Expected argument 'credential' to be a str")
        pulumi.set(__self__, "credential", credential)
        if expires and not isinstance(expires, str):
            raise TypeError("Expected argument 'expires' to be a str")
        pulumi.set(__self__, "expires", expires)
        if fields and not isinstance(fields, dict):
            raise TypeError("Expected argument 'fields' to be a dict")
        pulumi.set(__self__, "fields", fields)
        if filename and not isinstance(filename, str):
            raise TypeError("Expected argument 'filename' to be a str")
        pulumi.set(__self__, "filename", filename)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
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
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if valid_from and not isinstance(valid_from, str):
            raise TypeError("Expected argument 'valid_from' to be a str")
        pulumi.set(__self__, "valid_from", valid_from)
        if vault and not isinstance(vault, str):
            raise TypeError("Expected argument 'vault' to be a str")
        pulumi.set(__self__, "vault", vault)

    @property
    @pulumi.getter
    def attachments(self) -> Mapping[str, 'outputs.OutField']:
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter
    def category(self) -> str:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def credential(self) -> Optional[str]:
        return pulumi.get(self, "credential")

    @property
    @pulumi.getter
    def expires(self) -> Optional[str]:
        return pulumi.get(self, "expires")

    @property
    @pulumi.getter
    def fields(self) -> Mapping[str, 'outputs.OutField']:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def filename(self) -> Optional[str]:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

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
    def username(self) -> Optional[str]:
        return pulumi.get(self, "username")

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


class AwaitableGetAPICredentialResult(GetAPICredentialResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAPICredentialResult(
            attachments=self.attachments,
            category=self.category,
            credential=self.credential,
            expires=self.expires,
            fields=self.fields,
            filename=self.filename,
            hostname=self.hostname,
            notes=self.notes,
            references=self.references,
            sections=self.sections,
            tags=self.tags,
            title=self.title,
            type=self.type,
            username=self.username,
            uuid=self.uuid,
            valid_from=self.valid_from,
            vault=self.vault)


def get_api_credential(title: Optional[str] = None,
                       uuid: Optional[str] = None,
                       vault: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAPICredentialResult:
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
    __ret__ = pulumi.runtime.invoke('one-password-native-unoffical:index:GetAPICredential', __args__, opts=opts, typ=GetAPICredentialResult).value

    return AwaitableGetAPICredentialResult(
        attachments=__ret__.attachments,
        category=__ret__.category,
        credential=__ret__.credential,
        expires=__ret__.expires,
        fields=__ret__.fields,
        filename=__ret__.filename,
        hostname=__ret__.hostname,
        notes=__ret__.notes,
        references=__ret__.references,
        sections=__ret__.sections,
        tags=__ret__.tags,
        title=__ret__.title,
        type=__ret__.type,
        username=__ret__.username,
        uuid=__ret__.uuid,
        valid_from=__ret__.valid_from,
        vault=__ret__.vault)


@_utilities.lift_output_func(get_api_credential)
def get_api_credential_output(title: Optional[pulumi.Input[Optional[str]]] = None,
                              uuid: Optional[pulumi.Input[Optional[str]]] = None,
                              vault: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAPICredentialResult]:
    """
    Use this data source to access information about an existing resource.

    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str uuid: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str vault: The UUID of the vault the item is in.
    """
    ...
