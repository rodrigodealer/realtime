package services

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/rodrigodealer/realtime/models"
)

const (
	AppID     = "1203725573077981"
	AppSecret = "7fb858f4301f919ea56b97e70bbfe965"
	Fields    = "id,name,picture{url}"
)

func GetUid(entry models.FacebookUpdateEntry, token *models.FacebookToken, client FacebookClient, parentSpan opentracing.Span) FbResponse {
	if parentSpan != nil {
		span := opentracing.StartSpan("Get Uid", opentracing.ChildOf(parentSpan.Context()))
		defer span.Finish()
	}
	log.Printf("Getting information for: %s", entry.UID)
	response, _ := client.GetRequest(FacebookUIDUrl(entry.UID, token.Token))
	return response
}

func GetToken(client FacebookClient, parentSpan opentracing.Span) *models.FacebookToken {
	if parentSpan != nil {
		span := opentracing.StartSpan("Get token", opentracing.ChildOf(parentSpan.Context()))
		defer span.Finish()
	}

	response, _ := client.GetRequest(FacebookTokenUrl(AppID, AppSecret))
	facebookToken := &models.FacebookToken{}
	facebookToken.FromJson(response.Body)
	log.Printf("Got token using AppId: %s - %s", AppID, facebookToken.Token)

	return facebookToken
}

func FacebookTokenUrl(appId string, appSecret string) string {
	return fmt.Sprintf("https://graph.facebook.com/oauth/access_token?client_id=%s&client_secret=%s&grant_type=client_credentials", appId, appSecret)
}

func FacebookUIDUrl(uid string, token string) string {
	return fmt.Sprintf("https://graph.facebook.com/v2.10/%s?fields=%s&access_token=%s", uid, Fields, token)
}

func HttpResponseBodyToString(reader io.Reader) string {
	body, _ := ioutil.ReadAll(reader)
	return string(body)
}
