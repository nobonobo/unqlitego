package unqlitego

/*
 * Additional Database Functions to Aid Transaciotn From GKVLite
 */
import (
	"encoding/json"
)

type MarshalFunction func(interface{}) ([]byte, error)
type UnmarshalFunction func([]byte, interface{}) error

func (t *Database) Marshal() MarshalFunction {
	if t.marshal != nil {
		return t.marshal
	} else {
		return json.Marshal
	}
}

func (t *Database) SetMarshal(override MarshalFunction) {
	t.marshal = override
}

func (t *Database) Unmarshal() UnmarshalFunction {
	if t.unmarshal != nil {
		return t.unmarshal
	} else {
		return json.Unmarshal
	}
}

func (t *Database) SetUnmarshal(override UnmarshalFunction) {
	t.unmarshal = override
}

func (t *Database) SetObject(key string, object interface{}) error {
	byteObject, err := t.Marshal()(object)

	if err != nil {
		return err
	}

	err = t.Begin()
	if err != nil {
		return err
	}

	err = t.Store([]byte(key), byteObject)
	if err != nil {
		t.Rollback()
		return err
	}

	err = t.Commit()
	if err != nil {
		t.Rollback()
		return err
	}

	return nil
}

func (t *Database) GetObject(key string, object interface{}) error {

	byteObject, err := t.Fetch([]byte(key))
	if err != nil {
		if err == UnQLiteError(-6) {
			//Not Found is not an error in my world...
			return nil
		}
		return err
	}

	if byteObject != nil {
		err = t.Unmarshal()(byteObject, &object)
		if err != nil {
			return err
		}
	}

	return nil
}
