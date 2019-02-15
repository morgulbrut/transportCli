/*
Copyright © 2019 morgulbrut
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file or
 http://www.wtfpl.net/ for more details.
*/

package parseJSON

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ParseLocation(data []byte) RespLocation {
	var resp RespLocation

	err := json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Body\n%s\n", resp)

	return resp
}

func ParseStation(data []byte) RespStation {
	var resp RespStation

	err := json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Println("error:", err)
	}
	return resp
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
