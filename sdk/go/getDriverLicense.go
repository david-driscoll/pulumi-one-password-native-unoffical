// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDriverLicense(ctx *pulumi.Context, args *GetDriverLicenseArgs, opts ...pulumi.InvokeOption) (*GetDriverLicenseResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDriverLicenseResult
	err := ctx.Invoke("one-password-native-unofficial:index:GetDriverLicense", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDriverLicenseArgs struct {
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id *string `pulumi:"id"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetDriverLicenseResult struct {
	Address                *string                     `pulumi:"address"`
	Attachments            map[string]OutputAttachment `pulumi:"attachments"`
	Category               string                      `pulumi:"category"`
	ConditionsRestrictions *string                     `pulumi:"conditionsRestrictions"`
	Country                *string                     `pulumi:"country"`
	DateOfBirth            *string                     `pulumi:"dateOfBirth"`
	ExpiryDate             *string                     `pulumi:"expiryDate"`
	Fields                 map[string]OutputField      `pulumi:"fields"`
	FullName               *string                     `pulumi:"fullName"`
	Gender                 *string                     `pulumi:"gender"`
	Height                 *string                     `pulumi:"height"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id           string                   `pulumi:"id"`
	LicenseClass *string                  `pulumi:"licenseClass"`
	Notes        *string                  `pulumi:"notes"`
	Number       *string                  `pulumi:"number"`
	References   []OutputReference        `pulumi:"references"`
	Sections     map[string]OutputSection `pulumi:"sections"`
	State        *string                  `pulumi:"state"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title string            `pulumi:"title"`
	Urls  []OutputUrl       `pulumi:"urls"`
	Vault map[string]string `pulumi:"vault"`
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
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
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

func (o GetDriverLicenseResultOutput) Attachments() OutputAttachmentMapOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) map[string]OutputAttachment { return v.Attachments }).(OutputAttachmentMapOutput)
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

func (o GetDriverLicenseResultOutput) Fields() OutputFieldMapOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) map[string]OutputField { return v.Fields }).(OutputFieldMapOutput)
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

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetDriverLicenseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) string { return v.Id }).(pulumi.StringOutput)
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

func (o GetDriverLicenseResultOutput) References() OutputReferenceArrayOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) []OutputReference { return v.References }).(OutputReferenceArrayOutput)
}

func (o GetDriverLicenseResultOutput) Sections() OutputSectionMapOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) map[string]OutputSection { return v.Sections }).(OutputSectionMapOutput)
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

func (o GetDriverLicenseResultOutput) Urls() OutputUrlArrayOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) []OutputUrl { return v.Urls }).(OutputUrlArrayOutput)
}

func (o GetDriverLicenseResultOutput) Vault() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDriverLicenseResult) map[string]string { return v.Vault }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDriverLicenseResultOutput{})
}
