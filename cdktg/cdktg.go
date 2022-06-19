// Agile Threat Modeling as Code
package cdktg

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hupe1980/cdk-threagile-go/cdktg/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hupe1980/cdk-threagile-go/cdktg/internal"
)

type AnnotationMetadataEntryType string

const (
	AnnotationMetadataEntryType_INFO AnnotationMetadataEntryType = "INFO"
	AnnotationMetadataEntryType_WARN AnnotationMetadataEntryType = "WARN"
	AnnotationMetadataEntryType_ERROR AnnotationMetadataEntryType = "ERROR"
)

// Includes API for attaching annotations such as warning messages to constructs.
type Annotations interface {
	// Adds an { "error": <message> } metadata entry to this construct.
	//
	// The toolkit will fail synthesis when errors are reported.
	AddError(message *string)
	// Adds an info metadata entry to this construct.
	//
	// The CLI will display the info message when apps are synthesized.
	AddInfo(message *string)
	// Adds a warning metadata entry to this construct.
	//
	// The CLI will display the warning when an app is synthesized.
	// In a future release the CLI might introduce a --strict flag which
	// will then fail the synthesis if it encounters a warning.
	AddWarning(message *string)
}

// The jsii proxy struct for Annotations
type jsiiProxy_Annotations struct {
	_ byte // padding
}

// Returns the annotations API for a construct scope.
func Annotations_Of(scope constructs.IConstruct) Annotations {
	_init_.Initialize()

	var returns Annotations

	_jsii_.StaticInvoke(
		"cdktg.Annotations",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Annotations) AddError(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addError",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_Annotations) AddInfo(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addInfo",
		[]interface{}{message},
	)
}

func (a *jsiiProxy_Annotations) AddWarning(message *string) {
	_jsii_.InvokeVoid(
		a,
		"addWarning",
		[]interface{}{message},
	)
}

// Aspects can be applied to CDK tree scopes and can operate on the tree before synthesis.
type Aspects interface {
	// The list of aspects which were directly applied on this scope.
	All() *[]IAspect
	// Adds an aspect to apply this scope before synthesis.
	Add(aspect IAspect)
}

// The jsii proxy struct for Aspects
type jsiiProxy_Aspects struct {
	_ byte // padding
}

func (j *jsiiProxy_Aspects) All() *[]IAspect {
	var returns *[]IAspect
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}


