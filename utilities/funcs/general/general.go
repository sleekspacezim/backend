package general

import (
	"log"
	"math"

	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func GenerateVerificationCode() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	randomInt := rand.Intn(max-min) + min
	return randomInt
}

func ConvertIntToString(number int) string {
	return strconv.Itoa(number)
}

func ConvertStringToInt(number string) int {
	i, _ := strconv.Atoi(number)
	return i
}

func GenerateVerificationGracePeriod() time.Time {
	return time.Now().Add(time.Minute * 15)
}

func GetTimePassed(createdAt time.Time) string {
	timePassed := time.Since(createdAt)

	seconds := timePassed.Seconds()
	minutes := timePassed.Minutes()
	hours := timePassed.Hours()
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	if seconds < 60 {
		return ConvertIntToString(int(math.Floor(seconds))) + " seconds ago"
	} else if minutes < 60 {
		if math.Floor(minutes) == 1 {
			return ConvertIntToString(int(math.Floor(minutes))) + " minute ago"
		} else {
			return ConvertIntToString(int(math.Floor(minutes))) + " minutes ago"
		}
	} else if math.Floor(hours) < 24 {
		if math.Floor(hours) == 1 {
			return ConvertIntToString(int(math.Floor(hours))) + " hour ago"
		} else {
			return ConvertIntToString(int(math.Floor(hours))) + " hours ago"
		}
	} else if math.Floor(days) < 7 {
		if math.Floor(days) == 1 {
			return ConvertIntToString(int(math.Floor(days))) + " day ago"
		} else {
			return ConvertIntToString(int(math.Floor(days))) + " days ago"
		}
	} else if math.Floor(weeks) < 4 {
		if math.Floor(weeks) == 1 {
			return ConvertIntToString(int(math.Floor(weeks))) + " week ago"
		} else {
			return ConvertIntToString(int(math.Floor(weeks))) + " weeks ago"
		}
	} else if math.Floor(months) < 12 {
		if math.Floor(months) == 1 {
			return ConvertIntToString(int(math.Floor(months))) + " month ago"
		} else {
			return ConvertIntToString(int(math.Floor(months))) + " months ago"
		}
	} else {
		if math.Floor(years) == 1 {
			return ConvertIntToString(int(math.Floor(years))) + " year ago"
		} else {
			return ConvertIntToString(int(math.Floor(years))) + " years ago"
		}
	}
}

type EnvVariables struct {
	DatabaseDetails     string
	Email               string
	EmailPassword       string
	TokenSecret         string
	LocationIQToken     string
	SupabaseApiKey      string
	SupabaseReferenceId string
	AWSAcessKey         string
	AWSSecretKey        string
}

func GetEnvVariables() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	databaseDetails := os.Getenv("DATABASE_DETAILS")
	email := os.Getenv("EMAIL")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	tokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	locationIQToken := os.Getenv("LOCATION_IQ_ACCESS_TOKEN")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("AWS_SECRET_KEY")
	return EnvVariables{
		DatabaseDetails: databaseDetails,
		Email:           email,
		EmailPassword:   emailPassword,
		TokenSecret:     tokenSecret,
		LocationIQToken: locationIQToken,
		AWSAcessKey:     awsAccessKey,
		AWSSecretKey:    awsSecretKey,
	}
}
