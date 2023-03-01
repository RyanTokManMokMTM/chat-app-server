package errx

var errMessage map[InternalCode]string

func init() {
	errMessage = make(map[InternalCode]string)
	errMessage[SUCCESS] = "SUCCESS"
	errMessage[SERVER_COMMON_ERROR] = "SERVER INTERNAL ERROR"
	errMessage[REQ_PARAM_ERROR] = "REQUEST PARAMETER ERROR"
	errMessage[TOKEN_EXPIRED_ERROR] = "TOKEN HAS BEEN EXPIRED"
	errMessage[TOKEN_INVALID_ERROR] = "TOKEN HAS BEEN INVALID"
	errMessage[TOKEN_GENERATE_ERROR] = "TOKEN GENERATE FAILED"
	errMessage[DB_ERROR] = "DATABASE ERROR"
	errMessage[DB_AFFECTED_ZERO_ERROR] = "DATABASE AFFECTED 0 rows"
	errMessage[FILE_UPLOAD_FAILED] = "FILE UPLOAD FAILED"

	//USER API ERROR CODE
	errMessage[USER_SIGN_UP_FAILED] = "USER SIGN UP FAILED"
	errMessage[USER_SIGN_IN_FAILED] = "EMAIL OR PASSWORD INCORRECT"
	errMessage[EMAIL_HAS_BEEN_REGISTERED] = "EMAIL HAS BEEN REGISTERED"
	errMessage[USER_NOT_EXIST] = "USER NOT EXISTS"

	errMessage[IS_FRIEND_ALREADY] = "USER IS YOUR FRIEND ALREADY"
	errMessage[NOT_YET_FRIEND] = "USER IS NOT YOUR FRIEND"

	errMessage[GROUP_NOT_EXIST] = "GROUP NOT EXIST"
	errMessage[ALREADY_IN_GROUP] = "USER ALREADY JOINED GROUP"
	errMessage[NO_GROUP_DELETE_AUTHORITY] = "USER DO NOT HAVE AUTHORITY TO REMOVE"
	errMessage[NOT_JOIN_GROUP_YET] = "USER HAVEN'T JOINED THE GROUP"

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