// Returns the `Aspects` object associated with a construct scope.
func Aspects_Of(scope constructs.IConstruct) Aspects {
	_init_.Initialize()

	var returns Aspects

	_jsii_.StaticInvoke(
		"cdktg.Aspects",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Aspects) Add(aspect IAspect) {
	_jsii_.InvokeVoid(
		a,
		"add",
		[]interface{}{aspect},
	)
}

type Asset interface {
	constructs.Construct
	CiaTriad() CIATriad
	Description() *string
	// The tree node.
	Node() constructs.Node
	Usage() Usage
	Uuid() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Asset
type jsiiProxy_Asset struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Asset) CiaTriad() CIATriad {
	var returns CIATriad
	_jsii_.Get(
		j,
		"ciaTriad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) Usage() Usage {
	var returns Usage
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


func NewAsset_Override(a Asset, model constructs.Construct, id *string, props *AssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.Asset",
		[]interface{}{model, id, props},
		a,
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
func Asset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.Asset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Asset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AssetProps struct {
	CiaTriad CIATriad `field:"required" json:"ciaTriad" yaml:"ciaTriad"`
	Description *string `field:"required" json:"description" yaml:"description"`
	Usage Usage `field:"required" json:"usage" yaml:"usage"`
}

type AssetType string

const (
	AssetType_EXTERNAL_ENTITY AssetType = "EXTERNAL_ENTITY"
	AssetType_PROCESS AssetType = "PROCESS"
	AssetType_DATASTORE AssetType = "DATASTORE"
)

type Authentication string

const (
	Authentication_NONE Authentication = "NONE"
	Authentication_CREDENTIALS Authentication = "CREDENTIALS"
	Authentication_SESSION_ID Authentication = "SESSION_ID"
	Authentication_TOKEN Authentication = "TOKEN"
	Authentication_CLIENT_CERTIFICATE Authentication = "CLIENT_CERTIFICATE"
	Authentication_TWO_FACTOR Authentication = "TWO_FACTOR"
	Authentication_EXTERNALIZED Authentication = "EXTERNALIZED"
)

type Author interface {
	Homepage() *string
	Name() *string
}

// The jsii proxy struct for Author
type jsiiProxy_Author struct {
	_ byte // padding
}

func (j *jsiiProxy_Author) Homepage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homepage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Author) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewAuthor(props *AuthorProps) Author {
	_init_.Initialize()

	j := jsiiProxy_Author{}

	_jsii_.Create(
		"cdktg.Author",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAuthor_Override(a Author, props *AuthorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.Author",
		[]interface{}{props},
		a,
	)
}

type AuthorProps struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Homepage *string `field:"optional" json:"homepage" yaml:"homepage"`
}

type Authorization string

const (
	Authorization_NONE Authorization = "NONE"
	Authorization_TECHNICAL_USER Authorization = "TECHNICAL_USER"
	Authorization_ENDUSER_IDENTITY_PROPAGATION Authorization = "ENDUSER_IDENTITY_PROPAGATION"
)

type Availability string

const (
	Availability_ARCHIVE Availability = "ARCHIVE"
	Availability_OPERATIONAL Availability = "OPERATIONAL"
	Availability_IMPORTANT Availability = "IMPORTANT"
	Availability_CRITICAL Availability = "CRITICAL"
	Availability_MISSION_CRITICAL Availability = "MISSION_CRITICAL"
)

type BusinessCriticality string

const (
	BusinessCriticality_ARCHIVE BusinessCriticality = "ARCHIVE"
	BusinessCriticality_OPERATIONAL BusinessCriticality = "OPERATIONAL"
	BusinessCriticality_IMPORTANT BusinessCriticality = "IMPORTANT"
	BusinessCriticality_CRITICAL BusinessCriticality = "CRITICAL"
	BusinessCriticality_MISSION_CRITICAL BusinessCriticality = "MISSION_CRITICAL"
)

type CIATriad interface {
	Availability() Availability
	Confidentiality() Confidentiality
	Integrity() Integrity
	Justification() *string
}

// The jsii proxy struct for CIATriad
type jsiiProxy_CIATriad struct {
	_ byte // padding
}

func (j *jsiiProxy_CIATriad) Availability() Availability {
	var returns Availability
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CIATriad) Confidentiality() Confidentiality {
	var returns Confidentiality
	_jsii_.Get(
		j,
		"confidentiality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CIATriad) Integrity() Integrity {
	var returns Integrity
	_jsii_.Get(
		j,
		"integrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CIATriad) Justification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"justification",
		&returns,
	)
	return returns
}


func NewCIATriad(props *CIATriadProps) CIATriad {
	_init_.Initialize()

	j := jsiiProxy_CIATriad{}

	_jsii_.Create(
		"cdktg.CIATriad",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCIATriad_Override(c CIATriad, props *CIATriadProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.CIATriad",
		[]interface{}{props},
		c,
	)
}

type CIATriadProps struct {
	Availability Availability `field:"required" json:"availability" yaml:"availability"`
	Confidentiality Confidentiality `field:"required" json:"confidentiality" yaml:"confidentiality"`
	Integrity Integrity `field:"required" json:"integrity" yaml:"integrity"`
	Justification *string `field:"optional" json:"justification" yaml:"justification"`
}

type Communication interface {
	Authentication() Authentication
	Authorization() Authorization
	Description() *string
	Id() *string
	IpFiltered() *bool
	Protocol() Protocol
	Readonly() *bool
	Target() TechnicalAsset
	Usage() Usage
	Vpn() *bool
	Receive(assets ...DataAsset)
	Send(assets ...DataAsset)
}

// The jsii proxy struct for Communication
type jsiiProxy_Communication struct {
	_ byte // padding
}

func (j *jsiiProxy_Communication) Authentication() Authentication {
	var returns Authentication
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) Authorization() Authorization {
	var returns Authorization
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) IpFiltered() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ipFiltered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) Protocol() Protocol {
	var returns Protocol
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) Readonly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"readonly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) Target() TechnicalAsset {
	var returns TechnicalAsset
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) Usage() Usage {
	var returns Usage
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Communication) Vpn() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"vpn",
		&returns,
	)
	return returns
}


func NewCommunication(id *string, props *CommunicationProps) Communication {
	_init_.Initialize()

	j := jsiiProxy_Communication{}

	_jsii_.Create(
		"cdktg.Communication",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewCommunication_Override(c Communication, id *string, props *CommunicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.Communication",
		[]interface{}{id, props},
		c,
	)
}

func (c *jsiiProxy_Communication) Receive(assets ...DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"receive",
		args,
	)
}

func (c *jsiiProxy_Communication) Send(assets ...DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"send",
		args,
	)
}

type CommunicationOptions struct {
	Authentication Authentication `field:"required" json:"authentication" yaml:"authentication"`
	Authorization Authorization `field:"required" json:"authorization" yaml:"authorization"`
	Description *string `field:"required" json:"description" yaml:"description"`
	IpFiltered *bool `field:"required" json:"ipFiltered" yaml:"ipFiltered"`
	Protocol Protocol `field:"required" json:"protocol" yaml:"protocol"`
	Readonly *bool `field:"required" json:"readonly" yaml:"readonly"`
	Usage Usage `field:"required" json:"usage" yaml:"usage"`
	Vpn *bool `field:"required" json:"vpn" yaml:"vpn"`
}

