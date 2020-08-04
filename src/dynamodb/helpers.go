package dynamodb

import (
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func buildUpdateExpression(input map[string]interface{}, parentName, exprParentName string) string {

	var s []string

	for k, v := range input {

		keyName := strings.Join([]string{":", exprParentName, strings.ToLower(k)}, "")

		switch t := v.(type) {
		case string, bool, int:
			currAttrName := k
			if len(parentName) > 0 {
				currAttrName = strings.Join([]string{parentName, currAttrName}, ".")
			}
			s = append(s, strings.Join([]string{currAttrName, "=", keyName}, " "))
		case map[string]interface{}:
			nextExprParentName := strings.TrimPrefix(keyName, ":")
			nextParentName := k
			if len(parentName) > 0 {
				nextParentName = strings.Join([]string{parentName, k}, ".")
			}
			s = append(s, buildUpdateExpression(t, nextParentName, nextExprParentName))
		}

	}
	return strings.Join(s, ", ")
}

func buildExpressionAttributeValues(input map[string]interface{}, parentName string) map[string]*dynamodb.AttributeValue {

	m := make(map[string]*dynamodb.AttributeValue)

	for k, v := range input {

		keyName := strings.Join([]string{":", parentName, strings.ToLower(k)}, "")

		switch t := v.(type) {
		case map[string]interface{}:
			nextName := strings.TrimPrefix(keyName, ":")
			m[keyName] = &dynamodb.AttributeValue{
				M: buildExpressionAttributeValues(t, nextName),
			}
		case string:
			m[keyName] = &dynamodb.AttributeValue{
				S: aws.String(t),
			}
		case bool:
			m[keyName] = &dynamodb.AttributeValue{
				BOOL: aws.Bool(t),
			}
		case int:
			i := strconv.Itoa(t)
			m[keyName] = &dynamodb.AttributeValue{
				N: aws.String(i),
			}
			// case []interface{}:
			// 	nextName := strings.TrimPrefix(t, ":")
			// 	m[keyName] = &dynamodb.AttributeValue{
			// 		L: buildExpressionAttributeValues(t, nextName),
			// 	}
		}
	}
	return m
}
