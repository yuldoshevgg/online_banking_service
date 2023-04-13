package http

// Status ...
type Status struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

var (
	OK = Status{
		Code:        200,
		Status:      "OK",
		Description: "The request has succeeded",
	}
	Created = Status{
		Code:        201,
		Status:      "CREATED",
		Description: "The request has been fulfilled and has resulted in one or more new resources being created",
	}
	NoContent = Status{
		Code:        204,
		Status:      "NO_CONTENT",
		Description: "There is no content to send for this request, but the headers may be useful",
	}
	BadEnvironment = Status{
		Code:        400,
		Status:      "BAD_ENVIRONMENT",
		Description: "The service has an invalid environment value",
	}
	BadRequest = Status{
		Code:        400,
		Status:      "BAD_REQUEST",
		Description: "The server could not understand the request due to invalid syntax",
	}
	InvalidArgument = Status{
		Code:        400,
		Status:      "INVALID_ARGUMENT",
		Description: "Invalid argument value passed",
	}
	Unauthorized = Status{
		Code:        401,
		Status:      "UNAUTHORIZED",
		Description: "...",
	}
	Forbidden = Status{
		Code:        403,
		Status:      "FORBIDDEN",
		Description: "...",
	}
	TooManyRequests = Status{
		Code:        429,
		Status:      "TOO_MANY_REQUESTS",
		Description: "The user has sent too many requests in a given amount of time",
	}
	InternalServerError = Status{
		Code:        500,
		Status:      "INTERNAL_SERVER_ERROR",
		Description: "The server encountered an unexpected condition that prevented it from fulfilling the request",
	}
	GRPCError = Status{
		Code:        500,
		Status:      "GRPC_ERROR",
		Description: "The gRPC request failed",
	}
)
