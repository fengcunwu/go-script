package main

import (
	"bytes"
	"net/http"
)

func main() {
	input := `{
    "env": "pre",
    "configs": [
        {
            "module": "generic_module",
            "key": "async_log_ratio",
            "value": "100"
        },
        {
            "module": "generic_module",
            "key": "sync_log_ratio",
            "value": "39"
        },
        {
            "module": "generic_module",
            "key": "feed_log_ratio",
            "value": "100"
        }
    ]
}`

	reader := bytes.NewReader([]byte(input))
	_, err := http.Post("http://127.0.0.1:8000/v1/uve/sync_config", "application/json;charset=UTF-8", reader)
	if err != nil {
		panic(err)
	}
}
