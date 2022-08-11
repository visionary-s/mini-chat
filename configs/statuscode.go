package configs

const (
	OK                 = 200
	NotLoggedIn        = 1000
	ParameterIllegal   = 1001
	UnauthorizedUserId = 1002
	Unauthorized       = 1003
	ServerError        = 1004
	NotData            = 1005
	ModelAddError      = 1006
	ModelDeleteError   = 1007
	ModelStoreError    = 1008
	OperationFailure   = 1009
	RoutingNotExist    = 1010
)

func ErrorMessage(code uint32, message string) string {
	var statusMessage string
	msgMap := map[uint32]string{
		OK:                 "Success",
		NotLoggedIn:        "User not log in",
		ParameterIllegal:   "Invalid parameter",
		UnauthorizedUserId: "Unauthorized user",
		Unauthorized:       "Unauthorized",
		NotData:            "No data",
		ServerError:        "System Error",
		ModelAddError:      "Failed to add",
		ModelDeleteError:   "Failed to delete",
		ModelStoreError:    "Failed to save",
		OperationFailure:   "Failed to execute",
		RoutingNotExist:    "Missing routing",
	}

	if message == "" {
		if value, ok := msgMap[code]; ok {
			// 存在
			statusMessage = value
		} else {
			statusMessage = "Unknown ERR!"
		}
	} else {
		statusMessage = message
	}

	return statusMessage
}
