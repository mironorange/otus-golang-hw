package hw09structvalidator

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrIsNotStruct           = errors.New("is not struct")
	ErrHasUnknownValidator   = errors.New("has unknown validator")
	ErrHasInvalidValidator   = errors.New("has invalid validator")
	ErrWhanCreatingValidator = errors.New("when creating validator")

	ErrValidateLessValue       = errors.New("value is less than expected")
	ErrValidateOutOfScopeValue = errors.New("value out of scope")
	ErrValidateGreaterValue    = errors.New("value is greater than expected")
	ErrValidateStringLength    = errors.New("string length does not match expected")
	ErrValidateRegularMatch    = errors.New("value does not match regular expression")

	NilValidationError ValidationError
)

type (
	// Все валидаторы объедены общим интерфейсом.
	FieldValidator interface {
		ValidateField(field reflect.StructField, value reflect.Value) ValidationError
	}

	// Валидатор проверяет вхождение значения в множество.
	InStringFieldValidator struct {
		in map[string]struct{}
	}

	// Валидатор проверяет длину строки.
	LengthStringFieldValidator struct {
		length int
	}

	// Валидатор проверяет соответствие регулярному выражению
	RegexpStringFieldValidator struct {
		regexp *regexp.Regexp
	}

	// Валидатор проверяет вхождение значения в множество.
	InIntFieldValidator struct {
		in map[int]struct{}
	}

	// Валидатор проверяет что значение больше заданного в описании тега.
	MinIntFieldValidator struct {
		min int
	}

	// Валидатор проверяет что значение меньше заданного в описании тега.
	MaxIntFieldValidator struct {
		max int
	}

	// Ошибка валидации возникающая если проверка не прошла.
	ValidationError struct {
		Field string
		Err   error
	}

	// Ошибки валидации объединенные в массив
	ValidationErrors []ValidationError
)

func NewMinIntValidator(cond string) (FieldValidator, error) {
	var validator FieldValidator

	number, err := strconv.Atoi(cond)
	if err != nil {
		return validator, err
	}

	validator = &MinIntFieldValidator{
		min: number,
	}

	return validator, nil
}

func (fv *MinIntFieldValidator) ValidateField(field reflect.StructField, value reflect.Value) ValidationError {
	switch field.Type.Kind() {
	case reflect.Int:
		if value.Int() < int64(fv.min) {
			return ValidationError{
				Field: field.Name,
				Err:   ErrValidateLessValue,
			}
		}
	case reflect.Slice:
		for _, v := range value.Interface().([]int64) {
			if v < int64(fv.min) {
				return ValidationError{
					Field: field.Name,
					Err:   ErrValidateLessValue,
				}
			}
		}
	default:
		return ValidationError{
			Field: field.Name,
			Err:   ErrValidateLessValue,
		}
	}

	return NilValidationError
}

func NewInIntValidator(cond string) (FieldValidator, error) {
	var validator FieldValidator

	inValues := make(map[int]struct{})
	for _, value := range strings.Split(cond, ",") {
		val, err := strconv.Atoi(value)
		if err != nil {
			return validator, err
		}
		inValues[val] = struct{}{}
	}

	validator = &InIntFieldValidator{
		in: inValues,
	}

	return validator, nil
}

func (fv *InIntFieldValidator) ValidateField(field reflect.StructField, value reflect.Value) ValidationError {
	switch field.Type.Kind() {
	case reflect.Int:
		if _, ok := fv.in[int(value.Int())]; !ok {
			return ValidationError{
				Field: field.Name,
				Err:   ErrValidateOutOfScopeValue,
			}
		}
	case reflect.Slice:
		for _, v := range value.Interface().([]int) {
			if _, ok := fv.in[v]; !ok {
				return ValidationError{
					Field: field.Name,
					Err:   ErrValidateOutOfScopeValue,
				}
			}
		}
	default:
		return ValidationError{
			Field: field.Name,
			Err:   ErrValidateOutOfScopeValue,
		}
	}

	return NilValidationError
}

func NewMaxIntValidator(cond string) (FieldValidator, error) {
	var validator FieldValidator

	m, err := strconv.Atoi(cond)
	if err != nil {
		return validator, err
	}

	validator = &MaxIntFieldValidator{
		max: m,
	}

	return validator, nil
}

