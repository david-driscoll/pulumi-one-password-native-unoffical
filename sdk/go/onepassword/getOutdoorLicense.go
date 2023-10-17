// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOutdoorLicense(ctx *pulumi.Context, args *GetOutdoorLicenseArgs, opts ...pulumi.InvokeOption) (*GetOutdoorLicenseResult, error) {
	var rv GetOutdoorLicenseResult
	err := ctx.Invoke("onepassword:index:GetOutdoorLicense", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetOutdoorLicenseArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetOutdoorLicenseResult struct {
	ApprovedWildlife *string               `pulumi:"approvedWildlife"`
	Attachments      map[string]OutField   `pulumi:"attachments"`
	Category         string                `pulumi:"category"`
	Country          *string               `pulumi:"country"`
	Expires          *string               `pulumi:"expires"`
	Fields           map[string]OutField   `pulumi:"fields"`
	FullName         *string               `pulumi:"fullName"`
	MaximumQuota     *string               `pulumi:"maximumQuota"`
	Notes            *string               `pulumi:"notes"`
	References       map[string]OutField   `pulumi:"references"`
	Sections         map[string]OutSection `pulumi:"sections"`
	State            *string               `pulumi:"state"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid      string  `pulumi:"uuid"`
	ValidFrom *string `pulumi:"validFrom"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

func GetOutdoorLicenseOutput(ctx *pulumi.Context, args GetOutdoorLicenseOutputArgs, opts ...pulumi.InvokeOption) GetOutdoorLicenseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOutdoorLicenseResult, error) {
			args := v.(GetOutdoorLicenseArgs)
			r, err := GetOutdoorLicense(ctx, &args, opts...)
			var s GetOutdoorLicenseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOutdoorLicenseResultOutput)
}

type GetOutdoorLicenseOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetOutdoorLicenseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOutdoorLicenseArgs)(nil)).Elem()
}

type GetOutdoorLicenseResultOutput struct{ *pulumi.OutputState }

func (GetOutdoorLicenseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOutdoorLicenseResult)(nil)).Elem()
}

func (o GetOutdoorLicenseResultOutput) ToGetOutdoorLicenseResultOutput() GetOutdoorLicenseResultOutput {
	return o
}

func (o GetOutdoorLicenseResultOutput) ToGetOutdoorLicenseResultOutputWithContext(ctx context.Context) GetOutdoorLicenseResultOutput {
	return o
}

func (o GetOutdoorLicenseResultOutput) ApprovedWildlife() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) *string { return v.ApprovedWildlife }).(pulumi.StringPtrOutput)
}

func (o GetOutdoorLicenseResultOutput) Attachments() OutFieldMapOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) map[string]OutField { return v.Attachments }).(OutFieldMapOutput)
}

func (o GetOutdoorLicenseResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetOutdoorLicenseResultOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) *string { return v.Country }).(pulumi.StringPtrOutput)
}

func (o GetOutdoorLicenseResultOutput) Expires() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) *string { return v.Expires }).(pulumi.StringPtrOutput)
}

func (o GetOutdoorLicenseResultOutput) Fields() OutFieldMapOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) map[string]OutField { return v.Fields }).(OutFieldMapOutput)
}

func (o GetOutdoorLicenseResultOutput) FullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) *string { return v.FullName }).(pulumi.StringPtrOutput)
}

func (o GetOutdoorLicenseResultOutput) MaximumQuota() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) *string { return v.MaximumQuota }).(pulumi.StringPtrOutput)
}

func (o GetOutdoorLicenseResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetOutdoorLicenseResultOutput) References() OutFieldMapOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) map[string]OutField { return v.References }).(OutFieldMapOutput)
}

func (o GetOutdoorLicenseResultOutput) Sections() OutSectionMapOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) map[string]OutSection { return v.Sections }).(OutSectionMapOutput)
}

func (o GetOutdoorLicenseResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetOutdoorLicenseResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetOutdoorLicenseResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) string { return v.Title }).(pulumi.StringOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetOutdoorLicenseResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o GetOutdoorLicenseResultOutput) ValidFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) *string { return v.ValidFrom }).(pulumi.StringPtrOutput)
}

// The UUID of the vault the item is in.
func (o GetOutdoorLicenseResultOutput) Vault() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutdoorLicenseResult) string { return v.Vault }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOutdoorLicenseResultOutput{})
}
