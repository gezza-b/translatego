package main

//https://docs.aws.amazon.com/sdk-for-go/api/service/comprehend/

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	//"github.com/aws/aws-sdk-go/service/translate"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

const _eng string = "eng"

func GetLang(orig string) string {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)
	svc := comprehend.New(sess)
	result, err := svc.DetectDominantLanguage(&comprehend.DetectDominantLanguageInput{Text: aws.String(orig)})
	if err != nil {
		fmt.Println("GetLang:error: ", err)
	}
	fmt.Println(result)
	return _eng
}

func main() {
	// fmt.Println(GetLang("hi"))
	fmt.Println(GetLang("Guten Tag"))
}
