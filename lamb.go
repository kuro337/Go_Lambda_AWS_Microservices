package main

import (
        "github.com/aws/aws-lambda-go/lambda"
  "fmt"
	"github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/rekognition"
				"github.com/aws/aws-sdk-go/aws/awserr"

)

func getLabels() (*rekognition.DetectLabelsOutput, error) {
	svc := rekognition.New(session.New())

	input := &rekognition.DetectLabelsInput{
Image: &rekognition.Image{
	S3Object: &rekognition.S3Object{
			Bucket: aws.String("testbucket"),
			Name:   aws.String("happy-dog-wears-flowers.jpg"),
	},
},
MaxLabels:     aws.Int64(123),
MinConfidence: aws.Float64(70.000000),}

	result, err := svc.DetectLabels(input)
if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        case rekognition.ErrCodeInvalidS3ObjectException:
            fmt.Println(rekognition.ErrCodeInvalidS3ObjectException, aerr.Error())
        case rekognition.ErrCodeInvalidParameterException:
            fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
        case rekognition.ErrCodeImageTooLargeException:
            fmt.Println(rekognition.ErrCodeImageTooLargeException, aerr.Error())
        case rekognition.ErrCodeAccessDeniedException:
            fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
        case rekognition.ErrCodeInternalServerError:
            fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
        case rekognition.ErrCodeThrottlingException:
            fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
        case rekognition.ErrCodeProvisionedThroughputExceededException:
            fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
        case rekognition.ErrCodeInvalidImageFormatException:
            fmt.Println(rekognition.ErrCodeInvalidImageFormatException, aerr.Error())
        default:
            fmt.Println(aerr.Error())
        }
    } else {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
    }
    // return
}

fmt.Println(result)
        // return "Hello λ!", nil
				return result , nil
}

func main() {
        // Make the handler available for Remote Procedure Call by AWS Lambda

        lambda.Start(getLabels)
}


