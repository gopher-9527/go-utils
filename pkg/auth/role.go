package auth

func auth(user string, subsystem string, resource string, operation string) bool {
	// user
	// role
	// permission
	// resource
	// user -> role -> permission
	// user -> permission
	// user + resource -> role -> permission
	// subsystem -> role -> permission
	return true
}

// user + subsystem + resource + operation -> permitted
