package env

import (
	"os"
	"reflect"
	"testing"
)

func TestEnumerateEnvs(t *testing.T) {
	_ = os.Setenv("ENVX_HOME_DIR", "test_data")

	var expected = []string{
		"test_env_01",
	}

	result, err := EnumerateEnvs()
	if err != nil {
		t.Fatal()
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result (%s) does not match expected (%s)", result, expected)
	}

}

func TestReadEnv(t *testing.T) {
	_ = os.Setenv("ENVX_HOME_DIR", "test_data")
	var testEnv = "test_env_01"
	env, err := ReadEnv(testEnv)
	if err != nil {
		t.Fatal()
	}

	result := env.Vars

	var expected = map[string]string{
		"TESTVAR": "TESTVAL",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result (%s) does not match expected (%s)", result, expected)
	}

}
