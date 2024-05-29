package errors

import (
    "encoding/json"
    "fmt"
    "log"
)

func marshalApiError(e IBaseError, errorName string) []byte {
    response, err := json.MarshalIndent(e, "", "  ")

    if err != nil {
        log.Println(fmt.Sprintf("[ERROR] failed decoding '%s' struct because of: %e", errorName, err))
        return []byte("")
    }

    return response
}
