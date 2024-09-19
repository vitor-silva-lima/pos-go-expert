package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type config struct {
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	WebServerPort       int    `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret           string `mapstructure:"JWT_SECRET"`
	JWTExpiresInSeconds int    `mapstructure:"JWT_EXPIRES_IN_SECONDS"`
	JWTTokenAuth        *jwtauth.JWTAuth
}

func LoadConfig(path string) *config {
	config := &config{}
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}
	config.JWTTokenAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)
	return config
}
