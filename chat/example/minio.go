package main

import (
	"github.com/minio/minio-go/v6"
	"log"
	"os"
)

func main() {
	endpoint := "192.168.162.172:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// 初使化minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", minioClient) // minioClient初使化成功

	// 创建一个叫mymusic的存储桶。
	bucketName := "images"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	dir, _ := os.Getwd()
	// 上传一个zip文件。
	objectName := "10.jpeg"
	filePath := dir + "/10.jpeg"
	contentType := "images/jpeg"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType:contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}