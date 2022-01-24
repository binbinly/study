package minio

import (
	"context"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

//Config 配置
type Config struct {
	Endpoint     string
	AccessID     string
	SecretAccess string
	Token        string
	Bucket       string
	Region       string
	Secure       bool
}

//Storage 对象存储
type Storage struct {
	client *minio.Client
	cfg    *Config
	ctx    context.Context
}

//New 实例化对象存储
func New(cfg *Config) *Storage {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessID, cfg.SecretAccess, cfg.Token),
		Secure: cfg.Secure,
	})
	if err != nil {
		log.Fatalf("minio new client err: %v", err)
	}
	return &Storage{
		client: client,
		cfg:    cfg,
		ctx:    context.Background(),
	}
}

//Init 初始化
func (s *Storage) Init() error {
	exist, err := s.client.BucketExists(s.ctx, s.cfg.Bucket)
	if err != nil {
		return err
	}
	if !exist {
		err = s.client.MakeBucket(s.ctx, s.cfg.Bucket, minio.MakeBucketOptions{Region: s.cfg.Region})
		if err != nil {
			return err
		}
		log.Printf("Successfully created bucket: %v", s.cfg.Bucket)
	}
	return nil
}

//SignUrl 生成一个给HTTP PUT请求用的presigned URL
func (s *Storage) SignUrl(ctx context.Context, name string) (string, error) {
	u, err := s.client.PresignedPutObject(ctx, s.cfg.Bucket, name, time.Hour)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
