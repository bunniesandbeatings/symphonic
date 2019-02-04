package symphonic

type Config struct {
	//Port                  int    `long:"port" short:"p" description:"Server port" default:"8030" env:"PORT"`
	//ContentRoot           string `long:"content-root" short:"d" description:"content root directory" default:"./html"`
	//TokenPrivateKeyString string `long:"token-signing-key" description:"Private key used for signing tokens" env:"AUTH_TOKEN_PRIVATE_KEY"`
	//TokenPrivateKey       *ecdsa.PrivateKey
	//LoginTrackerSecret    string `long:"login-tracker-secret" description:"Secret used for hashing login tracking cookie" env:"AUTH_LOGIN_TRACKER_SECRET"`
	Debug bool `long:"debug" description:"Outputs more info than usual" env:"SYMPHONIC_DEBUG"`
	//AllowInsecure         bool   `long:"allow-insecure" description:"allows 'http' requests" env:"ALLOW_INSECURE"`
}
