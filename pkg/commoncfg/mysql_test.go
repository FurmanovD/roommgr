package commoncfg

import (
	"strconv"
	"testing"
	"testing/quick"

	"github.com/FurmanovD/go-kit/randomstring"
)

func TestMySqlConnectionString(t *testing.T) {

	testCases := map[string]interface{}{
		"DefaultHostOk": func(c SqlDBConfig) bool {
			// check default host
			c.Host = ""

			// make sure Port is > 0:
			if c.Port <= 0 {
				c.Port *= -1
			} else if c.Port == 0 {
				c.Port = 123
			}

			testStr := c.MySqlConnectionString()
			if testStr != c.User+":"+c.Password+"@tcp(127.0.0.1:"+strconv.Itoa(c.Port)+")/"+c.Database {
				t.Error("Invalid default host connection string built", testStr)
				return false
			}
			return true
		},
		"DefaultPortOk": func(c SqlDBConfig) bool {
			// check default port
			if c.Port > 0 {
				c.Port *= -1
			}

			// make sure Host is not empty
			if c.Host == "" {
				c.Host = randomstring.NonEmptyUTF8Printable(50, nil)
			}

			testStr := c.MySqlConnectionString()
			if testStr != c.User+":"+c.Password+"@tcp("+c.Host+":"+strconv.Itoa(DefaultMySQLPort)+")/"+c.Database {
				t.Error("Invalid default port connection string built", testStr)
				return false
			}

			return true
		},
		"AllIsFilled": func(c SqlDBConfig) bool {

			defaultMustBeSet := c.Host == "" || c.Port <= 0

			testStr := c.MySqlConnectionString()

			if defaultMustBeSet {
				if testStr == c.User+":"+c.Password+"@tcp("+c.Host+":"+strconv.Itoa(c.Port)+")/"+c.Database {
					t.Error("Invalid connection string built(default host or port expected)", testStr)
					return false
				}
			} else {
				if testStr != c.User+":"+c.Password+"@tcp("+c.Host+":"+strconv.Itoa(c.Port)+")/"+c.Database {
					t.Error("Invalid connection string built", testStr)
					return false
				}
			}

			return true
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if err := quick.Check(tc, nil); err != nil {
				t.Errorf("%v case failed with an error: %+v", name, err)
			}
		})
	}
}
