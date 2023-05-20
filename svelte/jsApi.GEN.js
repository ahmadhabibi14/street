
const axios = require("axios");


// rearrange response to be data first instead of axios error
function wrapErr( cb ) {
  return function( err ) {
    let data = ((err.response || {}).data || {})
    if( !data.error ) data.error = err.code
    data._axios = err
    cb( data )
  }
}

// rearrange response to be data first instead of axios error
function wrapOk( cb ) {
  return function( resp ) {
    let data = resp.data || {}
    data._axios = resp
    cb( data )
  }
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
/**
 * @typedef {Object} GuestDebugIn
 */
const GuestDebugIn = {
}
/**
 * @typedef {Object} GuestDebugOut
 * @property {Object} request
 */
const GuestDebugOut = {
  request: { // RequestCommon
  }, // RequestCommon
}
/**
 * @callback GuestDebugCallback
 * @param {GuestDebugOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestDebugIn} i
 * @param {GuestDebugCallback} cb
 * @returns {Promise}
 */
exports.GuestDebug = async function GuestDebug( i, cb ) {
  return await axios.post( '/guest/debug', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestForgotPasswordIn
 * @property {String} email
 */
const GuestForgotPasswordIn = {
  email: '', // string
}
/**
 * @typedef {Object} GuestForgotPasswordOut
 * @property {Object} ok
 * @property {String} resetPassUrl
 */
const GuestForgotPasswordOut = {
  ok: false, // bool
  resetPassUrl: '', // string
}
/**
 * @callback GuestForgotPasswordCallback
 * @param {GuestForgotPasswordOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestForgotPasswordIn} i
 * @param {GuestForgotPasswordCallback} cb
 * @returns {Promise}
 */
exports.GuestForgotPassword = async function GuestForgotPassword( i, cb ) {
  return await axios.post( '/guest/forgotPassword', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestLoginIn
 * @property {String} email
 * @property {String} password
 */
const GuestLoginIn = {
  email: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestLoginOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verificationSentAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 */
const GuestLoginOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
  }, // rqAuth.Users
}
/**
 * @callback GuestLoginCallback
 * @param {GuestLoginOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestLoginIn} i
 * @param {GuestLoginCallback} cb
 * @returns {Promise}
 */
exports.GuestLogin = async function GuestLogin( i, cb ) {
  return await axios.post( '/guest/login', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestRegisterIn
 * @property {String} email
 * @property {String} password
 */
const GuestRegisterIn = {
  email: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestRegisterOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verificationSentAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 * @property {String} verifyEmailUrl
 */
const GuestRegisterOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
  }, // rqAuth.Users
  verifyEmailUrl: '', // string
}
/**
 * @callback GuestRegisterCallback
 * @param {GuestRegisterOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestRegisterIn} i
 * @param {GuestRegisterCallback} cb
 * @returns {Promise}
 */
exports.GuestRegister = async function GuestRegister( i, cb ) {
  return await axios.post( '/guest/register', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestResendVerificationEmailIn
 * @property {String} email
 */
const GuestResendVerificationEmailIn = {
  email: '', // string
}
/**
 * @typedef {Object} GuestResendVerificationEmailOut
 * @property {Object} ok
 * @property {String} verifyEmailUrl
 */
const GuestResendVerificationEmailOut = {
  ok: false, // bool
  verifyEmailUrl: '', // string
}
/**
 * @callback GuestResendVerificationEmailCallback
 * @param {GuestResendVerificationEmailOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestResendVerificationEmailIn} i
 * @param {GuestResendVerificationEmailCallback} cb
 * @returns {Promise}
 */
exports.GuestResendVerificationEmail = async function GuestResendVerificationEmail( i, cb ) {
  return await axios.post( '/guest/resendVerificationEmail', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestResetPasswordIn
 * @property {String} secretCode
 * @property {String} hash
 * @property {String} password
 */
const GuestResetPasswordIn = {
  secretCode: '', // string
  hash: '', // string
  password: '', // string
}
/**
 * @typedef {Object} GuestResetPasswordOut
 * @property {Object} ok
 */
const GuestResetPasswordOut = {
  ok: false, // bool
}
/**
 * @callback GuestResetPasswordCallback
 * @param {GuestResetPasswordOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestResetPasswordIn} i
 * @param {GuestResetPasswordCallback} cb
 * @returns {Promise}
 */
exports.GuestResetPassword = async function GuestResetPassword( i, cb ) {
  return await axios.post( '/guest/resetPassword', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} GuestVerifyEmailIn
 * @property {String} secretCode
 * @property {String} hash
 */
const GuestVerifyEmailIn = {
  secretCode: '', // string
  hash: '', // string
}
/**
 * @typedef {Object} GuestVerifyEmailOut
 * @property {Object} ok
 * @property {String} email
 */
const GuestVerifyEmailOut = {
  ok: false, // bool
  email: '', // string
}
/**
 * @callback GuestVerifyEmailCallback
 * @param {GuestVerifyEmailOut} o
 * @returns {Promise}
 */
/**
 * @param  {GuestVerifyEmailIn} i
 * @param {GuestVerifyEmailCallback} cb
 * @returns {Promise}
 */
exports.GuestVerifyEmail = async function GuestVerifyEmail( i, cb ) {
  return await axios.post( '/guest/verifyEmail', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserLogoutIn
 */
const UserLogoutIn = {
}
/**
 * @typedef {Object} UserLogoutOut
 * @property {number} logoutAt
 */
const UserLogoutOut = {
  logoutAt: 0, // int64
}
/**
 * @callback UserLogoutCallback
 * @param {UserLogoutOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserLogoutIn} i
 * @param {UserLogoutCallback} cb
 * @returns {Promise}
 */
exports.UserLogout = async function UserLogout( i, cb ) {
  return await axios.post( '/user/logout', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}

/**
 * @typedef {Object} UserProfileIn
 */
const UserProfileIn = {
}
/**
 * @typedef {Object} UserProfileOut
 * @property {number} user.id
 * @property {String} user.email
 * @property {String} user.password
 * @property {number} user.createdAt
 * @property {number} user.createdBy
 * @property {number} user.updatedAt
 * @property {number} user.updatedBy
 * @property {number} user.deletedAt
 * @property {number} user.passwordSetAt
 * @property {String} user.secretCode
 * @property {number} user.secretCodeAt
 * @property {number} user.verificationSentAt
 * @property {number} user.verifiedAt
 * @property {number} user.lastLoginAt
 */
const UserProfileOut = {
  user: { // rqAuth.Users
    id: 0, // uint64
    email: '', // string
    password: '', // string
    createdAt: 0, // int64
    createdBy: 0, // uint64
    updatedAt: 0, // int64
    updatedBy: 0, // uint64
    deletedAt: 0, // int64
    passwordSetAt: 0, // int64
    secretCode: '', // string
    secretCodeAt: 0, // int64
    verificationSentAt: 0, // int64
    verifiedAt: 0, // int64
    lastLoginAt: 0, // int64
  }, // rqAuth.Users
}
/**
 * @callback UserProfileCallback
 * @param {UserProfileOut} o
 * @returns {Promise}
 */
/**
 * @param  {UserProfileIn} i
 * @param {UserProfileCallback} cb
 * @returns {Promise}
 */
exports.UserProfile = async function UserProfile( i, cb ) {
  return await axios.post( '/user/profile', i ).
    then( wrapOk( cb ) ).
    catch( wrapErr( cb ) )
}


// Code generated by 1_codegen_test.go DO NOT EDIT.
