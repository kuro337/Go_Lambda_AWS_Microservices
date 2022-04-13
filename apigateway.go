package main

import (
    "fmt"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/valyala/fastjson"
)

func HandleRequest(request events.APIGatewayProxyRequest) () {
    ApiResponse := events.APIGatewayProxyResponse{}
    // Switch for identifying the HTTP request
    switch request.HTTPMethod {
    case "POST":
        // Obtain the QueryStringParameter
        // name := request.QueryStringParameters["name"]
				name := request.Body
        if name != "" {
            ApiResponse = events.APIGatewayProxyResponse{Body: "Hey " + name + " welcome! ", StatusCode: 200}
        } else {
            ApiResponse = events.APIGatewayProxyResponse{Body: "Error: Query Parameter name missing", StatusCode: 500}
        }
				fmt.Println(name)

    // case "POST":    
    //     //validates json and returns error if not working
    //     err := fastjson.Validate(request.Body)

    //     if err != nil {
    //         body := "Error: Invalid JSON payload ||| " + fmt.Sprint(err) + " Body Obtained" + "||||" + request.Body
    //         ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 500}
    //     } else {
    //         ApiResponse = events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}
    //     }

    }
    // Response
    // return ApiResponse, nil
}

func main() {
    lambda.Start(HandleRequest)
}