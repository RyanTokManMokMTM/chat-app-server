package errx

var errMessage map[InternalCode]string

func init() {
	errMessage[SUCCESS] = "SUCCESS"
	errMessage[SERVER_COMMON_ERROR] = "SERVER INTERNAL ERROR"
	errMessage[REQ_PARAM_ERROR] = "REQUEST PARAMETER ERROR"
	errMessage[TOKEN_EXPIRED_ERROR] = "TOKEN HAS BEEN EXPIRED"
	errMessage[TOKEN_INVALID_ERROR] = "TOKEN HAS BEEN INVALID"
	errMessage[TOKEN_GENERATE_ERROR] = "TOKEN GENERATE FAILED"
	errMessage[DB_ERROR] = "DATABASE ERROR"
	errMessage[DB_AFFECTED_ZERO_ERROR] = "DATABASE AFFECTED 0 rows"

}

func MapErrMsg(code InternalCode) string {
	if msg, ok := errMessage[code]; ok {
		return msg
	}
	return "SERVER INTERNAL ERROR"
}

func IsErrorCode(code InternalCode) bool {
	_, ok := errMessage[code]
	return ok
}
