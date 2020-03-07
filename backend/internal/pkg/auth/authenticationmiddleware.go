package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// AuthenticationMiddleware extracts the user authentication from the request.
func AuthenticationMiddleware(issuer string, audience string) *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'iss' claim
			// fmt.Println("iss in token: ", token.Claims.(jwt.MapClaims)["iss"], ", iss expected: ", issuer)
			// fmt.Println("aud in token: ", token.Claims.(jwt.MapClaims)["aud"], ", aud expected: ", audience)

			verifyAudience(token, audience)

			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, true)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			// Verify 'aud' claim
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(audience, true)
			if !checkAud {
				return token, errors.New("Invalid audience.")
			}

			cert, err := getPemCert(issuer, token)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})
}

func verifyAudience(token *jwt.Token, expectedAudience string) bool {
	// func verifyIss(iss string, cmp string) bool {
	// 	if subtle.ConstantTimeCompare([]byte(iss), []byte(cmp)) != 0 {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }
	fmt.Println(token.Claims.(jwt.MapClaims)["aud"][0])
	return false
}

func getPemCert(issuer string, token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(issuer + ".well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
