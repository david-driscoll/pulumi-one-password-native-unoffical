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
    'GetDriverLicenseResult',
    'AwaitableGetDriverLicenseResult',
    'get_driver_license',
    'get_driver_license_output',
]

@pulumi.output_type
class GetDriverLicenseResult:
    def __init__(__self__, address=None, category=None, conditions_restrictions=None, country=None, date_of_birth=None, expiry_date=None, fields=None, full_name=None, gender=None, height=None, id=None, license_class=None, notes=None, number=None, sections=None, state=None, tags=None, title=None, uuid=None, vault=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if category and not isinstance(category, dict):
            raise TypeError("Expected argument 'category' to be a dict")
        pulumi.set(__self__, "category", category)
        if conditions_restrictions and not isinstance(conditions_restrictions, str):
            raise TypeError("Expected argument 'conditions_restrictions' to be a str")
        pulumi.set(__self__, "conditions_restrictions", conditions_restrictions)
        if country and not isinstance(country, str):
            raise TypeError("Expected argument 'country' to be a str")
        pulumi.set(__self__, "country", country)
        if date_of_birth and not isinstance(date_of_birth, str):
            raise TypeError("Expected argument 'date_of_birth' to be a str")
        pulumi.set(__self__, "date_of_birth", date_of_birth)
        if expiry_date and not isinstance(expiry_date, str):
            raise TypeError("Expected argument 'expiry_date' to be a str")
        pulumi.set(__self__, "expiry_date", expiry_date)
        if fields and not isinstance(fields, dict):
            raise TypeError("Expected argument 'fields' to be a dict")
        pulumi.set(__self__, "fields", fields)
        if full_name and not isinstance(full_name, str):
            raise TypeError("Expected argument 'full_name' to be a str")
        pulumi.set(__self__, "full_name", full_name)
        if gender and not isinstance(gender, str):
            raise TypeError("Expected argument 'gender' to be a str")
        pulumi.set(__self__, "gender", gender)
        if height and not isinstance(height, str):
            raise TypeError("Expected argument 'height' to be a str")
        pulumi.set(__self__, "height", height)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if license_class and not isinstance(license_class, str):
            raise TypeError("Expected argument 'license_class' to be a str")
        pulumi.set(__self__, "license_class", license_class)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if number and not isinstance(number, str):
            raise TypeError("Expected argument 'number' to be a str")
        pulumi.set(__self__, "number", number)
        if sections and not isinstance(sections, dict):
            raise TypeError("Expected argument 'sections' to be a dict")
        pulumi.set(__self__, "sections", sections)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
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

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def category(self) -> Optional[str]:
        return pulumi.get(self, "category")

    @property
    @pulumi.getter(name="conditionsRestrictions")
    def conditions_restrictions(self) -> Optional[str]:
        return pulumi.get(self, "conditions_restrictions")

    @property
    @pulumi.getter
    def country(self) -> Optional[str]:
        return pulumi.get(self, "country")

    @property
    @pulumi.getter(name="dateOfBirth")
    def date_of_birth(self) -> Optional[str]:
        return pulumi.get(self, "date_of_birth")

    @property
    @pulumi.getter(name="expiryDate")
    def expiry_date(self) -> Optional[str]:
        return pulumi.get(self, "expiry_date")

    @property
    @pulumi.getter
    def fields(self) -> Optional[Mapping[str, 'outputs.GetField']]:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> Optional[str]:
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter
    def gender(self) -> Optional[str]:
        return pulumi.get(self, "gender")

    @property
    @pulumi.getter
    def height(self) -> Optional[str]:
        return pulumi.get(self, "height")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="licenseClass")
    def license_class(self) -> Optional[str]:
        return pulumi.get(self, "license_class")

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
    def sections(self) -> Optional[Mapping[str, 'outputs.GetSection']]:
        return pulumi.get(self, "sections")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        return pulumi.get(self, "state")

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


class AwaitableGetDriverLicenseResult(GetDriverLicenseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDriverLicenseResult(
            address=self.address,
            category=self.category,
            conditions_restrictions=self.conditions_restrictions,
            country=self.country,
            date_of_birth=self.date_of_birth,
            expiry_date=self.expiry_date,
            fields=self.fields,
            full_name=self.full_name,
            gender=self.gender,
            height=self.height,
            id=self.id,
            license_class=self.license_class,
            notes=self.notes,
            number=self.number,
            sections=self.sections,
            state=self.state,
            tags=self.tags,
            title=self.title,
            uuid=self.uuid,
            vault=self.vault)


def get_driver_license(title: Optional[str] = None,
                       uuid: Optional[str] = None,
                       vault: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDriverLicenseResult:
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
    __ret__ = pulumi.runtime.invoke('onepassword:index:GetDriverLicense', __args__, opts=opts, typ=GetDriverLicenseResult).value

    return AwaitableGetDriverLicenseResult(
        address=__ret__.address,
        category=__ret__.category,
        conditions_restrictions=__ret__.conditions_restrictions,
        country=__ret__.country,
        date_of_birth=__ret__.date_of_birth,
        expiry_date=__ret__.expiry_date,
        fields=__ret__.fields,
        full_name=__ret__.full_name,
        gender=__ret__.gender,
        height=__ret__.height,
        id=__ret__.id,
        license_class=__ret__.license_class,
        notes=__ret__.notes,
        number=__ret__.number,
        sections=__ret__.sections,
        state=__ret__.state,
        tags=__ret__.tags,
        title=__ret__.title,
        uuid=__ret__.uuid,
        vault=__ret__.vault)


@_utilities.lift_output_func(get_driver_license)
def get_driver_license_output(title: Optional[pulumi.Input[Optional[str]]] = None,
                              uuid: Optional[pulumi.Input[Optional[str]]] = None,
                              vault: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDriverLicenseResult]:
    """
    Use this data source to access information about an existing resource.

    :param str title: The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
    :param str uuid: The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
    :param str vault: The UUID of the vault the item is in.
    """
    ...
