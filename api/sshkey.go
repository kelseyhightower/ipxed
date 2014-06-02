package api

import "encoding/json"

type SSHKey struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Key         string `json:"key"`
}

func (s SSHKey) Save() error {
	return PutToBucket("sshkeys", s.Name, &s)
}

func DeleteSSHKeyByName(name string) error {
	return DeleteFromBucket("sshkeys", name)
}

func GetSSHKeyByName(name string) (SSHKey, error) {
	var s SSHKey
	if err := GetFromBucket("sshkeys", name, &s); err != nil {
		return s, err
	}
	return s, nil
}

func GetSSHKeys() ([]SSHKey, error) {
	sshkeys := make([]SSHKey, 0)
	err := GetAllFromBucket("sshkeys", func(k, v []byte) error {
		var s SSHKey
		if err := json.Unmarshal(v, &s); err != nil {
			return err
		}
		sshkeys = append(sshkeys, s)
		return nil
	})
	if err != nil {
		return sshkeys, err
	}
	return sshkeys, nil
}
