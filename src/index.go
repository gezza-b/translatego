package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/translate"
)

const _eng string = "en"
const _germ string = "de"

func getLangAsString(result *comprehend.DetectDominantLanguageOutput) string {
	return *result.Languages[0].LanguageCode
}

func GetLang(orig string) string {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	svc := comprehend.New(sess)
	result, err := svc.DetectDominantLanguage(&comprehend.DetectDominantLanguageInput{Text: aws.String(orig)})
	if err != nil {
		fmt.Println("GetLang:error: ", err)
	}
	return getLangAsString(result)
}

func getToLang(fromLang string) string {
	if fromLang == _eng {
		return _germ
	}
	return _eng
}

func Translate(fromLang string, phrase string) string {
	var toLang = getToLang(fromLang)
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	svc := translate.New(sess)
	req, resp := svc.TextRequest(&translate.TextInput{SourceLanguageCode: aws.String(fromLang), TargetLanguageCode: aws.String(toLang), Text: aws.String(phrase)})
	err := req.Send()
	if err != nil {
		return ""
	}
	return resp.String()
}

func main() {
	lambda.Start(Handler)
}

// func Handler() (Response, error) {
// 	var phrase string = "Guten Tag"
// 	var fromLang string = GetLang(phrase)
// 	var output string = Translate(fromLang, phrase)
// 	fmt.Println(output)
// 	return Response{Message: output}, nil
// }
// type Session struct {
// 	New         bool   `json:"new"`
// 	SessionID   string `json:"sessionId"`
// 	Application struct {
// 		ApplicationID string `json:"applicationId"`
// 	} `json:"application"`
// 	Attributes map[string]interface{} `json:"attributes"`
// 	User       struct {
// 		UserID      string `json:"userId"`
// 		AccessToken string `json:"accessToken,omitempty"`
// 	} `json:"user"`
// }

// type Request struct {
// 	Version string  `json:"version"`
// 	Session Session `json:"session"`
// 	Body    ReqBody `json:"request"`
// 	Context Context `json:"context"`
// }
// type ReqBody struct {
// 	Type        string `json:"type"`
// 	RequestID   string `json:"requestId"`
// 	Timestamp   string `json:"timestamp"`
// 	Locale      string `json:"locale"`
// 	Intent      Intent `json:"intent,omitempty"`
// 	Reason      string `json:"reason,omitempty"`
// 	DialogState string `json:"dialogState,omitempty"`
// }
// type Response struct {
// 	Version string  `json:"version"`
// 	Body    ResBody `json:"response"`
// }

// // ResBody is the actual body of the response
// type ResBody struct {
// 	OutputSpeech     Payload ` json:"outputSpeech,omitempty"`
// 	ShouldEndSession bool    `json:"shouldEndSession"`
// }
// type Payload struct {
// 	Type string `json:"type,omitempty"`
// 	Text string `json:"text,omitempty"`
// }

// NewResponse builds a simple Alexa session response
// func NewResponse(speech string) Response {
// 	return Response{
// 		Version: "1.0",
// 		Body: ResBody{
// 			OutputSpeech: Payload{
// 				Type: "PlainText",
// 				Text: speech,
// 			},
// 			ShouldEndSession: true,
// 		},
// 	}
// }

// Handler is the lambda hander
//func Handler() (Response, error) {
func Handler(request alexa.Request) alexa.Response {
	var phrase = request.Body.Intent.Slots["Query"].Value

	//var phrase string = "Guten Tag"
	var fromLang string = GetLang(phrase)
	var output string = Translate(fromLang, phrase)
	// return NewResponse(output), nil
	return alexa.NewSimpleResponse(output, output)
}

func DispatchIntents(request alexa.Request) alexa.Response {
	var response alexa.Response
	response = Handler(request)
	return response
}
