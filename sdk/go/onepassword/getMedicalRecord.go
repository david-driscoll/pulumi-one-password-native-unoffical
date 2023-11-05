// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/Sacro/pulumi-onepassword/sdk/go/onepassword/medicalrecord"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMedicalRecord(ctx *pulumi.Context, args *GetMedicalRecordArgs, opts ...pulumi.InvokeOption) (*GetMedicalRecordResult, error) {
	var rv GetMedicalRecordResult
	err := ctx.Invoke("one-password-native:index:GetMedicalRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMedicalRecordArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetMedicalRecordResult struct {
	Attachments            map[string]OutField              `pulumi:"attachments"`
	Category               string                           `pulumi:"category"`
	Date                   *string                          `pulumi:"date"`
	Fields                 map[string]OutField              `pulumi:"fields"`
	HealthcareProfessional *string                          `pulumi:"healthcareProfessional"`
	Location               *string                          `pulumi:"location"`
	Medication             *medicalrecord.MedicationSection `pulumi:"medication"`
	Notes                  *string                          `pulumi:"notes"`
	Patient                *string                          `pulumi:"patient"`
	ReasonForVisit         *string                          `pulumi:"reasonForVisit"`
	References             map[string]OutField              `pulumi:"references"`
	Sections               map[string]OutSection            `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

func GetMedicalRecordOutput(ctx *pulumi.Context, args GetMedicalRecordOutputArgs, opts ...pulumi.InvokeOption) GetMedicalRecordResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMedicalRecordResult, error) {
			args := v.(GetMedicalRecordArgs)
			r, err := GetMedicalRecord(ctx, &args, opts...)
			var s GetMedicalRecordResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMedicalRecordResultOutput)
}

type GetMedicalRecordOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetMedicalRecordOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMedicalRecordArgs)(nil)).Elem()
}

type GetMedicalRecordResultOutput struct{ *pulumi.OutputState }

func (GetMedicalRecordResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMedicalRecordResult)(nil)).Elem()
}

func (o GetMedicalRecordResultOutput) ToGetMedicalRecordResultOutput() GetMedicalRecordResultOutput {
	return o
}

func (o GetMedicalRecordResultOutput) ToGetMedicalRecordResultOutputWithContext(ctx context.Context) GetMedicalRecordResultOutput {
	return o
}

func (o GetMedicalRecordResultOutput) Attachments() OutFieldMapOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) map[string]OutField { return v.Attachments }).(OutFieldMapOutput)
}

func (o GetMedicalRecordResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetMedicalRecordResultOutput) Date() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) *string { return v.Date }).(pulumi.StringPtrOutput)
}

func (o GetMedicalRecordResultOutput) Fields() OutFieldMapOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) map[string]OutField { return v.Fields }).(OutFieldMapOutput)
}

func (o GetMedicalRecordResultOutput) HealthcareProfessional() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) *string { return v.HealthcareProfessional }).(pulumi.StringPtrOutput)
}

func (o GetMedicalRecordResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GetMedicalRecordResultOutput) Medication() medicalrecord.MedicationSectionPtrOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) *medicalrecord.MedicationSection { return v.Medication }).(medicalrecord.MedicationSectionPtrOutput)
}

func (o GetMedicalRecordResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetMedicalRecordResultOutput) Patient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) *string { return v.Patient }).(pulumi.StringPtrOutput)
}

func (o GetMedicalRecordResultOutput) ReasonForVisit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) *string { return v.ReasonForVisit }).(pulumi.StringPtrOutput)
}

func (o GetMedicalRecordResultOutput) References() OutFieldMapOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) map[string]OutField { return v.References }).(OutFieldMapOutput)
}

func (o GetMedicalRecordResultOutput) Sections() OutSectionMapOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) map[string]OutSection { return v.Sections }).(OutSectionMapOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetMedicalRecordResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetMedicalRecordResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) string { return v.Title }).(pulumi.StringOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetMedicalRecordResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The UUID of the vault the item is in.
func (o GetMedicalRecordResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v GetMedicalRecordResult) string { return v.Vault }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMedicalRecordResultOutput{})
}
