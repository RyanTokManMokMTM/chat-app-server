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
	FILE_UPLOAD_FAILED     InternalCode = 100008
)

const (
	USER_SIGN_UP_FAILED       InternalCode = 200001
	USER_SIGN_IN_FAILED       InternalCode = 200002
	EMAIL_HAS_BEEN_REGISTERED InternalCode = 200003
	USER_NOT_EXIST            InternalCode = 200004
)
