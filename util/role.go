package util

const (
	HOKAGE   = "hokage"
	AKATSUKI = "akatsuki"
)

func IsValidRole(role string) bool {
	switch role {
	case HOKAGE, AKATSUKI:
		return true
	default:
		return false
	}
}
