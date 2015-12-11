package commands

// ConfigOptions is used to set viper defaults
var ConfigOptions = map[string]interface{}{
	"adduser": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "0.0.0.0",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": 8080,
			"help":  "Port to bind to",
		},
		"cert": map[string]interface{}{
			"value": "keys/cert.pem",
			"help":  "Certificate for https server",
		},
		"key": map[string]interface{}{
			"value": "keys/key.pem",
			"help":  "Key for https server",
		},
		"static": map[string]interface{}{
			"value": "static",
			"help":  "Directory which holds static content",
		},
		"recaptcha": map[string]interface{}{
			"value": "",
			"help":  "reCAPTCHA secret key",
		},
	},
}
