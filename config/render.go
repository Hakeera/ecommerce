package config

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer implementa o renderer de templates HTML para o Echo.
type TemplateRenderer struct {
	Templates *template.Template
}

// Render executa o template solicitado e escreve a saída no Writer.
func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	tmpl := t.Templates.Lookup(name)
	if tmpl == nil {
		log.Printf("❌ Template '%s' não encontrado!", name)
		for _, t := range t.Templates.Templates() {
			log.Printf("📄 Template disponível: %s", t.Name())
		}
		return fmt.Errorf("template %s não encontrado", name)
	}

	log.Printf("✅ Renderizando template: %s", name)
	log.Printf("📊 Dados: %+v", data)

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		log.Printf("❌ ERRO na execução do template '%s': %v", name, err)
		return fmt.Errorf("erro ao executar template %s: %v", name, err)
	}

	log.Printf("✅ Template '%s' executado com sucesso!", name)
	return nil
}

// TemplateFunctions define as funções de template personalizadas
var TemplateFunctions = template.FuncMap{
	// Função lower - converte string para minúscula
	"lower": strings.ToLower,

	// Função upper - converte string para maiúscula
	"upper": strings.ToUpper,

	// Função formatMoney - formata valores monetários
	"formatMoney": func(value int) string {
		reais := float64(value) / 100.0
		return fmt.Sprintf("R$ %.2f", reais)
	},

	// Função replace - substitui strings
	"replace": func(old, new, s string) string {
		return strings.ReplaceAll(s, old, new)
	},

	// Função contains - verifica se contém substring
	"contains": strings.Contains,
}