type CommunicationProps struct {
	Authentication Authentication `field:"required" json:"authentication" yaml:"authentication"`
	Authorization Authorization `field:"required" json:"authorization" yaml:"authorization"`
	Description *string `field:"required" json:"description" yaml:"description"`
	IpFiltered *bool `field:"required" json:"ipFiltered" yaml:"ipFiltered"`
	Protocol Protocol `field:"required" json:"protocol" yaml:"protocol"`
	Readonly *bool `field:"required" json:"readonly" yaml:"readonly"`
	Usage Usage `field:"required" json:"usage" yaml:"usage"`
	Vpn *bool `field:"required" json:"vpn" yaml:"vpn"`
	Target TechnicalAsset `field:"required" json:"target" yaml:"target"`
}

type Confidentiality string

const (
	Confidentiality_PUBLIC Confidentiality = "PUBLIC"
	Confidentiality_INTERNAL Confidentiality = "INTERNAL"
	Confidentiality_RESTRICTED Confidentiality = "RESTRICTED"
	Confidentiality_CONFIDENTIAL Confidentiality = "CONFIDENTIAL"
	Confidentiality_STRICTLY_CONFIDENTIAL Confidentiality = "STRICTLY_CONFIDENTIAL"
)

type DataAsset interface {
	Asset
	CiaTriad() CIATriad
	Description() *string
	// The tree node.
	Node() constructs.Node
	Origin() *string
	Owner() *string
	Quantity() Quantity
	Usage() Usage
	Uuid() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DataAsset
type jsiiProxy_DataAsset struct {
	jsiiProxy_Asset
}

func (j *jsiiProxy_DataAsset) CiaTriad() CIATriad {
	var returns CIATriad
	_jsii_.Get(
		j,
		"ciaTriad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAsset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAsset) Origin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAsset) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAsset) Quantity() Quantity {
	var returns Quantity
	_jsii_.Get(
		j,
		"quantity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAsset) Usage() Usage {
	var returns Usage
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAsset) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


func NewDataAsset(model constructs.Construct, id *string, props *DataAssetProps) DataAsset {
	_init_.Initialize()

	j := jsiiProxy_DataAsset{}

	_jsii_.Create(
		"cdktg.DataAsset",
		[]interface{}{model, id, props},
		&j,
	)

	return &j
}

func NewDataAsset_Override(d DataAsset, model constructs.Construct, id *string, props *DataAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.DataAsset",
		[]interface{}{model, id, props},
		d,
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
func DataAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.DataAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAssetProps struct {
	CiaTriad CIATriad `field:"required" json:"ciaTriad" yaml:"ciaTriad"`
	Description *string `field:"required" json:"description" yaml:"description"`
	Usage Usage `field:"required" json:"usage" yaml:"usage"`
	Quantity Quantity `field:"required" json:"quantity" yaml:"quantity"`
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
}

type DataFormat string

const (
	DataFormat_JSON DataFormat = "JSON"
	DataFormat_XML DataFormat = "XML"
	DataFormat_SERIALIZATION DataFormat = "SERIALIZATION"
	DataFormat_FILE DataFormat = "FILE"
	DataFormat_CSV DataFormat = "CSV"
)

type Encryption string

const (
	Encryption_NONE Encryption = "NONE"
	Encryption_TRANSPARENT Encryption = "TRANSPARENT"
	Encryption_SYMMETRIC_SHARED_KEY Encryption = "SYMMETRIC_SHARED_KEY"
	Encryption_ASYMMETRIC_SHARED_KEY Encryption = "ASYMMETRIC_SHARED_KEY"
	Encryption_ENDUSER_INDIVIDUAL_KEY Encryption = "ENDUSER_INDIVIDUAL_KEY"
)

// Represents an Aspect.
type IAspect interface {
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
}

// The jsii proxy for IAspect
type jsiiProxy_IAspect struct {
	_ byte // padding
}

func (i *jsiiProxy_IAspect) Visit(node constructs.IConstruct) {
	_jsii_.InvokeVoid(
		i,
		"visit",
		[]interface{}{node},
	)
}

type IManifest interface {
	Models() *map[string]*ModelManifest
	Version() *string
}

// The jsii proxy for IManifest
type jsiiProxy_IManifest struct {
	_ byte // padding
}

func (j *jsiiProxy_IManifest) Models() *map[string]*ModelManifest {
	var returns *map[string]*ModelManifest
	_jsii_.Get(
		j,
		"models",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManifest) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

type IModelSynthesizer interface {
	// Synthesize the associated model to the session.
	Synthesize(session ISynthesisSession)
}

// The jsii proxy for IModelSynthesizer
type jsiiProxy_IModelSynthesizer struct {
	_ byte // padding
}

func (i *jsiiProxy_IModelSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		i,
		"synthesize",
		[]interface{}{session},
	)
}

type ISynthesisSession interface {
	Manifest() Manifest
	// The output directory for this synthesis session.
	Outdir() *string
	SkipValidation() *bool
}

// The jsii proxy for ISynthesisSession
type jsiiProxy_ISynthesisSession struct {
	_ byte // padding
}

func (j *jsiiProxy_ISynthesisSession) Manifest() Manifest {
	var returns Manifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISynthesisSession) SkipValidation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipValidation",
		&returns,
	)
	return returns
}

type InScope interface {
	Scope
	IsInScope() *bool
	Justification() *string
}

// The jsii proxy struct for InScope
type jsiiProxy_InScope struct {
	jsiiProxy_Scope
}

func (j *jsiiProxy_InScope) IsInScope() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InScope) Justification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"justification",
		&returns,
	)
	return returns
}


