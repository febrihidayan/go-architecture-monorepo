package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
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

// ref: https://github.com/Golang-Coach/Lessons/blob/master/GoMailer/template_helper.go
func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	// ref: https://stackoverflow.com/a/18537419
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filepath := path.Join(pwd, templateFileName)

	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// encrypt data using aes
// https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
func ChiperEncrypt(stringToEncrypt string, keyString string) (string, error) {

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

// decrypt data using aes
// https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
func ChiperDecrypt(encryptedString string, keyString string) (string, error) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}

func TimestampToTime(s string) time.Time {
	i, _ := strconv.ParseInt(s, 10, 64)
	return time.Unix(i, 0)
}
