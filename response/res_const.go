package response

const (
	SUCCESS        = 2000
	INVALID_PARAMS = 3000
	FAIL_VALIDATE  = 4000
	FAIL_AUTH      = 5000
	OTHER_ERROR    = 6000

	// 数据库相关
	ERROR_EXIST_RECODE     = 10001
	ERROR_NOT_EXIST_RECODE = 10002
	ERROR_DB_OPER          = 10003
)
