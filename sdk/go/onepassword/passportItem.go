// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PassportItem struct {
	pulumi.CustomResourceState

	Category         pulumi.StringOutput    `pulumi:"category"`
	DateOfBirth      pulumi.StringPtrOutput `pulumi:"dateOfBirth"`
	ExpiryDate       pulumi.StringPtrOutput `pulumi:"expiryDate"`
	Fields           GetFieldMapOutput      `pulumi:"fields"`
	FullName         pulumi.StringPtrOutput `pulumi:"fullName"`
	Gender           pulumi.StringPtrOutput `pulumi:"gender"`
	Id               pulumi.StringOutput    `pulumi:"id"`
	IssuedOn         pulumi.StringPtrOutput `pulumi:"issuedOn"`
	IssuingAuthority pulumi.StringPtrOutput `pulumi:"issuingAuthority"`
	IssuingCountry   pulumi.StringPtrOutput `pulumi:"issuingCountry"`
	Nationality      pulumi.StringPtrOutput `pulumi:"nationality"`
	Notes            pulumi.StringPtrOutput `pulumi:"notes"`
	Number           pulumi.StringPtrOutput `pulumi:"number"`
	PlaceOfBirth     pulumi.StringPtrOutput `pulumi:"placeOfBirth"`
	Sections         GetSectionMapOutput    `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput    `pulumi:"title"`
	Type  pulumi.StringPtrOutput `pulumi:"type"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewPassportItem registers a new resource with the given unique name, arguments, and options.
func NewPassportItem(ctx *pulumi.Context,
	name string, args *PassportItemArgs, opts ...pulumi.ResourceOption) (*PassportItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	var resource PassportItem
	err := ctx.RegisterResource("onepassword:index:PassportItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPassportItem gets an existing PassportItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPassportItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PassportItemState, opts ...pulumi.ResourceOption) (*PassportItem, error) {
	var resource PassportItem
	err := ctx.ReadResource("onepassword:index:PassportItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PassportItem resources.
type passportItemState struct {
}

type PassportItemState struct {
}

func (PassportItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*passportItemState)(nil)).Elem()
}

type passportItemArgs struct {
	DateOfBirth      *string            `pulumi:"dateOfBirth"`
	ExpiryDate       *string            `pulumi:"expiryDate"`
	Fields           map[string]Field   `pulumi:"fields"`
	FullName         *string            `pulumi:"fullName"`
	Gender           *string            `pulumi:"gender"`
	IssuedOn         *string            `pulumi:"issuedOn"`
	IssuingAuthority *string            `pulumi:"issuingAuthority"`
	IssuingCountry   *string            `pulumi:"issuingCountry"`
	Nationality      *string            `pulumi:"nationality"`
	Notes            *string            `pulumi:"notes"`
	Number           *string            `pulumi:"number"`
	PlaceOfBirth     *string            `pulumi:"placeOfBirth"`
	Sections         map[string]Section `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	Type  *string `pulumi:"type"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a PassportItem resource.
type PassportItemArgs struct {
	DateOfBirth      pulumi.StringPtrInput
	ExpiryDate       pulumi.StringPtrInput
	Fields           FieldMapInput
	FullName         pulumi.StringPtrInput
	Gender           pulumi.StringPtrInput
	IssuedOn         pulumi.StringPtrInput
	IssuingAuthority pulumi.StringPtrInput
	IssuingCountry   pulumi.StringPtrInput
	Nationality      pulumi.StringPtrInput
	Notes            pulumi.StringPtrInput
	Number           pulumi.StringPtrInput
	PlaceOfBirth     pulumi.StringPtrInput
	Sections         SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	Type  pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (PassportItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passportItemArgs)(nil)).Elem()
}

type PassportItemInput interface {
	pulumi.Input

	ToPassportItemOutput() PassportItemOutput
	ToPassportItemOutputWithContext(ctx context.Context) PassportItemOutput
}

func (*PassportItem) ElementType() reflect.Type {
	return reflect.TypeOf((**PassportItem)(nil)).Elem()
}

func (i *PassportItem) ToPassportItemOutput() PassportItemOutput {
	return i.ToPassportItemOutputWithContext(context.Background())
}

func (i *PassportItem) ToPassportItemOutputWithContext(ctx context.Context) PassportItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PassportItemOutput)
}

