package main

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Configure sets command line options
func Configure() {
	viper.SetConfigName("config")
	viper.SetEnvPrefix("admin")
	viper.AutomaticEnv()
	viper.SetDefault("addr", "0.0.0.0")
	flag.String("addr", "0.0.0.0", "Address to listen on")
	viper.BindPFlag("addr", flag.Lookup("addr"))
	viper.SetDefault("port", 8080)
	flag.Int("port", 8080, "Port to listen on")
	viper.BindPFlag("port", flag.Lookup("port"))
	viper.SetDefault("cert", "keys/cert.pem")
	flag.String("cert", "keys/cert.pem", "Cert to use")
	viper.BindPFlag("cert", flag.Lookup("cert"))
	viper.SetDefault("key", "keys/key.pem")
	flag.String("key", "keys/key.pem", "Key to use")
	viper.BindPFlag("key", flag.Lookup("key"))
	viper.SetDefault("static", "static")
	flag.String("static", "static", "Static content dir")
	viper.BindPFlag("static", flag.Lookup("static"))

	flag.Parse()
}
