# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CustomerSectionArgs',
    'OrderSectionArgs',
    'PublisherSectionArgs',
]

@pulumi.input_type
class CustomerSectionArgs:
    def __init__(__self__, *,
                 company: Optional[pulumi.Input[str]] = None,
                 licensed_to: Optional[pulumi.Input[str]] = None,
                 registered_email: Optional[pulumi.Input[str]] = None):
        if company is not None:
            pulumi.set(__self__, "company", company)
        if licensed_to is not None:
            pulumi.set(__self__, "licensed_to", licensed_to)
        if registered_email is not None:
            pulumi.set(__self__, "registered_email", registered_email)

    @property
    @pulumi.getter
    def company(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "company")

    @company.setter
    def company(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "company", value)

    @property
    @pulumi.getter(name="licensedTo")
    def licensed_to(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "licensed_to")

    @licensed_to.setter
    def licensed_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "licensed_to", value)

    @property
    @pulumi.getter(name="registeredEmail")
    def registered_email(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "registered_email")

    @registered_email.setter
    def registered_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registered_email", value)


@pulumi.input_type
class OrderSectionArgs:
    def __init__(__self__, *,
                 order_number: Optional[pulumi.Input[str]] = None,
                 order_total: Optional[pulumi.Input[str]] = None,
                 purchase_date: Optional[pulumi.Input[str]] = None):
        if order_number is not None:
            pulumi.set(__self__, "order_number", order_number)
        if order_total is not None:
            pulumi.set(__self__, "order_total", order_total)
        if purchase_date is not None:
            pulumi.set(__self__, "purchase_date", purchase_date)

    @property
    @pulumi.getter(name="orderNumber")
    def order_number(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "order_number")

    @order_number.setter
    def order_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "order_number", value)

    @property
    @pulumi.getter(name="orderTotal")
    def order_total(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "order_total")

    @order_total.setter
    def order_total(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "order_total", value)

    @property
    @pulumi.getter(name="purchaseDate")
    def purchase_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "purchase_date")

    @purchase_date.setter
    def purchase_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purchase_date", value)


@pulumi.input_type
class PublisherSectionArgs:
    def __init__(__self__, *,
                 download_page: Optional[pulumi.Input[str]] = None,
                 publisher: Optional[pulumi.Input[str]] = None,
                 retail_price: Optional[pulumi.Input[str]] = None,
                 support_email: Optional[pulumi.Input[str]] = None,
                 website: Optional[pulumi.Input[str]] = None):
        if download_page is not None:
            pulumi.set(__self__, "download_page", download_page)
        if publisher is not None:
            pulumi.set(__self__, "publisher", publisher)
        if retail_price is not None:
            pulumi.set(__self__, "retail_price", retail_price)
        if support_email is not None:
            pulumi.set(__self__, "support_email", support_email)
        if website is not None:
            pulumi.set(__self__, "website", website)

    @property
    @pulumi.getter(name="downloadPage")
    def download_page(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "download_page")

    @download_page.setter
    def download_page(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "download_page", value)

    @property
    @pulumi.getter
    def publisher(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "publisher")

    @publisher.setter
    def publisher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "publisher", value)

    @property
    @pulumi.getter(name="retailPrice")
    def retail_price(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "retail_price")

    @retail_price.setter
    def retail_price(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retail_price", value)

    @property
    @pulumi.getter(name="supportEmail")
    def support_email(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "support_email")

    @support_email.setter
    def support_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "support_email", value)

    @property
    @pulumi.getter
    def website(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "website")

    @website.setter
    def website(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "website", value)


