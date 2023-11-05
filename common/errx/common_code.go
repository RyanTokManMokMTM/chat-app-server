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
	USER_STICKER_GROUP_EXIST  InternalCode = 200005
)

const (
	IS_FRIEND_ALREADY InternalCode = 300001
	NOT_YET_FRIEND    InternalCode = 300002
)

const (
	GROUP_NOT_EXIST    InternalCode = 400001
	ALREADY_IN_GROUP   InternalCode = 400002
	NO_GROUP_AUTHORITY InternalCode = 400003
	NOT_JOIN_GROUP_YET InternalCode = 400004
)

const (
	MESSAGE_NOT_EXIST           InternalCode = 500001
	NO_MESSAGE_DELETE_AUTHORITY InternalCode = 500002
)

const (
	STORY_NOT_EXIST      InternalCode = 600001
	STORY_CREATED_FAILED InternalCode = 600002
	STORY_NOT_AVAILABLE  InternalCode = 600003
)

const (
	STICKER_NOT_EXIST      InternalCode = 700001
	STICKER_CREATED_FAILED InternalCode = 700002
	STICKER_NOT_AVAILABLE  InternalCode = 700003
)
