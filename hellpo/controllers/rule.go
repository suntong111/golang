package controllers

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"time"
)

type RuleController struct {
	BaseController
}

func (c *RuleController) Get() {

	t := time.Now().Unix()
	// 生成一个MD5的哈希
	h := md5.New()
	// 将时间戳转换为byte，并写入哈希
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(t))
	h.Write([]byte(b))
	hex.EncodeToString(h.Sum(nil))




	c.ApiJson("查到信息",200, hex.EncodeToString(h.Sum(nil)))




}