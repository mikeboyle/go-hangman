package utils

import "strings"

func Contains(list []string, item string) bool {
  for _, element := range list {
    if strings.ToUpper(element) == strings.ToUpper(item) {
      return true
    }
  }
  return false
}
