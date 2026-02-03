package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId, // for some reason, although it is int64, userId will be stored as float64. So we have to convert it back while retrieving.
			"email":  email,
			"exp":    time.Now().Add(time.Hour * 2).Unix(), //token will expire after 2 hours.
			// expiry must be in unix time i.e the number 
			// of seconds elapsed since January 1, 1970 UTC
			// we must use Unix() for valid exp.
		},
	)

	//Getenv is used to access contents of the .env file.
	secretKey := os.Getenv("JWT_TOKEN")

	// converts the token into string form and returns.
	// SignedString() accepts any type of input key (converted to byte slice)
	return token.SignedString([]byte(secretKey))
}

// Verifies the token, if failed --> returns error,
// if passed returns the userId of the user that logged in.
func VerifyJWT(unverifiedToken string) (int64, error) {

	parsedToken, err := jwt.Parse(
		unverifiedToken,
		// this function is internally used by jwt.Parse() to obtain the
		// secretKey that we used to sign that token.
		func(token *jwt.Token) (interface{}, error) {
			// as an extra security step, we should also verify if the token was
			// signed using the same method we used to sign the token.
			// we signed using signingMethodHS256, which belongs to the family *jwt.SigningMethodHMAC
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("Unexpected signing method detected!!!...")
			}

			secretKey := os.Getenv("JWT_TOKEN")

			if secretKey == "" {
				fmt.Println("jwt token is not set!!")
			}

			//we need the key in byte converted mode
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		fmt.Println("jwt.Parse error!")
		return 0, errors.New("Couldnot parse Token...")
	}

	// at this point we know that the signing method and the secret key is valid.
	//we can use the parsedToken.Valid field to know if the token is valid or not.
	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, errors.New("Invalid token!")
	}

	// If we want we can use the parsedToken.Claims field and check if it is of the more
	// specific type jwt.MapClaims that we used. Then we can extract the result in a variable,
	// and through that can access the contents of the token. (in this case email, userId and exp)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token")
	}

	// userId is stored as float64, so we convert it to int64 and then return.
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
