package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

var (
	value float64
)

// OperationValue is the request DTO for the action supported by calculator service.
type OperationValue struct {
	Value float64 `json:"Value"`
}

// ResultValue is the response DTO for the action supported by calculator service.
type ResultValue struct {
	Result float64 `json:"Result"`
}

// Calculator is the controller for Calculator service.
type Calculator struct {
	beego.Controller
}

// Batch is the entry for /$batch.
func (c *Calculator) Batch() {
	log.Printf("Batch start.\n")
	// make sure it's multipart/mixed.
	mediaType, params, err := mime.ParseMediaType(
		c.Ctx.Request.Header.Get("Content-Type"),
	)
	if err != nil {
		log.Printf("ParseMediaType() failed, err =%v\n", err)
		c.returnError()
	}
	if mediaType != "multipart/mixed" {
		log.Printf("Not supported media type. %s\n", mediaType)
		c.returnError()
	}
	// get the request.
	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		log.Printf("Read request failed, err = %v\n", err)
		c.returnError()
	}
	mr := multipart.NewReader(strings.NewReader(string(body)), params["boundary"])
	// for each part convert it to request.
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			c.returnResult()
			break
		}
		if err != nil {
			log.Printf("NextPart() failed, err = %v\n", err)
			c.returnError()
			break
		}
		// readbuffer := bytes.NewBuffer([]byte("GET / HTTP/1.1\r\nheader:foo\r\n\r\n"))
		// reader := bufio.NewReader(readbuffer)
		// req, err := http.ReadRequest(reader)
		// log.Printf("method: %s, host: %s\n", req.Method, req.Host)

		_bufio := bufio.NewReader(p)
		for {
			b, err := _bufio.ReadByte()
			if err == io.EOF {
				break
			}
			if b == '\r' {
				fmt.Printf("(\r")
			} else if b == '\n' {
				fmt.Printf(")\n")
			} else {
				fmt.Printf("%c", b)
			}
		}
		// localRequest, err := http.ReadRequest(_bufio)
		// if err != io.EOF && err != nil {
		// 	log.Printf("ReadRequest() failed, err = %v\n", err)
		// 	c.returnError()
		// 	break
		// }
		// log.Printf("method: %s, host: %s\n", localRequest.Method, localRequest.Host)
	}
}

// Reset the value to 0.0f.
func (c *Calculator) Reset() {
	log.Printf("Reset.\n")
	value = 0
	c.returnResult()
}

// Add to the addition operation.
func (c *Calculator) Add() {
	if opValue, err := c.getOperationValue(); err != nil {
		c.returnError()
	} else {
		log.Printf("%f + %f = ", value, opValue)
		value += opValue
		log.Printf("%f\n", value)
		c.returnResult()
	}
}

// Sub do the subtraction operation.
func (c *Calculator) Sub() {
	if opValue, err := c.getOperationValue(); err != nil {
		c.returnError()
	} else {
		log.Printf("%f - %f = ", value, opValue)
		value -= opValue
		log.Printf("%f\n", value)
		c.returnResult()
	}
}

// Mul do the multiplication operation.
func (c *Calculator) Mul() {
	if opValue, err := c.getOperationValue(); err != nil {
		c.returnError()
	} else {
		log.Printf("%f * %f = ", value, opValue)
		value *= opValue
		log.Printf("%f\n", value)
		c.returnResult()
	}
}

// Div do the division operation.
func (c *Calculator) Div() {
	if opValue, err := c.getOperationValue(); err != nil || opValue == 0 {
		c.returnError()
	} else {
		log.Printf("%f / %f = ", value, opValue)
		value /= opValue
		log.Printf("%f\n", value)
		c.returnResult()
	}
}

func (c *Calculator) getOperationValue() (float64, error) {
	request := &OperationValue{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, request); err != nil {
		return 0, err
	}
	return request.Value, nil
}

// returnResult return the result as JSON.
func (c *Calculator) returnResult() {
	ret := &ResultValue{
		Result: value,
	}
	c.Data["json"] = ret
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.ServeJSON()
	return
}

// returnError return the error response.
func (c *Calculator) returnError() {
	c.Ctx.Output.SetStatus(http.StatusBadRequest)
	c.ServeJSON()
	return
}