func (fv *MaxIntFieldValidator) ValidateField(field reflect.StructField, value reflect.Value) ValidationError {
	switch field.Type.Kind() {
	case reflect.Int:
		if value.Int() > int64(fv.max) {
			return ValidationError{
				Field: field.Name,
				Err:   ErrValidateGreaterValue,
			}
		}
	case reflect.Slice:
		for _, v := range value.Interface().([]int64) {
			if v > int64(fv.max) {
				return ValidationError{
					Field: field.Name,
					Err:   ErrValidateGreaterValue,
				}
			}
		}
	default:
		return ValidationError{
			Field: field.Name,
			Err:   ErrValidateGreaterValue,
		}
	}

	return NilValidationError
}

func NewLengthStringValidator(cond string) (FieldValidator, error) {
	var validator FieldValidator

	length, err := strconv.Atoi(cond)
	if err != nil {
		return validator, err
	}

	validator = &LengthStringFieldValidator{
		length: length,
	}

	return validator, nil
}

func (fv *LengthStringFieldValidator) ValidateField(field reflect.StructField, value reflect.Value) ValidationError {
	switch field.Type.Kind() {
	case reflect.String:
		if len(value.String()) != fv.length {
			return ValidationError{
				Field: field.Name,
				Err:   ErrValidateStringLength,
			}
		}
	case reflect.Slice:
		for _, v := range value.Interface().([]string) {
			if len(v) != fv.length {
				return ValidationError{
					Field: field.Name,
					Err:   ErrValidateStringLength,
				}
			}
		}
	default:
		return ValidationError{
			Field: field.Name,
			Err:   ErrValidateStringLength,
		}
	}

	return NilValidationError
}

func NewInStringValidator(values string) (FieldValidator, error) {
	var validator FieldValidator

	inValues := make(map[string]struct{})
	for _, val := range strings.Split(values, ",") {
		inValues[val] = struct{}{}
	}

	validator = &InStringFieldValidator{
		in: inValues,
	}

	return validator, nil
}

func (fv *InStringFieldValidator) ValidateField(field reflect.StructField, value reflect.Value) ValidationError {
	switch field.Type.Kind() {
	case reflect.String:
		if _, ok := fv.in[value.String()]; !ok {
			return ValidationError{
				Field: field.Name,
				Err:   ErrValidateOutOfScopeValue,
			}
		}
	case reflect.Slice:
		for _, v := range value.Interface().([]string) {
			if _, ok := fv.in[v]; !ok {
				return ValidationError{
					Field: field.Name,
					Err:   ErrValidateOutOfScopeValue,
				}
			}
		}
	default:
		return ValidationError{
			Field: field.Name,
			Err:   ErrValidateOutOfScopeValue,
		}
	}

	return NilValidationError
}

func NewRegexpStringFieldValidator(exp string) (FieldValidator, error) {
	var validator FieldValidator

	var reg, err = regexp.Compile(exp)
	if err != nil {
		return validator, err
	}

	validator = &RegexpStringFieldValidator{
		regexp: reg,
	}

	return validator, nil
}

func (fv *RegexpStringFieldValidator) ValidateField(field reflect.StructField, value reflect.Value) ValidationError {
	switch field.Type.Kind() {
	case reflect.String:
		if !fv.regexp.MatchString(value.String()) {
			return ValidationError{
				Field: field.Name,
				Err:   ErrValidateRegularMatch,
			}
		}
	case reflect.Slice:
		for _, v := range value.Interface().([]string) {
			if !fv.regexp.MatchString(v) {
				return ValidationError{
					Field: field.Name,
					Err:   ErrValidateRegularMatch,
				}
			}
		}
	default:
		return ValidationError{
			Field: field.Name,
			Err:   ErrValidateRegularMatch,
		}
	}

	return NilValidationError
}

func prepareNotSupportFieldValidator(f reflect.StructField) ([]FieldValidator, error) {
	var validators []FieldValidator

	_, ok := f.Tag.Lookup("validate")
	if !ok {
		return validators, nil
	}

	return validators, ErrHasUnknownValidator
}

func prepareStringFieldValidator(f reflect.StructField) ([]FieldValidator, error) {
	var validators []FieldValidator

	tags, ok := f.Tag.Lookup("validate")
	if !ok {
		return validators, nil
	}

	validators = make([]FieldValidator, 0)
	for _, validationRules := range strings.Split(tags, "|") {
		if ruleСondition := strings.Split(validationRules, ":"); len(ruleСondition) == 2 {
			switch ruleСondition[0] {
			case "len":
				validator, err := NewLengthStringValidator(ruleСondition[1])
				if err != nil {
					return validators, err
				}
				validators = append(validators, validator)
			case "in":
				validator, err := NewInStringValidator(ruleСondition[1])
				if err != nil {
					return validators, err
				}
				validators = append(validators, validator)
			case "regexp":
				validator, err := NewRegexpStringFieldValidator(ruleСondition[1])
				if err != nil {
					return validators, err
				}
				validators = append(validators, validator)
			default:
				return validators, ErrHasUnknownValidator
			}
		}
	}

	return validators, nil
}

