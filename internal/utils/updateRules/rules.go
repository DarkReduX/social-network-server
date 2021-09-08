package updateRules

var ProfileChangeRule = map[string]bool{
	"id":              false,
	"username":        false,
	"password":        true,
	"avatar_link":     true,
	"last_activity":   false,
	"created_at":      false,
	"created_from_ip": false,
	"deleted_at":      false,
	"is_activate":     false,
}
