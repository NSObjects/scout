/*
 *
 * encoder.go
 * encoder
 *
 * Created by lintao on 2022/6/23 5:16 PM
 * Copyright © 2020-2022 LINTAO. All rights reserved.
 *
 */

package encoder

type Encoder interface {
	GetName() string
	Encode(data string) (string, error)
}

var Encoders map[string]Encoder

func init() {
	Encoders = make(map[string]Encoder)
	Encoders[DefaultURLEncoder.GetName()] = DefaultURLEncoder
}

func Apply(encoderName, data string) (string, error) {
	ec, ok := Encoders[encoderName]
	if !ok {
		return "", nil
	}
	ret, err := ec.Encode(data)
	if err != nil {
		return "", err
	}
	return ret, nil
}
