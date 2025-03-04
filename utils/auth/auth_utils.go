package auth_utils

import (
	"api/config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/timewise-team/timewise-models/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
	"regexp"
	"time"
)

type GoogleOauthData struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

var cfg, err = config.LoadConfig()

var GoogleOauth = oauth2.Config{
	ClientID:     cfg.GoogleOauth.ClientID,
	ClientSecret: cfg.GoogleOauth.ClientSecret,
	RedirectURL:  "",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

type TokenInfo struct {
	ExpiresIn int64 `json:"expires_in"`
	// other fields as necessary
}

func VerifyGoogleToken(code string) ([]byte, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=" + code)
	if err != nil {
		return nil, err
	}

	// this is done to prevent memory leakage
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// returns data of verified google user
	return data, nil
}
func CheckGoogleTokenExpiry(accessToken string) error {
	response, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v1/tokeninfo?access_token=%s", accessToken))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("invalid or expired token")
	}

	var tokenInfo TokenInfo
	if err := json.NewDecoder(response.Body).Decode(&tokenInfo); err != nil {
		return err
	}

	// Check if the token is expired
	if tokenInfo.ExpiresIn <= 0 {
		return errors.New("token has expired")
	}

	return nil
}
func GenerateJWTToken(user models.TwUser, secretKey string) (string, int, error) {
	// Định nghĩa thời gian hết hạn cho token (ví dụ: 2 giờ)
	expirationTime := time.Now().Add(168 * time.Hour).Unix()

	// Tạo claims cho JWT
	claims := jwt.MapClaims{
		"userid": user.ID,
		"email":  user.Email,
		"exp":    expirationTime,
	}

	// Tạo token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token với secretKey
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", 0, err
	}

	// Tính thời gian hết hạn
	expiresIn := int(expirationTime - time.Now().Unix())

	// Trả về token, thời gian hết hạn
	return tokenString, expiresIn, nil
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func GenerateInvitationToken(workspaceId int, action string, secretKey string, email string, role string, isMember bool) (string, error) {
	claims := jwt.MapClaims{
		"email":        email,
		"workspace_id": workspaceId,
		"role":         role,
		"action":       action, // accept hoặc decline
		"is_member":    isMember,
		"exp":          time.Now().Add(24 * time.Hour).Unix(), // Token có thời hạn 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateScheduleInvitationToken(workspaceUserId int, action string, secretKey string, scheduleId int) (string, error) {
	claims := jwt.MapClaims{
		"schedule_id":       scheduleId,
		"workspace_user_id": workspaceUserId,
		"action":            action,                                // accept hoặc decline
		"exp":               time.Now().Add(24 * time.Hour).Unix(), // Token có thời hạn 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ParseInvitationToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
func GenerateLinkEmailToken(currentUid string, email string, action string, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": currentUid,
		"email":   email,
		"action":  action,                                // accept hoặc decline
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token có thời hạn 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
