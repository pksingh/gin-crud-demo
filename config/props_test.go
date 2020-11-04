package config

import (
	"testing"
	"time"
)

func TestGetProps(t *testing.T) {
	type args struct {
		resourceDir  string
		shouldSetEnv bool
		setEnvTo     string
		searchKey    string
		defPropVal   string
	}

	tests := []struct {
		name     string
		args     args
		expValue string
	}{
		{
			name: "load props from app.properties file and make sure it works",
			args: args{
				resourceDir:  "./testData",
				shouldSetEnv: false,
				setEnvTo:     "",
				defPropVal:   "",
				searchKey:    "sample.prop1",
			},
			expValue: "test",
		},

		{
			name: "load props from app.properties file and make sure app-dev overriding works",
			args: args{
				resourceDir:  "./testData",
				shouldSetEnv: true,
				setEnvTo:     "dev",
				defPropVal:   "",
				searchKey:    "sample.prop1",
			},
			expValue: "value1",
		},

		{
			name: "load props from app.properties file and make sure app-preprod overriding works",
			args: args{
				resourceDir:  "./testData",
				shouldSetEnv: true,
				setEnvTo:     "preprod",
				defPropVal:   "",
				searchKey:    "sample.prop2",
			},
			expValue: "value2",
		},

		{
			name: "load props from app.properties file and make sure app-prod overriding works",
			args: args{
				resourceDir:  "./testData",
				shouldSetEnv: true,
				setEnvTo:     "prod",
				defPropVal:   "",
				searchKey:    "sample.prop2",
			},
			expValue: "im damn serious",
		},

		{
			name: "inject incorrect folder path of properties file and make sure that default values are used",
			args: args{
				resourceDir:  "./unknown-dir", // injecting wrong properties' folder as input variable
				shouldSetEnv: true,
				setEnvTo:     "dev",
				defPropVal:   "imDefault",
				searchKey:    "sample.prop1",
			},
			expValue: "imDefault",
		},

		{
			name: "inject incorrect env variable and make sure that default values are used",
			args: args{
				resourceDir:  "./testData",
				shouldSetEnv: true,
				setEnvTo:     "some-fake-env", // injecting wrong environment variable
				defPropVal:   "imOtherDefault",
				searchKey:    "sample.prop1",
			},
			expValue: "imOtherDefault",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.shouldSetEnv {
				t.Setenv(envDetectorKey, tt.args.setEnvTo)
			}
			p := GetProps(tt.args.resourceDir)
			got := p.GetString(tt.args.searchKey, tt.args.defPropVal)
			if got != tt.expValue {
				t.Errorf("Reading prop on key %s resulted in %s, we expected %s",
					tt.args.searchKey, got, tt.expValue)
			}
		})
	}
}

func TestGetPropsArr(t *testing.T) {
	t.Setenv(envDetectorKey, "dev")
	p := GetProps("./testData")
	arrKey := "sample.prop3"
	expValue := "item1,item2, item3, item4 ,   item5"
	if v, ok := p.Get(arrKey); ok && v != expValue {
		t.Errorf("Reading prop on key %s resulted in empty value, we expected %s",
			arrKey, expValue)
	}
}

func TestGetPropsDur(t *testing.T) {
	t.Setenv(envDetectorKey, "preprod")
	p := GetProps("./testData")
	key := "http.timeout.sample"
	expValue := 15 * time.Second
	o := p.MustGetParsedDuration(key)
	if o != expValue {
		t.Errorf("Reading prop on key %s resulted in %d value, we expected %s",
			key, o, expValue)
	}
}

func TestGetPropsBool(t *testing.T) {
	p := GetProps("./testData")
	key := "i.like.go"
	o := p.MustGetBool(key)
	if !o {
		t.Errorf("Reading prop on key %s resulted in %t value, we expected %t",
			key, o, true)
	}
}

func TestGetSecretFromEnv(t *testing.T) {
	t.Setenv(envDetectorKey, "preprod")
	expected := "some-random-password"

	// In real world scenarios, Kube does this behind the scenes when we configure it in Rancher
	t.Setenv("MP-PAYMENTS-PAYIN-DB-PASSWORD", expected)

	p := GetProps("./testData")
	key := "postgres.database.password"
	result := p.MustGetString(key)
	if result != expected {
		t.Errorf("Reading prop on key %s resulted in %s value, we expected %s",
			key, result, expected)
	}
}
