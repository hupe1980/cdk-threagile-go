package cdktg

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdktg.AbuseCase",
		reflect.TypeOf((*AbuseCase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_AbuseCase{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.AbuseCaseProps",
		reflect.TypeOf((*AbuseCaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.AnnotationMetadataEntryType",
		reflect.TypeOf((*AnnotationMetadataEntryType)(nil)).Elem(),
		map[string]interface{}{
			"INFO": AnnotationMetadataEntryType_INFO,
			"WARN": AnnotationMetadataEntryType_WARN,
			"ERROR": AnnotationMetadataEntryType_ERROR,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.Annotations",
		reflect.TypeOf((*Annotations)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addError", GoMethod: "AddError"},
			_jsii_.MemberMethod{JsiiMethod: "addInfo", GoMethod: "AddInfo"},
			_jsii_.MemberMethod{JsiiMethod: "addWarning", GoMethod: "AddWarning"},
		},
		func() interface{} {
			return &jsiiProxy_Annotations{}
		},
	)
	_jsii_.RegisterClass(
		"cdktg.Aspects",
		reflect.TypeOf((*Aspects)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberProperty{JsiiProperty: "all", GoGetter: "All"},
		},
		func() interface{} {
			return &jsiiProxy_Aspects{}
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Authentication",
		reflect.TypeOf((*Authentication)(nil)).Elem(),
		map[string]interface{}{
			"NONE": Authentication_NONE,
			"CREDENTIALS": Authentication_CREDENTIALS,
			"SESSION_ID": Authentication_SESSION_ID,
			"TOKEN": Authentication_TOKEN,
			"CLIENT_CERTIFICATE": Authentication_CLIENT_CERTIFICATE,
			"TWO_FACTOR": Authentication_TWO_FACTOR,
			"EXTERNALIZED": Authentication_EXTERNALIZED,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.Author",
		reflect.TypeOf((*Author)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "homepage", GoGetter: "Homepage"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Author{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.AuthorProps",
		reflect.TypeOf((*AuthorProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.Authorization",
		reflect.TypeOf((*Authorization)(nil)).Elem(),
		map[string]interface{}{
			"NONE": Authorization_NONE,
			"TECHNICAL_USER": Authorization_TECHNICAL_USER,
			"ENDUSER_IDENTITY_PROPAGATION": Authorization_ENDUSER_IDENTITY_PROPAGATION,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Availability",
		reflect.TypeOf((*Availability)(nil)).Elem(),
		map[string]interface{}{
			"ARCHIVE": Availability_ARCHIVE,
			"OPERATIONAL": Availability_OPERATIONAL,
			"IMPORTANT": Availability_IMPORTANT,
			"CRITICAL": Availability_CRITICAL,
			"MISSION_CRITICAL": Availability_MISSION_CRITICAL,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.BusinessCriticality",
		reflect.TypeOf((*BusinessCriticality)(nil)).Elem(),
		map[string]interface{}{
			"ARCHIVE": BusinessCriticality_ARCHIVE,
			"OPERATIONAL": BusinessCriticality_OPERATIONAL,
			"IMPORTANT": BusinessCriticality_IMPORTANT,
			"CRITICAL": BusinessCriticality_CRITICAL,
			"MISSION_CRITICAL": BusinessCriticality_MISSION_CRITICAL,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.CIATriad",
		reflect.TypeOf((*CIATriad)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availability", GoGetter: "Availability"},
			_jsii_.MemberProperty{JsiiProperty: "confidentiality", GoGetter: "Confidentiality"},
			_jsii_.MemberProperty{JsiiProperty: "integrity", GoGetter: "Integrity"},
			_jsii_.MemberProperty{JsiiProperty: "justification", GoGetter: "Justification"},
		},
		func() interface{} {
			return &jsiiProxy_CIATriad{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.CIATriadProps",
		reflect.TypeOf((*CIATriadProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.Communication",
		reflect.TypeOf((*Communication)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authorization", GoGetter: "Authorization"},
			_jsii_.MemberProperty{JsiiProperty: "caller", GoGetter: "Caller"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "ipFiltered", GoGetter: "IpFiltered"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "receives", GoMethod: "Receives"},
			_jsii_.MemberMethod{JsiiMethod: "sends", GoMethod: "Sends"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "vpn", GoGetter: "Vpn"},
		},
		func() interface{} {
			return &jsiiProxy_Communication{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.CommunicationOptions",
		reflect.TypeOf((*CommunicationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktg.CommunicationProps",
		reflect.TypeOf((*CommunicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.Confidentiality",
		reflect.TypeOf((*Confidentiality)(nil)).Elem(),
		map[string]interface{}{
			"PUBLIC": Confidentiality_PUBLIC,
			"INTERNAL": Confidentiality_INTERNAL,
			"RESTRICTED": Confidentiality_RESTRICTED,
			"CONFIDENTIAL": Confidentiality_CONFIDENTIAL,
			"STRICTLY_CONFIDENTIAL": Confidentiality_STRICTLY_CONFIDENTIAL,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.DataAsset",
		reflect.TypeOf((*DataAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ciaTriad", GoGetter: "CiaTriad"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "quantity", GoGetter: "Quantity"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_DataAsset{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.DataAssetProps",
		reflect.TypeOf((*DataAssetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.DataBreachProbability",
		reflect.TypeOf((*DataBreachProbability)(nil)).Elem(),
		map[string]interface{}{
			"IMPROBABLE": DataBreachProbability_IMPROBABLE,
			"POSSIBLE": DataBreachProbability_POSSIBLE,
			"PROBABLE": DataBreachProbability_PROBABLE,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.DataFormat",
		reflect.TypeOf((*DataFormat)(nil)).Elem(),
		map[string]interface{}{
			"JSON": DataFormat_JSON,
			"XML": DataFormat_XML,
			"SERIALIZATION": DataFormat_SERIALIZATION,
			"FILE": DataFormat_FILE,
			"CSV": DataFormat_CSV,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Encryption",
		reflect.TypeOf((*Encryption)(nil)).Elem(),
		map[string]interface{}{
			"NONE": Encryption_NONE,
			"TRANSPARENT": Encryption_TRANSPARENT,
			"SYMMETRIC_SHARED_KEY": Encryption_SYMMETRIC_SHARED_KEY,
			"ASYMMETRIC_SHARED_KEY": Encryption_ASYMMETRIC_SHARED_KEY,
			"ENDUSER_INDIVIDUAL_KEY": Encryption_ENDUSER_INDIVIDUAL_KEY,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.ExploitationImpact",
		reflect.TypeOf((*ExploitationImpact)(nil)).Elem(),
		map[string]interface{}{
			"LOW": ExploitationImpact_LOW,
			"MEDIUM": ExploitationImpact_MEDIUM,
			"HIGH": ExploitationImpact_HIGH,
			"VERY_HIGH": ExploitationImpact_VERY_HIGH,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.ExploitationLikelihood",
		reflect.TypeOf((*ExploitationLikelihood)(nil)).Elem(),
		map[string]interface{}{
			"UNLIKELY": ExploitationLikelihood_UNLIKELY,
			"LIKELY": ExploitationLikelihood_LIKELY,
			"VERY_LIKELY": ExploitationLikelihood_VERY_LIKELY,
			"FREQUENT": ExploitationLikelihood_FREQUENT,
		},
	)
	_jsii_.RegisterInterface(
		"cdktg.IAspect",
		reflect.TypeOf((*IAspect)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			return &jsiiProxy_IAspect{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktg.ICustomSynthesis",
		reflect.TypeOf((*ICustomSynthesis)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
		},
		func() interface{} {
			return &jsiiProxy_ICustomSynthesis{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktg.IManifest",
		reflect.TypeOf((*IManifest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "models", GoGetter: "Models"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_IManifest{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktg.IModelSynthesizer",
		reflect.TypeOf((*IModelSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			return &jsiiProxy_IModelSynthesizer{}
		},
	)
	_jsii_.RegisterInterface(
		"cdktg.ISynthesisSession",
		reflect.TypeOf((*ISynthesisSession)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "skipValidation", GoGetter: "SkipValidation"},
		},
		func() interface{} {
			return &jsiiProxy_ISynthesisSession{}
		},
	)
	_jsii_.RegisterClass(
		"cdktg.Image",
		reflect.TypeOf((*Image)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filePath", GoGetter: "FilePath"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
		},
		func() interface{} {
			return &jsiiProxy_Image{}
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Integrity",
		reflect.TypeOf((*Integrity)(nil)).Elem(),
		map[string]interface{}{
			"ARCHIVE": Integrity_ARCHIVE,
			"OPERATIONAL": Integrity_OPERATIONAL,
			"IMPORTANT": Integrity_IMPORTANT,
			"CRITICAL": Integrity_CRITICAL,
			"MISSION_CRITICAL": Integrity_MISSION_CRITICAL,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Machine",
		reflect.TypeOf((*Machine)(nil)).Elem(),
		map[string]interface{}{
			"PHYSICAL": Machine_PHYSICAL,
			"VIRTUAL": Machine_VIRTUAL,
			"CONTAINER": Machine_CONTAINER,
			"SERVERLESS": Machine_SERVERLESS,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.Manifest",
		reflect.TypeOf((*Manifest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "buildManifest", GoMethod: "BuildManifest"},
			_jsii_.MemberMethod{JsiiMethod: "forModel", GoMethod: "ForModel"},
			_jsii_.MemberProperty{JsiiProperty: "models", GoGetter: "Models"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberMethod{JsiiMethod: "writeToFile", GoMethod: "WriteToFile"},
		},
		func() interface{} {
			j := jsiiProxy_Manifest{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IManifest)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdktg.Model",
		reflect.TypeOf((*Model)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAbuseCases", GoMethod: "AddAbuseCases"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addQuestion", GoMethod: "AddQuestion"},
			_jsii_.MemberMethod{JsiiMethod: "addSecurityRequirements", GoMethod: "AddSecurityRequirements"},
			_jsii_.MemberMethod{JsiiMethod: "addTag", GoMethod: "AddTag"},
			_jsii_.MemberMethod{JsiiMethod: "addTags", GoMethod: "AddTags"},
			_jsii_.MemberProperty{JsiiProperty: "author", GoGetter: "Author"},
			_jsii_.MemberProperty{JsiiProperty: "businessCriticality", GoGetter: "BusinessCriticality"},
			_jsii_.MemberProperty{JsiiProperty: "businessOverview", GoGetter: "BusinessOverview"},
			_jsii_.MemberProperty{JsiiProperty: "date", GoGetter: "Date"},
			_jsii_.MemberProperty{JsiiProperty: "managementSummary", GoGetter: "ManagementSummary"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "technicalOverview", GoGetter: "TechnicalOverview"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "trackRisk", GoMethod: "TrackRisk"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_Model{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.ModelAnnotation",
		reflect.TypeOf((*ModelAnnotation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktg.ModelManifest",
		reflect.TypeOf((*ModelManifest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktg.ModelProps",
		reflect.TypeOf((*ModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.ModelSynthesizer",
		reflect.TypeOf((*ModelSynthesizer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberProperty{JsiiProperty: "model", GoGetter: "Model"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
		},
		func() interface{} {
			j := jsiiProxy_ModelSynthesizer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IModelSynthesizer)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.OutOfScopeProps",
		reflect.TypeOf((*OutOfScopeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.Overview",
		reflect.TypeOf((*Overview)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "images", GoGetter: "Images"},
		},
		func() interface{} {
			return &jsiiProxy_Overview{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.OverviewProps",
		reflect.TypeOf((*OverviewProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.Project",
		reflect.TypeOf((*Project)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outdir", GoGetter: "Outdir"},
			_jsii_.MemberProperty{JsiiProperty: "skipValidation", GoGetter: "SkipValidation"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Project{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.ProjectProps",
		reflect.TypeOf((*ProjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"UNKNOEN": Protocol_UNKNOEN,
			"HTTP": Protocol_HTTP,
			"HTTPS": Protocol_HTTPS,
			"WS": Protocol_WS,
			"WSS": Protocol_WSS,
			"REVERSE_PROXY_WEB_PROTOCOL": Protocol_REVERSE_PROXY_WEB_PROTOCOL,
			"REVERSE_PROXY_WEB_PROTOCOL_ENCRYPTED": Protocol_REVERSE_PROXY_WEB_PROTOCOL_ENCRYPTED,
			"MQTT": Protocol_MQTT,
			"JDBC": Protocol_JDBC,
			"JDBC_ENCRYPTED": Protocol_JDBC_ENCRYPTED,
			"ODBC": Protocol_ODBC,
			"ODBC_ENCRYPTED": Protocol_ODBC_ENCRYPTED,
			"SQL_ACCESS_PROTOCOL": Protocol_SQL_ACCESS_PROTOCOL,
			"SQL_ACCESS_PROTOCOL_ENCRYPTED": Protocol_SQL_ACCESS_PROTOCOL_ENCRYPTED,
			"NOSQL_ACCESS_PROTOCOL": Protocol_NOSQL_ACCESS_PROTOCOL,
			"NOSQL_ACCESS_PROTOCOL_ENCRYPTED": Protocol_NOSQL_ACCESS_PROTOCOL_ENCRYPTED,
			"BINARY": Protocol_BINARY,
			"BINARY_ENCRYPTED": Protocol_BINARY_ENCRYPTED,
			"TEXT": Protocol_TEXT,
			"TEXT_ENCRYPTED": Protocol_TEXT_ENCRYPTED,
			"SSH": Protocol_SSH,
			"SSH_TUNNEL": Protocol_SSH_TUNNEL,
			"SMTP": Protocol_SMTP,
			"SMTP_ENCRYPTED": Protocol_SMTP_ENCRYPTED,
			"POP3": Protocol_POP3,
			"POP3_ENCRYPTED": Protocol_POP3_ENCRYPTED,
			"IMAP": Protocol_IMAP,
			"IMAP_ENCRYPTED": Protocol_IMAP_ENCRYPTED,
			"FTP": Protocol_FTP,
			"FTPS": Protocol_FTPS,
			"SFTP": Protocol_SFTP,
			"SCP": Protocol_SCP,
			"LDAP": Protocol_LDAP,
			"LDAPS": Protocol_LDAPS,
			"JMS": Protocol_JMS,
			"NFS": Protocol_NFS,
			"SMB": Protocol_SMB,
			"SMB_ENCRYPTED": Protocol_SMB_ENCRYPTED,
			"LOCAL_FILE_ACCESS": Protocol_LOCAL_FILE_ACCESS,
			"NRPE": Protocol_NRPE,
			"XMPP": Protocol_XMPP,
			"IIOP": Protocol_IIOP,
			"IIOP_ENCRYPTED": Protocol_IIOP_ENCRYPTED,
			"JRMP": Protocol_JRMP,
			"JRMP_ENCRYPTED": Protocol_JRMP_ENCRYPTED,
			"IN_PROCESS_LIBRARY_CALL": Protocol_IN_PROCESS_LIBRARY_CALL,
			"CONTAINER_SPAWNING": Protocol_CONTAINER_SPAWNING,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Quantity",
		reflect.TypeOf((*Quantity)(nil)).Elem(),
		map[string]interface{}{
			"VERY_FEW": Quantity_VERY_FEW,
			"FEW": Quantity_FEW,
			"MANY": Quantity_MANY,
			"VERY_MANY": Quantity_VERY_MANY,
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.Question",
		reflect.TypeOf((*Question)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.Resource",
		reflect.TypeOf((*Resource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_Resource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.ResourceProps",
		reflect.TypeOf((*ResourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.Risk",
		reflect.TypeOf((*Risk)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dataBreachProbability", GoGetter: "DataBreachProbability"},
			_jsii_.MemberProperty{JsiiProperty: "dataBreachTechnicalAssets", GoGetter: "DataBreachTechnicalAssets"},
			_jsii_.MemberProperty{JsiiProperty: "exploitationImpact", GoGetter: "ExploitationImpact"},
			_jsii_.MemberProperty{JsiiProperty: "exploitationLikelihood", GoGetter: "ExploitationLikelihood"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "mostRelevantCommunicationLink", GoGetter: "MostRelevantCommunicationLink"},
			_jsii_.MemberProperty{JsiiProperty: "mostRelevantDataAsset", GoGetter: "MostRelevantDataAsset"},
			_jsii_.MemberProperty{JsiiProperty: "mostRelevantSharedRuntime", GoGetter: "MostRelevantSharedRuntime"},
			_jsii_.MemberProperty{JsiiProperty: "mostRelevantTechnicalAsset", GoGetter: "MostRelevantTechnicalAsset"},
			_jsii_.MemberProperty{JsiiProperty: "mostRelevantTrustBoundary", GoGetter: "MostRelevantTrustBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "severity", GoGetter: "Severity"},
		},
		func() interface{} {
			return &jsiiProxy_Risk{}
		},
	)
	_jsii_.RegisterClass(
		"cdktg.RiskCategory",
		reflect.TypeOf((*RiskCategory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberMethod{JsiiMethod: "addIdentifiedRisk", GoMethod: "AddIdentifiedRisk"},
			_jsii_.MemberProperty{JsiiProperty: "asvs", GoGetter: "Asvs"},
			_jsii_.MemberProperty{JsiiProperty: "cheatSheat", GoGetter: "CheatSheat"},
			_jsii_.MemberProperty{JsiiProperty: "check", GoGetter: "Check"},
			_jsii_.MemberProperty{JsiiProperty: "cwe", GoGetter: "Cwe"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "detectionLogic", GoGetter: "DetectionLogic"},
			_jsii_.MemberProperty{JsiiProperty: "falsePositives", GoGetter: "FalsePositives"},
			_jsii_.MemberProperty{JsiiProperty: "function", GoGetter: "Function"},
			_jsii_.MemberMethod{JsiiMethod: "identifiedAtDataAsset", GoMethod: "IdentifiedAtDataAsset"},
			_jsii_.MemberMethod{JsiiMethod: "identifiedAtSharedRuntime", GoMethod: "IdentifiedAtSharedRuntime"},
			_jsii_.MemberMethod{JsiiMethod: "identifiedAtTechnicalAsset", GoMethod: "IdentifiedAtTechnicalAsset"},
			_jsii_.MemberMethod{JsiiMethod: "identifiedAtTrustBoundary", GoMethod: "IdentifiedAtTrustBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "impact", GoGetter: "Impact"},
			_jsii_.MemberProperty{JsiiProperty: "mitigation", GoGetter: "Mitigation"},
			_jsii_.MemberProperty{JsiiProperty: "modelFailurePossibleReason", GoGetter: "ModelFailurePossibleReason"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "riskAssessment", GoGetter: "RiskAssessment"},
			_jsii_.MemberProperty{JsiiProperty: "stride", GoGetter: "Stride"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_RiskCategory{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.RiskCategoryProps",
		reflect.TypeOf((*RiskCategoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.RiskFunction",
		reflect.TypeOf((*RiskFunction)(nil)).Elem(),
		map[string]interface{}{
			"BUSINESS_SIDE": RiskFunction_BUSINESS_SIDE,
			"ARCHITECTURE": RiskFunction_ARCHITECTURE,
			"DEVELOPMENT": RiskFunction_DEVELOPMENT,
			"OPERATIONS": RiskFunction_OPERATIONS,
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.RiskOptions",
		reflect.TypeOf((*RiskOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdktg.RiskProps",
		reflect.TypeOf((*RiskProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.RiskTracking",
		reflect.TypeOf((*RiskTracking)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "checkedBy", GoGetter: "CheckedBy"},
			_jsii_.MemberProperty{JsiiProperty: "date", GoGetter: "Date"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "justification", GoGetter: "Justification"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "ticket", GoGetter: "Ticket"},
		},
		func() interface{} {
			return &jsiiProxy_RiskTracking{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.RiskTrackingProps",
		reflect.TypeOf((*RiskTrackingProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.RiskTrackingStatus",
		reflect.TypeOf((*RiskTrackingStatus)(nil)).Elem(),
		map[string]interface{}{
			"UNCHECKED": RiskTrackingStatus_UNCHECKED,
			"IN_DISCUSSION": RiskTrackingStatus_IN_DISCUSSION,
			"ACCEPTED": RiskTrackingStatus_ACCEPTED,
			"IN_PROGRESS": RiskTrackingStatus_IN_PROGRESS,
			"MITIGATED": RiskTrackingStatus_MITIGATED,
			"FALSE_POSITIVE": RiskTrackingStatus_FALSE_POSITIVE,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.Scope",
		reflect.TypeOf((*Scope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "justification", GoGetter: "Justification"},
			_jsii_.MemberProperty{JsiiProperty: "out", GoGetter: "Out"},
		},
		func() interface{} {
			return &jsiiProxy_Scope{}
		},
	)
	_jsii_.RegisterClass(
		"cdktg.SecurityRequirement",
		reflect.TypeOf((*SecurityRequirement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_SecurityRequirement{}
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.SecurityRequirementProps",
		reflect.TypeOf((*SecurityRequirementProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.Severity",
		reflect.TypeOf((*Severity)(nil)).Elem(),
		map[string]interface{}{
			"LOW": Severity_LOW,
			"MEDIUM": Severity_MEDIUM,
			"ELEVATED": Severity_ELEVATED,
			"HIGH": Severity_HIGH,
			"CRITICAL": Severity_CRITICAL,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.SharedRuntime",
		reflect.TypeOf((*SharedRuntime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "runs", GoMethod: "Runs"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_SharedRuntime{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.SharedRuntimeProps",
		reflect.TypeOf((*SharedRuntimeProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.Size",
		reflect.TypeOf((*Size)(nil)).Elem(),
		map[string]interface{}{
			"SYSTEM": Size_SYSTEM,
			"SERVICE": Size_SERVICE,
			"APPLICATION": Size_APPLICATION,
			"COMPONENT": Size_COMPONENT,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Stride",
		reflect.TypeOf((*Stride)(nil)).Elem(),
		map[string]interface{}{
			"SPOOFING": Stride_SPOOFING,
			"TAMPERING": Stride_TAMPERING,
			"REPUDIATION": Stride_REPUDIATION,
			"INFORMATION_DISCLOSURE": Stride_INFORMATION_DISCLOSURE,
			"DENIAL_OF_SERVICE": Stride_DENIAL_OF_SERVICE,
			"ELEVATION_OF_PRIVILEGE": Stride_ELEVATION_OF_PRIVILEGE,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.TechnicalAsset",
		reflect.TypeOf((*TechnicalAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ciaTriad", GoGetter: "CiaTriad"},
			_jsii_.MemberMethod{JsiiMethod: "communicatesWith", GoMethod: "CommunicatesWith"},
			_jsii_.MemberProperty{JsiiProperty: "customDevelopedParts", GoGetter: "CustomDevelopedParts"},
			_jsii_.MemberProperty{JsiiProperty: "dataFormatsAccepted", GoGetter: "DataFormatsAccepted"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "humanUse", GoGetter: "HumanUse"},
			_jsii_.MemberProperty{JsiiProperty: "internet", GoGetter: "Internet"},
			_jsii_.MemberProperty{JsiiProperty: "machine", GoGetter: "Machine"},
			_jsii_.MemberProperty{JsiiProperty: "multiTenant", GoGetter: "MultiTenant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberMethod{JsiiMethod: "processes", GoMethod: "Processes"},
			_jsii_.MemberProperty{JsiiProperty: "redundant", GoGetter: "Redundant"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberMethod{JsiiMethod: "stores", GoMethod: "Stores"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "technology", GoGetter: "Technology"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_TechnicalAsset{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.TechnicalAssetProps",
		reflect.TypeOf((*TechnicalAssetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.TechnicalAssetType",
		reflect.TypeOf((*TechnicalAssetType)(nil)).Elem(),
		map[string]interface{}{
			"EXTERNAL_ENTITY": TechnicalAssetType_EXTERNAL_ENTITY,
			"PROCESS": TechnicalAssetType_PROCESS,
			"DATASTORE": TechnicalAssetType_DATASTORE,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Technology",
		reflect.TypeOf((*Technology)(nil)).Elem(),
		map[string]interface{}{
			"UNKNOWN": Technology_UNKNOWN,
			"CLIENT_SYSTEM": Technology_CLIENT_SYSTEM,
			"BROWSER": Technology_BROWSER,
			"DESKTOP": Technology_DESKTOP,
			"MOBILE_APP": Technology_MOBILE_APP,
			"DEVOPS_CLIENT": Technology_DEVOPS_CLIENT,
			"WEB_SERVER": Technology_WEB_SERVER,
			"WEB_APPLICATION": Technology_WEB_APPLICATION,
			"APPLICATION_SERVER": Technology_APPLICATION_SERVER,
			"DATABASE": Technology_DATABASE,
			"FILE_SERVER": Technology_FILE_SERVER,
			"LOCAL_FILE_SERVER": Technology_LOCAL_FILE_SERVER,
			"ERP": Technology_ERP,
			"CMS": Technology_CMS,
			"WEB_SERVICE_REST": Technology_WEB_SERVICE_REST,
			"WEB_SERVICE_SOAP": Technology_WEB_SERVICE_SOAP,
			"EJB": Technology_EJB,
			"SEARCH_INDEX": Technology_SEARCH_INDEX,
			"SEARCH_ENGINE": Technology_SEARCH_ENGINE,
			"SERVICE_REGISTRY": Technology_SERVICE_REGISTRY,
			"REVERSE_PROXY": Technology_REVERSE_PROXY,
			"LOAD_BALANCER": Technology_LOAD_BALANCER,
			"BUILD_PIPELINE": Technology_BUILD_PIPELINE,
			"SOURCECODE_REPOSITORY": Technology_SOURCECODE_REPOSITORY,
			"ARTIFACT_REGISTRY": Technology_ARTIFACT_REGISTRY,
			"CODE_INSPECTION_PLATFORM": Technology_CODE_INSPECTION_PLATFORM,
			"MONITORING": Technology_MONITORING,
			"LDAP_SERVER": Technology_LDAP_SERVER,
			"CONTAINER_PLATFORM": Technology_CONTAINER_PLATFORM,
			"BATCH_PROCESSING": Technology_BATCH_PROCESSING,
			"EVENT_LISTENER": Technology_EVENT_LISTENER,
			"IDENTITIY_PROVIDER": Technology_IDENTITIY_PROVIDER,
			"IDENTITY_STORE_LDAP": Technology_IDENTITY_STORE_LDAP,
			"IDENTITY_STORE_DATABASE": Technology_IDENTITY_STORE_DATABASE,
			"TOOL": Technology_TOOL,
			"CLI": Technology_CLI,
			"TASK": Technology_TASK,
			"FUNCTION": Technology_FUNCTION,
			"GATEWAY": Technology_GATEWAY,
			"IOT_DEVICE": Technology_IOT_DEVICE,
			"MESSAGE_QUEUE": Technology_MESSAGE_QUEUE,
			"STREAM_PROCESSING": Technology_STREAM_PROCESSING,
			"SERVICE_MESH": Technology_SERVICE_MESH,
			"DATA_LAKE": Technology_DATA_LAKE,
			"REPORT_ENGINE": Technology_REPORT_ENGINE,
			"AI": Technology_AI,
			"MAIL_SERVER": Technology_MAIL_SERVER,
			"VAULT": Technology_VAULT,
			"HASM": Technology_HASM,
			"WAF": Technology_WAF,
			"IDS": Technology_IDS,
			"IPS": Technology_IPS,
			"SCHEDULER": Technology_SCHEDULER,
			"MAINFRAME": Technology_MAINFRAME,
			"BLOCK_STORAGE": Technology_BLOCK_STORAGE,
			"LIBRARY": Technology_LIBRARY,
		},
	)
	_jsii_.RegisterClass(
		"cdktg.TrustBoundary",
		reflect.TypeOf((*TrustBoundary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addTechnicalAssets", GoMethod: "AddTechnicalAssets"},
			_jsii_.MemberMethod{JsiiMethod: "addTrustBoundary", GoMethod: "AddTrustBoundary"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_TrustBoundary{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Resource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.TrustBoundaryProps",
		reflect.TypeOf((*TrustBoundaryProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.TrustBoundaryType",
		reflect.TypeOf((*TrustBoundaryType)(nil)).Elem(),
		map[string]interface{}{
			"NETWORK_ON_PREM": TrustBoundaryType_NETWORK_ON_PREM,
			"NETWORK_DEDICATED_HOSTER": TrustBoundaryType_NETWORK_DEDICATED_HOSTER,
			"NETWORK_VIRTUAL_LAN": TrustBoundaryType_NETWORK_VIRTUAL_LAN,
			"NETWORK_CLOUD_PROVIDER": TrustBoundaryType_NETWORK_CLOUD_PROVIDER,
			"NETWORK_CLOUD_SECURITY_GROUP": TrustBoundaryType_NETWORK_CLOUD_SECURITY_GROUP,
			"NETWORK_POLICY_NAMESPACE_ISOLATION": TrustBoundaryType_NETWORK_POLICY_NAMESPACE_ISOLATION,
			"EXECUTION_ENVIRONMENT": TrustBoundaryType_EXECUTION_ENVIRONMENT,
		},
	)
	_jsii_.RegisterEnum(
		"cdktg.Usage",
		reflect.TypeOf((*Usage)(nil)).Elem(),
		map[string]interface{}{
			"BUSINESS": Usage_BUSINESS,
			"DEVOPS": Usage_DEVOPS,
		},
	)
}
