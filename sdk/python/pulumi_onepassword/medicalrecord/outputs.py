# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'MedicationSection',
]

@pulumi.output_type
class MedicationSection(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "medicationNotes":
            suggest = "medication_notes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MedicationSection. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MedicationSection.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MedicationSection.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dosage: Optional[str] = None,
                 medication: Optional[str] = None,
                 medication_notes: Optional[str] = None):
        if dosage is not None:
            pulumi.set(__self__, "dosage", dosage)
        if medication is not None:
            pulumi.set(__self__, "medication", medication)
        if medication_notes is not None:
            pulumi.set(__self__, "medication_notes", medication_notes)

    @property
    @pulumi.getter
    def dosage(self) -> Optional[str]:
        return pulumi.get(self, "dosage")

    @property
    @pulumi.getter
    def medication(self) -> Optional[str]:
        return pulumi.get(self, "medication")

    @property
    @pulumi.getter(name="medicationNotes")
    def medication_notes(self) -> Optional[str]:
        return pulumi.get(self, "medication_notes")


