// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWirelessRouter(ctx *pulumi.Context, args *GetWirelessRouterArgs, opts ...pulumi.InvokeOption) (*GetWirelessRouterResult, error) {
	var rv GetWirelessRouterResult
	err := ctx.Invoke("onepassword:index:GetWirelessRouter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetWirelessRouterArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetWirelessRouterResult struct {
	AirPortId               *string      `pulumi:"airPortId"`
	AttachedStoragePassword *string      `pulumi:"attachedStoragePassword"`
	BaseStationName         *string      `pulumi:"baseStationName"`
	BaseStationPassword     *string      `pulumi:"baseStationPassword"`
	Category                *string      `pulumi:"category"`
	Fields                  []GetField   `pulumi:"fields"`
	Id                      *string      `pulumi:"id"`
	NetworkName             *string      `pulumi:"networkName"`
	Notes                   *string      `pulumi:"notes"`
	Sections                []GetSection `pulumi:"sections"`
	ServerIpAddress         *string      `pulumi:"serverIpAddress"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title *string `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid *string `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault                   *string `pulumi:"vault"`
	WirelessNetworkPassword *string `pulumi:"wirelessNetworkPassword"`
	WirelessSecurity        *string `pulumi:"wirelessSecurity"`
}

func GetWirelessRouterOutput(ctx *pulumi.Context, args GetWirelessRouterOutputArgs, opts ...pulumi.InvokeOption) GetWirelessRouterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWirelessRouterResult, error) {
			args := v.(GetWirelessRouterArgs)
			r, err := GetWirelessRouter(ctx, &args, opts...)
			var s GetWirelessRouterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWirelessRouterResultOutput)
}

type GetWirelessRouterOutputArgs struct {
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetWirelessRouterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWirelessRouterArgs)(nil)).Elem()
}

type GetWirelessRouterResultOutput struct{ *pulumi.OutputState }

func (GetWirelessRouterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWirelessRouterResult)(nil)).Elem()
}

func (o GetWirelessRouterResultOutput) ToGetWirelessRouterResultOutput() GetWirelessRouterResultOutput {
	return o
}

func (o GetWirelessRouterResultOutput) ToGetWirelessRouterResultOutputWithContext(ctx context.Context) GetWirelessRouterResultOutput {
	return o
}

func (o GetWirelessRouterResultOutput) AirPortId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.AirPortId }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) AttachedStoragePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.AttachedStoragePassword }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) BaseStationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.BaseStationName }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) BaseStationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.BaseStationPassword }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) Fields() GetFieldArrayOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) []GetField { return v.Fields }).(GetFieldArrayOutput)
}

func (o GetWirelessRouterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) NetworkName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.NetworkName }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) Sections() GetSectionArrayOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) []GetSection { return v.Sections }).(GetSectionArrayOutput)
}

func (o GetWirelessRouterResultOutput) ServerIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.ServerIpAddress }).(pulumi.StringPtrOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetWirelessRouterResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetWirelessRouterResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetWirelessRouterResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

// The UUID of the vault the item is in.
func (o GetWirelessRouterResultOutput) Vault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.Vault }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) WirelessNetworkPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.WirelessNetworkPassword }).(pulumi.StringPtrOutput)
}

func (o GetWirelessRouterResultOutput) WirelessSecurity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWirelessRouterResult) *string { return v.WirelessSecurity }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWirelessRouterResultOutput{})
}
