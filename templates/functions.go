package templates

import (
	"encoding/json"
	"fmt"
	"html/template"
	"time"

	radio "github.com/R-a-dio/valkyrie"
	"github.com/R-a-dio/valkyrie/errors"
	"golang.org/x/exp/constraints"
)

type divisible interface {
	constraints.Integer | constraints.Float
}

func TemplateFuncs() template.FuncMap {
	return fnMap
}

var fnMap = map[string]any{
	"printjson":                   PrintJSON,
	"safeHTML":                    SafeHTML,
	"safeHTMLAttr":                SafeHTMLAttr,
	"Until":                       time.Until,
	"Since":                       time.Since,
	"Now":                         time.Now,
	"TimeagoDuration":             TimeagoDuration,
	"PrettyDuration":              TimeagoDuration,
	"HumanDuration":               HumanDuration,
	"Div":                         func(a, b int) int { return a / b },
	"CalculateSubmissionCooldown": radio.CalculateSubmissionCooldown,
}

func PrintJSON(v any) (template.HTML, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	return template.HTML("<pre>" + string(b) + "</pre>"), err
}

func SafeHTML(v any) (template.HTML, error) {
	s, ok := v.(string)
	if !ok {
		return "", errors.E(errors.InvalidArgument)
	}
	return template.HTML(s), nil
}

func SafeHTMLAttr(v any) (template.HTMLAttr, error) {
	s, ok := v.(string)
	if !ok {
		return "", errors.E(errors.InvalidArgument)
	}
	return template.HTMLAttr(s), nil
}

func TimeagoDuration(d time.Duration) string {
	if d > 0 { // future duration
		if d <= time.Minute {
			return "in less than a min"
		}
		if d < time.Minute*2 {
			return fmt.Sprintf("in %.0f min", d.Minutes())
		}
		return fmt.Sprintf("in %.0f mins", d.Minutes())
	} else { // past duration
		d = d.Abs()
		if d <= time.Minute {
			return "less than a min ago"
		}
		if d < time.Minute*2 {
			return fmt.Sprintf("%.0f min ago", d.Minutes())
		}
		return fmt.Sprintf("%.0f mins ago", d.Minutes())
	}
}

func HumanDuration(d time.Duration) string {
	return d.Truncate(time.Second).String()
}
