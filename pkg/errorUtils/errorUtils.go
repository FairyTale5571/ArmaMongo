package errorUtils

import "errors"

var (
	ErrorCantCreateBucket   = errors.New("can't create bucket")
	ErrorCantUnmarshallBolt = errors.New("can't unmarshall bolt")
	ErrorNotFound           = errors.New("not found")
)
