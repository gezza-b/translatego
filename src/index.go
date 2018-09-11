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
}

func DispatchIntents(request alexa.Request) alexa.Response {
	var intent = request.Body.Intent.Name
	var response alexa.Response
	switch intent {
	case "TranslateIntent":
		response = handleTranslate(request)
	case "TranslateFromDeIntent":
		response = handleTranslateFromDe(request)
	case alexa.CancelIntent:
		response = handleCancel(request)
	}
	fmt.Println("response::: ", response)
	return response
}

func handleTranslate(request alexa.Request) alexa.Response {
	var phrase = request.Body.Intent.Slots["Query"].Value
	fmt.Printf("PHRASE:: ", phrase)
	fmt.Println("Intent:: ", request.Body.Intent)
	fmt.Printf("Slots:: ", request.Body.Intent.Slots["Query"])
	fmt.Printf("SlotVal:: ", request.Body.Intent.Slots["Query"].Value)
	var fromLang string = GetLang(phrase)
	var output string = Translate(fromLang, phrase)
	return alexa.NewSimpleResponse(output, output)
}

func handleTranslateFromDe(request alexa.Request) alexa.Response {
	var phrase = request.Body.Intent.Slots["Query"].Value
	fmt.Printf("PHRASE:: ", phrase)
	fmt.Println("Intent:: ", request.Body.Intent)
	fmt.Printf("Slots:: ", request.Body.Intent.Slots["Query"])
	fmt.Printf("SlotVal:: ", request.Body.Intent.Slots["Query"].Value)
	var fromLang string = _germ
	var output string = Translate(fromLang, phrase)
	return alexa.NewSimpleResponse(output, output)
}

func handleCancel(request alexa.Request) alexa.Response {
	var output string = "Good bye"
	return alexa.NewSimpleResponse(output, output)
}

//https://github.com/arienmalec/alexa-go
