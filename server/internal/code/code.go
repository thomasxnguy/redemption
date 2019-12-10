package code

import (
	"bytes"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/redemption/server/internal/config"
	"github.com/trustwallet/redemption/server/pkg/codegen"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"html/template"
	"time"
)

func CreateLinks(count int, provider string, assets redemption.Assets) (redemption.Links, error) {
	if len(assets.Assets) == 0 {
		return nil, errors.E("invalid assets")
	}
	logger.Info("generating links...", logger.Params{"count": count, "provider": provider})
	codes, err := generateCodes(count)
	if err != nil {
		return nil, err
	}
	links := make(redemption.Links, 0)
	for _, code := range codes {
		url, err := getUrl(code, provider)
		if err != nil {
			return nil, err
		}
		links = append(links, &redemption.Link{
			Link:           url,
			Code:           code,
			Provider:       provider,
			Valid:          true,
			ExpirationDate: time.Time{},
			CreatedDate:    time.Now(),
			Asset:          assets,
		})
	}
	logger.Info("links generated", logger.Params{"count": count, "provider": provider})
	return links, nil
}

func generateCodes(count int) ([]string, error) {
	if count == 0 {
		return nil, errors.E("invalid code number")
	}
	return codegen.GenerateCodes(count, &codegen.Options{
		Charset: codegen.CharsetAlphanumeric,
		Prefix:  config.Configuration.Code.Prefix,
		Pattern: config.Configuration.Code.Pattern,
	})
}

func getUrl(code, provider string) (string, error) {
	tpl := template.New("url")
	tpl, err := tpl.Parse(config.Configuration.Link.Url)
	if err != nil {
		return "", err
	}

	data := struct {
		Code     string
		Provider string
	}{
		Code:     code,
		Provider: provider,
	}

	var out bytes.Buffer
	err = tpl.Execute(&out, data)
	if err != nil {
		return "", err
	}
	u := out.String()
	return u, nil
}
