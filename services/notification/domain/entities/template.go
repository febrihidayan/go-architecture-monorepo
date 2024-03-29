package entities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/hashicorp/go-multierror"
)

type Template struct {
	ID        common.ID
	Name      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TemplateDto struct {
	ID        *common.ID
	Name      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TemplateLanguage struct {
	EN string `json:"en"`
	ID string `json:"id"`
}

type TemplateDataDTO struct {
	Title TemplateLanguage `json:"title"`
	Body  TemplateLanguage `json:"body"`
}

type TemplateData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type TemplateQueryParams struct {
	Search string
}

const (
	TemplateLangID = "id"
	TemplateLangEN = "en"
)

func NewTemplate(x TemplateDto, finds ...*Template) *Template {

	result := Template{
		ID:        common.NewID(),
		Name:      x.Name,
		Data:      x.Data,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	if x.ID != nil {
		result.ID = *x.ID
	}

	for _, item := range finds {
		result.CreatedAt = item.CreatedAt
	}

	return &result
}

func (x *Template) Validate() (err *multierror.Error) {
	if x.Name == "" {
		err = multierror.Append(err, lang.Trans("filled", "Name"))
	}
	if x.Data == "" {
		err = multierror.Append(err, lang.Trans("filled", "Data"))
	}

	return
}

func (x *Template) GetData() (data interface{}) {
	json.Unmarshal([]byte(x.Data), &data)

	return
}

func (x *Template) GetTemplateData(data interface{}, lang string) (result TemplateData) {
	var (
		replaced   bytes.Buffer
		resultData TemplateDataDTO
	)

	parser, _ := template.New("").Option("missingkey=error").Parse(x.Data)

	parser.Execute(&replaced, data)
	json.Unmarshal(replaced.Bytes(), &resultData)

	switch lang {
	case TemplateLangID:
		result.Title = resultData.Title.ID
		result.Body = resultData.Body.ID
	default:
		result.Title = resultData.Title.EN
		result.Body = resultData.Body.EN
	}

	return
}

func (x *Template) GetTemplateMaps(data interface{}, lang string) map[string]string {
	var (
		replaced   bytes.Buffer
		resultData map[string]interface{}
		result     = make(map[string]string, 0)
	)

	parser, _ := template.New("").Option("missingkey=error").Parse(x.Data)

	parser.Execute(&replaced, data)
	json.Unmarshal(replaced.Bytes(), &resultData)

	for key, items := range resultData {
		for langCode, val := range items.(map[string]interface{}) {
			switch lang {
			case langCode:
				result[key] = fmt.Sprintf("%s", val)
			default:
				result[TemplateLangEN] = fmt.Sprintf("%s", val)
			}
		}
	}

	return result
}
