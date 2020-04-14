package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/cdn"
	"github.com/qiniu/api.v7/v7/storage"
)

func upload() error {
	rt.upload = append(rt.upload, "index.html")
	for _, d := range rt.upload {
		if err := uploadFile2CDN(d); err != nil {
			return err
		}
	}

	return refreshIndex()
}

func uploadFile2CDN(path string) error {
	putPolicy := storage.PutPolicy{
		Scope: rc.Conf.CDN.Bucket,
	}

	mac := qbox.NewMac(rc.Conf.CDN.AccessKey, rc.Conf.CDN.SecretKey)
	cfg := storage.Config{}

	cfg.Zone = &storage.ZoneHuabei
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	client := http.Client{}

	resumeUploader := storage.NewResumeUploaderEx(&cfg, &storage.Client{Client: &client})
	upToken := putPolicy.UploadToken(mac)

	ret := storage.PutRet{}

	err := resumeUploader.PutFile(context.Background(), &ret, upToken, path, path, nil)
	if err != nil {
		return err
	}

	return nil
}

func refreshIndex() error {
	mac := qbox.NewMac(rc.Conf.CDN.AccessKey, rc.Conf.CDN.SecretKey)
	cdnManager := cdn.NewCdnManager(mac)
	url := fmt.Sprintf("%s/index.html", rc.Conf.Url)
	urlsToRefresh := []string{
		url,
	}

	_, err := cdnManager.RefreshUrls(urlsToRefresh)
	if err != nil {
		return err
	}

	fmt.Printf("%s cdn refresh complete \n", url)
	return nil
}
