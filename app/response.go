package app

import (
    "net/http"
    "encoding/json"
)


type ResponseWriter interface {
    http.ResponseWriter
    WriteJson(interface{}) error
    EncodeJson(interface{}) ([]byte, error)
    getContentCopy() []byte
}


type responseWriter struct {
    http.ResponseWriter
    contentCopy []byte
    wrote       bool
}

func (self *responseWriter) EncodeJson(v interface{}) ([]byte, error) {
    b, err := json.Marshal(v)
    if err != nil {
        return nil, err
    }
    return b, nil
}

func (self *responseWriter) WriteHeader(code int) {
    if self.Header().Get("Content-Type") == "" {
        self.Header().Set("Content-Type", "application/json; charset=utf-8")
    }

    self.ResponseWriter.WriteHeader(code)

    self.wrote = true
}

func (self *responseWriter) WriteJson(v interface{}) error {
    b, err := self.EncodeJson(v)
    if err != nil {
        return err
    }

    // save content
    self.contentCopy = b

    if !self.wrote {
        self.WriteHeader(200)
    }

    _, err = self.Write(b)
    if err != nil {
        return err
    }

    return nil
}

func (self *responseWriter) getContentCopy() ([]byte) {
    return self.contentCopy
}

// change `http.ResponseWriter` to `ResponseWriter`
func AdaptResponse(w http.ResponseWriter) ResponseWriter {
    return &responseWriter{ResponseWriter: w}
}
