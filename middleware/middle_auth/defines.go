package middle_auth

const (
	// admin token
	AdminJwtKey    = "AdminAuthorization" // 取jwtKey
	AdminJwtSecret = "ADMIN_JWT_SECRET"   // jwt-setcret
	AdminUserKey   = "admin_id"           // 存在gin.Context的key

	// inner token
	InnerJwtKey    = "Inner-Authorization"
	InnerJwtSecret = "INNER_JWT_SECRET"
	InnerSystemKey = "system_id" // 存在gin.Context的key

	// user token
	UserJwtKey    = "Token"
	UserJwtSecret = "INNER_JWT_SECRET"
	UserAuthKey   = "user_id" // 存在gin.Context的key
)
