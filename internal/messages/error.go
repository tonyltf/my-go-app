package messages

type ErrorMessage struct {
	MISSING_EXCHANGE_RATE string
}

var Error ErrorMessage

func main() {
	Error.MISSING_EXCHANGE_RATE = "Exchange rate is missing"
}
