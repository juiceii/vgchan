package common

// Maximum lengths of various string input fields
const (
	MaxLenName         = 50
	MaxLenEmail        = 100
	MaxLenAuth         = 50
	MaxLenPostPassword = 50
	MaxLenSubject      = 100
	MaxLenBody         = 2000
	MaxLenPassword     = 50
	MaxLenUserID       = 20
	MaxLenBoardID      = 3
	MaxLenBoardTitle   = 100
	MaxLenNotice       = 500
	MaxLenRules        = 5000
	MaxLenEightball    = 2000
)

const (
	// LenSession defines the length of an unpadded base64-encoded account
	// login session token
	LenSession = 171
)

// Commonly used errors
var (
	ErrNameTooLong         = ErrTooLong("name")
	ErrSubjectTooLong      = ErrTooLong("subject")
	ErrPostPasswordTooLong = ErrTooLong("post password")
	ErrBodyTooLong         = ErrTooLong("post body")
)

// ErrTooLong is passed, when a field exceeds the maximum string length for
// that specific field
type ErrTooLong string

func (e ErrTooLong) Error() string {
	return string(e) + " too long"
}