func NewInScope(justification *string) InScope {
	_init_.Initialize()

	j := jsiiProxy_InScope{}

	_jsii_.Create(
		"cdktg.InScope",
		[]interface{}{justification},
		&j,
	)

	return &j
}

func NewInScope_Override(i InScope, justification *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.InScope",
		[]interface{}{justification},
		i,
	)
}

type Integrity string

const (
	Integrity_ARCHIVE Integrity = "ARCHIVE"
	Integrity_OPERATIONAL Integrity = "OPERATIONAL"
	Integrity_IMPORTANT Integrity = "IMPORTANT"
	Integrity_CRITICAL Integrity = "CRITICAL"
	Integrity_MISSION_CRITICAL Integrity = "MISSION_CRITICAL"
)

type Machine string

const (
	Machine_PHYSICAL Machine = "PHYSICAL"
	Machine_VIRTUAL Machine = "VIRTUAL"
	Machine_CONTAINER Machine = "CONTAINER"
	Machine_SERVERLESS Machine = "SERVERLESS"
)

type Manifest interface {
	IManifest
	Models() *map[string]*ModelManifest
	Outdir() *string
	Version() *string
	BuildManifest() IManifest
	ForModel(model Model) *ModelManifest
	WriteToFile()
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	jsiiProxy_IManifest
}

func (j *jsiiProxy_Manifest) Models() *map[string]*ModelManifest {
	var returns *map[string]*ModelManifest
	_jsii_.Get(
		j,
		"models",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Manifest) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Manifest) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewManifest(version *string, outdir *string) Manifest {
	_init_.Initialize()

	j := jsiiProxy_Manifest{}

	_jsii_.Create(
		"cdktg.Manifest",
		[]interface{}{version, outdir},
		&j,
	)

	return &j
}

func NewManifest_Override(m Manifest, version *string, outdir *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.Manifest",
		[]interface{}{version, outdir},
		m,
	)
}

func Manifest_FromPath(dir *string) Manifest {
	_init_.Initialize()

	var returns Manifest

	_jsii_.StaticInvoke(
		"cdktg.Manifest",
		"fromPath",
		[]interface{}{dir},
		&returns,
	)

	return returns
}

func Manifest_FileName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktg.Manifest",
		"fileName",
		&returns,
	)
	return returns
}

