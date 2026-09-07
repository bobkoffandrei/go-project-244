package formatters

import (
        "github.com/bobkoffandrei/go-project-244/models"
)


func GetFormatter(format string) func([]models.Node) string {
    switch format {
    case "plain":
        return FormatPlain
    case "stylish":
        return FormatStylish
    case "json":
		return FormatJSON

    default:
        return FormatStylish
    }
}

