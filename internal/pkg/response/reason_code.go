package response

const UNKNOWN_ERROR = "UNKNOWN_ERROR"
const VALIDATE_ERROR = "VALIDATE_ERROR"
const USER_NOT_FOUND = "USER_NOT_FOUND"

var MappingReasonCode = map[string]int32{
	// 全局相关 -------------------------------------------------
	UNKNOWN_ERROR:  1000001,
	VALIDATE_ERROR: 1000002,

	// 用户相关 -------------------------------------------------
	USER_NOT_FOUND: 2000001,
}
