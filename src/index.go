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

func Handler(request alexa.Request) (alexa.Response, error) {
	return DispatchIntents(request), nil
	// var phrase = request.Body.Intent.Slots["Query"].Value
	// //var phrase string = "Guten Tag"
	// var fromLang string = GetLang(phrase)
	// var output string = Translate(fromLang, phrase)
	// // return NewResponse(output), nil
	// return alexa.NewSimpleResponse(output, output), nil
}

func DispatchIntents(request alexa.Request) alexa.Response {
	var phrase = "Good day"
	// var phrase = request.Body.Intent.Slots["Query"].Value
	fmt.Printf("PHRASE: ", phrase)
	var fromLang string = GetLang(phrase)
	var output string = Translate(fromLang, phrase)
	return alexa.NewSimpleResponse(output, output)
}
