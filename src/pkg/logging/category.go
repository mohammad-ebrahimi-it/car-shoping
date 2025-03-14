package logging

type Category string

type SubCategory string

type ExtraKey string

const (
	Internal        Category = "Internal"
	General         Category = "General"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
)

const (
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"

	Select   SubCategory = "Select"
	Rollback SubCategory = "Rollback"
	Update   SubCategory = "Update"
	Delete   SubCategory = "Delete"
	Insert   SubCategory = "Insert"

	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"

	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
)
