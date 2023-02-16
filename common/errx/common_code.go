package errx

type InternalCode uint

const (
	SUCCESS                InternalCode = 20
	SERVER_COMMON_ERROR    InternalCode = 100001
	REQ_PARAM_ERROR        InternalCode = 100002
	TOKEN_EXPIRED_ERROR    InternalCode = 100003
	TOKEN_GENERATE_ERROR   InternalCode = 100004
	TOKEN_INVALID_ERROR    InternalCode = 100005
	DB_ERROR               InternalCode = 100006
	DB_AFFECTED_ZERO_ERROR InternalCode = 100007
)
