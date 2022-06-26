package auth

type AuthorizationServiceInterface interface {
	IsAllowed(action string, subject string) bool
	HasRole(roles ...string) bool
}
