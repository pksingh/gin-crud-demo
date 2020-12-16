package config

import (
	"testing"
	"time"
)

func TestGetPropsArr(t *testing.T) {
	t.Setenv(envKey, "test")
	p := GetProps("../resources")
	arrKey := "sample.prop3"
	expValue := "item1,item2, item3, item4 ,   item5"
	if v, ok := p.Get(arrKey); ok && v != expValue {
		t.Errorf("Reading prop on key %s resulted in empty value, we expected %s",
			arrKey, expValue)
	}
}

func TestGetPropsDur(t *testing.T) {
	t.Setenv(envKey, "test")
	p := GetProps("../resources")
	key := "http.timeout.sample"
	expValue := 15 * time.Second
	o := p.MustGetParsedDuration(key)
	if o != expValue {
		t.Errorf("Reading prop on key %s resulted in %d value, we expected %s",
			key, o, expValue)
	}
}

func TestGetPropsBool(t *testing.T) {
	t.Setenv(envKey, "test")
	p := GetProps("../resources")
	key := "i.like.go"
	o := p.MustGetBool(key)
	if !o {
		t.Errorf("Reading prop on key %s resulted in %t value, we expected %t",
			key, o, true)
	}
}

func TestGetSecretFromEnv(t *testing.T) {
	t.Setenv(envKey, "test")
	expected := "some-random-password"

	// In real world scenarios, Kube does this behind the scenes when we configure it in Rancher
	t.Setenv("MP-PAYMENTS-PAYIN-DB-PASSWORD", expected)

	p := GetProps("../resources")
	key := "postgres.database.password"
	result := p.MustGetString(key)
	if result != expected {
		t.Errorf("Reading prop on key %s resulted in %s value, we expected %s",
			key, result, expected)
	}
}
