package fps

import "github.com/project-flogo/core/data/coerce"

// Note: Not being used currently.
type Image struct {
	Data      []byte `md:"data"`
	ImageType string `md:"imageType"`
	Size      int64  `md:"size"`
}

func (i *Image) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data":      i.Data,
		"imageType": i.ImageType,
		"size":      i.Size,
	}
}

func (i *Image) FromMap(values map[string]interface{}) error {

	var err error
	i.Data, err = coerce.ToBytes(values["data"])
	if err != nil {
		return err
	}
	i.ImageType, err = coerce.ToString(values["imageType"])
	if err != nil {
		return err
	}
	i.Size, err = coerce.ToInt64(values["size"])
	if err != nil {
		return err
	}

	return nil
}
