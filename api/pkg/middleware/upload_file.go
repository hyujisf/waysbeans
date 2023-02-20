package middleware

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("image")

		if err := c.Request().ParseMultipartForm(1024); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		uploadedFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer uploadedFile.Close()

		if _, err = io.Copy(uploadedFile, src); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// create empty context
		var ctx = context.Background()

		// setup cloudinary credentials
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		// create new instance of cloudinary object using cloudinary credentials
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

		// Upload file to Cloudinary
		resp, err := cld.Upload.Upload(ctx, uploadedFile.Name(), uploader.UploadParams{Folder: "waysbeans"})
		if err != nil {
			fmt.Println(err.Error())
		}

		c.Set("dataFile", resp.SecureURL)
		return next(c)
	}
}
