package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailValidate(t *testing.T) {
	assert.True(t, EmailValidate("google@gmail.com"))
	assert.False(t, EmailValidate("google@gmail@.com"))
	assert.False(t, EmailValidate("google@gmail_com"))
	assert.False(t, EmailValidate("ç$?§/az@gmail.comg"))
}

var phoneTests = []struct {
	phone    string
	expected bool
}{
	{"09211522539", true},
	{"092115225390", false},
	{"08211522539", false},
	{"19211522539", false},
	{"09811522539", true},
}

func TestPhoneValidation(t *testing.T) {
	for _, tt := range phoneTests {
		t.Run(tt.phone, func(t *testing.T) {

			if PhoneRegex(tt.phone) != tt.expected {
				t.Errorf("got %s, want ", tt.phone)
			}
		})
	}
}
