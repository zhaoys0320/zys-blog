package common

import (
	"zys-boke-master/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate()
}
