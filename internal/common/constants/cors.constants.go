package constants

func AllowedMethods() []string {
	return []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
}
func AllowedHeaders() []string {
	return []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}
}
func AllowedOrigins() []string {
	return []string{"https://*", "http://*"}
}
func ExposedHeaders() []string {
	return []string{"Link"}
}

const (
	MaxAge           int  = 300
	AllowCredentials bool = true
)
