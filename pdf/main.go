package pdf

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

type fullName struct {
	Title string `json:"title,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
}

func (fullName *fullName) fullName() string {
	return strings.Title(fullName.First + " " + fullName.Last)
}

type name struct {
	Name fullName `json:"name,omitempty"`
}

type results struct {
	Results []name `json:"results,omitempty"`
}

func generateRandomNames() []name {
	HTTPresponse, HTTPError := http.Get("https://randomuser.me/api/?results=10&gender=male&nat=gb&inc=name")
	if HTTPError != nil {
		log.Fatal(HTTPError)
	}

	body, readErr := ioutil.ReadAll(HTTPresponse.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var randomNames results
	jsonErr := json.Unmarshal([]byte(body), &randomNames)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return randomNames.Results
}

func drawCrewCard(pdf *gofpdf.Fpdf, x, y float64, role string, image string, name string) {
	var opt gofpdf.ImageOptions
	pdf.Rect(x, y, 86, 46, "D")

	pdf.SetFont("Arial", "B", 10)
	pdf.Text(x+2, y+33, strings.Title(name))

	pdf.SetFont("Arial", "", 10)
	pdf.Rect(x+72, y+7, 9, 6, "D")
	pdf.Rect(x+72, y+13, 9, 6, "D")
	pdf.Rect(x+67, y+19, 14, 6, "D")
	pdf.Rect(x+72, y+25, 9, 6, "D")
	pdf.Rect(x+72, y+31, 9, 6, "D")

	pdf.SetFont("Arial", "B", 16)
	pdf.Text(x+26, y+6, role)

	pdf.SetFont("Arial", "", 10)
	pdf.Text(x+27, y+11, "Numero de misiones")
	pdf.Text(x+27, y+17, "Cazas derribados")
	pdf.Text(x+27, y+23, "Puntos de mision")
	pdf.Text(x+27, y+29, "Corazon Purpura")
	pdf.Text(x+57, y+35, "Fatiga")
	pdf.ImageOptions(image, x+1, y+1, 24, 24, false, opt, 0, "")
}

func GeneratePDF() []byte {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	var randomNames = generateRandomNames()
	drawCrewCard(pdf, 20, 11, "Bombardero", "images/bombardero.jpg", randomNames[0].Name.fullName())
	drawCrewCard(pdf, 110, 11, "Navegante", "images/navegante.jpg", randomNames[1].Name.fullName())
	drawCrewCard(pdf, 20, 64, "Piloto", "images/piloto.jpg", randomNames[2].Name.fullName())
	drawCrewCard(pdf, 110, 64, "Copiloto", "images/copiloto.jpg", randomNames[3].Name.fullName())
	drawCrewCard(pdf, 20, 117, "Ingeniero", "images/ingeniero.jpg", randomNames[4].Name.fullName())
	drawCrewCard(pdf, 110, 117, "Operador de Radio", "images/operador_de_radio.jpg", randomNames[5].Name.fullName())
	drawCrewCard(pdf, 20, 168, "Ametrallador Ventral", "images/ametrallador_ventral.jpg", randomNames[6].Name.fullName())
	drawCrewCard(pdf, 110, 168, "Ametrallador Babor", "images/ametrallador_babor.jpg", randomNames[7].Name.fullName())
	drawCrewCard(pdf, 20, 221, "Ametrallador Estribor", "images/ametrallador_estribor.jpg", randomNames[8].Name.fullName())
	drawCrewCard(pdf, 110, 221, "Ametrallador de Cola", "images/ametrallador_de_cola.jpg", randomNames[9].Name.fullName())

	var result bytes.Buffer
	var io = io.Writer(&result)

	err := pdf.Output(io)

	if err != nil {
		log.Fatal(err)
	}

	return result.Bytes()

}
