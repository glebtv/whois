/*
 * Go module for domain whois
 * https://www.likexian.com/
 *
 * Copyright 2014-2018, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package whois

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestWhois(t *testing.T) {
	result := Whois("rsmon")
	assert.NotEqual(t, nil, result.Error)
	assert.Equal(t, "", result.Raw)
	assert.NotEqual(t, nil, result.Expires)

	result = Whois("rsmon.ru")
	assert.Equal(t, nil, result.Error)
	assert.NotEqual(t, "", result.Raw)
	assert.NotEqual(t, nil, result.Expires)

	//result, err = Whois("likexian.com", "127.0.0.1")
	//assert.Equal(t, nil, Resut.Error)
	//assert.Equal(t, "", result)

	//result, err := Query("rocketmon.com", "com.whois-servers.net")
	//assert.Equal(t, nil, err)
	//assert.NotEqual(t, "", result.Raw)
	//assert.NotEqual(t, nil, result.Expires)

	result = Whois("gleb.tv")
	assert.Equal(t, nil, result.Error)
	assert.NotEqual(t, "", result.Raw)
	assert.NotEqual(t, nil, result.Expires)
}
