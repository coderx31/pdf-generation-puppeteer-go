package main

import (
	"encoding/base64"
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const footerTemplate = `
	<style> .footer{ padding: 20px; font-size: 8px; width: 297mm; }</style>
	<div class="footer">
		<p>Page Number: <span class="pageNumber"></span></p>
	</div>
`
const headerTemplate = `
	<style>.pdf-header{ text-align: left; width: 297mm; margin-left: 10mm } img { width: auto; height: 10mm}</style>
	<div class=pdf-header>
		<img src=%s alt=medium-logo>
	</div>
`

const HTMLPath = "tmp/html/%s.html"
const PdfPath = "tmp/pdfs/%s.pdf"

const PeopleTemplate = "templates/new-people.mustache"

type Person struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Age       int
	City      string
}

func main() {
	fmt.Println("PDF Generation with puppeteer")

	people := make([]Person, 0)

	for i := 0; i < 50; i++ {
		person := Person{
			ID:        uuid.New(),
			FirstName: fmt.Sprintf("User %d", i),
			LastName:  "lastname",
			Age:       25,
			City:      "Negombo",
		}

		people = append(people, person)
	}

	data := map[string]interface{}{
		"filename": "people",
		"title":    "pdf testing doc",
		"people":   people,
	}

	templatePath := generateAbsolutePath(PeopleTemplate)
	htmlFilepath, err := generateHtmlFile(data, templatePath)
	if err != nil {
		log.Fatalf("err while html file generation: %v", err.Error())
	}

	pdfSavingPth := generateAbsolutePath(fmt.Sprintf(PdfPath, "people"))
	err = generatePDF(htmlFilepath, pdfSavingPth)
	if err != nil {
		log.Fatalf("error while pdf generation: %v", err.Error())
	}

}

func imageToBase64(imgPath string) (string, error) {
	bytes, err := os.ReadFile(imgPath)
	if err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(bytes)
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Str = "data:image/jpeg;base64," + base64Str
	case "image/png":
		base64Str = "data:image/png;base64," + base64Str
	default:
		base64Str = ""
	}

	return base64Str, nil
}

func saveFile(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, 0o600)
	return err
}

func generateAbsolutePath(path string) string {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	return absolutePath
}

func generateHtmlFile(data map[string]interface{}, templatePath string) (string, error) {
	templatePath = generateAbsolutePath(templatePath)
	htmlOutput, err := mustache.RenderFile(templatePath, data)
	if err != nil {
		return "", err
	}

	filename := data["filename"]
	filePath := generateAbsolutePath(fmt.Sprintf(HTMLPath, filename))

	err = saveFile(filePath, []byte(htmlOutput))
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func generatePDF(htmlPath, pdfPath string) error {
	imgPath := generateAbsolutePath("templates/images/medium-logo.png")
	imgStr, err := imageToBase64(imgPath)
	if err != nil {
		return err
	}
	headerImgTemplate := fmt.Sprintf(headerTemplate, imgStr)
	command := "puppeteer"
	commandArgs := []string{"print", htmlPath, pdfPath, "--sandbox", "false", "--wait-until", "networkidle0",
		"--display-header-footer", "true", "--footer-template", footerTemplate, "--header-template", headerImgTemplate,
		"--margin-top", "30mm", "--margin-bottom", "50mm"}

	cmd := exec.Command(command, commandArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	log.Println("puppeteer cli output", string(output))
	return nil
}
