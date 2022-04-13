<link href="style.css" rel="stylesheet">

# Go Development on AWS

## Setup

<ul>
<li>To deploy code to Lambda , we need to zip the go executable to a Linux Binary </li>

<li>Launching Linux Machine using Docker  </li>

</br>

```
Run from powershell

PWD is current directory , run from where go file is - amazonlinux launches al2 os

docker run -i -t --name mygo -v ${PWD}:/usr/test  amazonlinux /bin/bash

$PWD maps the VM to the current directory we are in. When we launch the Linux machine and create files at the root dir , they will show up where we created the docker container from

To exit from Container, exit out of Linux

To go back in to machine , start container , find container name and
// starting
docker start mygo
// attaching
docker attach containername



```

</br>

</br>

```
Once VM launched , install Go on it

yum install golang -y


1. Download build-lambda-zip tool from Github

go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip

2. Download lambda library
go get github.com/aws/aws-lambda-go/lambda
GOOS=linux go build hello.go
zip function.zip hello

We have our file function.zip



```

</br>

<li> Initial Go App Deploying to Lambda </li>

</br>

```
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
  "fmt"
)

func hello() (string, error) {
  fmt.Println("hello aws");
	return "Hello λ!", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda

	lambda.Start(hello)
}

Create above file main.go
Run following commands to create executable
go get github.com/aws/aws-lambda-go/lambda
GOOS=linux go build main.go

This will create main from main.go
Give file 644 permissions
or
chmod +x main

Then create zip from executable and send to /usr/test
zip -j main.zip main



Send file to /usr/test


```

</br>

<li><k>Setting up Golang to interact with AWS from local</k>  </li>

</br>

```
First, install Go SDK
go get -u github.com/aws/aws-sdk-go/...

go mod init aws // creates go.mod

Create our code file : aws.go :

package main

import (
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/rekognition"
"fmt"
"github.com/aws/aws-sdk-go/aws/awserr"
)

func main() {
        svc := rekognition.New(session.New())
        input := &rekognition.DetectLabelsInput{
    Image: &rekognition.Image{
        S3Object: &rekognition.S3Object{
            Bucket: aws.String("testbucketchinmay"),
            Name:   aws.String("happy-dog-wears-flowers.jpg"),
        },
    },
    MaxLabels:     aws.Int64(123),
    MinConfidence: aws.Float64(70.000000),
}

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
    return
}

fmt.Println(result)
}

NOTE : Our credentials file has to be at /root/.aws/credentials

Credentials :
[default]
aws_access_key_id = abcd
aws_secret_access_key = 1234

Also , the region should be set up

export AWS_REGION=us-east-1

echo $AWS_REGION should show us-east-1

Now when we run
go run aws.go

Output is :

{
  LabelModelVersion: "2.0",
  Labels: [
    {
      Confidence: 98.97029113769531,
      Instances: [{
          BoundingBox: {
            Height: 0.7660114169120789,
            Left: 0.05057369917631149,
            Top: 0.2242119461297989,
            Width: 0.9429351091384888
          },
          Confidence: 98.97029113769531
        }],
      Name: "Dog",
      Parents: [
        {
          Name: "Pet"
        },
        {
          Name: "Canine"
        },
        {
          Name: "Animal"
        },
        {
          Name: "Mammal"
        }
      ]
    },
    {
      Confidence: 98.97029113769531,
      Instances: [],
      Name: "Pet",
      Parents: [{
          Name: "Animal"
        }]
    },
    {
      Confidence: 98.97029113769531,
      Instances: [],
      Name: "Canine",
      Parents: [{
          Name: "Mammal"
        },{
          Name: "Animal"
        }]
    },
    {
      Confidence: 98.97029113769531,
      Instances: [],
      Name: "Animal",
      Parents: []
    },
    {
      Confidence: 98.97029113769531,
      Instances: [],
      Name: "Mammal",
      Parents: [{
          Name: "Animal"
        }]
    },
    {
      Confidence: 88.8868408203125,
      Instances: [],
      Name: "Appenzeller",
      Parents: [
        {
          Name: "Dog"
        },
        {
          Name: "Pet"
        },
        {
          Name: "Canine"
        },
        {
          Name: "Animal"
        },
        {
          Name: "Mammal"
        }
      ]
    },
    {
      Confidence: 85.79913330078125,
      Instances: [],
      Name: "Plant",
      Parents: []
    },
    {
      Confidence: 81.39485931396484,
      Instances: [],
      Name: "Flower",
      Parents: [{
          Name: "Plant"
        }]
    },
    {
      Confidence: 81.39485931396484,
      Instances: [],
      Name: "Blossom",
      Parents: [{
          Name: "Plant"
        }]
    }
  ]
}

```

</br>
<li>  <a href="https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go">Good Article for How-To of Concurrency / Parallel Execution Golang</a> </li>

<li> Channels can be read only or write only </li>

</br>

```
Read Only Channel

func read ( ch <- chan int ) {
  // read only
}

func write ( ch chan <- int ) {
  // writeable
}

```

</br>
<li>  </li>
<li>  </li>
<li>  </li>

<li>  </li>
<li>  </li>
<li>  </li>
<li>  </li>

<li>  </li>
<li>  </li>
<li>  </li>
<li>  </li>

<li>  </li>
<li>  </li>
<li>  </li>
<li>  </li>

</ul>
