package storage

import (
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

var S3Client *s3.Client
var bucketName string = "sleekspace"
var bucketUrl string = "https://" + bucketName + ".s3.eu-north-1.amazonaws.com/"

func InitializeS3() {
	customProvider := credentials.NewStaticCredentialsProvider(
		generalUtilities.GetEnvVariables().AWSAcessKey,
		generalUtilities.GetEnvVariables().AWSSecretKey,
		"",
	)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(customProvider),
		config.WithRegion("eu-north-1"),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("s3 connected")
	S3Client = s3.NewFromConfig(cfg)
}

func UploadFile(base64FileSrc string, name string, c *gin.Context) <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		i := strings.Index(base64FileSrc, ",")
		decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader((base64FileSrc[i+1:])))
		uploader := manager.NewUploader(S3Client)
		_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: &bucketName,
			Key:    &name,
			Body:   decoder,
		})
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "failed to upload file on cloud storage"})
		}
		result <- bucketUrl + name
	}()
	return result
}

func DeleteFile(name string, c *gin.Context) <-chan bool {
	result := make(chan bool)
	go func() {
		defer close(result)
		input := &s3.DeleteObjectInput{
			Bucket: &bucketName,
			Key:    &name,
		}
		_, err := S3Client.DeleteObject(context.TODO(), input)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "failed to delete file on cloud storage"})
		}
		result <- true
	}()
	return result
}

func DeleteFiles(fileNames []string, c *gin.Context) <-chan bool {
	result := make(chan bool)
	go func() {
		defer close(result)
		for i := 0; i < len(fileNames); i++ {
			input := &s3.DeleteObjectInput{
				Bucket: &bucketName,
				Key:    &fileNames[i],
			}
			_, err := S3Client.DeleteObject(context.TODO(), input)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"error": "failed to delete file on cloud storage"})
			}
		}
		result <- true
	}()
	return result
}

func uploadFileFromAList(
	j int,
	base64FilesSrc []propertyUtilities.MediaFile,
	c *gin.Context,
	wg *sync.WaitGroup,
	ch chan<- map[string]string,
) {
	fileUrlMap := make(map[string]string)
	defer wg.Done()
	i := strings.Index(base64FilesSrc[j].FileSrc, ",")
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader((base64FilesSrc[j].FileSrc[i+1:])))
	uploader := manager.NewUploader(S3Client)
	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &base64FilesSrc[j].Name,
		Body:   decoder,
	})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to upload file on cloud storage"})
	}

	fileUrlMap[base64FilesSrc[j].Name] = bucketUrl + base64FilesSrc[j].Name
	ch <- fileUrlMap
}

func UploadFiles(base64FilesSrc []propertyUtilities.MediaFile, c *gin.Context) map[string]string {
	mergedMap := make(map[string]string)
	resultCh := make(chan map[string]string, len(base64FilesSrc))
	var wg sync.WaitGroup
	if len(base64FilesSrc) > 0 {
		for j := 0; j < len(base64FilesSrc); j++ {
			wg.Add(1)
			go uploadFileFromAList(j, base64FilesSrc, c, &wg, resultCh)
		}
		go func() {
			wg.Wait()
			close(resultCh)
		}()
		for m := range resultCh {
			for name, uri := range m {
				mergedMap[name] = uri
			}
		}
	}
	return mergedMap
}
