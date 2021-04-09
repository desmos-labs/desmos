package keeper_test

import "github.com/desmos-labs/desmos/x/links/types"

func (suite *KeeperTestSuite) Test_SetPort() {
	suite.Run("valid setup", func() {
		suite.SetupTest()
		suite.k.SetPort(suite.ctx, types.PortID)
		err := suite.k.BindPort(suite.ctx, types.PortID)
		if err != nil {
			panic(err)
		}
		port := suite.k.GetPort(suite.ctx)
		suite.Require().Equal(port, types.PortID)
	})
}
