package controller

import (
	"io"
	"io/ioutil"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/astaxie/beego"
)

var (
	value float64
)

// OperationValue is the request DTO for the action supported by calculator service.
type OperationValue struct {
	Value float64	`json:"Value"`
}

// ResultValue is the response DTO for the action supported by calculator service.
type ResultValue struct {
	Result float64	`json:"Result"`
}

// Calculator is the controller for Calculator service.
type Calculator struct {
	beego.Controller
}

// Batch is the entry for /$batch.
func (c *Calculator) Batch() {
	fmt.Printf("Batch start.\n")
	m, err := c.Ctx.Request.MultipartReader()
	if err != nil {
		fmt.Printf("MultiPartReader() failed, err = %v\n", err)
		c.returnError()
	}
	for {
		p, err := m.NextPart()
		if err == io.EOF {
			c.returnResult()
			return
		}
		if err != nil {
			fmt.Printf("NextPart() failed, err = %v\n", err)
		}
		content, err := ioutil.ReadAll(p)
		if err != nil {
			fmt.Printf("ReadAll() failed, err = %v\n", err)
		}
		fmt.Printf("%q\n", content)
	}
	c.returnResult()
}

// Reset the value to 0.0f.
func (c *Calculator) Reset() {
	fmt.Printf("Reset.\n")
	value = 0;
	c.returnResult()
}

// Add to the addition operation.
func (c *Calculator) Add() {
	if opValue, err := c.getOperationValue(); err != nil {
		c.returnError()
	} else {
		fmt.Printf("%f + %f = ", value, opValue)
		value += opValue
		fmt.Printf("%f\n", value)
		c.returnResult()
	}
}

// Sub do the subtraction operation.
func (c *Calculator) Sub() {
	if opValue, err := c.getOperationValue(); err != nil {
		c.returnError()
	} else {
		fmt.Printf("%f - %f = ", value, opValue)
		value -= opValue
		fmt.Printf("%f\n", value)
		c.returnResult()
	}	
}

// Mul do the multiplication operation.
func (c *Calculator) Mul() {
	if opValue, err := c.getOperationValue(); err != nil {
		c.returnError()
	} else {
		fmt.Printf("%f * %f = ", value, opValue)
		value *= opValue
		fmt.Printf("%f\n", value)
		c.returnResult()
	}	
}

// Div do the division operation.
func (c *Calculator) Div() {
	if opValue, err := c.getOperationValue(); err != nil || opValue == 0 {
		c.returnError()
	} else {
		fmt.Printf("%f / %f = ", value, opValue)
		value /= opValue
		fmt.Printf("%f\n", value)
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