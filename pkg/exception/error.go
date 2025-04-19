package exception

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ERROR_CODE_INVALID_TOKEN = &Error{
		Code:    1000,
		Message: "Invalid token",
	}
	ERROR_CODE_EXPIRED_TOKEN = &Error{
		Code:    1001,
		Message: "Expired token",
	}
	ERROR_CODE_UNAUTHORIZED = &Error{
		Code:    1002,
		Message: "Unauthorized",
	}
	ERROR_CODE_FORBIDDEN = &Error{
		Code:    1003,
		Message: "Forbidden",
	}
	ERROR_INTERNAL_SERVER = &Error{
		Code:    1004,
		Message: "Internal server error",
	}

	// ERROR_CODE_NOT_FOUND is returned when the requested resource does not exist.
	ERROR_CODE_NOT_FOUND = &Error{
		Code:    2000,
		Message: "Resource not found",
	}
	ERROR_CODE_BAD_REQUEST = &Error{
		Code:    2001,
		Message: "Bad request",
	}
	ERROR_CODE_INTERNAL_ERROR = &Error{
		Code:    2002,
		Message: "Internal server error",
	}
	ERROR_CODE_SERVICE_UNAVAILABLE = &Error{
		Code:    2003,
		Message: "Service temporarily unavailable",
	}

	ERROR_CODE_VALIDATION_FAILED = &Error{
		Code:    3000,
		Message: "Validation failed",
	}
	ERROR_CODE_RATE_LIMIT_EXCEEDED = &Error{
		Code:    3001,
		Message: "Rate limit exceeded",
	}
	ERROR_GENERATE_TOKEN = &Error{
		Code:    3001,
		Message: "Rate limit exceeded",
	}

	ERROR_EMAIL_EXISTED = &Error{
		Code:    4000,
		Message: "Email already existed",
	}
	ERROR_NO_DOCUMENT = &Error{
		Code:    4001,
		Message: "No document found",
	}
	ERROR_HASH_PASSWORD = &Error{
		Code:    4002,
		Message: "Hash password failed",
	}
	ERROR_COMPARE_PASSWORD = &Error{
		Code:    4003,
		Message: "Compare password failed",
	}
	ERROR_INVALID_PASSWORD = &Error{
		Code:    4004,
		Message: "Invalid password",
	}
	ERROR_CODE_USERNAME_EXISTED = &Error{
		Code:    4005,
		Message: "Username already existed",
	}
	ERROR_INVALID_CREDENTIAL = &Error{
		Code:    4006,
		Message: "Invalid credential",
	}
	ERROR_INSERT_TOKEN = &Error{
		Code:    4007,
		Message: "Insert token failed",
	}
	ERROR_INVALID_USER_ID = &Error{
		Code:    4008,
		Message: "Invalid user id",
	}
	ERROR_DELETE_TOKEN = &Error{
		Code:    4009,
		Message: "Delete token failed",
	}
	ERROR_USER_NOT_FOUND = &Error{
		Code:    4010,
		Message: "User not found",
	}
)
