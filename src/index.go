package main

//https://docs.aws.amazon.com/sdk-for-go/api/service/comprehend/

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

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

type Response struct {
	Version string  `json:"version"`
	Body    ResBody `json:"response"`
}

// ResBody is the actual body of the response
type ResBody struct {
	OutputSpeech     Payload ` json:"outputSpeech,omitempty"`
	ShouldEndSession bool    `json:"shouldEndSession"`
}
type Payload struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

// NewResponse builds a simple Alexa session response
func NewResponse(speech string) Response {
	return Response{
		Version: "1.0",
		Body: ResBody{
			OutputSpeech: Payload{
				Type: "PlainText",
				Text: speech,
			},
			ShouldEndSession: true,
		},
	}
}

// Handler is the lambda hander
func Handler() (Response, error) {
	var phrase string = "Guten Tag"
	var fromLang string = GetLang(phrase)
	var output string = Translate(fromLang, phrase)
	fmt.Println(output)
	return NewResponse(output), nil
}
