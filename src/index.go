package main

//https://docs.aws.amazon.com/sdk-for-go/api/service/comprehend/

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/translate"
)

const _eng string = "eng"
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
	fmt.Println("toLang: ", toLang)
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	svc := translate.New(sess)
	svc.Translate(&translate.Translate{SourceLanguageCode: aws.String(fromLang), TargetLanguageCode: aws.String(toLang), TranslatedText: aws.String(phrase)})
	// result, err := svc.TranslatedText(&translate.TranslatedText{SourceLanguageCode: aws.String(fromLang), TargetLanguageCode: aws.String(toLang), TranslatedText: aws.String(phrase)})

	//result, err := svc.DetectDominantLanguage(&comprehend.DetectDominantLanguageInput{Text: aws.String(orig)})
	fmt.Println("T::", svc)
	return toLang
}

func main() {
	// fmt.Println(GetLang("hi"))
	var phrase string = "Guten Tag"
	var fromLang string = GetLang(phrase)
	fmt.Println("from: ", fromLang)
	Translate(fromLang, phrase)
}
