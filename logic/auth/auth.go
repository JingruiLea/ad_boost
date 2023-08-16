package auth

import (
	"github.com/JingruiLea/ad_boost/common/logs"
	"github.com/JingruiLea/ad_boost/dal/redis_dal"
	"github.com/oceanengine/ad_open_sdk_go"
	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

type ApiOpenApiOauth2AccessTokenPostRequestExample struct {
	Oauth2AccessTokenRequest Oauth2AccessTokenRequest `json:"Oauth2AccessTokenRequest,omitempty"`
}

const APPID = 1773197267842080
const Secret = "88ac6c359042864c88fd2e7ec93ff785faa234ee"
const AuthCode = "8243c18d89cde3052bc0e1806c493135faf6a625"

const AccessTokenKey = "access_token"
const RefreshTokenKey = "refresh_token"

func Auth(ctx context.Context) (accessToken, refreshToken string, err error) {
	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	var request ApiOpenApiOauth2AccessTokenPostRequestExample
	request.Oauth2AccessTokenRequest.AppId = PtrInt64(APPID)
	request.Oauth2AccessTokenRequest.Secret = Secret
	request.Oauth2AccessTokenRequest.AuthCode = AuthCode
	resp, _, err := apiClient.Oauth2AccessTokenApi().
		Post(ctx).
		Oauth2AccessTokenRequest(request.Oauth2AccessTokenRequest).
		Execute()
	logs.CtxDebugf(ctx, ToJsonString(resp))
	if err != nil {
		logs.CtxErrorf(ctx, "Auth apiClient.Oauth2AccessTokenApi error: %v", err)
		return "", "", err
	}
	err = saveAtRtToRedis(ctx, *resp.Data.AccessToken, *resp.Data.RefreshToken)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAccessToken saveAtRtToRedis error: %v", err)
		return *resp.Data.AccessToken, *resp.Data.RefreshToken, err
	}
	return *resp.Data.AccessToken, *resp.Data.RefreshToken, nil
}

type ApiOpenApiOauth2RefreshTokenPostRequestExample struct {
	Oauth2RefreshTokenRequest Oauth2RefreshTokenRequest `json:"Oauth2RefreshTokenRequest,omitempty"`
}

func RefreshToken(ctx context.Context) (at, rt string, err error) {
	configuration := config.NewConfiguration()
	apiClient := ad_open_sdk_go.Init(configuration)
	apiClient.SetLogEnable(true)

	rtOld, err := redis_dal.GetRedisClient().Get(ctx, RefreshTokenKey).Result()
	if err != nil {
		logs.CtxErrorf(ctx, "GetAccessToken Get error: %v", err)
		return "", "", err
	}

	var request ApiOpenApiOauth2RefreshTokenPostRequestExample
	request.Oauth2RefreshTokenRequest.AppId = PtrInt64(APPID)
	request.Oauth2RefreshTokenRequest.Secret = Secret
	request.Oauth2RefreshTokenRequest.RefreshToken = rtOld
	resp, _, err := apiClient.Oauth2RefreshTokenApi().
		Post(ctx).
		Oauth2RefreshTokenRequest(request.Oauth2RefreshTokenRequest).
		Execute()
	logs.CtxDebugf(ctx, ToJsonString(resp))
	if err != nil {
		logs.CtxErrorf(ctx, "RefreshToken apiClient.Oauth2RefreshTokenApi error: %v", err)
		return
	}
	err = saveAtRtToRedis(ctx, *resp.Data.AccessToken, *resp.Data.RefreshToken)
	if err != nil {
		logs.CtxErrorf(ctx, "GetAccessToken saveAtRtToRedis error: %v", err)
		return *resp.Data.AccessToken, *resp.Data.RefreshToken, err
	}
	return *resp.Data.AccessToken, *resp.Data.RefreshToken, nil
}

func GetAccessToken(ctx context.Context) (token string, err error) {
	at, err := redis_dal.GetRedisClient().Get(ctx, AccessTokenKey).Result()
	if errors.Is(err, redis.Nil) {
		at, _, err := Auth(ctx)
		if err != nil {
			logs.CtxErrorf(ctx, "GetAccessToken Auth error: %v", err)
			return "", err
		}
		return at, nil
	}
	if err != nil {
		logs.CtxErrorf(ctx, "GetAccessToken Get error: %v", err)
		return "", err
	}
	return at, nil
}

func saveAtRtToRedis(ctx context.Context, at, rt string) error {
	err := redis_dal.GetRedisClient().Set(ctx, AccessTokenKey, at, 24*time.Hour).Err()
	if err != nil {
		logs.CtxErrorf(ctx, "saveAtRtToRedis Set AccessToken error: %v", err)
		return err
	}
	err = redis_dal.GetRedisClient().Set(ctx, RefreshTokenKey, rt, 24*time.Hour*30).Err()
	if err != nil {
		logs.CtxErrorf(ctx, "saveAtRtToRedis Set RefreshToken error: %v", err)
		return err
	}
	return nil
}