func Manifest_ModelsFolder() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdktg.Manifest",
		"modelsFolder",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_Manifest) BuildManifest() IManifest {
	var returns IManifest

	_jsii_.Invoke(
		m,
		"buildManifest",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Manifest) ForModel(model Model) *ModelManifest {
	var returns *ModelManifest

	_jsii_.Invoke(
		m,
		"forModel",
		[]interface{}{model},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Manifest) WriteToFile() {
	_jsii_.InvokeVoid(
		m,
		"writeToFile",
		nil, // no parameters
	)
}

type Model interface {
	constructs.Construct
	Author() Author
	BusinessCriticality() BusinessCriticality
	Date() *time.Time
	// The tree node.
	Node() constructs.Node
	Synthesizer() IModelSynthesizer
	SetSynthesizer(val IModelSynthesizer)
	Title() *string
	Version() *string
	AddOverride(path *string, value interface{})
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Model
type jsiiProxy_Model struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Model) Author() Author {
	var returns Author
	_jsii_.Get(
		j,
		"author",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) BusinessCriticality() BusinessCriticality {
	var returns BusinessCriticality
	_jsii_.Get(
		j,
		"businessCriticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) Date() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"date",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) Synthesizer() IModelSynthesizer {
	var returns IModelSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Model) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewModel(project constructs.Construct, id *string, props *ModelProps) Model {
	_init_.Initialize()

	j := jsiiProxy_Model{}

	_jsii_.Create(
		"cdktg.Model",
		[]interface{}{project, id, props},
		&j,
	)

	return &j
}

func NewModel_Override(m Model, project constructs.Construct, id *string, props *ModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.Model",
		[]interface{}{project, id, props},
		m,
	)
}

func (j *jsiiProxy_Model) SetSynthesizer(val IModelSynthesizer) {
	_jsii_.Set(
		j,
		"synthesizer",
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
func Model_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.Model",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Model_IsModel(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.Model",
		"isModel",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Model) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_Model) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ModelAnnotation struct {
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	Level AnnotationMetadataEntryType `field:"required" json:"level" yaml:"level"`
	Message *string `field:"required" json:"message" yaml:"message"`
	Stacktrace *[]*string `field:"optional" json:"stacktrace" yaml:"stacktrace"`
}

type ModelManifest struct {
	Annotations *[]*ModelAnnotation `field:"required" json:"annotations" yaml:"annotations"`
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	Name *string `field:"required" json:"name" yaml:"name"`
	SanitizedName *string `field:"required" json:"sanitizedName" yaml:"sanitizedName"`
	SynthesizedModelPath *string `field:"required" json:"synthesizedModelPath" yaml:"synthesizedModelPath"`
	WorkingDirectory *string `field:"required" json:"workingDirectory" yaml:"workingDirectory"`
}

type ModelProps struct {
	// Author of the model.
	Author Author `field:"required" json:"author" yaml:"author"`
	// Business criticality of the target.
	BusinessCriticality BusinessCriticality `field:"required" json:"businessCriticality" yaml:"businessCriticality"`
	// Version of the Threagile toolkit.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Date of the model.
	Date *time.Time `field:"optional" json:"date" yaml:"date"`
	// Title of the model.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

type ModelSynthesizer interface {
	IModelSynthesizer
	Model() Model
	SetModel(val Model)
	// Synthesize the associated model to the session.
	Synthesize(session ISynthesisSession)
}

// The jsii proxy struct for ModelSynthesizer
type jsiiProxy_ModelSynthesizer struct {
	jsiiProxy_IModelSynthesizer
}

func (j *jsiiProxy_ModelSynthesizer) Model() Model {
	var returns Model
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}


func NewModelSynthesizer(model Model, continueOnErrorAnnotations *bool) ModelSynthesizer {
	_init_.Initialize()

	j := jsiiProxy_ModelSynthesizer{}

	_jsii_.Create(
		"cdktg.ModelSynthesizer",
		[]interface{}{model, continueOnErrorAnnotations},
		&j,
	)

	return &j
}

func NewModelSynthesizer_Override(m ModelSynthesizer, model Model, continueOnErrorAnnotations *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.ModelSynthesizer",
		[]interface{}{model, continueOnErrorAnnotations},
		m,
	)
}

func (j *jsiiProxy_ModelSynthesizer) SetModel(val Model) {
	_jsii_.Set(
		j,
		"model",
		val,
	)
}

func (m *jsiiProxy_ModelSynthesizer) Synthesize(session ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		[]interface{}{session},
	)
}

type OutOfScope interface {
	Scope
	IsInScope() *bool
	Justification() *string
}

// The jsii proxy struct for OutOfScope
type jsiiProxy_OutOfScope struct {
	jsiiProxy_Scope
}

func (j *jsiiProxy_OutOfScope) IsInScope() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OutOfScope) Justification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"justification",
		&returns,
	)
	return returns
}


func NewOutOfScope(justification *string) OutOfScope {
	_init_.Initialize()

	j := jsiiProxy_OutOfScope{}

	_jsii_.Create(
		"cdktg.OutOfScope",
		[]interface{}{justification},
		&j,
	)

	return &j
}

func NewOutOfScope_Override(o OutOfScope, justification *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.OutOfScope",
		[]interface{}{justification},
		o,
	)
}

type OutOfScopeProps struct {
	OutOfScope *bool `field:"required" json:"outOfScope" yaml:"outOfScope"`
	Justification *string `field:"optional" json:"justification" yaml:"justification"`
}

type Project interface {
	constructs.Construct
	Manifest() Manifest
	// The tree node.
	Node() constructs.Node
	// The output directory into which models will be synthesized.
	Outdir() *string
	// Whether to skip the validation during synthesis of the app.
	SkipValidation() *bool
	// Synthesizes the model to the output directory.
	Synth()
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Project
type jsiiProxy_Project struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Project) Manifest() Manifest {
	var returns Manifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) SkipValidation() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipValidation",
		&returns,
	)
	return returns
}


