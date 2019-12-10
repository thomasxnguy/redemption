package message

import (
	"bytes"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"math"
	"strconv"
)

type message struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Image       string `yaml:"image"`
}

func GetMessage() (message, error) {
	var msg message
	yamlFile, err := ioutil.ReadFile("message.yaml")
	if err != nil {
		return message{}, errors.E(err, "GetMessage read file error")
	}
	err = yaml.Unmarshal(yamlFile, &msg)
	if err != nil {
		return message{}, errors.E(err, "GetMessage unmarshal file error")
	}
	return msg, nil
}

func (m message) GetTitle() string {
	return m.Title
}

func (m message) GetImage() string {
	return m.Image
}

func (m message) GetDescription(assets []*redemption.Asset, decimals uint) string {
	values := ""
	for i, asset := range assets {
		a := floatValue(asset.Amount, decimals)
		value := strconv.FormatFloat(a, 'f', -1, 64)
		if i == 0 {
			values += fmt.Sprintf("%s %s", value, asset.TokenId)
			continue
		}
		if i == len(assets)-1 {
			values += fmt.Sprintf(" and %s %s", value, asset.TokenId)
			continue
		}
		values += fmt.Sprintf(", %s %s", value, asset.TokenId)
	}
	msg, err := replaceValues(m.Description, values)
	if err != nil {
		return m.Description
	}
	return msg
}

func replaceValues(text, values string) (string, error) {
	tpl := template.New("url")
	tpl, err := tpl.Parse(text)
	if err != nil {
		return "", err
	}

	data := struct {
		Values string
	}{Values: values}

	var out bytes.Buffer
	err = tpl.Execute(&out, data)
	if err != nil {
		return "", err
	}
	u := out.String()
	return u, nil
}

func floatValue(value int64, decimals uint) float64 {
	pow := math.Pow(10, float64(decimals))
	return float64(value) / pow
}
