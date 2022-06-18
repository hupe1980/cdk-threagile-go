package cdktg

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdktg.Asset",
		reflect.TypeOf((*Asset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ciaTriad", GoGetter: "CiaTriad"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_Asset{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.AssetProps",
		reflect.TypeOf((*AssetProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdktg.AssetType",
		reflect.TypeOf((*AssetType)(nil)).Elem(),
		map[string]interface{}{
			"EXTERNAL_ENTITY": AssetType_EXTERNAL_ENTITY,
			"PROCESS": AssetType_PROCESS,
			"DATASTORE": AssetType_DATASTORE,
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
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "ipFiltered", GoGetter: "IpFiltered"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "readonly", GoGetter: "Readonly"},
			_jsii_.MemberMethod{JsiiMethod: "received", GoMethod: "Received"},
			_jsii_.MemberMethod{JsiiMethod: "sent", GoMethod: "Sent"},
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_DataAsset{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Asset)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.DataAssetProps",
		reflect.TypeOf((*DataAssetProps)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"cdktg.InScope",
		reflect.TypeOf((*InScope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "isInScope", GoGetter: "IsInScope"},
			_jsii_.MemberProperty{JsiiProperty: "justification", GoGetter: "Justification"},
		},
		func() interface{} {
			j := jsiiProxy_InScope{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Scope)
			return &j
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
		"cdktg.Model",
		reflect.TypeOf((*Model)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "author", GoGetter: "Author"},
			_jsii_.MemberProperty{JsiiProperty: "businessCriticality", GoGetter: "BusinessCriticality"},
			_jsii_.MemberProperty{JsiiProperty: "date", GoGetter: "Date"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_Model{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.ModelProps",
		reflect.TypeOf((*ModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdktg.OutOfScope",
		reflect.TypeOf((*OutOfScope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "isInScope", GoGetter: "IsInScope"},
			_jsii_.MemberProperty{JsiiProperty: "justification", GoGetter: "Justification"},
		},
		func() interface{} {
			j := jsiiProxy_OutOfScope{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Scope)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.OutOfScopeProps",
		reflect.TypeOf((*OutOfScopeProps)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"cdktg.Scope",
		reflect.TypeOf((*Scope)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "isInScope", GoGetter: "IsInScope"},
			_jsii_.MemberProperty{JsiiProperty: "justification", GoGetter: "Justification"},
		},
		func() interface{} {
			return &jsiiProxy_Scope{}
		},
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
	_jsii_.RegisterClass(
		"cdktg.TechnicalAsset",
		reflect.TypeOf((*TechnicalAsset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assetType", GoGetter: "AssetType"},
			_jsii_.MemberProperty{JsiiProperty: "ciaTriad", GoGetter: "CiaTriad"},
			_jsii_.MemberMethod{JsiiMethod: "communicatedWith", GoMethod: "CommunicatedWith"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "humanUse", GoGetter: "HumanUse"},
			_jsii_.MemberProperty{JsiiProperty: "internet", GoGetter: "Internet"},
			_jsii_.MemberProperty{JsiiProperty: "machine", GoGetter: "Machine"},
			_jsii_.MemberProperty{JsiiProperty: "multiTenant", GoGetter: "MultiTenant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberMethod{JsiiMethod: "processed", GoMethod: "Processed"},
			_jsii_.MemberProperty{JsiiProperty: "redundant", GoGetter: "Redundant"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberMethod{JsiiMethod: "stored", GoMethod: "Stored"},
			_jsii_.MemberProperty{JsiiProperty: "technology", GoGetter: "Technology"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_TechnicalAsset{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Asset)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdktg.TechnicalAssetProps",
		reflect.TypeOf((*TechnicalAssetProps)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "uuid", GoGetter: "Uuid"},
		},
		func() interface{} {
			j := jsiiProxy_TrustBoundary{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
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
