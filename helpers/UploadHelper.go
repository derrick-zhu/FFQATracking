package helpers

import (
	"FFQATracking/utils"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/astaxie/beego"
)

// SaveAttachFile write the uploaded file into back-end server
func SaveAttachFile(c *http.Request, htmlID, baseDir string) (string, error) {

	var filePath string

	for {
		beego.Debug(htmlID, " ", baseDir)
		infile, fheader, inErr := c.FormFile(htmlID)
		if inErr != nil {
			beego.Error(inErr)
			break
		}
		defer infile.Close()

		// generate hash file name (MD5)
		hasher := md5.New()
		_, err := io.WriteString(hasher, fheader.Filename+strconv.FormatInt(utils.TimeTickSince1970(), 10))
		if err != nil {
			beego.Error(err)
		}
		hashBytes := hasher.Sum(nil)[:]
		tmpFileName := hex.EncodeToString(hashBytes)

		fileName := tmpFileName + filepath.Ext(fheader.Filename)
		filePath = filepath.Join(baseDir, fileName)
		beego.Info(filePath)

		outfile, outErr := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if outErr != nil {
			beego.Error(outErr)
			break
		}
		defer outfile.Close()

		io.Copy(outfile, infile)
		break
	}

	return filePath, nil
}

// DeleteAttachFile remove a file with given file path
func DeleteAttachFile(filePath string) error {

	err := os.Remove(filePath)
	if err != nil {
		beego.Error(err)
	}

	return err
}
