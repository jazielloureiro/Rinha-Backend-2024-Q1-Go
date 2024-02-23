package helper

type ServiceError int

const (
	AccountNotFoundError     ServiceError = 0
	InsufficientBalanceError ServiceError = 1
)

var errorMessages = map[ServiceError]string{
	AccountNotFoundError:     "there's not account for the given id",
	InsufficientBalanceError: "insufficient balance",
}

func (e ServiceError) Error() string {
	return errorMessages[e]
}
