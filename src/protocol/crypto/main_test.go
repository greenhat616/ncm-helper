package crypto

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWEAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockIUtil(ctrl)
	u := &Util{}
	// Reflect the original func
	m.EXPECT().PKCS7Padding(gomock.Any(), gomock.Any()).DoAndReturn(u.PKCS7Padding).AnyTimes()
	m.EXPECT().base62Encode(gomock.Any()).DoAndReturn(u.base62Encode).AnyTimes()
	m.EXPECT().reverse(gomock.Any()).DoAndReturn(u.reverse).AnyTimes()
	m.EXPECT().charCodeAt(gomock.Any(), gomock.Any()).DoAndReturn(u.charCodeAt).AnyTimes()
	// Mock generate Random Bytes
	m.EXPECT().GenRandomBytes(gomock.Eq(16)).Return([]byte{
		100,
		153,
		158,
		112,
		8,
		249,
		210,
		196,
		9,
		92,
		13,
		101,
		30,
		225,
		193,
		202,
	}, nil).AnyTimes()
	util = m
	assert := assert.New(t)
	// have a try
	if b, _ := util.GenRandomBytes(16); len(b) != 16 {
		t.Error("can't generate specific length of bytes.")
		t.Fail()
		return
	}
	// TestWEAPI
	params, enSecKey, err := WEAPI("{ nickname: 'test' }")
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Logf("params: %s", params)
		t.Logf("enSecKey: %s", enSecKey)
		assert.Equal("0PzH++wpg/l5lIb2L9gmeQ7QiQTfrjCSorIF/LCeY0ZIeAGjRQPF8eTV53gxTM6e", string(params), "params should be the same")
		assert.Equal("0cc0bf7500f05eb4c5a5979ce2a82e722c786c04945a692977cfc1c9c8e0297836ba06fa9acbc357d82b01049e5f683a718037728331d5de4d1b85c258619efb9dfbeaa4ec30ee1f6ce0eed8e5933286120ed25b08da93d5f651512ad45f57d115f6d024137a7bc52ddaa0f347f51c7d6a19a23ab574037443c9829aee32a434", string(enSecKey), "enSecKey should be the same")
	}
}

func TestLinuxAPI(t *testing.T) {
	eParams, err := LinuxAPI("{\"a\":1}")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	t.Logf("eparams: %s", eParams)
	assert := assert.New(t)
	assert.Equal("6270E70FBBC25054048A37935D9EAAC6", string(eParams), "eParams must be the same")
}

// TODO: Test Decrypt func, not found a valid test data source.

func TestEAPI(t *testing.T) {
	params, err := EAPI("/api/cellphone/existence/check", "{\"a\":1}")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	t.Logf("params: %s", params)
	assert := assert.New(t)
	assert.Equal("AF047AB9ACC436C08101E8542E2D2378AB110DB4C7B0F4BC32B33E80214D1C932042444B7488CCCD560ED501F1C1E0D45431D9540836813A074FCC4F407C8541B072375589492FA0A18D43B82350430D211D5D5093DA6E1E5528ABADF7FBB357", string(params), "params must be the same")
}
