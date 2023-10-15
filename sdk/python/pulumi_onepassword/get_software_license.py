# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from . import softwarelicense as _softwarelicense
from ._enums import *

__all__ = [
    'GetSoftwareLicenseResult',
    'AwaitableGetSoftwareLicenseResult',
    'get_software_license',
    'get_software_license_output',
]

@pulumi.output_type
class GetSoftwareLicenseResult:
    def __init__(__self__, category=None, customer=None, fields=None, id=None, license_key=None, notes=None, order=None, publisher=None, sections=None, tags=None, title=None, uuid=None, vault=None, version=None):
        if category and not isinstance(category, dict):
            raise TypeError("Expected argument 'category' to be a dict")
        pulumi.set(__self__, "category", category)
        if customer and not isinstance(customer, dict):
            raise TypeError("Expected argument 'customer' to be a dict")
        pulumi.set(__self__, "customer", customer)
        if fields and not isinstance(fields, list):
            raise TypeError("Expected argument 'fields' to be a list")
        pulumi.set(__self__, "fields", fields)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if license_key and not isinstance(license_key, str):
            raise TypeError("Expected argument 'license_key' to be a str")
        pulumi.set(__self__, "license_key", license_key)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if order and not isinstance(order, dict):
            raise TypeError("Expected argument 'order' to be a dict")
        pulumi.set(__self__, "order", order)
        if publisher and not isinstance(publisher, dict):
            raise TypeError("Expected argument 'publisher' to be a dict")
        pulumi.set(__self__, "publisher", publisher)
        if sections and not isinstance(sections, list):
            raise TypeError("Expected argument 'sections' to be a list")
        pulumi.set(__self__, "sections", sections)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vault and not isinstance(vault, str):
            raise TypeError("Expected argument 'vault' to be a str")
        pulumi.set(__self__, "vault", vault)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def customer(self) -> Optional['_softwarelicense.outputs.CustomerSection']:
        return pulumi.get(self, "customer")

    @property
    @pulumi.getter
    def fields(self) -> Optional[Sequence['outputs.GetField']]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="licenseKey")
    def license_key(self) -> Optional[str]:
        return pulumi.get(self, "license_key")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def order(self) -> Optional['_softwarelicense.outputs.OrderSection']:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def publisher(self) -> Optional['_softwarelicense.outputs.PublisherSection']:
        return pulumi.get(self, "publisher")

    @property
    @pulumi.getter
    def sections(self) -> Optional[Sequence['outputs.GetSection']]:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        An array of strings of the tags assigned to the item.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def title(self) -> Optional[str]:
        """
        The title of the item.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def uuid(self) -> Optional[str]:
        """
        The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vault(self) -> Optional[str]:
        """
        The UUID of the vault the item is in.
        """
        return pulumi.get(self, "vault")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        return pulumi.get(self, "version")


class AwaitableGetSoftwareLicenseResult(GetSoftwareLicenseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSoftwareLicenseResult(
            category=self.category,
            customer=self.customer,
            fields=self.fields,
            id=self.id,
            license_key=self.license_key,
            notes=self.notes,
            order=self.order,
            publisher=self.publisher,
            sections=self.sections,
            tags=self.tags,
            title=self.title,
            uuid=self.uuid,
            vault=self.vault,
            version=self.version)


def get_software_license(title: Optional[str] = None,
                         uuid: Optional[str] = None,
                         vault: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSoftwareLicenseResult:
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
    __ret__ = pulumi.runtime.invoke('onepassword:index:GetSoftwareLicense', __args__, opts=opts, typ=GetSoftwareLicenseResult).value

    return AwaitableGetSoftwareLicenseResult(
        category=__ret__.category,
        customer=__ret__.customer,
        fields=__ret__.fields,
        id=__ret__.id,
        license_key=__ret__.license_key,
        notes=__ret__.notes,
        order=__ret__.order,
        publisher=__ret__.publisher,
        sections=__ret__.sections,
        tags=__ret__.tags,
        title=__ret__.title,
        uuid=__ret__.uuid,
        vault=__ret__.vault,
        version=__ret__.version)


@_utilities.lift_output_func(get_software_license)
def get_software_license_output(title: Optional[pulumi.Input[Optional[str]]] = None,
                                uuid: Optional[pulumi.Input[Optional[str]]] = None,
                                vault: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSoftwareLicenseResult]:
    """
    Use this data source to access information about an existing resource.

    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str uuid: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str vault: The UUID of the vault the item is in.
    """
    ...
