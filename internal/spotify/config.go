package spotify

type Config struct {
	ID     string `mapstruture:"ID" env:"ID"`
	Secret string `mapstruture:"SECRET" env:"SECRET"`
}
