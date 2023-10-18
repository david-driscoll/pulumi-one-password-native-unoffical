// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDriverLicense(ctx *pulumi.Context, args *GetDriverLicenseArgs, opts ...pulumi.InvokeOption) (*GetDriverLicenseResult, error) {
	var rv GetDriverLicenseResult
	err := ctx.Invoke("one-password-native:index:GetDriverLicense", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDriverLicenseArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetDriverLicenseResult struct {
	Address                *string               `pulumi:"address"`
	Attachments            map[string]OutField   `pulumi:"attachments"`
	Category               string                `pulumi:"category"`
	ConditionsRestrictions *string               `pulumi:"conditionsRestrictions"`
	Country                *string               `pulumi:"country"`
	DateOfBirth            *string               `pulumi:"dateOfBirth"`
	ExpiryDate             *string               `pulumi:"expiryDate"`
	Fields                 map[string]OutField   `pulumi:"fields"`
	FullName               *string               `pulumi:"fullName"`
	Gender                 *string               `pulumi:"gender"`
	Height                 *string               `pulumi:"height"`
	LicenseClass           *string               `pulumi:"licenseClass"`
	Notes                  *string               `pulumi:"notes"`
	Number                 *string               `pulumi:"number"`
	References             map[string]OutField   `pulumi:"references"`
	Sections               map[string]OutSection `pulumi:"sections"`
	State                  *string               `pulumi:"state"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

func GetDriverLicenseOutput(ctx *pulumi.Context, args GetDriverLicenseOutputArgs, opts ...pulumi.InvokeOption) GetDriverLicenseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDriverLicenseResult, error) {
			args := v.(GetDriverLicenseArgs)
			r, err := GetDriverLicense(ctx, &args, opts...)
			var s GetDriverLicenseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDriverLicenseResultOutput)
}

type GetDriverLicenseOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetDriverLicenseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDriverLicenseArgs)(nil)).Elem()
}

type GetDriverLicenseResultOutput struct{ *pulumi.OutputState }

func (GetDriverLicenseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDriverLicenseResult)(nil)).Elem()
}

func (o GetDriverLicenseResultOutput) ToGetDriverLicenseResultOutput() GetDriverLicenseResultOutput {
	return o
}

func (o GetDriverLicenseResultOutput) ToGetDriverLicenseResultOutputWithContext(ctx context.Context) GetDriverLicenseResultOutput {
	return o
}

func (o GetDriverLicenseResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) Attachments() OutFieldMapOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) map[string]OutField { return v.Attachments }).(OutFieldMapOutput)
}

func (o GetDriverLicenseResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetDriverLicenseResultOutput) ConditionsRestrictions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.ConditionsRestrictions }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.Country }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) DateOfBirth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.DateOfBirth }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.ExpiryDate }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) Fields() OutFieldMapOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) map[string]OutField { return v.Fields }).(OutFieldMapOutput)
}

func (o GetDriverLicenseResultOutput) FullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.FullName }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) Gender() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.Gender }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) Height() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.Height }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) LicenseClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.LicenseClass }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) Number() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.Number }).(pulumi.StringPtrOutput)
}

func (o GetDriverLicenseResultOutput) References() OutFieldMapOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) map[string]OutField { return v.References }).(OutFieldMapOutput)
}

func (o GetDriverLicenseResultOutput) Sections() OutSectionMapOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) map[string]OutSection { return v.Sections }).(OutSectionMapOutput)
}

func (o GetDriverLicenseResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetDriverLicenseResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetDriverLicenseResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) string { return v.Title }).(pulumi.StringOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetDriverLicenseResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// The UUID of the vault the item is in.
func (o GetDriverLicenseResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) string { return v.Vault }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDriverLicenseResultOutput{})
}
