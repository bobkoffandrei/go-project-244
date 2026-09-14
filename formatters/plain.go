package formatters

import (
    "fmt"
    "strings"
    "code/models"
)

func FormatPlain(nodes []models.Node) string {
    result := FormatPlainWithPath(nodes, "")
    return strings.TrimSuffix(result, "\n")
}

func FormatPlainWithPath(nodes []models.Node, path string) string {
    var result string
    
    for _, node := range nodes {

        currentPath := path
        if currentPath == "" {
            currentPath = node.Key
        } else {
            currentPath = path + "." + node.Key
        }
        
        switch node.Type {
        case models.NESTED:

            result += FormatPlainWithPath(node.Children, currentPath)
            
        case models.UNCHANGED:
           
            
        case models.ADDED:
            result += formatAddedPlain(currentPath, node.Value)
            
        case models.REMOVED:
            result += formatRemovedPlain(currentPath)
            
        case models.CHANGED:
            result += formatChangedPlain(currentPath, node.OldValue, node.Value)
        }
    }
    
    return result
}


func formatAddedPlain(path string, value interface{}) string {
    return fmt.Sprintf("Property '%s' was added with value: %s\n", 
        path, formatValuePlain(value))
}


func formatRemovedPlain(path string) string {
    return fmt.Sprintf("Property '%s' was removed\n", path)
}


func formatChangedPlain(path string, oldValue, newValue interface{}) string {
    return fmt.Sprintf("Property '%s' was updated. From %s to %s\n", 
        path, formatValuePlain(oldValue), formatValuePlain(newValue))
}


func formatValuePlain(value interface{}) string {
    if value == nil {
        return "null"
    }
    

    if isComplexValue(value) {
        return "[complex value]"
    }
    

    if str, ok := value.(string); ok {
        return "'" + str + "'"
    }
    

    if _, ok := value.(bool); ok {
        return fmt.Sprintf("%v", value)
    }
    

    if _, ok := value.(float64); ok {
        return fmt.Sprintf("%v", value)
    }
    if _, ok := value.(int); ok {
        return fmt.Sprintf("%v", value)
    }
    

    return fmt.Sprintf("%v", value)
}


func isComplexValue(value interface{}) bool {
    if value == nil {
        return false
    }

    if _, ok := value.(map[string]any); ok {
        return true
    }
    if _, ok := value.(map[interface{}]interface{}); ok {
        return true
    }
    

    if _, ok := value.([]interface{}); ok {
        return true
    }
    
    return false
}