// PassportItemArrayInput is an input type that accepts PassportItemArray and PassportItemArrayOutput values.
// You can construct a concrete instance of `PassportItemArrayInput` via:
//
//	PassportItemArray{ PassportItemArgs{...} }
type PassportItemArrayInput interface {
	pulumi.Input

	ToPassportItemArrayOutput() PassportItemArrayOutput
	ToPassportItemArrayOutputWithContext(context.Context) PassportItemArrayOutput
}

type PassportItemArray []PassportItemInput

func (PassportItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PassportItem)(nil)).Elem()
}

func (i PassportItemArray) ToPassportItemArrayOutput() PassportItemArrayOutput {
	return i.ToPassportItemArrayOutputWithContext(context.Background())
}

func (i PassportItemArray) ToPassportItemArrayOutputWithContext(ctx context.Context) PassportItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PassportItemArrayOutput)
}

// PassportItemMapInput is an input type that accepts PassportItemMap and PassportItemMapOutput values.
// You can construct a concrete instance of `PassportItemMapInput` via:
//
//	PassportItemMap{ "key": PassportItemArgs{...} }
type PassportItemMapInput interface {
	pulumi.Input

	ToPassportItemMapOutput() PassportItemMapOutput
	ToPassportItemMapOutputWithContext(context.Context) PassportItemMapOutput
}

type PassportItemMap map[string]PassportItemInput

func (PassportItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PassportItem)(nil)).Elem()
}

func (i PassportItemMap) ToPassportItemMapOutput() PassportItemMapOutput {
	return i.ToPassportItemMapOutputWithContext(context.Background())
}

func (i PassportItemMap) ToPassportItemMapOutputWithContext(ctx context.Context) PassportItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PassportItemMapOutput)
}

type PassportItemOutput struct{ *pulumi.OutputState }

func (PassportItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PassportItem)(nil)).Elem()
}

func (o PassportItemOutput) ToPassportItemOutput() PassportItemOutput {
	return o
}

func (o PassportItemOutput) ToPassportItemOutputWithContext(ctx context.Context) PassportItemOutput {
	return o
}

type PassportItemArrayOutput struct{ *pulumi.OutputState }

func (PassportItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PassportItem)(nil)).Elem()
}

func (o PassportItemArrayOutput) ToPassportItemArrayOutput() PassportItemArrayOutput {
	return o
}

func (o PassportItemArrayOutput) ToPassportItemArrayOutputWithContext(ctx context.Context) PassportItemArrayOutput {
	return o
}

func (o PassportItemArrayOutput) Index(i pulumi.IntInput) PassportItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PassportItem {
		return vs[0].([]*PassportItem)[vs[1].(int)]
	}).(PassportItemOutput)
}

type PassportItemMapOutput struct{ *pulumi.OutputState }

func (PassportItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PassportItem)(nil)).Elem()
}

func (o PassportItemMapOutput) ToPassportItemMapOutput() PassportItemMapOutput {
	return o
}

func (o PassportItemMapOutput) ToPassportItemMapOutputWithContext(ctx context.Context) PassportItemMapOutput {
	return o
}

func (o PassportItemMapOutput) MapIndex(k pulumi.StringInput) PassportItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PassportItem {
		return vs[0].(map[string]*PassportItem)[vs[1].(string)]
	}).(PassportItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PassportItemInput)(nil)).Elem(), &PassportItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*PassportItemArrayInput)(nil)).Elem(), PassportItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PassportItemMapInput)(nil)).Elem(), PassportItemMap{})
	pulumi.RegisterOutputType(PassportItemOutput{})
	pulumi.RegisterOutputType(PassportItemArrayOutput{})
	pulumi.RegisterOutputType(PassportItemMapOutput{})
}
