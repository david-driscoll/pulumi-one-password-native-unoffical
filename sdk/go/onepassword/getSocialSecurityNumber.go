// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSocialSecurityNumber(ctx *pulumi.Context, args *GetSocialSecurityNumberArgs, opts ...pulumi.InvokeOption) (*GetSocialSecurityNumberResult, error) {
	var rv GetSocialSecurityNumberResult
	err := ctx.Invoke("onepassword:index:GetSocialSecurityNumber", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSocialSecurityNumberArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetSocialSecurityNumberResult struct {
	Category *string      `pulumi:"category"`
	Fields   []GetField   `pulumi:"fields"`
	Id       *string      `pulumi:"id"`
	Name     *string      `pulumi:"name"`
	Notes    *string      `pulumi:"notes"`
	Number   *string      `pulumi:"number"`
	Sections []GetSection `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

func GetSocialSecurityNumberOutput(ctx *pulumi.Context, args GetSocialSecurityNumberOutputArgs, opts ...pulumi.InvokeOption) GetSocialSecurityNumberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSocialSecurityNumberResult, error) {
			args := v.(GetSocialSecurityNumberArgs)
			r, err := GetSocialSecurityNumber(ctx, &args, opts...)
			var s GetSocialSecurityNumberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSocialSecurityNumberResultOutput)
}

type GetSocialSecurityNumberOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetSocialSecurityNumberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSocialSecurityNumberArgs)(nil)).Elem()
}

type GetSocialSecurityNumberResultOutput struct{ *pulumi.OutputState }

func (GetSocialSecurityNumberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSocialSecurityNumberResult)(nil)).Elem()
}

func (o GetSocialSecurityNumberResultOutput) ToGetSocialSecurityNumberResultOutput() GetSocialSecurityNumberResultOutput {
	return o
}

func (o GetSocialSecurityNumberResultOutput) ToGetSocialSecurityNumberResultOutputWithContext(ctx context.Context) GetSocialSecurityNumberResultOutput {
	return o
}

func (o GetSocialSecurityNumberResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o GetSocialSecurityNumberResultOutput) Fields() GetFieldArrayOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) []GetField { return v.Fields }).(GetFieldArrayOutput)
}

func (o GetSocialSecurityNumberResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GetSocialSecurityNumberResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetSocialSecurityNumberResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetSocialSecurityNumberResultOutput) Number() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) *string { return v.Number }).(pulumi.StringPtrOutput)
}

func (o GetSocialSecurityNumberResultOutput) Sections() GetSectionArrayOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) []GetSection { return v.Sections }).(GetSectionArrayOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetSocialSecurityNumberResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetSocialSecurityNumberResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetSocialSecurityNumberResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

// The UUID of the vault the item is in.
func (o GetSocialSecurityNumberResultOutput) Vault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSocialSecurityNumberResult) *string { return v.Vault }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSocialSecurityNumberResultOutput{})
}
