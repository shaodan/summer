package summer

import "encoding/json"

func Get(string ) {

}


type Content interface {
	OK() (bool, error)
}

type Response struct {
	data []byte
}

func (r *Response) Json(c *Content) error {
	err := json.Unmarshal(r.data, c)
	if err != nil {
		return err
	}
	return nil
}
