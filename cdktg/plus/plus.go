package plus

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hupe1980/cdk-threagile-go/cdktg/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hupe1980/cdk-threagile-go/cdktg"
	"github.com/hupe1980/cdk-threagile-go/cdktg/plus/internal"
)

type Browser interface {
	cdktg.TechnicalAsset
	CiaTriad() cdktg.CIATriad
	CustomDevelopedParts() *bool
	DataFormatsAccepted() *[]cdktg.DataFormat
	Description() *string
	Encryption() cdktg.Encryption
	HumanUse() *bool
	Internet() *bool
	Machine() cdktg.Machine
	MultiTenant() *bool
	// The tree node.
	Node() constructs.Node
	Owner() *string
	Redundant() *bool
	Scope() cdktg.Scope
	Size() cdktg.Size
	Tags() *[]*string
	Technology() cdktg.Technology
	TrustBoundary() cdktg.TrustBoundary
	SetTrustBoundary(val cdktg.TrustBoundary)
	Type() cdktg.TechnicalAssetType
	Usage() cdktg.Usage
	Uuid() *string
	CommunicatesWith(id *string, target cdktg.TechnicalAsset, options *cdktg.CommunicationOptions) cdktg.Communication
	IsTrafficForwarding() *bool
	IsWebApplication() *bool
	IsWebService() *bool
	Processes(assets ...cdktg.DataAsset)
	Stores(assets ...cdktg.DataAsset)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Browser
type jsiiProxy_Browser struct {
	internal.Type__cdktgTechnicalAsset
}

func (j *jsiiProxy_Browser) CiaTriad() cdktg.CIATriad {
	var returns cdktg.CIATriad
	_jsii_.Get(
		j,
		"ciaTriad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) CustomDevelopedParts() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"customDevelopedParts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) DataFormatsAccepted() *[]cdktg.DataFormat {
	var returns *[]cdktg.DataFormat
	_jsii_.Get(
		j,
		"dataFormatsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Encryption() cdktg.Encryption {
	var returns cdktg.Encryption
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) HumanUse() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"humanUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Internet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"internet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Machine() cdktg.Machine {
	var returns cdktg.Machine
	_jsii_.Get(
		j,
		"machine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) MultiTenant() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"multiTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Redundant() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"redundant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Scope() cdktg.Scope {
	var returns cdktg.Scope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Size() cdktg.Size {
	var returns cdktg.Size
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Technology() cdktg.Technology {
	var returns cdktg.Technology
	_jsii_.Get(
		j,
		"technology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) TrustBoundary() cdktg.TrustBoundary {
	var returns cdktg.TrustBoundary
	_jsii_.Get(
		j,
		"trustBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Type() cdktg.TechnicalAssetType {
	var returns cdktg.TechnicalAssetType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Usage() cdktg.Usage {
	var returns cdktg.Usage
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Browser) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


func NewBrowser(scope constructs.Construct, id *string, props *BrowserProps) Browser {
	_init_.Initialize()

	j := jsiiProxy_Browser{}

	_jsii_.Create(
		"cdktg.plus.Browser",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBrowser_Override(b Browser, scope constructs.Construct, id *string, props *BrowserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.plus.Browser",
		[]interface{}{scope, id, props},
		b,
	)
}

func (j *jsiiProxy_Browser) SetTrustBoundary(val cdktg.TrustBoundary) {
	_jsii_.Set(
		j,
		"trustBoundary",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Browser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.plus.Browser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Browser) CommunicatesWith(id *string, target cdktg.TechnicalAsset, options *cdktg.CommunicationOptions) cdktg.Communication {
	var returns cdktg.Communication

	_jsii_.Invoke(
		b,
		"communicatesWith",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Browser) IsTrafficForwarding() *bool {
	var returns *bool

	_jsii_.Invoke(
		b,
		"isTrafficForwarding",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Browser) IsWebApplication() *bool {
	var returns *bool

	_jsii_.Invoke(
		b,
		"isWebApplication",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Browser) IsWebService() *bool {
	var returns *bool

	_jsii_.Invoke(
		b,
		"isWebService",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Browser) Processes(assets ...cdktg.DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"processes",
		args,
	)
}

func (b *jsiiProxy_Browser) Stores(assets ...cdktg.DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"stores",
		args,
	)
}

func (b *jsiiProxy_Browser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BrowserProps struct {
	CiaTriad cdktg.CIATriad `field:"required" json:"ciaTriad" yaml:"ciaTriad"`
	Scope cdktg.Scope `field:"required" json:"scope" yaml:"scope"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
}

type StorageType string

const (
	// Cloud Provider (storage buckets or similar).
	StorageType_CLOUD_PROVIDER StorageType = "CLOUD_PROVIDER"
	// Container Platform (orchestration platform managed storage).
	StorageType_CONTAINER_PLATFORM StorageType = "CONTAINER_PLATFORM"
	// Database (SQL-DB, NoSQL-DB, object store or similar).
	StorageType_DATABASE StorageType = "DATABASE"
	// Filesystem (local or remote).
	StorageType_FILESYSTEM StorageType = "FILESYSTEM"
	// In-Memory (no persistent storage of secrets).
	StorageType_IN_MEMORY StorageType = "IN_MEMORY"
	// Service Registry.
	StorageType_SERVICE_REGISTRY StorageType = "SERVICE_REGISTRY"
)

type Vault interface {
	cdktg.TechnicalAsset
	CiaTriad() cdktg.CIATriad
	ConfigurationSecrets() cdktg.DataAsset
	CustomDevelopedParts() *bool
	DataFormatsAccepted() *[]cdktg.DataFormat
	Description() *string
	Encryption() cdktg.Encryption
	HumanUse() *bool
	Internet() *bool
	Machine() cdktg.Machine
	MultiTenant() *bool
	// The tree node.
	Node() constructs.Node
	Owner() *string
	Redundant() *bool
	Scope() cdktg.Scope
	Size() cdktg.Size
	Tags() *[]*string
	Technology() cdktg.Technology
	TrustBoundary() cdktg.TrustBoundary
	SetTrustBoundary(val cdktg.TrustBoundary)
	Type() cdktg.TechnicalAssetType
	Usage() cdktg.Usage
	Uuid() *string
	VaultStorage() cdktg.TechnicalAsset
	CommunicatesWith(id *string, target cdktg.TechnicalAsset, options *cdktg.CommunicationOptions) cdktg.Communication
	IsTrafficForwarding() *bool
	IsUsedBy(client cdktg.TechnicalAsset)
	IsWebApplication() *bool
	IsWebService() *bool
	Processes(assets ...cdktg.DataAsset)
	Stores(assets ...cdktg.DataAsset)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Vault
type jsiiProxy_Vault struct {
	internal.Type__cdktgTechnicalAsset
}

func (j *jsiiProxy_Vault) CiaTriad() cdktg.CIATriad {
	var returns cdktg.CIATriad
	_jsii_.Get(
		j,
		"ciaTriad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) ConfigurationSecrets() cdktg.DataAsset {
	var returns cdktg.DataAsset
	_jsii_.Get(
		j,
		"configurationSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) CustomDevelopedParts() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"customDevelopedParts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) DataFormatsAccepted() *[]cdktg.DataFormat {
	var returns *[]cdktg.DataFormat
	_jsii_.Get(
		j,
		"dataFormatsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Encryption() cdktg.Encryption {
	var returns cdktg.Encryption
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) HumanUse() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"humanUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Internet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"internet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Machine() cdktg.Machine {
	var returns cdktg.Machine
	_jsii_.Get(
		j,
		"machine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) MultiTenant() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"multiTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Redundant() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"redundant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Scope() cdktg.Scope {
	var returns cdktg.Scope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Size() cdktg.Size {
	var returns cdktg.Size
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Technology() cdktg.Technology {
	var returns cdktg.Technology
	_jsii_.Get(
		j,
		"technology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) TrustBoundary() cdktg.TrustBoundary {
	var returns cdktg.TrustBoundary
	_jsii_.Get(
		j,
		"trustBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Type() cdktg.TechnicalAssetType {
	var returns cdktg.TechnicalAssetType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Usage() cdktg.Usage {
	var returns cdktg.Usage
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Vault) VaultStorage() cdktg.TechnicalAsset {
	var returns cdktg.TechnicalAsset
	_jsii_.Get(
		j,
		"vaultStorage",
		&returns,
	)
	return returns
}


func NewVault(scope constructs.Construct, id *string, props *VaultProps) Vault {
	_init_.Initialize()

	j := jsiiProxy_Vault{}

	_jsii_.Create(
		"cdktg.plus.Vault",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVault_Override(v Vault, scope constructs.Construct, id *string, props *VaultProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.plus.Vault",
		[]interface{}{scope, id, props},
		v,
	)
}

func (j *jsiiProxy_Vault) SetTrustBoundary(val cdktg.TrustBoundary) {
	_jsii_.Set(
		j,
		"trustBoundary",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Vault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.plus.Vault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Vault) CommunicatesWith(id *string, target cdktg.TechnicalAsset, options *cdktg.CommunicationOptions) cdktg.Communication {
	var returns cdktg.Communication

	_jsii_.Invoke(
		v,
		"communicatesWith",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Vault) IsTrafficForwarding() *bool {
	var returns *bool

	_jsii_.Invoke(
		v,
		"isTrafficForwarding",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Vault) IsUsedBy(client cdktg.TechnicalAsset) {
	_jsii_.InvokeVoid(
		v,
		"isUsedBy",
		[]interface{}{client},
	)
}

func (v *jsiiProxy_Vault) IsWebApplication() *bool {
	var returns *bool

	_jsii_.Invoke(
		v,
		"isWebApplication",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Vault) IsWebService() *bool {
	var returns *bool

	_jsii_.Invoke(
		v,
		"isWebService",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Vault) Processes(assets ...cdktg.DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"processes",
		args,
	)
}

func (v *jsiiProxy_Vault) Stores(assets ...cdktg.DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		v,
		"stores",
		args,
	)
}

func (v *jsiiProxy_Vault) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VaultProps struct {
	Authtentication cdktg.Authentication `field:"required" json:"authtentication" yaml:"authtentication"`
	MultiTenant *bool `field:"required" json:"multiTenant" yaml:"multiTenant"`
	StorageType StorageType `field:"required" json:"storageType" yaml:"storageType"`
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	TrustBoundary cdktg.TrustBoundary `field:"optional" json:"trustBoundary" yaml:"trustBoundary"`
	Vendor *string `field:"optional" json:"vendor" yaml:"vendor"`
}

