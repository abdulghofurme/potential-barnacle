package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/url"

	_ "embed"

	"abdulghofur.me/pshamo-go/config"
	"abdulghofur.me/pshamo-go/helper"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

//go:embed storage_credential.json
var credential string

func NewBucket() *storage.BucketHandle {
	config := &firebase.Config{
		StorageBucket: config.MyEnv.STORAGE_BUCKET,
	}

	opt := option.WithCredentialsJSON([]byte(credential))
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	return bucket
}

func NewStorage() *Storage {
	return &Storage{
		bucket: NewBucket(),
	}
}

type Storage struct {
	bucket *storage.BucketHandle
}

func (st *Storage) Upload(file io.Reader, fileName, id string) string {
	ctx := context.Background()
	object := st.bucket.Object(fmt.Sprintf("%v/%v", config.MyEnv.STORAGE_FOLDER_NAME, fileName))

	// Optional: set a metageneration-match precondition to avoid potential race
	// conditions and data corruptions. The request to update is aborted if the
	// object's metageneration does not match your precondition.
	attrs, err := object.Attrs(ctx)
	helper.PanicIfErrof(err)
	object = object.If(storage.Conditions{MetagenerationMatch: attrs.Metageneration})

	// Update the object to set the metadata.
	objectAttrsToUpdate := storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{
			"firebaseStorageDownloadTokens": id,
		},
	}

	_, err = object.Update(ctx, objectAttrsToUpdate)
	helper.PanicIfErrof(err)
	//   object = object.If(storage.Conditions{DoesNotExist: true})
	// If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	// attrs, err := o.Attrs(ctx)
	// if err != nil {
	//      return fmt.Errorf("object.Attrs: %w", err)
	// }
	// o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	wc := object.NewWriter(ctx)
	_, err = io.Copy(wc, file)
	helper.PanicIfErrof(err)

	err = wc.Close()
	helper.PanicIfErrof(err)

	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%v/o/%v?alt=media&token=%v", object.BucketName(), url.QueryEscape(object.ObjectName()), id)
	return url
}
