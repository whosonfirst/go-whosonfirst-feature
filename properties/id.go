package properties

import (
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-feature"
	"github.com/whosonfirst/go-whosonfirst-feature/constants"
)

func Id(body []byte) (int64, error) {

	rsp := gjson.GetBytes(body, "properties.wof:id")

	if !rsp.Exists() {
		return 0, feature.PropertyNotFoundError("wof:id")
	}

	wof_id := rsp.Int()

	if wof_id < 0 {

		switch wof_id {
		case constants.MULTIPLE_PARENTS, constants.MULTIPLE_NEIGHBOURHOODS, constants.ITS_COMPLICATED, constants.UNKNOWN:
			// pass
		default:
			return 0, fmt.Errorf("Invalid or unrecognized ID value (%d)", wof_id)
		}
	}

	return wof_id, nil
}
