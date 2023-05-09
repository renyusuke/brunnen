package varsx

import (
	"errors"
)

// 系统相关
var (
	ErrSysBadRequest   = errors.New("网络异常,请稍后重试")
	ErrSysTokenExpired = errors.New("token失效")
	ErrSysAuthFailed   = errors.New("权限校验失败")
	ErrRequestLimit    = errors.New("请求频率限制")
	ErrTimeout         = errors.New("请求超时")
	ErrIPLimit         = errors.New("IP限制")
	ErrMissAcctPwd     = errors.New("账号或密码错误") // 账号或密码错误
	ErrAcctDisabled    = errors.New("账号已停用")   // 账号已停用
	ErrImageSizeLimit  = errors.New("图片大小限制")
	ErrImageSuffix     = errors.New("图片后缀错误")
)

// 通讯相关
var (
	ErrNewsCallBinding             = errors.New("该通话已经绑定，请先解绑")
	ErrNewsNoSupportedCarrier      = errors.New("无支持的电话运营商")
	ErrNewsNoSupportedMsg          = errors.New("无支持的短信运营商")
	ErrNewsBoundCustomerService    = errors.New("该客服已经绑定，请先解绑")
	ErrNoMsgService                = errors.New("没有对应短信服务商")
	ErrTheBussAlreadyBound         = errors.New("该坐席已经绑定,请先解绑")        // 该坐席已经绑定,请先解绑
	ErrTheBussCannotBindOther      = errors.New("不能解绑他人的坐席")           // 不能解绑他人的坐席
	ErrTheServiceMustOffer         = errors.New("切换为备用线路的服务商必须为未开启状态") // 不能解绑他人的坐席
	ErrTheServiceStatusDifferent   = errors.New("数据库和三方服务状态不一致，请重试")   // 不能解绑他人的坐席
	ErrTheServiceRequestFailed     = errors.New("请求第三方接口失败,置忙置闲失败")    // 不能解绑他人的坐席
	ErrTheAuthFailed               = errors.New("校验失败")                // 不能解绑他人的坐席
	ErrIllegalAction               = errors.New("非法操作，请检查您的操作")        // action 非法，请检查您的action
	ErrMainServiceCanNotBeDisabled = errors.New("主线服务无法禁用")            // 主线服务无法禁用
	ErrNoPhoneNumber               = errors.New("没有电话号码")              // 没有电话号码
	ErrSendMessageFailed           = errors.New("发送消息失败")
	ErrCheckBalanceBusyBee         = errors.New("检测BusyBee失败") // 检测BusyBee失败
	ErrCheckBalanceTexCell         = errors.New("检测TexCell失败") // 检测TexCell失败
	ErrCheckBalanceJoey            = errors.New("检测Joey失败")    // 检测Joey失败
)

// 场面相关
var ()

// 户口相关
var (
	ErrMemberMistakeOldPwd = errors.New("旧密码错误")
	ErrMemberNotParent     = errors.New("非子户口的上级,无法调整关系")
	ErrMemberMistakeFlag   = errors.New("户口Flag错误")
)

// 银头相关
var (
	ErrLedgerChipsNameDuplicated   = errors.New("籌碼名稱重複")
	ErrLedgerJunketsChipsNotEnough = errors.New("泥碼不足")
	ErrLedgerRegularChipsNotEnough = errors.New("现金码不足")
	ErrLedgerChipsNotExist         = errors.New("籌碼不存在")
	ErrLedgerHallNotExist          = errors.New("場館不存在")
	ErrLedgerRoundStillNotClose    = errors.New("還未全部收工")
	ErrLedgerIdsNecessary          = errors.New("绑定交收方案必须至少传入一笔筹码id")
	ErrLedgerChipsLockFail         = errors.New("筹码加锁失败")
	ErrLedgerHallLockFail          = errors.New("场馆加锁失败")
	ErrLedgerDoesNotExist          = errors.New("该银头数据部存在")
)

// 礼物积分服务相关
var ()

// 报表相关
var ()

// 钱包或者Marker相关 8000-8999
var (
	ErrMarkerCreditSubHas        = errors.New("该子户口存在未归还的上级授信")
	ErrWalletQuery               = errors.New("用户钱包余额查询失败")
	ErrWalletDoesNotExist        = errors.New("该钱包不存在")
	ErrMarkerLimit               = errors.New("负额上限不足")
	ErrWalletInsufficientBalance = errors.New("钱包余额不足") // 钱包余额不足
)

// 权限校验 9000-9999
var (
	ErrAuthFail = errors.New("权限校验失败")
)
