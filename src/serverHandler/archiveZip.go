package serverHandler

import (
	"../serverError"
	"archive/zip"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
)

func writeZip(zw *zip.Writer, f *os.File, fInfo os.FileInfo, archivePath string) error {
	if archivePath[0] == '/' {
		archivePath = archivePath[1:]
	}
	if fInfo.IsDir() {
		archivePath += "/"
	}

	var size int64
	if !fInfo.IsDir() {
		size = fInfo.Size()
	}

	w, err := zw.Create(archivePath)
	if err != nil {
		return err
	}

	if size == 0 || f == nil {
		return nil
	}

	_, err = io.Copy(w, f)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) zip(w http.ResponseWriter, r *http.Request, pageData *pageData) {
	zipWriter := zip.NewWriter(w)
	defer func() {
		err := zipWriter.Close()
		serverError.LogError(err)
	}()

	filename := url.PathEscape(pageData.ItemName + ".zip")

	header := w.Header()
	header.Set("Content-Type", "application/zip")
	header.Set("Content-Disposition", "attachment; filename*=UTF-8''"+filename)
	header.Set("Cache-Control", "public, max-age=0")
	w.WriteHeader(http.StatusOK)

	h.visitFs(
		h.root+pageData.handlerRequestPath,
		pageData.rawRequestPath,
		"",
		func(f *os.File, fInfo os.FileInfo, relPath string) {
			go h.logArchive(filename, relPath, r)
			err := writeZip(zipWriter, f, fInfo, relPath)
			if serverError.LogError(err) {
				runtime.Goexit()
			}
		},
	)
}