func prepareIntFieldValidator(f reflect.StructField) ([]FieldValidator, error) {
	var validators []FieldValidator

	tags, ok := f.Tag.Lookup("validate")
	if !ok {
		return validators, nil
	}

	validators = make([]FieldValidator, 0)
	for _, validationRules := range strings.Split(tags, "|") {
		if ruleСondition := strings.Split(validationRules, ":"); len(ruleСondition) == 2 {
			switch ruleСondition[0] {
			case "min":
				validator, err := NewMinIntValidator(ruleСondition[1])
				if err != nil {
					return validators, err
				}
				validators = append(validators, validator)
			case "in":
				validator, err := NewInIntValidator(ruleСondition[1])
				if err != nil {
					return validators, err
				}
				validators = append(validators, validator)
			case "max":
				validator, err := NewMaxIntValidator(ruleСondition[1])
				if err != nil {
					return validators, err
				}
				validators = append(validators, validator)
			default:
				return validators, ErrHasUnknownValidator
			}
		}
	}

	return validators, nil
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Err)
}

func (v ValidationErrors) Error() string {
	var buffer bytes.Buffer

	if len(v) > 0 {
		buffer.WriteString("Errors:\n")
	}
	for _, err := range v {
		buffer.WriteString(fmt.Sprintf("- %s\n", err.Error()))
	}

	return buffer.String()
}

func Validate(v interface{}) error {
	// Переданное тип переданного значения не является валидной структурой или указателем на нее
	r := reflect.Indirect(reflect.ValueOf(v))
	if r.Type().Kind() != reflect.Struct {
		return ErrIsNotStruct
	}

	validators := make(map[string][]FieldValidator, 0)
	for i := 0; i < r.NumField(); i++ {
		structField := r.Type().Field(i)
		fieldValue := r.Field(i)
		switch structField.Type.Kind() {
		case reflect.String:
			fieldValidators, err := prepareStringFieldValidator(structField)
			if err == nil && len(fieldValidators) > 0 {
				validators[structField.Name] = fieldValidators
			} else if err != nil {
				return ErrWhanCreatingValidator
			}
		case reflect.Int:
			fieldValidators, err := prepareIntFieldValidator(structField)
			if err == nil && len(fieldValidators) > 0 {
				validators[structField.Name] = fieldValidators
			} else if err != nil {
				return ErrWhanCreatingValidator
			}
		case reflect.Slice:
			if fieldValue.Len() > 0 {
				switch fieldValue.Index(0).Type().Kind() {
				case reflect.String:
					fieldValidators, err := prepareStringFieldValidator(structField)
					if err == nil && len(fieldValidators) > 0 {
						validators[structField.Name] = fieldValidators
					} else if err != nil {
						return ErrWhanCreatingValidator
					}
				case reflect.Int:
					fieldValidators, err := prepareIntFieldValidator(structField)
					if err == nil && len(fieldValidators) > 0 {
						validators[structField.Name] = fieldValidators
					} else if err != nil {
						return ErrWhanCreatingValidator
					}
				default:
					if _, err := prepareNotSupportFieldValidator(structField); err != nil {
						return ErrWhanCreatingValidator
					}
				}
			}
		case reflect.Ptr:
			if _, err := prepareNotSupportFieldValidator(structField); err != nil {
				return ErrHasUnknownValidator
			}
		case reflect.Struct:
			if _, err := prepareNotSupportFieldValidator(structField); err != nil {
				return ErrHasUnknownValidator
			}
		default:
			if _, err := prepareNotSupportFieldValidator(structField); err != nil {
				return ErrHasUnknownValidator
			}
		}
	}

	vErrors := make(ValidationErrors, 0)
	for name, fieldValidators := range validators {
		field, _ := r.Type().FieldByName(name)
		value := r.FieldByName(name)
		for _, v := range fieldValidators {
			if err := v.ValidateField(field, value); err != NilValidationError {
				vErrors = append(vErrors, err)
			}
		}
	}

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil
}
