package pesapal

// PesapalError represents the error object returned by Pesapal v3 API responses.
type PesapalError struct {
	ErrorType   interface{} `json:"error_type"`
	Code        interface{} `json:"code"`
	Message     string      `json:"message"`
	CallBackURL string      `json:"call_back_url"`
}

// AuthRequest is the payload for Pesapal v3 authentication (RequestToken).
type AuthRequest struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

// AuthResponse is the response returned by Pesapal v3 authentication.
type AuthResponse struct {
	Token      string        `json:"token"`
	ExpiryDate string        `json:"expiryDate"`
	Error      *PesapalError `json:"error"`
	Status     string        `json:"status"`
	Message    string        `json:"message"`
}

// IPNRegistrationRequest registers an IPN URL with Pesapal v3.
type IPNRegistrationRequest struct {
	URL                 string `json:"url"`
	IPNNotificationType string `json:"ipn_notification_type"`
}

// IPNRegistrationResponse is returned after registering an IPN URL.
type IPNRegistrationResponse struct {
	URL                            string        `json:"url"`
	CreatedDate                    string        `json:"created_date"`
	IPNID                          string        `json:"ipn_id"`
	NotificationType               int           `json:"notification_type"`
	IPNNotificationTypeDescription string        `json:"ipn_notification_type_description"`
	IPNStatus                      int           `json:"ipn_status"`
	IPNStatusDescription           string        `json:"ipn_status_description"`
	Error                          *PesapalError `json:"error"`
	Status                         string        `json:"status"`
}

// BillingAddress is the customer address object used in SubmitOrderRequest.
type BillingAddress struct {
	PhoneNumber  string `json:"phone_number,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	MiddleName   string `json:"middle_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Line1        string `json:"line_1,omitempty"`
	Line2        string `json:"line_2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	PostalCode   int    `json:"postal_code,omitempty"`
	ZipCode      int    `json:"zip_code,omitempty"`
}

// SubmitOrderRequest creates a payment order in Pesapal v3.
type SubmitOrderRequest struct {
	ID              string         `json:"id"`
	Currency        string         `json:"currency"`
	Amount          float64        `json:"amount"`
	Description     string         `json:"description"`
	CallbackURL     string         `json:"callback_url"`
	RedirectMode    string         `json:"redirect_mode,omitempty"`
	CancellationURL string         `json:"cancellation_url,omitempty"`
	NotificationID  string         `json:"notification_id"`
	Branch          string         `json:"branch,omitempty"`
	BillingAddress  BillingAddress `json:"billing_address"`
}

// SubmitOrderResponse is returned after submitting an order.
type SubmitOrderResponse struct {
	OrderTrackingID   string        `json:"order_tracking_id"`
	MerchantReference string        `json:"merchant_reference"`
	RedirectURL       string        `json:"redirect_url"`
	Error             *PesapalError `json:"error"`
	Status            string        `json:"status"`
	Message           string        `json:"message"`
}

// TransactionStatusResponse is returned by GetTransactionStatus.
type TransactionStatusResponse struct {
	PaymentMethod            string        `json:"payment_method"`
	Amount                   float64       `json:"amount"`
	CreatedDate              string        `json:"created_date"`
	ConfirmationCode         string        `json:"confirmation_code"`
	PaymentStatusDescription string        `json:"payment_status_description"`
	Description              string        `json:"description"`
	Message                  string        `json:"message"`
	PaymentAccount           string        `json:"payment_account"`
	CallBackURL              string        `json:"call_back_url"`
	StatusCode               int           `json:"status_code"`
	MerchantReference        string        `json:"merchant_reference"`
	PaymentStatusCode        string        `json:"payment_status_code"`
	Currency                 string        `json:"currency"`
	Error                    *PesapalError `json:"error"`
	Status                   string        `json:"status"`
}
