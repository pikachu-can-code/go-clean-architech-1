package common

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Env struct {
	DBConnectionStr   string
	DBPassword        string
	HttpPort          string
	S3BucketName      string
	S3Region          string
	S3APIKey          string
	S3Secret          string
	S3Domain          string
	SecretKeyJWT      string
	FirebaseService   string
	BaseEmailPassword string
	DefaultEndpoint   string
}

func checkEnvFile(file string) error {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	return err
}

func getEnvVar(key string) string {
	return viper.GetString(key)
}

func Init(dirFile string) Env {
	valDeployed := os.Getenv("DEPLOYED")
	isDeployed := false
	var env Env
	if valDeployed != "" {
		var err error
		isDeployed, err = strconv.ParseBool(valDeployed)
		if err != nil {
			log.Fatalf("error get env: %v", err)
		}
	}
	if !isDeployed {
		if err := checkEnvFile(dirFile); err != nil {
			log.Fatalf("Error read env %s", err)
		}

		env.DBConnectionStr = getEnvVar("DB_CONNECTION")
		env.DBPassword = getEnvVar("DB_PASSWORD")
		env.HttpPort = getEnvVar("HTTP_PORT")
		env.S3BucketName = getEnvVar("S3_BUCKET_NAME")
		env.S3Region = getEnvVar("S3_REGION")
		env.S3APIKey = getEnvVar("S3_API_KEY")
		env.S3Secret = getEnvVar("S3_SECRET")
		env.S3Domain = getEnvVar("S3_DOMAIN")
		env.SecretKeyJWT = getEnvVar("SECRET_KEY_JWT")
		env.FirebaseService = getEnvVar("FIREBASE_SERVICE")
		env.BaseEmailPassword = getEnvVar("EMAIL")
		env.DefaultEndpoint = getEnvVar("DEFAULT_ENDPOINT")
	} else {
		env.DBConnectionStr = os.Getenv("DB_CONNECTION")
		env.DBPassword = os.Getenv("DB_PASSWORD")
		env.HttpPort = os.Getenv("HTTP_PORT")
		env.S3BucketName = os.Getenv("S3_BUCKET_NAME")
		env.S3Region = os.Getenv("S3_REGION")
		env.S3APIKey = os.Getenv("S3_API_KEY")
		env.S3Secret = os.Getenv("S3_SECRET")
		env.S3Domain = os.Getenv("S3_DOMAIN")
		env.SecretKeyJWT = os.Getenv("SECRET_KEY_JWT")
		env.FirebaseService = os.Getenv("FIREBASE_SERVICE")
		env.BaseEmailPassword = os.Getenv("EMAIL")
		env.DefaultEndpoint = os.Getenv("DEFAULT_ENDPOINT")

		// dbPass, err := ioutil.ReadFile(env.DBPassword)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// env.DBPassword = string(dbPass)
		// s3, err := ioutil.ReadFile(env.S3Secret)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// env.S3Secret = string(s3)
		// jwt, err := ioutil.ReadFile(env.SecretKeyJWT)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// env.SecretKeyJWT = string(jwt)
		// mailPass, err := ioutil.ReadFile(env.BaseEmailPassword)
		// if err != nil {
		// 	log.Fatal(mailPass)
		// }
		// env.BaseEmailPassword = string(mailPass)
	}
	return env
}
