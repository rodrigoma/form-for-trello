package gdrive

import (
	"context"
	"fmt"
	"formfortrello/setting"
	"formfortrello/utils"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//const maxUploadSize2 = 10 << 20

var service *drive.Service

func Setup() {
	config, err := google.ConfigFromJSON([]byte(setting.GoogleSetting.Credentials), drive.DriveFileScope)

	if err != nil {
		log.Fatalf("Error: Reading Config From JSON:  %v\n", err)
	}

	client := config.Client(context.Background(), setting.GoogleTokenJson)

	service, err = drive.New(client)

	if err != nil {
		log.Fatalf("Cannot create the Google Drive service: %v\n", err)
	}
}

func UploadGDrive(r *http.Request) (fileGD *drive.File, err error){
	//if err := r.ParseMultipartForm(maxUploadSize2); err != nil {
	//	return "", errors.New("FILE_TOO_BIG")
	//}

	// Step 1. Open the file
	//----------- ISSO AQUI FEZ FUNCIONAR, ALGO RELACIONADO AO BYTES
	fileAtt, headerAtt, err := r.FormFile("attachment")
	if err != nil {
		return nil, err
	}
	defer fileAtt.Close()
	//---------

	filenameGD := mountFilenameGD(r.PostForm, headerAtt.Filename)
	fmt.Println("FILENAME: " + filenameGD)

	// Step 2. Get the Google Drive service
	//service, err := getService()

	// Step 4. Create the file, upload its content and give permission
	fileGD, err = createFile(filenameGD, fileAtt, setting.GoogleDriveSetting.FolderId)
	if err != nil {
		panic(fmt.Sprintf("Could not create file: %v\n", err))
	}

	return fileGD, err
}

func mountFilenameGD(form url.Values, attachmentName string) (filename string) {
	dt := strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					utils.FormatDate(form.Get("date")), "-", ""), " ", "-"), ":", "")

	return strings.ToUpper( dt + "-" +
		strings.ReplaceAll(form.Get("minister"), " ", "_") + "-" +
		strings.ReplaceAll(form.Get("event"), " ", "_")) + "-" + attachmentName
}

func createFile(name string, content io.Reader, parentId string) (*drive.File, error) {
	f := &drive.File{
		Name:     name,
		Parents:  []string{parentId},
	}

	file, err := service.Files.Create(f).Media(content).Do()
	if err != nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}

	createPermission(file.Id)

	return file, nil
}

func createPermission(fileId string) {
	p := &drive.Permission{
		AllowFileDiscovery: false,
		Role:	"reader",
		Type:	"anyone",
	}

	_, err := service.Permissions.Create(fileId, p).Do()
	if err != nil {
		log.Println("Could not give permission to file: " + err.Error())
	}
}