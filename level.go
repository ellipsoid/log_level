package log_level

import "fmt"

type Level int

const (
	TRACE Level = iota
	DEBUG Level = iota
	INFO  Level = iota
	WARN  Level = iota
	ERROR Level = iota
  INVALID Level = iota // DO NOT USE
)

// GetLevel accepts the name of a logging level and returns
// the appropriate Level object. If no matching Level object 
// is found, GetLevel returns an error.
func GetLevel(levelStr string) (Level, error) {
  switch {
  case levelStr == "TRACE":
    return TRACE, nil
  case levelStr == "DEBUG":
    return DEBUG, nil
  case levelStr == "INFO":
    return INFO, nil
  case levelStr == "WARN":
    return WARN, nil
  case levelStr == "ERROR":
    return ERROR, nil
  }

  return INVALID, fmt.Errorf("Unrecognized Log Level: %v", levelStr)
}
