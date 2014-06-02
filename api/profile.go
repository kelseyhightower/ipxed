package api

import "encoding/json"

type Profile struct {
	Name            string `json:"name"`
	CloudConfig     string `json:"cloud_config"`
	Console         string `json:"console"`
	CoreosAutologin string `json:"coreos_autologin"`
	RootFstype      string `json:"rootfstype"`
	Root            string `json:"root"`
	SSHKey          string `json:"sshkey"`
	Version         string `json:"version"`
}

func (p Profile) Save() error {
	return PutToBucket("profiles", p.Name, &p)
}

func DeleteProfileByName(name string) error {
	return DeleteFromBucket("profiles", name)
}

func GetProfileByName(name string) (Profile, error) {
	var p Profile
	if err := GetFromBucket("profiles", name, &p); err != nil {
		return p, err
	}
	return p, nil
}

func GetProfiles() ([]Profile, error) {
	profiles := make([]Profile, 0)
	err := GetAllFromBucket("profiles", func(k, v []byte) error {
		var p Profile
		if err := json.Unmarshal(v, &p); err != nil {
			return err
		}
		profiles = append(profiles, p)
		return nil
	})
	if err != nil {
		return profiles, err
	}
	return profiles, nil
}
