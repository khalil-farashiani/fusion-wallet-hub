package config

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/wallet/internal/repository/mysql"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

const testYAMLFileContent = `---
type: yml
http_server:
  port: 8088

mysql:
  port: 3308
  host: localhost
  db_name: wallet_db
  username: test
  password: toohardpasswordtofind
application:
  graceful_shutdown_timeout: 10000000000
`

type loadTest struct {
	name         string
	EnvVariables map[string]string
	fileName     string
	input        []byte
	expected     Config
}

func newLoadTestCases() []loadTest {
	return []loadTest{
		{
			name:         "test with yaml file without env vars",
			EnvVariables: map[string]string{},
			fileName:     "*.yaml",
			input:        []byte(testYAMLFileContent),
			expected: Config{
				Application: Application{
					GracefulShutdownTimeout: time.Second * 10,
				},
				HTTPServer: HTTPServer{
					Port: 8088,
				},
				Mysql: mysql.Config{
					Username: "test",
					Password: "toohardpasswordtofind",
					Port:     3308,
					Host:     "localhost",
					DBName:   "wallet_db",
				},
			},
		},
		{
			name: "test with env vars",
			EnvVariables: map[string]string{
				"WALLET_HTTP..SERVER_PORT":                        "8088",
				"WALLET_MYSQL_PASSWORD":                           "toohardpasswordtofind",
				"WALLET_MYSQL_PORT":                               "3308",
				"WALLET_MYSQL_HOST":                               "192.168.10.202",
				"WALLET_MYSQL_USERNAME":                           "test",
				"WALLET_MYSQL_DB..NAME":                           "wallet_db",
				"WALLETـAPPLICATION..GRACEFUL..SHUTDOWN..TIMEOUT": "10000000000",
			},
			fileName: "*.yaml",
			input:    nil,
			expected: Config{
				Application: Application{
					GracefulShutdownTimeout: time.Second * 10,
				},
				HTTPServer: HTTPServer{
					Port: 8088,
				},
				Mysql: mysql.Config{
					Username: "test",
					Password: "toohardpasswordtofind",
					Port:     3308,
					Host:     "192.168.10.202",
					DBName:   "wallet_db",
				},
			},
		},
	}
}

func createTempFile(testCase loadTest, t *testing.T) string {
	// Create a temporary file to act as a YAML configuration file.
	tempFile, err := os.CreateTemp("./", testCase.fileName)
	if err != nil {
		t.Fatalf("test case:%s\nFailed to create temp file:%s", testCase.name, err)
	}
	// Write some content to the temp file.
	if _, err := tempFile.Write(testCase.input); err != nil {
		t.Fatal("Failed to write to temp file:", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatal("Failed to close temp file:", err)
	}
	return tempFile.Name()
}

func setEnvVarTestLoad(envVars map[string]string) {
	for key, val := range envVars {
		os.Setenv(key, val)
	}
}

func unsetEnvVarTestLoad(envVars []map[string]string) {
	for _, varMap := range envVars {
		for key, _ := range varMap {
			os.Unsetenv(key)
		}
	}
}

func removeTempFile(fileNames []string) {
	for _, name := range fileNames {
		os.Remove(name)
	}
}

func TestLoad(t *testing.T) {
	//arrange
	var fileNames []string
	var envVars []map[string]string
	testCases := newLoadTestCases()
	for _, testCase := range testCases {
		fileName := createTempFile(testCase, t)
		fileNames = append(fileNames, fileName)
		setEnvVarTestLoad(testCase.EnvVariables)
		envVars = append(envVars, testCase.EnvVariables)

		//act
		config := Load(fileName)

		//assert
		//because of ensure that all struct equal we don't compare struct directly or with deepEqual
		assert.Equal(t, testCase.expected.Mysql.Port, config.Mysql.Port)
		assert.Equal(t, testCase.expected.Mysql.Host, config.Mysql.Host)
		assert.Equal(t, testCase.expected.Mysql.DBName, config.Mysql.DBName)
		assert.Equal(t, testCase.expected.Mysql.Password, config.Mysql.Password)
		assert.Equal(t, testCase.expected.Mysql.Username, config.Mysql.Username)
		assert.Equal(t, testCase.expected.Application.GracefulShutdownTimeout, config.Application.GracefulShutdownTimeout)
		assert.Equal(t, testCase.expected.HTTPServer.Port, config.HTTPServer.Port)
	}
	defer removeTempFile(fileNames)
	defer unsetEnvVarTestLoad(envVars)
}
