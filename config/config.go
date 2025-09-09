package config

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

var AppConfig *ConfigStructure

// Struct Env values
type ConfigStructure struct {
	Mode             string `mapstructure:"MODE"`
	Port             string `mapstructure:"PORT"`
	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
}

// ฟังก์ชัน bindingEnv ใช้เพื่อบอก Viper ว่า field ใน struct นี้จะต้องอ่านจาก environment ตัวไหน
func bindingEnv() error {
	// ดึง type ของ struct ConfigStructure
	fields := reflect.TypeOf(ConfigStructure{})

	// วนลูปทีละ field ใน struct
	for i := range fields.NumField() {
		// ดึงค่า tag "mapstructure" ของ field นั้น เช่น "DB_HOST"
		tag := fields.Field(i).Tag.Get("mapstructure")

		// Bind กับ environment variable โดยใช้ viper
		if err := viper.BindEnv(strings.ToLower(tag)); err != nil {
			return err
		}
	}
	return nil
}

// LoadConfig loads configuration from a .env file (for local)
func LoadConfig(path string) (*ConfigStructure, error) {
	// Set up the path and name for the optional .env file for local development.
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Attempt to read the .env file. This is now optional.
	if err := viper.ReadInConfig(); err != nil {
		// If the error is simply "config file not found", that's okay on the server.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Info: .env file not found. Using system environment variables only. This is expected on Deployment.")
		} else {
			// If it's a different kind of error (e.g., a malformed .env file),
			// we should report it as a real error.
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Binding viper env by manual.
	if err := bindingEnv(); err != nil {
		return nil, fmt.Errorf("error binding env: %w", err)
	}

	// Unmarshal all the loaded configurations (from file and/or env) into our struct.
	var configStruct ConfigStructure
	if err := viper.Unmarshal(&configStruct); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &configStruct, nil
}
