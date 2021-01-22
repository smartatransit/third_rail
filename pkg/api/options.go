package api

type Options struct {
	TwitterClientID     string `long:"twitter-client-id" env:"TWITTER_CLIENT_ID" description:"the client id for the twitter account"`
	TwitterClientSecret string `long:"twitter-client-secret" env:"TWITTER_CLIENT_SECRET" description:"the client secret for the twitter account"`
	MartaAPIKey         string `long:"marta-api-key" env:"MARTA_API_KEY" description:"marta api key"`
	TwitterCacheTTL     int    `long:"twitter-cache-ttl" env:"TWITTER_CACHE_TTL" default:"15" description:"how long we keep the twitter responses" required:"true"`
	MartaCacheTTL       int    `long:"marta-cache-ttl" env:"MARTA_CACHE_TTL" default:"15" description:"how long we keep the marta responses" required:"true"`
	DBConnectionString  string `long:"db-connection-string" env:"DB_CONNECTION_STRING"`
	AdminAPIKey         string `long:"admin-api-key" env:"ADMIN_API_KEY" description:"admin api key"`
	Servicedomain       string `long:"service-domain" env:"SERVICE_DOMAIN" description:"the domain at which the service is served" default:"third-rail.services.smartatransit.com"`
	RailRunner          bool   `long:"rail-runner" env:"RAIL_RUNNER" description:"enable the rail runner"`
}
