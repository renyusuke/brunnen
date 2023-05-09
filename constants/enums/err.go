package enums

// 2、定义Code
const (
	ErrSysBadRequest   int64 = 1100 // 请求错误
	ErrSysTokenExpired int64 = 1101 // Token失效
	ErrSysAuthFailed   int64 = 1102 // 权限校验失败
	ErrRequestLimit    int64 = 1103 // 请求频率限制
	ErrTimeout         int64 = 1104 // 请求超时
	ErrIPLimit         int64 = 1105 // IP限制
	ErrImageSizeLimit  int64 = 1106 // 图片大小限制
	ErrImageSuffix     int64 = 1107 // 图片后缀错误
)
