package utils

import (
	"fmt"
	"net/http"
)

type ResponseType string

const (
	ResponseSuccess ResponseType = "success"
	ResponseError   ResponseType = "error"
	ResponseInfo    ResponseType = "info"
	ResponseWarning ResponseType = "warning"
)

func SendHTMLResponse(w http.ResponseWriter, title string, responseType ResponseType, message string, details string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var cssClass string
	var displayType string

	switch responseType {
	case ResponseSuccess:
		cssClass = "success"
		displayType = "Sucesso"
	case ResponseError:
		cssClass = "danger"
		displayType = "Erro"
	case ResponseInfo:
		cssClass = "info"
		displayType = "Informação"
	case ResponseWarning:
		cssClass = "warning"
		displayType = "Aviso"
	default:
		cssClass = "info"
		displayType = "Resultado"
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="pt-BR">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>%s</title>
  <link rel="stylesheet" href="/static/styles/resultado.style.css">
</head>
<body>
  <div class="card %s">
    <h2>%s</h2>
    <p>%s</p>
    %s
    <a href="/" class="btn">Voltar ao Início</a>
  </div>
</body>
</html>`, title, cssClass, displayType, message, details)

	fmt.Fprint(w, html)
}

func SendDetailedResponse(w http.ResponseWriter, title string, responseType ResponseType, message string, details map[string]string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var cssClass string
	var displayType string

	switch responseType {
	case ResponseSuccess:
		cssClass = "success"
		displayType = "Sucesso"
	case ResponseError:
		cssClass = "danger"
		displayType = "Erro"
	case ResponseInfo:
		cssClass = "info"
		displayType = "Informação"
	default:
		cssClass = "info"
		displayType = "Resultado"
	}

	detailsHTML := ""
	if len(details) > 0 {
		detailsHTML = "<ul class=\"details\">"
		for key, value := range details {
			detailsHTML += fmt.Sprintf("<li><strong>%s:</strong> %s</li>", key, value)
		}
		detailsHTML += "</ul>"
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="pt-BR">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>%s</title>
  <link rel="stylesheet" href="/static/styles/resultado.style.css">
</head>
<body>
  <div class="card %s">
    <h2>%s</h2>
    <p>%s</p>
    %s
    <a href="/" class="btn">Voltar ao Início</a>
  </div>
</body>
</html>`, title, cssClass, displayType, message, detailsHTML)

	fmt.Fprint(w, html)
}
