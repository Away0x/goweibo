package jwttoken

import (
  "errors"
  "github.com/dgrijalva/jwt-go"
  "goweibo/core/pkg/timeutils"
  "time"
)

type Config struct {
  SecretKey string
  AccessTokenLifeTime time.Duration
  RefreshTokenLifeTime time.Duration
}

type AppJWTClaims struct {
  jwt.StandardClaims
  UserID uint
}

// AppJWTInfo jwt token info
type AppJWTInfo struct {
  Token     string `json:"token"`
  ExpiresIn string `json:"expires_in"`
}

var (
  tokenClaimsErr   = errors.New("token claims parse error")
  tokenInvalidErr  = errors.New("token invalid error")
  tokenParseErr    = errors.New("token error")
  tokenNotFoundErr = errors.New("token not found")
  config            *Config
)

func Setup(c *Config) {
  config = c
}

func (a *AppJWTClaims) SetExpiredAt(t time.Duration) {
  now := time.Now()
  a.IssuedAt = now.Unix()
  a.ExpiresAt = now.Add(t).Unix()
}

func NewAppJWTInfo(t string, a *AppJWTClaims) *AppJWTInfo {
  return &AppJWTInfo{
    Token:     t,
    ExpiresIn: timeutils.FormatTime(time.Unix(a.ExpiresAt, 0)),
  }
}

func createToken(uid uint, t time.Duration) (string, *AppJWTClaims, error) {
  claims := &AppJWTClaims{UserID: uid}
  claims.SetExpiredAt(t)
  return createTokenWithClaims(claims)
}

func createTokenWithClaims(claims *AppJWTClaims) (string, *AppJWTClaims, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  signedToken, err := token.SignedString([]byte(config.SecretKey))
  if err != nil {
    return "", nil, err
  }

  return signedToken, claims, err
}

func VerifyToken(tokenStr string) (*AppJWTClaims, error) {
  token, err := jwt.ParseWithClaims(tokenStr, &AppJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(config.SecretKey), nil
  })

  if err != nil {
    return nil, tokenParseErr
  }

  claims, ok := token.Claims.(*AppJWTClaims)
  if !ok {
    return nil, tokenClaimsErr
  }

  if err := token.Claims.Valid(); err != nil {
    return nil, tokenInvalidErr
  }

  return claims, nil
}

func CreateToken(u uint) (accessTokenInfo *AppJWTInfo, refreshTokenInfo *AppJWTInfo, err error) {
  accessToken, accessClaims, err := createToken(u, config.AccessTokenLifeTime)
  if accessToken == "" || err != nil {
    return nil, nil, err
  }

  refreshToken, refreshTokenClaims, err := createToken(u, config.RefreshTokenLifeTime)
  if refreshToken == "" || err != nil {
    return nil, nil, err
  }

  accessTokenInfo = NewAppJWTInfo(accessToken, accessClaims)
  refreshTokenInfo = NewAppJWTInfo(refreshToken, refreshTokenClaims)
  return
}

func RefreshToken(refreshTokenStr string) (accessTokenInfo *AppJWTInfo, err error) {
  refreshClaims, err := VerifyToken(refreshTokenStr)
  if err != nil {
    return nil, err
  }

  accessToken, accessClaims, err := createToken(refreshClaims.UserID, config.AccessTokenLifeTime)
  if accessToken == "" || err != nil {
    return nil, err
  }

  return NewAppJWTInfo(accessToken, accessClaims), nil
}
