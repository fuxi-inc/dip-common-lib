package constants

const (
	Success int64 = 0 //success

	ErrnoServerError int64 = 50000

	ErrnoParamsError               int64 = 40000 //统一的参数错误
	ErrnoAuthEmpty                 int64 = 40001 //auth为空
	ErrnoAuthValidError            int64 = 40002 //auth验证异常
	ErrnoConfirmationCanNotBeEmpty int64 = 40003 //Confirmation不可为空
	ErrnoNotFoundError             int64 = 40004 //DO无效（找不到DO）
	ErrnoPermissionError           int64 = 40005 //权限错误
	ErrnoRepeatedError             int64 = 40006 //重复错误

	ErrnoDomainResolutionError    int64 = 30000 // DNS权威解析失败
	ErrnoDOAttributeNotFoundError int64 = 30001 // DO属性不存在

)
