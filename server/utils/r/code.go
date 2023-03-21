package r

// 错误码汇总
const (
	OK   = 0
	FAIL = 500

	// code = 9000... 通用错误
	ERROR_REQUEST_PARAM = 9001
	ERROR_REQUEST_PAGE  = 9002
	ERROR_INVALID_PARAM = 9003
	ERROR_DB_OPE        = 9004

	// code = 9100... 文件相关错误
	ERROR_FILE_UPLOAD  = 9100
	ERROR_FILE_RECEIVE = 9101

	// code = 1000... 用户模块错误
	ERROR_USER_NAME_USED    = 1001
	ERROR_PASSWORD_WRONG    = 1002
	ERROR_USER_NOT_EXIST    = 1003
	ERROR_USER_NO_RIGHT     = 1009
	ERROR_OLD_PASSWORD      = 1010
	ERROR_EMAIL_SEND        = 1011
	ERROR_EMAIL_HAS_SEND    = 1012
	ERROR_VERIFICATION_CODE = 1013

	// code = 1200... 鉴权相关错误
	ERROR_TOKEN_NOT_EXIST  = 1201
	ERROR_TOKEN_RUNTIME    = 1202
	ERROR_TOKEN_WRONG      = 1203
	ERROR_TOKEN_TYPE_WRONG = 1204
	ERROR_TOKEN_CREATE     = 1205
	ERROR_TOKEN_DENIED     = 1206
	FORCE_OFFLINE          = 1207
	LOGOUT                 = 1208

	// code = 2000... 文章模块错误
	ERROR_ART_NOT_EXIST = 2001

	// code = 3000... 分类模块错误
	ERROR_CATE_NAME_USED = 3001
	ERROR_CATE_NOT_EXIST = 3002
	ERROR_CATE_ART_EXIST = 3003

	// code = 4000... 标签模块的错误
	ERROR_TAG_EXIST     = 4001
	ERROR_TAG_NOT_EXIST = 4002
	ERROR_TAG_ART_EXIST = 4003

	// code = 5000... 评论模块错误
	ERROR_COMMENT_NOT_EXIST = 5001

	// code = 6000... 权限模块错误
	ERROR_RESOURCE_NAME_EXIST   = 6001
	ERROR_RESOURCE_NOT_EXIST    = 6002
	ERROR_RESOURCE_USED_BY_ROLE = 6003
	ERROR_RESOURCE_HAS_CHLIDREN = 6004
	ERROR_MENU_NAME_EXIST       = 6005
	ERROR_MENU_NOT_EXIST        = 6006
	ERROR_MENU_USED_BY_ROLE     = 6007
	ERROR_MENU_HAS_CHILDREN     = 6008
	ERROR_ROLE_NAME_EXIST       = 60010
	ERROR_ROLE_NOT_EXIST        = 60011

	// code = 7000... 页面模块错误
	ERROR_PAGE_NAME_EXIST = 7001
)
