package api

import "goweibo/core/context"

type testResp struct {
	Hello string `json:"hello"`
}

// Test api test
// @Summary api test
// @Tags Test
// @Accept json
// @Produce json
// @Param key query string false "test key"
// @Success 200 {object} context.CommonResponse{data=testResp}
// @Security ApiKeyAuth
// @Router /test [get]
func Test(c *context.AppContext) error {
	return c.AWSuccessJSON(testResp{
		Hello: c.QueryParam("key"),
	})
}
