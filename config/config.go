package config

import "flag"
import "os"
import "strings"

var defaults map[string]string

func config() {
	defaults = make(map[string]string)

	defaults["port"] = "8080"
}

func toEnv(key string) string {
	// normalize the key we look up
	key = strings.ToUpper(key)
	key = strings.Replace(key, "-", "_", -1)

	return key
}

func toArg(key string) string {
	key = strings.ToLower(key)
	key = strings.Replace(key, "_", "-", -1)

	return key
}

// Get fetch a provided config variable or its default value
func Get(key string) string {
	// Look for the key in an environment variable
	envVar := os.Getenv(toEnv(key))

	// Look for the key in a flag
	flagVar := flag.String(toArg(key), "", "port on which to listen for HTTP requests")
	flag.Parse()

	// The variable provided by a flag gets precedence
	if *flagVar != "" {
		return *flagVar
	}

	// No flag found, try the environment variable
	if envVar != "" {
		return envVar
	}

	// No flag or environment var, just return the default
	return defaults[key]
}
