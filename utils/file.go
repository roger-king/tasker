package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func DownloadZip(fp string, url string) error {
	resp, err := http.Get(url)

	if err != nil {
		logrus.Error(err)
		return err
	}

	defer resp.Body.Close()
	fPath, err := filepath.Abs(fp)

	if err != nil {
		return err
	}

	out, err := os.Create(fPath + "/build.zip")

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
