package cloud_handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/valueobject"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/http/response"
)

func (x *CloudHttpHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx       = context.Background()
		maxUpload = int64(x.Cfg.MaxUpload) * 1024 * 1024
	)

	jwtToken, errJwt := utils.DecodeJwtToken(r.Header.Get("Authorization"))
	if errJwt != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errJwt})
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{err})
		return
	}

	if handler.Size > maxUpload {
		errMessageSize := fmt.Sprintf("max file size %vMB", x.Cfg.MaxUpload)
		utils.RespondWithError(w, http.StatusBadRequest, []error{errors.New(errMessageSize)})
		return
	}

	mime, errMime := valueobject.NewFileType(handler.Header.Get("Content-Type"))
	if errMime != nil {
		utils.RespondWithError(w, http.StatusBadRequest, []error{errMime})
		return
	}

	tempFile, err := os.CreateTemp(filepath.Dir(handler.Filename)+"/temp-assets", "*."+mime.String())
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	payload := entities.CloudDto{
		Name:      handler.Filename,
		CreatedBy: jwtToken.Subject,
		File: entities.File{
			Origin:      handler.Filename,
			Name:        tempFile.Name(),
			Size:        int(handler.Size),
			MimeType:    mime.String(),
			ContentType: mime.Type(),
			Directory:   x.Cfg.Aws.Directory,
		},
	}

	log.Println("payload:", payload)

	cloud, errCreate := x.CloudUsecase.Create(ctx, payload)
	if err != nil {
		utils.RespondWithError(w, exceptions.MapToHttpStatusCode(errCreate.Status), errCreate.Errors.Errors)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, response.MapCloudListResponse(cloud))
}
