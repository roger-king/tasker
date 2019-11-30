package utils

import (
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func DownloadZip(filepath string, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func UnZip(path string) {
	logrus.Info("testing")
}