func NewProject(props *ProjectProps) Project {
	_init_.Initialize()

	j := jsiiProxy_Project{}

	_jsii_.Create(
		"cdktg.Project",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewProject_Override(p Project, props *ProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.Project",
		[]interface{}{props},
		p,
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
func Project_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.Project",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) Synth() {
	_jsii_.InvokeVoid(
		p,
		"synth",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Project) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ProjectProps struct {
	// The directory to output the threadgile model.
	Outdir *string `field:"optional" json:"outdir" yaml:"outdir"`
	// Whether to skip the validation during synthesis of the project.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
}

type Protocol string

const (
	Protocol_UNKNOEN Protocol = "UNKNOEN"
	Protocol_HTTP Protocol = "HTTP"
	Protocol_HTTPS Protocol = "HTTPS"
	Protocol_WS Protocol = "WS"
	Protocol_WSS Protocol = "WSS"
	Protocol_REVERSE_PROXY_WEB_PROTOCOL Protocol = "REVERSE_PROXY_WEB_PROTOCOL"
	Protocol_REVERSE_PROXY_WEB_PROTOCOL_ENCRYPTED Protocol = "REVERSE_PROXY_WEB_PROTOCOL_ENCRYPTED"
	Protocol_MQTT Protocol = "MQTT"
	Protocol_JDBC Protocol = "JDBC"
	Protocol_JDBC_ENCRYPTED Protocol = "JDBC_ENCRYPTED"
	Protocol_ODBC Protocol = "ODBC"
	Protocol_ODBC_ENCRYPTED Protocol = "ODBC_ENCRYPTED"
	Protocol_SQL_ACCESS_PROTOCOL Protocol = "SQL_ACCESS_PROTOCOL"
	Protocol_SQL_ACCESS_PROTOCOL_ENCRYPTED Protocol = "SQL_ACCESS_PROTOCOL_ENCRYPTED"
	Protocol_NOSQL_ACCESS_PROTOCOL Protocol = "NOSQL_ACCESS_PROTOCOL"
	Protocol_NOSQL_ACCESS_PROTOCOL_ENCRYPTED Protocol = "NOSQL_ACCESS_PROTOCOL_ENCRYPTED"
	Protocol_BINARY Protocol = "BINARY"
	Protocol_BINARY_ENCRYPTED Protocol = "BINARY_ENCRYPTED"
	Protocol_TEXT Protocol = "TEXT"
	Protocol_TEXT_ENCRYPTED Protocol = "TEXT_ENCRYPTED"
	Protocol_SSH Protocol = "SSH"
	Protocol_SSH_TUNNEL Protocol = "SSH_TUNNEL"
	Protocol_SMTP Protocol = "SMTP"
	Protocol_SMTP_ENCRYPTED Protocol = "SMTP_ENCRYPTED"
	Protocol_POP3 Protocol = "POP3"
	Protocol_POP3_ENCRYPTED Protocol = "POP3_ENCRYPTED"
	Protocol_IMAP Protocol = "IMAP"
	Protocol_IMAP_ENCRYPTED Protocol = "IMAP_ENCRYPTED"
	Protocol_FTP Protocol = "FTP"
	Protocol_FTPS Protocol = "FTPS"
	Protocol_SFTP Protocol = "SFTP"
	Protocol_SCP Protocol = "SCP"
	Protocol_LDAP Protocol = "LDAP"
	Protocol_LDAPS Protocol = "LDAPS"
	Protocol_JMS Protocol = "JMS"
	Protocol_NFS Protocol = "NFS"
	Protocol_SMB Protocol = "SMB"
	Protocol_SMB_ENCRYPTED Protocol = "SMB_ENCRYPTED"
	Protocol_LOCAL_FILE_ACCESS Protocol = "LOCAL_FILE_ACCESS"
	Protocol_NRPE Protocol = "NRPE"
	Protocol_XMPP Protocol = "XMPP"
	Protocol_IIOP Protocol = "IIOP"
	Protocol_IIOP_ENCRYPTED Protocol = "IIOP_ENCRYPTED"
	Protocol_JRMP Protocol = "JRMP"
	Protocol_JRMP_ENCRYPTED Protocol = "JRMP_ENCRYPTED"
	Protocol_IN_PROCESS_LIBRARY_CALL Protocol = "IN_PROCESS_LIBRARY_CALL"
	Protocol_CONTAINER_SPAWNING Protocol = "CONTAINER_SPAWNING"
)

type Quantity string

const (
	Quantity_VERY_FEW Quantity = "VERY_FEW"
	Quantity_FEW Quantity = "FEW"
	Quantity_MANY Quantity = "MANY"
	Quantity_VERY_MANY Quantity = "VERY_MANY"
)

type Scope interface {
	IsInScope() *bool
	Justification() *string
}

// The jsii proxy struct for Scope
type jsiiProxy_Scope struct {
	_ byte // padding
}

func (j *jsiiProxy_Scope) IsInScope() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Scope) Justification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"justification",
		&returns,
	)
	return returns
}


func NewScope_Override(s Scope, justification *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.Scope",
		[]interface{}{justification},
		s,
	)
}

type Size string

const (
	Size_SYSTEM Size = "SYSTEM"
	Size_SERVICE Size = "SERVICE"
	Size_APPLICATION Size = "APPLICATION"
	Size_COMPONENT Size = "COMPONENT"
)

