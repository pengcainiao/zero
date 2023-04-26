package private

import (
	"context"
	"errors"
	"strings"

	"github.com/963204765/httpclient/httplib"
	"gitlab.flyele.vip/server-side/go-zero/v2/core/env"
	"gitlab.flyele.vip/server-side/go-zero/v2/core/logx"
	"gitlab.flyele.vip/server-side/go-zero/v2/private/ssh"
)

// PrivateDeploy 私有化部署
type PrivateDeploy struct {
	AuthCode string
}

// getMasterSerialNumberAndUUID 获取主板序列号和uuid
func getMasterSerialNumberAndUUID() (string, string) {
	client := ssh.SSHClient{Addr: "120.77.248.81:22", User: "root", Keyfile: "/home/chenbin/.ssh/id_rsa"}
	serialNumber, err := client.Command("dmidecode -s system-serial-number")
	if err != nil {
		logx.NewTraceLogger(context.Background()).Err(err).Msg("获取主板序列号失败")
	}

	systemUUID, err := client.Command("dmidecode -s system-uuid")
	if err != nil {
		logx.NewTraceLogger(context.Background()).Err(err).Msg("获取系统uuid失败")
	}

	serialNumber = strings.Trim(serialNumber, "\n")
	systemUUID = strings.Trim(systemUUID, "\n")

	return serialNumber, systemUUID
}

// getMasterSerialNumberAndUUIDByHTTP 获取主板序列号和uuid http获取
func getMasterSerialNumberAndUUIDByHTTP() (string, string) {
	req := httplib.NewBeegoRequest(env.PrivateAuthHTTPHOST+"/system/info", "GET")
	var res struct {
		Code int `json:"code"`
		Data struct {
			SerialNumber string `json:"serial_number"`
			SystemUUID   string `json:"system_uuid"`
		} `json:"data"`
	}
	_ = req.ToJSON(&res)
	return res.Data.SerialNumber, res.Data.SystemUUID
}

// getEnterpriseInfo 获取企业信息
func (p PrivateDeploy) getEnterpriseInfo() (Enterprise, error) {
	enterprise := Enterprise{AuthCode: p.AuthCode}
	if err := enterprise.DecryptAuthCode(); err != nil {
		return enterprise, errors.New("解密授权码失败")
	}
	return enterprise, nil
}

// CheckSerialNumber 验证主板序列号
func (p PrivateDeploy) CheckSerialNumber() bool {
	p.getEnterpriseInfo()
	enterprise, err := p.getEnterpriseInfo()
	if err != nil {
		return false
	}

	var (
		serialNumber string
		systemUUID   string
	)
	if env.PrivateAuthMode == "ssh" {
		serialNumber, systemUUID = getMasterSerialNumberAndUUID()
	} else {
		serialNumber, systemUUID = getMasterSerialNumberAndUUIDByHTTP()
	}

	return enterprise.SerialNumber == serialNumber && enterprise.SystemUUID == systemUUID
}

// GetExpireDate 获取有效日期
func (p PrivateDeploy) GetExpireDate() (string, error) {
	p.getEnterpriseInfo()
	enterprise, err := p.getEnterpriseInfo()
	if err != nil {
		return "", err
	}
	return enterprise.ExpireDate, nil
}

// GetAuthUsers 获取授权用户数
func (p PrivateDeploy) GetAuthUsers() (int, error) {
	p.getEnterpriseInfo()
	enterprise, err := p.getEnterpriseInfo()
	if err != nil {
		return 0, err
	}
	return enterprise.AuthUsers, nil
}
