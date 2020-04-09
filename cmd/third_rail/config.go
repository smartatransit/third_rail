package main

type options struct {
	TwitterClientID     string `long:"twitter-client-id" env:"TWITTER_CLIENT_ID" description:"the client id for the twitter acount"`
	TwitterClientSecret string `long:"twitter-client-secret" env:"TWITTER_CLIENT_SECRET" description:"the client secret for the twitter acount"`
	MartaAPIKey         string `long:"marta-api-key" env:"MARTA_API_KEY" description:"marta api key"`
	TwitterCacheTTL     int    `long:"twitter-cache-ttl" env:"TWITTER_CACHE_TTL" default:"15" description:"how long we keep the twitter responses" required:"true"`
	MartaCacheTTL       int    `long:"marta-cache-ttl" env:"MARTA_CACHE_TTL" default:"15" description:"how long we keep the marta responses" required:"true"`
}
