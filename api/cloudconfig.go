package api

import "encoding/json"

type CloudConfig struct {
	Name   string `json:"name"`
	Config string `json:"config"`
}

func (c CloudConfig) Save() error {
	return PutToBucket("cloudconfigs", c.Name, &c)
}

func DeleteCloudConfigByName(name string) error {
	return DeleteFromBucket("cloudconfigs", name)
}

func GetCloudConfigByName(name string) (CloudConfig, error) {
	var c CloudConfig
	if err := GetFromBucket("cloudconfigs", name, &c); err != nil {
		return c, err
	}
	return c, nil
}

func GetCloudConfigs() ([]CloudConfig, error) {
	cloudconfigs := make([]CloudConfig, 0)
	err := GetAllFromBucket("cloudconfigs", func(k, v []byte) error {
		var c CloudConfig
		if err := json.Unmarshal(v, &c); err != nil {
			return err
		}
		cloudconfigs = append(cloudconfigs, c)
		return nil
	})
	if err != nil {
		return cloudconfigs, err
	}
	return cloudconfigs, nil
}
