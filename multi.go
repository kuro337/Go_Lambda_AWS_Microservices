package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/rekognition"
				"github.com/aws/aws-sdk-go/aws/awserr"

)

// Pseudocode for Concurrent Rekognition Image Analysis 

// Bucket name is defined as a const
const bucketName string = "testbucketchinmay"

func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler() {
	 // Frontend sends us an array containing all image names
    // ["image1" , "image2" , "image3"]

    // 2 lines below create an array slice of length 0 and then populate it with image names
    images:= make([]string , 0)
    images= append(images , "chuttersnap-TSgwbumanuE-unsplash.jpg" , "business-team-meeting-boardroom.jpg" , "happy-dog-wears-flowers.jpg")
    // images= append(images , "chuttersnap-TSgwbumanuE-unsplash.jpg" )
    
    // Create 1 channel 
    ch := make ( chan *rekognition.DetectLabelsOutput , len(images))

		// Creating Wait Group so we can know when our code finishes executing in GoRoutine

	//	var wg sync.WaitGroup

	//	wg.Add(len(images))

		fmt.Println("length of images" , len(images))
		
    // Run our func concurrently for every image in the array slice 
    for _,img:= range images {
        // dest:= fmt.Sprintf("s3://%s/%s" , bucketName,img)
		
        go getImageResults(img , ch)   
}

    fmt.Println("Waiting for goroutines to finish...")

		// Create results slice to store results of every image computation
		results := make([]*rekognition.DetectLabelsOutput , len(images))


		// Collect results
		// Send every item from our channel to results array (will contain analysis results)
		for i:=range results {
			results[i] = <- ch 
		}

	//  wg.Wait()
	// 	close(ch)
		fmt.Println("Done!")

// Log Results 
fmt.Printf("Results: %+v\n", results)

}



// This function takes a string and passes to the channel 
func getImageResults (s string , ch chan <- *rekognition.DetectLabelsOutput  )  {
    
		// Passes image name to getLabels and sends results of that to ch 
		res , err := getLabels(s)
		fmt.Println("finished getlabels")
		fmt.Println(res , err)
		fmt.Println("starting sending to channel")
		 ch <- res
		 fmt.Println("finished sending to channel")

    
    // ch <- s 
 
 }

 // Above function should receive the string , compute the results using rekognition API ,
 // then send back to channel

 // Rekognition SDK Calls

 func getLabels(objName string) (*rekognition.DetectLabelsOutput, error) {
	svc := rekognition.New(session.New())

	input := &rekognition.DetectLabelsInput{
Image: &rekognition.Image{
	S3Object: &rekognition.S3Object{
			Bucket: aws.String("testbucketchinmay"),
			Name:   aws.String(objName),
	},
},
MaxLabels:     aws.Int64(10),
MinConfidence: aws.Float64(90.000000),}
fmt.Println("calling detectLabels")

	result, err := svc.DetectLabels(input)
	fmt.Println("completed calling detectLabels")
if err != nil {
	fmt.Println("got errors")
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

// fmt.Println(result)
        // return "Hello λ!", nil
				return result , nil
}

