package formatters

import (
    "fmt"
    "sort"
    "strings"
    "github.com/bobkoffandrei/go-project-244/models"
)

func FormatStylish(nodes []models.Node) string {
	return FormatStylishWithDepth(nodes, 0)
}

func FormatStylishWithDepth(nodes []models.Node, depth int) string {
    var result string
    indent := strings.Repeat("    ", depth)
    
    for _, node := range nodes {
        switch node.Type {
        case NESTED:

            result += fmt.Sprintf("%s  %s: {\n", indent, node.Key)
            result += FormatStylishWithDepth(node.Children, depth+1)
            result += fmt.Sprintf("%s  }\n", indent)
            
        case UNCHANGED:
            result += fmt.Sprintf("%s  %s: %v\n", indent, node.Key, node.Value)
            
        case ADDED:

            if isMap(node.Value) {
                result += fmt.Sprintf("%s+ %s: {\n", indent, node.Key)
                result += formatMap(node.Value.(map[string]any), indent+"    ")
                result += fmt.Sprintf("%s  }\n", indent)
            } else {
                result += fmt.Sprintf("%s+ %s: %v\n", indent, node.Key, node.Value)
            }
            
        case REMOVED:
            if isMap(node.OldValue) {
                result += fmt.Sprintf("%s- %s: {\n", indent, node.Key)
                result += formatMap(node.OldValue.(map[string]any), indent+"    ")
                result += fmt.Sprintf("%s  }\n", indent)
            } else {
                result += fmt.Sprintf("%s- %s: %v\n", indent, node.Key, node.OldValue)
            }
            
        case CHANGED:

            if isMap(node.OldValue) {
                result += fmt.Sprintf("%s- %s: {\n", indent, node.Key)
                result += formatMap(node.OldValue.(map[string]any), indent+"    ")
                result += fmt.Sprintf("%s  }\n", indent)
            } else {
                result += fmt.Sprintf("%s- %s: %v\n", indent, node.Key, node.OldValue)
            }
            
            if isMap(node.Value) {
                result += fmt.Sprintf("%s+ %s: {\n", indent, node.Key)
                result += formatMap(node.Value.(map[string]any), indent+"    ")
                result += fmt.Sprintf("%s  }\n", indent)
            } else {
                result += fmt.Sprintf("%s+ %s: %v\n", indent, node.Key, node.Value)
            }
        }
    }
    
    return result
}

func formatMap(m map[string]any, indent string) string {
    var result string
    
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    for _, key := range keys {
        value := m[key]
        if isMap(value) {
            result += fmt.Sprintf("%s%s: {\n", indent, key)
            result += formatMap(value.(map[string]any), indent+"  ")
            result += fmt.Sprintf("%s}\n", indent)
        } else {
            result += fmt.Sprintf("%s%s: %v\n", indent, key, value)
        }
    }
    
    return result
}

func isMap(v any) bool {
    _, ok := v.(map[string]any)
    return ok
}

const (
    UNCHANGED = "unchanged"
    ADDED     = "added"
    REMOVED   = "removed"
    CHANGED   = "changed"
    NESTED    = "nested"
)