package config

import "testing"

func TestConfig_StoreConfigFile(t *testing.T) {
	c := &Config{
		Enviroment: make([]map[string]string, 0),
		Command:    make(map[string]string),
	}

	c.Enviroment = append(c.Enviroment, map[string]string{"test_enviroment": "test_enviroment"})
	c.Command = map[string]string{"test_command": "test_command"}

	err := c.StoreConfigFile("xwc.yml")
	if err != nil {
		t.Fatal(err)
	}
}

func TestConfig_LoadConfigFile(t *testing.T) {
	c := &Config{}
	err := c.LoadConfigFile("xwc.yml")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(c.Enviroment, c.Command)
}

func TestConfig_IsExists(t *testing.T) {
	c := &Config{}
	exists := c.IsExists("xwc.yml")
	if !exists {
		exists = c.IsExists("xwc.yaml")
	}

	if !exists {
		t.Fatal("file isn't exists")
	}
}
