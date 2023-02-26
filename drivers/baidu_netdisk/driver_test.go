package baidu_netdisk

import (
	"testing"
)

func TestGetQuota(t *testing.T) {
	// 创建一个 BaiduNetdisk 实例，假设已经正确配置了账号信息
	bn := NewBaiduNetdisk("access_token", "refresh_token", "client_id", "client_secret")

	// 调用 GetQuota() 方法获取网盘配额信息
	resp, err := bn.GetQuota()
	if err != nil {
		t.Errorf("GetQuota() error: %v", err)
	}

	// 判断获取到的配额信息是否正确
	if resp.Used < 0 || resp.Total < 0 || resp.Used > resp.Total {
		t.Errorf("GetQuota() returned invalid response: %+v", resp)
	}
}
