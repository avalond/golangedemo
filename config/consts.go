package config

func ConstonsKey(key string) string {
	switch key {
	case "server_port":
		return "22222"

	default:
		return ""
	}
}