type TechnicalAsset interface {
	Asset
	AssetType() AssetType
	CiaTriad() CIATriad
	Description() *string
	Encryption() Encryption
	HumanUse() *bool
	Internet() *bool
	Machine() Machine
	MultiTenant() *bool
	// The tree node.
	Node() constructs.Node
	Owner() *string
	Redundant() *bool
	Scope() Scope
	Size() Size
	Technology() Technology
	Usage() Usage
	Uuid() *string
	CommunicateWith(id *string, target TechnicalAsset, options *CommunicationOptions) Communication
	Process(assets ...DataAsset)
	Store(assets ...DataAsset)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TechnicalAsset
type jsiiProxy_TechnicalAsset struct {
	jsiiProxy_Asset
}

func (j *jsiiProxy_TechnicalAsset) AssetType() AssetType {
	var returns AssetType
	_jsii_.Get(
		j,
		"assetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) CiaTriad() CIATriad {
	var returns CIATriad
	_jsii_.Get(
		j,
		"ciaTriad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Encryption() Encryption {
	var returns Encryption
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) HumanUse() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"humanUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Internet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"internet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Machine() Machine {
	var returns Machine
	_jsii_.Get(
		j,
		"machine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) MultiTenant() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"multiTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Redundant() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"redundant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Scope() Scope {
	var returns Scope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Size() Size {
	var returns Size
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Technology() Technology {
	var returns Technology
	_jsii_.Get(
		j,
		"technology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Usage() Usage {
	var returns Usage
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TechnicalAsset) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


func NewTechnicalAsset(model constructs.Construct, id *string, props *TechnicalAssetProps) TechnicalAsset {
	_init_.Initialize()

	j := jsiiProxy_TechnicalAsset{}

	_jsii_.Create(
		"cdktg.TechnicalAsset",
		[]interface{}{model, id, props},
		&j,
	)

	return &j
}

func NewTechnicalAsset_Override(t TechnicalAsset, model constructs.Construct, id *string, props *TechnicalAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.TechnicalAsset",
		[]interface{}{model, id, props},
		t,
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
func TechnicalAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.TechnicalAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TechnicalAsset) CommunicateWith(id *string, target TechnicalAsset, options *CommunicationOptions) Communication {
	var returns Communication

	_jsii_.Invoke(
		t,
		"communicateWith",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TechnicalAsset) Process(assets ...DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"process",
		args,
	)
}

func (t *jsiiProxy_TechnicalAsset) Store(assets ...DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"store",
		args,
	)
}

func (t *jsiiProxy_TechnicalAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type TechnicalAssetProps struct {
	CiaTriad CIATriad `field:"required" json:"ciaTriad" yaml:"ciaTriad"`
	Description *string `field:"required" json:"description" yaml:"description"`
	Usage Usage `field:"required" json:"usage" yaml:"usage"`
	AssetType AssetType `field:"required" json:"assetType" yaml:"assetType"`
	Encryption Encryption `field:"required" json:"encryption" yaml:"encryption"`
	HumanUse *bool `field:"required" json:"humanUse" yaml:"humanUse"`
	Internet *bool `field:"required" json:"internet" yaml:"internet"`
	Machine Machine `field:"required" json:"machine" yaml:"machine"`
	MultiTenant *bool `field:"required" json:"multiTenant" yaml:"multiTenant"`
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	Redundant *bool `field:"required" json:"redundant" yaml:"redundant"`
	Size Size `field:"required" json:"size" yaml:"size"`
	Technology Technology `field:"required" json:"technology" yaml:"technology"`
	Scope Scope `field:"optional" json:"scope" yaml:"scope"`
	TrustBoundary TrustBoundary `field:"optional" json:"trustBoundary" yaml:"trustBoundary"`
}

type Technology string

const (
	Technology_UNKNOWN Technology = "UNKNOWN"
	Technology_CLIENT_SYSTEM Technology = "CLIENT_SYSTEM"
	Technology_BROWSER Technology = "BROWSER"
	Technology_DESKTOP Technology = "DESKTOP"
	Technology_MOBILE_APP Technology = "MOBILE_APP"
	Technology_DEVOPS_CLIENT Technology = "DEVOPS_CLIENT"
	Technology_WEB_SERVER Technology = "WEB_SERVER"
	Technology_WEB_APPLICATION Technology = "WEB_APPLICATION"
	Technology_APPLICATION_SERVER Technology = "APPLICATION_SERVER"
	Technology_DATABASE Technology = "DATABASE"
	Technology_FILE_SERVER Technology = "FILE_SERVER"
	Technology_LOCAL_FILE_SERVER Technology = "LOCAL_FILE_SERVER"
	Technology_ERP Technology = "ERP"
	Technology_CMS Technology = "CMS"
	Technology_WEB_SERVICE_REST Technology = "WEB_SERVICE_REST"
	Technology_WEB_SERVICE_SOAP Technology = "WEB_SERVICE_SOAP"
	Technology_EJB Technology = "EJB"
	Technology_SEARCH_INDEX Technology = "SEARCH_INDEX"
	Technology_SEARCH_ENGINE Technology = "SEARCH_ENGINE"
	Technology_SERVICE_REGISTRY Technology = "SERVICE_REGISTRY"
	Technology_REVERSE_PROXY Technology = "REVERSE_PROXY"
	Technology_LOAD_BALANCER Technology = "LOAD_BALANCER"
	Technology_BUILD_PIPELINE Technology = "BUILD_PIPELINE"
	Technology_SOURCECODE_REPOSITORY Technology = "SOURCECODE_REPOSITORY"
	Technology_ARTIFACT_REGISTRY Technology = "ARTIFACT_REGISTRY"
	Technology_CODE_INSPECTION_PLATFORM Technology = "CODE_INSPECTION_PLATFORM"
	Technology_MONITORING Technology = "MONITORING"
	Technology_LDAP_SERVER Technology = "LDAP_SERVER"
	Technology_CONTAINER_PLATFORM Technology = "CONTAINER_PLATFORM"
	Technology_BATCH_PROCESSING Technology = "BATCH_PROCESSING"
	Technology_EVENT_LISTENER Technology = "EVENT_LISTENER"
	Technology_IDENTITIY_PROVIDER Technology = "IDENTITIY_PROVIDER"
	Technology_IDENTITY_STORE_LDAP Technology = "IDENTITY_STORE_LDAP"
	Technology_IDENTITY_STORE_DATABASE Technology = "IDENTITY_STORE_DATABASE"
	Technology_TOOL Technology = "TOOL"
	Technology_CLI Technology = "CLI"
	Technology_TASK Technology = "TASK"
	Technology_FUNCTION Technology = "FUNCTION"
	Technology_GATEWAY Technology = "GATEWAY"
	Technology_IOT_DEVICE Technology = "IOT_DEVICE"
	Technology_MESSAGE_QUEUE Technology = "MESSAGE_QUEUE"
	Technology_STREAM_PROCESSING Technology = "STREAM_PROCESSING"
	Technology_SERVICE_MESH Technology = "SERVICE_MESH"
	Technology_DATA_LAKE Technology = "DATA_LAKE"
	Technology_REPORT_ENGINE Technology = "REPORT_ENGINE"
	Technology_AI Technology = "AI"
	Technology_MAIL_SERVER Technology = "MAIL_SERVER"
	Technology_VAULT Technology = "VAULT"
	Technology_HASM Technology = "HASM"
	Technology_WAF Technology = "WAF"
	Technology_IDS Technology = "IDS"
	Technology_IPS Technology = "IPS"
	Technology_SCHEDULER Technology = "SCHEDULER"
	Technology_MAINFRAME Technology = "MAINFRAME"
	Technology_BLOCK_STORAGE Technology = "BLOCK_STORAGE"
	Technology_LIBRARY Technology = "LIBRARY"
)

type TrustBoundary interface {
	constructs.Construct
	Description() *string
	// The tree node.
	Node() constructs.Node
	Type() TrustBoundaryType
	Uuid() *string
	AddTechnicalAssets(assets ...TechnicalAsset)
	AddTrustBoundary(boundary TrustBoundary)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TrustBoundary
type jsiiProxy_TrustBoundary struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TrustBoundary) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrustBoundary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrustBoundary) Type() TrustBoundaryType {
	var returns TrustBoundaryType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TrustBoundary) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


func NewTrustBoundary(model constructs.Construct, id *string, props *TrustBoundaryProps) TrustBoundary {
	_init_.Initialize()

	j := jsiiProxy_TrustBoundary{}

	_jsii_.Create(
		"cdktg.TrustBoundary",
		[]interface{}{model, id, props},
		&j,
	)

	return &j
}

func NewTrustBoundary_Override(t TrustBoundary, model constructs.Construct, id *string, props *TrustBoundaryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.TrustBoundary",
		[]interface{}{model, id, props},
		t,
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
func TrustBoundary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.TrustBoundary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TrustBoundary) AddTechnicalAssets(assets ...TechnicalAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"addTechnicalAssets",
		args,
	)
}

func (t *jsiiProxy_TrustBoundary) AddTrustBoundary(boundary TrustBoundary) {
	_jsii_.InvokeVoid(
		t,
		"addTrustBoundary",
		[]interface{}{boundary},
	)
}

func (t *jsiiProxy_TrustBoundary) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type TrustBoundaryProps struct {
	Description *string `field:"required" json:"description" yaml:"description"`
	Type TrustBoundaryType `field:"required" json:"type" yaml:"type"`
}

type TrustBoundaryType string

const (
	TrustBoundaryType_NETWORK_ON_PREM TrustBoundaryType = "NETWORK_ON_PREM"
	TrustBoundaryType_NETWORK_DEDICATED_HOSTER TrustBoundaryType = "NETWORK_DEDICATED_HOSTER"
	TrustBoundaryType_NETWORK_VIRTUAL_LAN TrustBoundaryType = "NETWORK_VIRTUAL_LAN"
	TrustBoundaryType_NETWORK_CLOUD_PROVIDER TrustBoundaryType = "NETWORK_CLOUD_PROVIDER"
	TrustBoundaryType_NETWORK_CLOUD_SECURITY_GROUP TrustBoundaryType = "NETWORK_CLOUD_SECURITY_GROUP"
	TrustBoundaryType_NETWORK_POLICY_NAMESPACE_ISOLATION TrustBoundaryType = "NETWORK_POLICY_NAMESPACE_ISOLATION"
	TrustBoundaryType_EXECUTION_ENVIRONMENT TrustBoundaryType = "EXECUTION_ENVIRONMENT"
)

type Usage string

const (
	Usage_BUSINESS Usage = "BUSINESS"
	Usage_DEVOPS Usage = "DEVOPS"
)

