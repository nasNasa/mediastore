package storagehandler

import (
	"fmt"
	"io"
	"mediaStorer/pkg/claim"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (h Handler) writeFileHandler(e echo.Context) error {
	claims := claim.GetClaimsFromEchoContext(e)
	fmt.Println("claims userid", claims.UserId)
	//bind data

	name := e.FormValue("name")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := e.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileId, uErr := h.fileSvc.UploadFile(src, name, claims.UserId)
	if uErr != nil {
		fmt.Println("upload error", err)
		return err
	}

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	gfile, _ := h.fileSvc.DownloadFile(fileId, 10)

	e.Response().Write(gfile.Bytes())
	return e.JSON(http.StatusOK, gfile.Bytes())

	// return e.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s.</p>", file.Filename, name))
}

func (h Handler) downloadFileHandler(e echo.Context) error {
	claims := claim.GetClaimsFromEchoContext(e)
	fmt.Println("claims userid", claims.UserId)
	//bind data

	// file, serr := h.fileSvc.DownloadFile("648858328e1f369df731d10b", 10)
	// if serr != nil {
	// 	fmt.Println("download file errror", serr)
	// 	return serr
	// }
	// fmt.Println(file.String())

	//return
	return e.JSON(http.StatusOK, "file.String()")
}
