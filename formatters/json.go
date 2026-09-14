package formatters

import (
    "fmt"
    "code/models"
		"encoding/json"
)



func FormatJSON(nodes []models.Node) string {
    diff := buildJSONDiff(nodes)
    result := map[string]interface{}{
        "diff": diff,   // ← обёртка с ключом "diff"
    }
    jsonBytes, err := json.MarshalIndent(result, "", "  ")
    if err != nil {
        return fmt.Sprintf(`{"error": "failed to marshal JSON: %v"}`, err)
    }
    return string(jsonBytes)
}

type JSONDiff struct {
	Type     string      `json:"type,omitempty"`
	Key      string      `json:"key,omitempty"`
	Value    interface{} `json:"value,omitempty"`
	OldValue interface{} `json:"oldValue,omitempty"`
	Children []JSONDiff  `json:"children,omitempty"`
}

func buildJSONDiff(nodes []models.Node) []JSONDiff {
	var result []JSONDiff

	for _, node := range nodes {
		diff := JSONDiff{
			Type: node.Type,
			Key:  node.Key,
		}

		switch node.Type {
		case models.NESTED:
			diff.Children = buildJSONDiff(node.Children)

		case models.UNCHANGED:
			diff.Value = node.Value

		case models.ADDED:
			diff.Value = node.Value

		case models.REMOVED:
			diff.OldValue = node.OldValue

		case models.CHANGED:
			diff.OldValue = node.OldValue
			diff.Value = node.Value
		}

		result = append(result, diff)
	}

	return result
}