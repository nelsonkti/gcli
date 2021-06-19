package helper

import (
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/labstack/echo/v4"
	"sync"
)

type Binder struct {
	once          sync.Once
	validate      *validator.Validate
	DefaultBinder echo.DefaultBinder
}

var uni *ut.UniversalTranslator

func (m *Binder) Validate(v interface{}) error {
	m.lazyInit()
	err := m.validate.Struct(v)
	zh := zh.New()
	en := en.New()
	uni = ut.New(en, zh)

	trans, _ := uni.GetTranslator("zh")
	if err != nil {
		_ = zh_translations.RegisterDefaultTranslations(m.validate, trans)
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			return errors.New(e.Translate(trans))
		}

	}
	return err
}

func (m *Binder) lazyInit() {
	m.once.Do(func() {
		m.validate = validator.New()
	})
}

// 数据绑定以及校验
func (m *Binder) Bind(v interface{}, c echo.Context) error {
	if err := m.DefaultBinder.Bind(v, c); err != nil {
		return err
	}

	if err := m.Validate(v); err != nil {
		return err
	}

	return nil
}
