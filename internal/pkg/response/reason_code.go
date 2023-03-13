package response

const UNKNOWN_ERROR = "UNKNOWN_ERROR"
const USER_NOT_FOUND = "USER_NOT_FOUND"

var MappingReasonCode = map[string]int32{
	// 全局相关 -------------------------------------------------
	UNKNOWN_ERROR: 1000001,
	// 用户相关 -------------------------------------------------
	USER_NOT_FOUND: 2000001,
}
