package main
//https://docs.aws.amazon.com/sdk-for-go/api/service/comprehend/

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
    //"github.com/aws/aws-sdk-go/service/translate"
    "github.com/aws/aws-sdk-go/service/comprehend"
) 


const _eng string = "eng" 

func GetLang(orig string) string {
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String("ap-southeast-2")},
    )
    svc := comprehend.New(sess)
    fmt.Println(svc)
//comprehend.detectDominantLanguage(params, function (err, data) {
    params := comprehend.DetectDominantLanguageInput { Text: aws.String(orig) } 
    svc.DetectDominantLanguage(*params)
    //req, resp := svc.DetectDominantLanguage(orig)
    return _eng
}


func main() {
    fmt.Println( GetLang("hi") )
}
