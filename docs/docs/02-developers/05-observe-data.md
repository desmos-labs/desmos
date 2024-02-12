---
id: observe-data
title: Observing data
sidebar_label: Observing data
slug: observe-data
---

# Observing new data

## Introduction
Aside from querying data, you can also observe new data as its inserted inside the chain itself. In this way, you will be notified as soon as a transaction is properly executed without having to constantly polling the chain state by yourself. 

## Websocket  
All the live data observation is done though the usage of a [websocket](https://en.wikipedia.org/wiki/WebSocket). The endpoint of such websocket is the following: 

```
ws://lcd-endpoint/websocket

# Example
# ws://morpheus.desmos.network/websocket
```

### Events
In order to subscribe to specific events, you will need to send one or more messages to the websocket once you opened a connection to it. Such messages need to contain the following JSON object and must be string encoded: 

```json
{
  "jsonrpc": "2.0",
  "method": "subscribe",
  "id": "0",
  "params": {
    "query": "tm.event='eventCategory' AND eventType.eventAttribute='attributeValue'"
  }
}
``` 

The `query` field can have the following values: 

* `tm.event='NewBlock'` if you want to observe each new block that is created (even empty ones);
* `tm.event='Tx'` if you want to subscribe to all new transactions;
* `message.action='<action>'` if you want to subscribe to events emitted when a specific message is sent to the chain. 
  In this case, please refer to the `Message action` section on each transaction message 
  specification page to know what is the type associated to each message.
message: type MsgsTestSuite struct {
	suite.Suite
	contract      sdk.AccAddress
	deployer      sdk.AccAddress
	deployerStr   string
	withdrawerStr string
}

func TestMsgsTestSuite(t *testing.T) {
	suite.Run(t, new(MsgsTestSuite))
}

func (suite *MsgsTestSuite) SetupTest() {
	deployer := "cosmos1"
	withdraw := "cosmos2"
	suite.contract = sdk.AccAddress([]byte("juno15u3dt79t6sxxa3x3kpkhzsy56edaa5a66wvt3kxmukqjz2sx0hes5sn38g"))
	suite.deployer = sdk.AccAddress([]byte(deployer))
	suite.deployerStr = suite.deployer.String()
	suite.withdrawerStr = sdk.AccAddress([]byte(withdraw)).String()
}

func (suite *MsgsTestSuite) TestMsgRegisterFeeShareGetters() {
	msgInvalid := MsgRegisterFeeShare{}
	msg := NewMsgRegisterFeeShare(
		suite.contract,
		suite.deployer,
		suite.deployer,
	)
	suite.Require().Equal(RouterKey, msg.Route())
	suite.Require().Equal(TypeMsgRegisterFeeShare, msg.Type())
	suite.Require().NotNil(msgInvalid.GetSignBytes())
	suite.Require().NotNil(msg.GetSigners())
}

func (suite *MsgsTestSuite) TestMsgRegisterFeeShareNew() {
	testCases := []struct {
		msg        string
		contract   string
		deployer   string
		withdraw   string
		expectPass bool
	}{
		{
			"pass",
			suite.contract.String(),
			suite.deployerStr,
			suite.withdrawerStr,
			true,
		},
		{
			"pass - empty withdrawer address",
			suite.contract.String(),
			suite.deployerStr,
			"",
			true,
		},
		{
			"pass - same withdrawer and deployer address",
			suite.contract.String(),
			suite.deployerStr,
			suite.deployerStr,
			true,
		},
		{
			"invalid contract address",
			"",
			suite.deployerStr,
			suite.withdrawerStr,
			false,
		},
		{
			"invalid deployer address",
			suite.contract.String(),
			"",
			suite.withdrawerStr,
			false,
		},
		{
			"invalid withdraw address",
			suite.contract.String(),
			suite.deployerStr,
			"withdraw",
			false,
		},
	}

	for i, tc := range testCases {
		tx := MsgRegisterFeeShare{
			ContractAddress:   tc.contract,
			DeployerAddress:   tc.deployer,
			WithdrawerAddress: tc.withdraw,
		}
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s", i, tc.msg)
			suite.Require().Contains(err.Error(), tc.msg)
		}
	}
}

func (suite *MsgsTestSuite) TestMsgCancelFeeShareGetters() {
	msgInvalid := MsgCancelFeeShare{}
	msg := NewMsgCancelFeeShare(
		suite.contract,
		sdk.AccAddress(suite.deployer.Bytes()),
	)
	suite.Require().Equal(RouterKey, msg.Route())
	suite.Require().Equal(TypeMsgCancelFeeShare, msg.Type())
	suite.Require().NotNil(msgInvalid.GetSignBytes())
	suite.Require().NotNil(msg.GetSigners())
}

func (suite *MsgsTestSuite) TestMsgCancelFeeShareNew() {
	testCases := []struct {
		msg        string
		contract   string
		deployer   string
		expectPass bool
	}{
		{
			"msg cancel contract fee - pass",
			suite.contract.String(),
			suite.deployerStr,
			true,
		},
	}

	for i, tc := range testCases {
		tx := MsgCancelFeeShare{
			ContractAddress: tc.contract,
			DeployerAddress: tc.deployer,
		}
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
			suite.Require().Contains(err.Error(), tc.msg)
		}
	}
}

func (suite *MsgsTestSuite) TestMsgUpdateFeeShareGetters() {
	msgInvalid := MsgUpdateFeeShare{}
	msg := NewMsgUpdateFeeShare(
		suite.contract,
		sdk.AccAddress(suite.deployer.Bytes()),
		sdk.AccAddress(suite.deployer.Bytes()),
	)
	suite.Require().Equal(RouterKey, msg.Route())
	suite.Require().Equal(TypeMsgUpdateFeeShare, msg.Type())
	suite.Require().NotNil(msgInvalid.GetSignBytes())
	suite.Require().NotNil(msg.GetSigners())
}

func (suite *MsgsTestSuite) TestMsgUpdateFeeShareNew() {
	testCases := []struct {
		msg        string
		contract   string
		deployer   string
		withdraw   string
		expectPass bool
	}{
		{
			"msg update fee - pass",
			suite.contract.String(),
			suite.deployerStr,
			suite.withdrawerStr,
			true,
		},
		{
			"invalid contract address",
			"",
			suite.deployerStr,
			suite.withdrawerStr,
			false,
		},
		{
			"invalid withdraw address",
			suite.contract.String(),
			suite.deployerStr,
			"withdraw",
			false,
		},
		{
			"change fee withdrawer to deployer - pass",
			suite.contract.String(),
			suite.deployerStr,
			suite.deployerStr,
			true,
		},
	}

	for i, tc := range testCases {
		tx := MsgUpdateFeeShare{
			ContractAddress:   tc.contract,
			DeployerAddress:   tc.deployer,
			WithdrawerAddress: tc.withdraw,
		}
		err := tx.ValidateBasic()

		if tc.expectPass {
			suite.Require().NoError(err, "valid test %d failed: %s, %v", i, tc.msg)
		} else {
			suite.Require().Error(err, "invalid test %d passed: %s, %v", i, tc.msg)
			suite.Require().Contains(err.Error(), tc.msg)
		}
	}
}

Please note that if you want to subscribe to multiple events you will need to send multiple query messages upon connecting to the websocket. 

#### Example
```json
{
  "jsonrpc": "2.0",
  "method": "subscribe",
  "id": "0",
  "params": {
    "query": "message.action='save_profile'"
  }
}
```
