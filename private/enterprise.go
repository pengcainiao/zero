package private

//
//import (
//	"context"
//	"encoding/base64"
//	"encoding/json"
//	"strconv"
//
//	"github.com/forgoer/openssl"
//	"github.com/pengcainiao/core/logx"
//)
//
//var (
//	// aesCryptKey AES加密key
//	aesCryptKey = "feixiang@#flyele"
//	// aesCryptIv AES加密iv
//	aesCryptIv = "flyele@#feixiang"
//)
//
//// Enterprise 定义管理员信息
//type Enterprise struct {
//	ID           string `json:"enterprise_id,omitempty" db:"id"`            // 企业id
//	Name         string `json:"name,omitempty" db:"name"`                   // 账号
//	SerialNumber string `json:"serial_number,omitempty" db:"serial_number"` // 主板序列号
//	SystemUUID   string `json:"system_uuid,omitempty" db:"system_uuid"`     // 系统uuid
//	ExpireDate   string `json:"expire_date,omitempty" db:"expire_date"`     // 有效日期
//	AuthUsers    int    `json:"auth_users,omitempty" db:"auth_users"`       // 授权用户
//	AuthCode     string `json:"auth_code,omitempty" db:"auth_code"`         // 授权码
//}
//
//// DecryptAuthCode 解密授权码
//func (e *Enterprise) DecryptAuthCode() error {
//	dst, err := base64.StdEncoding.DecodeString(e.AuthCode)
//	if err != nil {
//		logx.NewTraceLogger(context.Background()).Err(err).Msg("base64解码失败")
//		return err
//	}
//
//	authCode, err := openssl.AesCBCDecrypt(dst, []byte(aesCryptKey), []byte(aesCryptIv), openssl.PKCS7_PADDING)
//	if err != nil {
//		logx.NewTraceLogger(context.Background()).Err(err).Msg("授权码解密失败")
//		return err
//	}
//
//	var info []string
//	err = json.Unmarshal(authCode, &info)
//	if err != nil {
//		logx.NewTraceLogger(context.Background()).Err(err).Msg("json反序列化失败")
//		return err
//	}
//
//	e.ID = info[0]
//	e.Name = info[1]
//	e.SerialNumber = info[2]
//	e.SystemUUID = info[3]
//	e.ExpireDate = info[4]
//	e.AuthUsers, err = strconv.Atoi(info[5])
//	if err != nil {
//		logx.NewTraceLogger(context.Background()).Err(err).Msg("转换授权人数失败")
//		return err
//	}
//	return nil
//}
