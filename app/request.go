package app

import (
    "errors"
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/julienschmidt/httprouter"
)


type Request struct {
    *http.Request
    Env map[string]interface{}
    PathParams map[string]string
}

// DecodeJsonPayload reads the request body and decodes the JSON using json.Unmarshal.
func (self *Request) DecodeJsonPayload(v interface{}) error {
    content, err := ioutil.ReadAll(self.Body)
    self.Body.Close()

    if err != nil {
        return err
    }

    if len(content) == 0 {
        return errors.New("empty body")
    }

    err = json.Unmarshal(content, v)

    if err != nil {
        return err
    }

    return nil
}

// change `*http.Request` to `*Request`
func AdaptRequest(r *http.Request, ps httprouter.Params) *Request {
    ps_map := map[string]string{}

    for _, param := range ps{
        ps_map[param.Key] = param.Value
    }

    return &Request{
        r,
        map[string]interface{}{},
        ps_map,
    }
}
