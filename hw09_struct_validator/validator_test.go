package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestIntFieldValidator(t *testing.T) {
	var validators []FieldValidator

	// Валидация полей структуры User
	user := User{
		Age: 21,
	}

	// Если значение является указателем, то вернуть значение за указателем
	userValue := reflect.Indirect(reflect.ValueOf(user))

	// Узнать тип атрибута
	ageFieldValue := userValue.FieldByName("Age")
	ageStructField, _ := userValue.Type().FieldByName("Age")
	validators, _ = prepareIntFieldValidator(ageStructField)

	require.True(t, ageFieldValue.Type().Kind() == reflect.Int)
	require.Len(t, validators, 2)
	for _, validator := range validators {
		require.True(t, validator.ValidateField(ageStructField, ageFieldValue) == NilValidationError)
	}

	// Валидация полей структуры Response
	response := Response{
		Code: 301,
	}
	responseValue := reflect.Indirect(reflect.ValueOf(response))
	// Узнать тип атрибута
	codeFieldValue := responseValue.FieldByName("Code")
	codeStructField, _ := responseValue.Type().FieldByName("Code")
	validators, _ = prepareIntFieldValidator(codeStructField)
	require.Len(t, validators, 1)
	for _, validator := range validators {
		require.Error(t, validator.ValidateField(codeStructField, codeFieldValue))
	}
}

func TestStringFieldValidator(t *testing.T) {
	var validators []FieldValidator
	user := User{
		ID:     "007",
		Phones: []string{"71111111111", "7111111111"},
		Email:  "007agent",
		Role:   "troll",
	}

	// Если значение является указателем, то вернуть значение за указателем
	userValue := reflect.Indirect(reflect.ValueOf(user))

	// Узнать тип атрибута
	idFieldValue := userValue.FieldByName("ID")
	idStructField, _ := userValue.Type().FieldByName("ID")
	require.True(t, idFieldValue.Type().Kind() == reflect.String)
	validators, _ = prepareStringFieldValidator(idStructField)
	for _, validator := range validators {
		require.Error(t, validator.ValidateField(idStructField, idFieldValue))
	}

	// Узнать то, что атрибут является слайсом и тип его элементов
	phonesFieldValue := userValue.FieldByName("Phones")
	phonesStructField, _ := userValue.Type().FieldByName("Phones")
	require.True(t, phonesFieldValue.Type().Kind() == reflect.Slice)
	if phonesFieldValue.Type().Kind() == reflect.Slice && phonesFieldValue.Len() > 0 {
		require.True(t, phonesFieldValue.Index(0).Type().Kind() == reflect.String)
	}
	validators, _ = prepareStringFieldValidator(phonesStructField)
	for _, validator := range validators {
		require.Error(t, validator.ValidateField(phonesStructField, phonesFieldValue))
	}

	// Проверить на соответсвие регулярному выражению
	emailFieldValue := userValue.FieldByName("Email")
	emailStructField, _ := userValue.Type().FieldByName("Email")
	validators, _ = prepareStringFieldValidator(emailStructField)
	for _, validator := range validators {
		require.Error(t, validator.ValidateField(emailStructField, emailFieldValue))
	}

	// // Обойти все элементы слайса (через его приведение к исходному типу)
	// phones := phonesField.Interface().([]string)
	// for _, v := range phones {
	// 	fmt.Println(v)
	// }

	// Убедиться в том, что можно работать с значением, основанное на простом типе
	roleFieldValue := userValue.FieldByName("Role")
	roleStructField, _ := userValue.Type().FieldByName("Role")
	require.True(t, roleFieldValue.Type().Kind() == reflect.String)
	validators, _ = prepareStringFieldValidator(roleStructField)
	for _, validator := range validators {
		require.Error(t, validator.ValidateField(roleStructField, roleFieldValue))
	}
}

func TestValidate(t *testing.T) {
	type invalidTag struct {
		age int `validate:"in:admin,stuff"`
	}

	type invalidType struct {
		char byte `validate:"in:admin,stuff"`
	}

	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          "qwerty",
			expectedErr: ErrIsNotStruct,
		},
		{
			in:          123,
			expectedErr: ErrIsNotStruct,
		},
		{
			in: invalidTag{
				age: 5,
			},
			expectedErr: ErrWhanCreatingValidator,
		},
		{
			in: invalidType{
				char: '!',
			},
			expectedErr: ErrHasUnknownValidator,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			require.ErrorIs(t, tt.expectedErr, Validate(tt.in))
		})
	}

	userWithErrors := User{
		ID:     "007",
		Phones: []string{"71111111111", "7111111111"},
		Email:  "007agent",
		Role:   "troll",
	}

	var e ValidationErrors

	validateErrors := Validate(userWithErrors)
	require.ErrorAs(t, validateErrors, &e)
}
