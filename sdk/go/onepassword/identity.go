// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-onepassword/sdk/go/onepassword/identity"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Identity struct {
	pulumi.CustomResourceState

	Address         identity.AddressPtrOutput         `pulumi:"address"`
	Category        pulumi.StringOutput               `pulumi:"category"`
	Fields          GetFieldArrayOutput               `pulumi:"fields"`
	Id              pulumi.StringOutput               `pulumi:"id"`
	Identification  identity.IdentificationPtrOutput  `pulumi:"identification"`
	InternetDetails identity.InternetDetailsPtrOutput `pulumi:"internetDetails"`
	Notes           pulumi.StringPtrOutput            `pulumi:"notes"`
	Sections        GetSectionArrayOutput             `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput `pulumi:"title"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewIdentity registers a new resource with the given unique name, arguments, and options.
func NewIdentity(ctx *pulumi.Context,
	name string, args *IdentityArgs, opts ...pulumi.ResourceOption) (*Identity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	if isZero(args.Category) {
		args.Category = pulumi.String("Item")
	}
	var resource Identity
	err := ctx.RegisterResource("onepassword:index:Identity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentity gets an existing Identity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityState, opts ...pulumi.ResourceOption) (*Identity, error) {
	var resource Identity
	err := ctx.ReadResource("onepassword:index:Identity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Identity resources.
type identityState struct {
}

type IdentityState struct {
}

func (IdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityState)(nil)).Elem()
}

type identityArgs struct {
	Address         *identity.Address         `pulumi:"address"`
	Category        string                    `pulumi:"category"`
	Fields          []Field                   `pulumi:"fields"`
	Identification  *identity.Identification  `pulumi:"identification"`
	InternetDetails *identity.InternetDetails `pulumi:"internetDetails"`
	Notes           *string                   `pulumi:"notes"`
	Sections        []Section                 `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a Identity resource.
type IdentityArgs struct {
	Address         identity.AddressPtrInput
	Category        pulumi.StringInput
	Fields          FieldArrayInput
	Identification  identity.IdentificationPtrInput
	InternetDetails identity.InternetDetailsPtrInput
	Notes           pulumi.StringPtrInput
	Sections        SectionArrayInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityArgs)(nil)).Elem()
}

type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(ctx context.Context) IdentityOutput
}

func (*Identity) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *Identity) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i *Identity) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

// IdentityArrayInput is an input type that accepts IdentityArray and IdentityArrayOutput values.
// You can construct a concrete instance of `IdentityArrayInput` via:
//
//	IdentityArray{ IdentityArgs{...} }
type IdentityArrayInput interface {
	pulumi.Input

	ToIdentityArrayOutput() IdentityArrayOutput
	ToIdentityArrayOutputWithContext(context.Context) IdentityArrayOutput
}

type IdentityArray []IdentityInput

func (IdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Identity)(nil)).Elem()
}

func (i IdentityArray) ToIdentityArrayOutput() IdentityArrayOutput {
	return i.ToIdentityArrayOutputWithContext(context.Background())
}

func (i IdentityArray) ToIdentityArrayOutputWithContext(ctx context.Context) IdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityArrayOutput)
}

// IdentityMapInput is an input type that accepts IdentityMap and IdentityMapOutput values.
// You can construct a concrete instance of `IdentityMapInput` via:
//
//	IdentityMap{ "key": IdentityArgs{...} }
type IdentityMapInput interface {
	pulumi.Input

	ToIdentityMapOutput() IdentityMapOutput
	ToIdentityMapOutputWithContext(context.Context) IdentityMapOutput
}

type IdentityMap map[string]IdentityInput

func (IdentityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Identity)(nil)).Elem()
}

func (i IdentityMap) ToIdentityMapOutput() IdentityMapOutput {
	return i.ToIdentityMapOutputWithContext(context.Background())
}

func (i IdentityMap) ToIdentityMapOutputWithContext(ctx context.Context) IdentityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityMapOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

type IdentityArrayOutput struct{ *pulumi.OutputState }

func (IdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Identity)(nil)).Elem()
}

func (o IdentityArrayOutput) ToIdentityArrayOutput() IdentityArrayOutput {
	return o
}

func (o IdentityArrayOutput) ToIdentityArrayOutputWithContext(ctx context.Context) IdentityArrayOutput {
	return o
}

func (o IdentityArrayOutput) Index(i pulumi.IntInput) IdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Identity {
		return vs[0].([]*Identity)[vs[1].(int)]
	}).(IdentityOutput)
}

type IdentityMapOutput struct{ *pulumi.OutputState }

func (IdentityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Identity)(nil)).Elem()
}

func (o IdentityMapOutput) ToIdentityMapOutput() IdentityMapOutput {
	return o
}

func (o IdentityMapOutput) ToIdentityMapOutputWithContext(ctx context.Context) IdentityMapOutput {
	return o
}

func (o IdentityMapOutput) MapIndex(k pulumi.StringInput) IdentityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Identity {
		return vs[0].(map[string]*Identity)[vs[1].(string)]
	}).(IdentityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), &Identity{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityArrayInput)(nil)).Elem(), IdentityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityMapInput)(nil)).Elem(), IdentityMap{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityArrayOutput{})
	pulumi.RegisterOutputType(IdentityMapOutput{})
}
