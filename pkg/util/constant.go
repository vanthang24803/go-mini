package util

import "time"

const ACCESS_TOKEN = "access_token"
const REFRESH_TOKEN = "refresh_token"

const JWT_SECRET_KEY = "JWT_SECRET_KEY"
const JWT_REFRESH_KEY = "JWT_REFRESH_KEY"
const ACCESS_TOKEN_EXPIRATION = time.Hour * 24 * 7
const REFRESH_TOKEN_EXPIRATION = time.Hour * 24 * 30

const ENV = "GO_ENV"
const APP_NAME = "APP_NAME"
const APP_PORT = "PORT"

const PRODUCTION = "production"
const DEVELOPMENT = "development"
const TEST = "test"

const DB_HOST = "DB_HOST"
const DB_PORT = "DB_PORT"
const DB_USERNAME = "DB_USERNAME"
const DB_PASSWORD = "DB_PASSWORD"
const DB_NAME = "DB_NAME"

const REDIS_HOST = "REDIS_HOST"
const REDIS_PORT = "REDIS_PORT"
const REDIS_PASSWORD = "REDIS_PASSWORD"
