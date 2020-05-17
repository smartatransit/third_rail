package api

type Options struct {
	TwitterClientID     string `long:"twitter-client-id" env:"TWITTER_CLIENT_ID" description:"the client id for the twitter account"`
	TwitterClientSecret string `long:"twitter-client-secret" env:"TWITTER_CLIENT_SECRET" description:"the client secret for the twitter account"`
	MartaAPIKey         string `long:"marta-api-key" env:"MARTA_API_KEY" description:"marta api key"`
	TwitterCacheTTL     int    `long:"twitter-cache-ttl" env:"TWITTER_CACHE_TTL" default:"15" description:"how long we keep the twitter responses" required:"true"`
	MartaCacheTTL       int    `long:"marta-cache-ttl" env:"MARTA_CACHE_TTL" default:"15" description:"how long we keep the marta responses" required:"true"`
	DbHost              string `long:"db-host" env:"DB_HOST" description:"the host for the smarta database"`
	DbPort              string `long:"db-port" env:"DB_PORT" description:"the host port for the smarta database"`
	DbName              string `long:"db-name" env:"DB_NAME" description:"the name for the smarta database"`
	DbUsername          string `long:"db-username" env:"DB_USERNAME" description:"the username for the smarta database"`
	DbPassword          string `long:"db-password" env:"DB_PASSWORD" description:"the password for the smarta database"`
}
