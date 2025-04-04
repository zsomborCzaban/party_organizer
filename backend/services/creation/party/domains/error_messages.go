package domains

const (
	PartyNotFound       = "party with the given details was not found"
	BadRequest          = "request does not contain necessary inputs"
	FailedValidation    = "request failed validation"
	InternalServerError = "internal server error"
	InvalidCredentials  = "invalid or insufficient credentials"
	DeletedUser         = "unable to find organizer, maybe the account is deleted"
	NoAccessToParty     = "you cannot access this party"
)
