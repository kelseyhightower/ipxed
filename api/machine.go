package api

import "encoding/json"

type Machine struct {
	Name       string `json:"name"`
	MacAddress string `json:"macaddress"`
	Profile    string `json:"profile"`
}

func (m Machine) Save() error {
	return PutToBucket("machines", m.Name, &m)
}

func DeleteMachineByName(name string) error {
	return DeleteFromBucket("machines", name)
}

func GetMachineByName(name string) (Machine, error) {
	var m Machine
	if err := GetFromBucket("machines", name, &m); err != nil {
		return m, err
	}
	return m, nil
}

func GetMachines() ([]Machine, error) {
	machines := make([]Machine, 0)
	err := GetAllFromBucket("machines", func(k, v []byte) error {
		var m Machine
		if err := json.Unmarshal(v, &m); err != nil {
			return err
		}
		machines = append(machines, m)
		return nil
	})
	if err != nil {
		return machines, err
	}
	return machines, nil
}
