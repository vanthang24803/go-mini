package constant

import "time"

const (
	ACCESS_TOKEN  = "access_token"
	REFRESH_TOKEN = "refresh_token"
	VERSION       = 1

	ROLE_USER                = "user"
	ROLE_ADMIN               = "admin"
	ROLE_ROOT                = "root"
	JWT_SECRET_KEY           = "JWT_SECRET_KEY"
	JWT_REFRESH_KEY          = "JWT_REFRESH_KEY"
	ACCESS_TOKEN_EXPIRATION  = time.Hour * 24 * 7
	REFRESH_TOKEN_EXPIRATION = time.Hour * 24 * 30

	ENV      = "GO_ENV"
	APP_NAME = "APP_NAME"
	APP_PORT = "PORT"

	PRODUCTION  = "production"
	DEVELOPMENT = "development"
	TEST        = "test"

	DB_HOST          = "DB_HOST"
	DB_PORT          = "DB_PORT"
	DB_USERNAME      = "DB_USERNAME"
	DB_PASSWORD      = "DB_PASSWORD"
	MONGODB_DATABASE = "MONGODB_DATABASE"
	MONGODB_URI      = "MONGODB_URI"

	REDIS_HOST     = "REDIS_HOST"
	REDIS_PORT     = "REDIS_PORT"
	REDIS_PASSWORD = "REDIS_PASSWORD"
)
