# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AddressSectionArgs',
    'IdentificationSectionArgs',
    'InternetDetailsSectionArgs',
]

@pulumi.input_type
class AddressSectionArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 business: Optional[pulumi.Input[str]] = None,
                 cell: Optional[pulumi.Input[str]] = None,
                 default_phone: Optional[pulumi.Input[str]] = None,
                 home: Optional[pulumi.Input[str]] = None):
        if address is not None:
            pulumi.set(__self__, "address", address)
        if business is not None:
            pulumi.set(__self__, "business", business)
        if cell is not None:
            pulumi.set(__self__, "cell", cell)
        if default_phone is not None:
            pulumi.set(__self__, "default_phone", default_phone)
        if home is not None:
            pulumi.set(__self__, "home", home)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def business(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "business")

    @business.setter
    def business(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "business", value)

    @property
    @pulumi.getter
    def cell(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cell")

    @cell.setter
    def cell(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cell", value)

    @property
    @pulumi.getter(name="defaultPhone")
    def default_phone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default_phone")

    @default_phone.setter
    def default_phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_phone", value)

    @property
    @pulumi.getter
    def home(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "home")

    @home.setter
    def home(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "home", value)


@pulumi.input_type
class IdentificationSectionArgs:
    def __init__(__self__, *,
                 birth_date: Optional[pulumi.Input[str]] = None,
                 company: Optional[pulumi.Input[str]] = None,
                 department: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 gender: Optional[pulumi.Input[str]] = None,
                 initial: Optional[pulumi.Input[str]] = None,
                 job_title: Optional[pulumi.Input[str]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 occupation: Optional[pulumi.Input[str]] = None):
        if birth_date is not None:
            pulumi.set(__self__, "birth_date", birth_date)
        if company is not None:
            pulumi.set(__self__, "company", company)
        if department is not None:
            pulumi.set(__self__, "department", department)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if gender is not None:
            pulumi.set(__self__, "gender", gender)
        if initial is not None:
            pulumi.set(__self__, "initial", initial)
        if job_title is not None:
            pulumi.set(__self__, "job_title", job_title)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if occupation is not None:
            pulumi.set(__self__, "occupation", occupation)

    @property
    @pulumi.getter(name="birthDate")
    def birth_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "birth_date")

    @birth_date.setter
    def birth_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "birth_date", value)

    @property
    @pulumi.getter
    def company(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "company")

    @company.setter
    def company(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "company", value)

    @property
    @pulumi.getter
    def department(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "department")

    @department.setter
    def department(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "department", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter
    def gender(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gender")

    @gender.setter
    def gender(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gender", value)

    @property
    @pulumi.getter
    def initial(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "initial")

    @initial.setter
    def initial(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initial", value)

    @property
    @pulumi.getter(name="jobTitle")
    def job_title(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "job_title")

    @job_title.setter
    def job_title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_title", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter
    def occupation(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "occupation")

    @occupation.setter
    def occupation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "occupation", value)


@pulumi.input_type
class InternetDetailsSectionArgs:
    def __init__(__self__, *,
                 aol_aim: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 forum_signature: Optional[pulumi.Input[str]] = None,
                 icq: Optional[pulumi.Input[str]] = None,
                 msn: Optional[pulumi.Input[str]] = None,
                 reminder_answer: Optional[pulumi.Input[str]] = None,
                 reminder_question: Optional[pulumi.Input[str]] = None,
                 skype: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 website: Optional[pulumi.Input[str]] = None,
                 yahoo: Optional[pulumi.Input[str]] = None):
        if aol_aim is not None:
            pulumi.set(__self__, "aol_aim", aol_aim)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if forum_signature is not None:
            pulumi.set(__self__, "forum_signature", forum_signature)
        if icq is not None:
            pulumi.set(__self__, "icq", icq)
        if msn is not None:
            pulumi.set(__self__, "msn", msn)
        if reminder_answer is not None:
            pulumi.set(__self__, "reminder_answer", reminder_answer)
        if reminder_question is not None:
            pulumi.set(__self__, "reminder_question", reminder_question)
        if skype is not None:
            pulumi.set(__self__, "skype", skype)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if website is not None:
            pulumi.set(__self__, "website", website)
        if yahoo is not None:
            pulumi.set(__self__, "yahoo", yahoo)

    @property
    @pulumi.getter(name="aolAim")
    def aol_aim(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "aol_aim")

    @aol_aim.setter
    def aol_aim(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aol_aim", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="forumSignature")
    def forum_signature(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "forum_signature")

    @forum_signature.setter
    def forum_signature(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forum_signature", value)

    @property
    @pulumi.getter
    def icq(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "icq")

    @icq.setter
    def icq(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icq", value)

    @property
    @pulumi.getter
    def msn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "msn")

    @msn.setter
    def msn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "msn", value)

    @property
    @pulumi.getter(name="reminderAnswer")
    def reminder_answer(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "reminder_answer")

    @reminder_answer.setter
    def reminder_answer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reminder_answer", value)

    @property
    @pulumi.getter(name="reminderQuestion")
    def reminder_question(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "reminder_question")

    @reminder_question.setter
    def reminder_question(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reminder_question", value)

    @property
    @pulumi.getter
    def skype(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "skype")

    @skype.setter
    def skype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "skype", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def website(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "website")

    @website.setter
    def website(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "website", value)

    @property
    @pulumi.getter
    def yahoo(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "yahoo")

    @yahoo.setter
    def yahoo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "yahoo", value)


