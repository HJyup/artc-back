package auth

import (
	"artc-back/config"
	"artc-back/types"
	"artc-back/utils"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

type contextKey string

const UserKey contextKey = "userID"
const IsAcceptedKey contextKey = "isAccepted"
const IsReviewerKey contextKey = "isReviewer"

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)

		u, err := store.GetUserById(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		if u.IsAccepted != claims["isAccepted"].(bool) {
			log.Printf("isAccepted mismatch: %v", err)
			permissionDenied(w)
			return
		}

		if u.IsReviewer != claims["isReviewer"].(bool) {
			log.Printf("isReviewer mismatch: %v", err)
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		ctx = context.WithValue(ctx, IsAcceptedKey, u.IsAccepted)
		ctx = context.WithValue(ctx, IsReviewerKey, u.IsReviewer)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func CreateJWT(secret []byte, userID string, isAccepted bool, isReviewer bool) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     userID,
		"isAccepted": isAccepted,
		"isReviewer": isReviewer,
		"expiresAt":  time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func _(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
