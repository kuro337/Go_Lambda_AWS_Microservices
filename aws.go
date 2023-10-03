package main

import ("fmt")

// Pseudocode for Concurrent Rekognition Image Analysis 

// Bucket name is defined as a const
const bucketName string = "testbucket"

func main() {
     // Frontend sends us an array containing all image names
    // ["image1" , "image2" , "image3"]

    // 2 lines below create an array slice of length 0 and then populate it with image names
    images:= make([]string , 0)
    images= append(images , "img1.jpg" , "img2.jpg" , "img3.jpg")
    
    // Create 1 channel 
    ch := make ( chan string )

    // Run our func concurrently for every image in the array slice 
    for _,img:= range images {
        dest:= fmt.Sprintf("s3://%s/%s" , bucketName,img)
        go getImageResults(dest , ch)   
}

// Collect results

// Create results slice to store results of every image computation
results := make([]string , len(images))

// Send every item from our channel to results array (will contain analysis results)
for i:=range results {
    results[i] = <- ch 
}

// Log Results 
fmt.Printf("Results: %+v\n", results)

}

// This function takes a string and passes to the channel 
func getImageResults (s string , ch chan <- string)  {
    //getImage()
    
    ch <- s 
 
 }

 // Above function should receive the string , compute the results using rekognition API ,
 // then send back to channel
