// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v2

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 用户信息未找到
func IsProfileNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_PROFILE_NOT_FOUND.String() && e.Code == 404
}

// 用户信息未找到
func ErrorProfileNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorProfileReason_PROFILE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 获取数据失败
func IsGetDate(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_GET_DATE.String() && e.Code == 500
}

// 获取数据失败
func ErrorGetDate(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorProfileReason_GET_DATE.String(), fmt.Sprintf(format, args...))
}

// 编辑用户异常
func IsFailedToEditProfile(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_FAILED_TO_EDIT_PROFILE.String() && e.Code == 400
}

// 编辑用户异常
func ErrorFailedToEditProfile(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_FAILED_TO_EDIT_PROFILE.String(), fmt.Sprintf(format, args...))
}

// 手机号不能为空
func IsPhoneIsNotEmpty(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_PHONE_IS_NOT_EMPTY.String() && e.Code == 400
}

// 手机号不能为空
func ErrorPhoneIsNotEmpty(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_PHONE_IS_NOT_EMPTY.String(), fmt.Sprintf(format, args...))
}

// 删除用户信息失败
func IsFailedDeleteProfile(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_FAILED_DELETE_PROFILE.String() && e.Code == 500
}

// 删除用户信息失败
func ErrorFailedDeleteProfile(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorProfileReason_FAILED_DELETE_PROFILE.String(), fmt.Sprintf(format, args...))
}

// 历史密码验证失败
func IsHistoryPrivacyPasswordVerificationFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_HISTORY_PRIVACY_PASSWORD_VERIFICATION_FAIL.String() && e.Code == 500
}

// 历史密码验证失败
func ErrorHistoryPrivacyPasswordVerificationFail(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorProfileReason_HISTORY_PRIVACY_PASSWORD_VERIFICATION_FAIL.String(), fmt.Sprintf(format, args...))
}

// 输入邮箱与用户绑定邮箱不一致
func IsParamsEmailNotIsUserBind(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_PARAMS_EMAIL_NOT_IS_USER_BIND.String() && e.Code == 400
}

// 输入邮箱与用户绑定邮箱不一致
func ErrorParamsEmailNotIsUserBind(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_PARAMS_EMAIL_NOT_IS_USER_BIND.String(), fmt.Sprintf(format, args...))
}

// 输入手机与用户绑定手机不一致
func IsParamsPhoneNotIsUserBind(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_PARAMS_PHONE_NOT_IS_USER_BIND.String() && e.Code == 400
}

// 输入手机与用户绑定手机不一致
func ErrorParamsPhoneNotIsUserBind(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_PARAMS_PHONE_NOT_IS_USER_BIND.String(), fmt.Sprintf(format, args...))
}

// 超过批处理的最大数量
func IsExceededTheMaximumBatchProcessingQuantity(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_EXCEEDED_THE_MAXIMUM_BATCH_PROCESSING_QUANTITY.String() && e.Code == 400
}

// 超过批处理的最大数量
func ErrorExceededTheMaximumBatchProcessingQuantity(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_EXCEEDED_THE_MAXIMUM_BATCH_PROCESSING_QUANTITY.String(), fmt.Sprintf(format, args...))
}

// 批量导入失败
func IsBatchInsertError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_BATCH_INSERT_ERROR.String() && e.Code == 400
}

// 批量导入失败
func ErrorBatchInsertError(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_BATCH_INSERT_ERROR.String(), fmt.Sprintf(format, args...))
}

// 数据不能为空
func IsDataIsNotEmpty(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_DATA_IS_NOT_EMPTY.String() && e.Code == 400
}

// 数据不能为空
func ErrorDataIsNotEmpty(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_DATA_IS_NOT_EMPTY.String(), fmt.Sprintf(format, args...))
}

// 用户已经实名
func IsUserHasRealName(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_USER_HAS_REAL_NAME.String() && e.Code == 400
}

// 用户已经实名
func ErrorUserHasRealName(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_USER_HAS_REAL_NAME.String(), fmt.Sprintf(format, args...))
}

// 身份证好已被他人绑定
func IsIdNumberHasBeenUsed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_ID_NUMBER_HAS_BEEN_USED.String() && e.Code == 400
}

// 身份证好已被他人绑定
func ErrorIdNumberHasBeenUsed(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_ID_NUMBER_HAS_BEEN_USED.String(), fmt.Sprintf(format, args...))
}

// 实名认证失败
func IsUserRealNameErr(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorProfileReason_USER_REAL_NAME_ERR.String() && e.Code == 400
}

// 实名认证失败
func ErrorUserRealNameErr(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorProfileReason_USER_REAL_NAME_ERR.String(), fmt.Sprintf(format, args...))
}
