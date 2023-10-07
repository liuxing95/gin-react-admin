package response

import "github.com/liuxing95/gin-react-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
