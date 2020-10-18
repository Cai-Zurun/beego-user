package utils

const (
	SUCCESS			= 200
	FAIL			= 400
	UNAUTHORIZED	= 401
)

func Response(code int, message string, data interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"code": code,
		"message": message,
		"data": data,
	}
	return res
}
