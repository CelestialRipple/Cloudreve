package auth

import (
	"github.com/HFO4/cloudreve/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSignURI(t *testing.T) {
	asserts := assert.New(t)
	General = HMACAuth{SecretKey: []byte(util.RandStringRunes(256))}

	// 成功
	{
		sign, err := SignURI("/api/v3/something?id=1", 0)
		asserts.NoError(err)
		queries := sign.Query()
		asserts.Equal("1", queries.Get("id"))
		asserts.NotEmpty(queries.Get("sign"))
	}

	// URI解码失败
	{
		sign, err := SignURI("://dg.;'f]gh./'", 0)
		asserts.Error(err)
		asserts.Nil(sign)
	}
}

func TestCheckURI(t *testing.T) {
	asserts := assert.New(t)
	General = HMACAuth{SecretKey: []byte(util.RandStringRunes(256))}

	// 成功
	{
		sign, err := SignURI("/api/ok?if=sdf&fd=go", time.Now().Unix()+10)
		asserts.NoError(err)
		asserts.NoError(CheckURI(sign))
	}

	// 过期
	{
		sign, err := SignURI("/api/ok?if=sdf&fd=go", time.Now().Unix()-1)
		asserts.NoError(err)
		asserts.Error(CheckURI(sign))
	}
}
