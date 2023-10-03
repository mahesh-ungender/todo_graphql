package utils

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
)

var nameGenerator namegenerator.Generator

func init() {
	nameGenerator = namegenerator.NewNameGenerator(time.Now().UTC().UnixNano())
}

// MaxInt64 returns the max of two int64s
// Can't believe I have to define this.
func MaxInt64(a int64, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

// CheckNullAndSet b if b is not null, else it returns a
func CheckNullAndSet(a *string, b *string) *string {
	if b != nil {
		return b
	}
	return a
}

// Check str for nil and set it to sql.Nill
func CheckStrNilSetToSQLNull(str string) sql.NullString {
	if len(str) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}

// CheckEmptyExperience str not null or empty, else it returns 0 experience
func CheckEmptyExp(str *string, defaultValue string) *string {
	if str != nil && *str != "" {
		return str
	}
	emptyExp := defaultValue
	return &emptyExp
}

// GetRandomUUID returns a random uuid
func GetRandomUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func RandomName() string {
	return nameGenerator.Generate()
}

// Remove removes the string at index i in a slice.
// Note that it doesn't preserve the order. So don't use when order is important
func Remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

// SerializeStringSliceToString serializes a slice of string to comma separate string
func SerializeStringSliceToString(list []string) string {
	res := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(list)), ","), "[]")
	return res
}

// StringToStringSlice returns a slice of string from a single string of comma separated strings
func StringToStringSlice(s *string) []string {
	var res []string
	if s != nil {
		res = strings.Split(*s, ",")
	}
	return res
}

// Contains given string is present or not
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
