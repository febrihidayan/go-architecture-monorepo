package utils

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func TimeUTC() time.Time {
	return time.Now().UTC()
}

/**
 * create password hash
 */
func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

/**
 * validation password hash
 */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/**
 * map query params to struct
 */
func MapQueryParams(r *http.Request, data interface{}) error {
	s := reflect.ValueOf(data).Elem()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		tp := f.Type().Kind()

		name := s.Type().Field(i).Tag.Get("query")
		if name == "" {
			return errors.New("query name tag is empty")
		}

		value := r.URL.Query().Get(name)
		if value == "" {
			value = s.Type().Field(i).Tag.Get("default")
		}

		if value != "" {

			switch tp {
			case reflect.Slice:
				values := strings.Split(value, ",")
				f.Set(reflect.ValueOf(values))
			case reflect.Int:
				intValue, _ := strconv.Atoi(value)
				f.Set(reflect.ValueOf(intValue))
			case reflect.Bool:
				boolValue, _ := strconv.ParseBool(value)
				f.Set(reflect.ValueOf(boolValue))
			default:
				f.Set(reflect.ValueOf(value).Convert(f.Type()))
			}
		}
	}

	return nil
}

// Contains reports whether substr is within s.
func ContainsString(s []string, str string) bool {
	for _, v := range s {
		if strings.Contains(strings.ToLower(v), strings.ToLower(str)) {
			return true
		}
	}

	return false
}
