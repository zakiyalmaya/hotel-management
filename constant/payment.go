package constant

type (
	PaymentMethod int
	PaymentStatus int
)

const (
	PaymentMethodCreditCard PaymentMethod = iota + 1
	PaymentMethodBankTransfer
	PaymentMethodCash
)

var mapPaymentMethod = map[PaymentMethod]string{
	PaymentMethodCreditCard:   "Credit Card",
	PaymentMethodBankTransfer: "Bank Transfer",
	PaymentMethodCash:         "Cash",
}

func (pm PaymentMethod) Enum() string {
	if val, ok := mapPaymentMethod[pm]; ok {
		return val
	}

	return "Unknown"
}

func (pm PaymentMethod) Validation() bool {
	if _, ok := mapPaymentMethod[pm]; !ok {
		return false
	}

	return true
}

const (
	PaymentStatusPending   PaymentStatus = iota + 1 // Payment has been initiated but not completed yet
	PaymentStatusCompleted                          // Payment has been successfully processed
	PaymentStatusFailed                             // Payment attempt failed due to an error
	PaymentStatusCanceled                           // Payment has been canceled by the customer or the system
	PaymentStatusRefunded                           // Payment has been refunded to the customer
)

var mapPayemntStatus = map[PaymentStatus]string{
	PaymentStatusPending:   "Pending",
	PaymentStatusCompleted: "Completed",
	PaymentStatusFailed:    "Failed",
	PaymentStatusCanceled:  "Canceled",
	PaymentStatusRefunded:  "Refunded",
}

func (ps PaymentStatus) Enum() string {
	if val, ok := mapPayemntStatus[ps]; ok {
		return val
	}

	return "Unknown"
}

func (ps PaymentStatus) Validation() bool {
	if _, ok := mapPayemntStatus[ps]; !ok {
		return false
	}

	return true
}