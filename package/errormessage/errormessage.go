package errormessage

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("not found")

//ErrBadRequestPromoType bad request
var ErrBadRequestPromoType = errors.New("bad request, promo type must be BUY1GET1, BUYXPAYY, or DISCOUNT")

//ErrBadRequest bad request
var ErrBadRequest = errors.New("bad request")

//ErrNotAuthorized bad request
var ErrNotAuthorized = errors.New("user not authorized")

//ErrNotEnoughStock cannot add
var ErrNotEnoughStock = errors.New("not enough stock")

//ErrUserExist user exist
var ErrUserExist = errors.New("user already exists")

//ErrProductExist product exist
var ErrProductExist = errors.New("product already exists")
var ErrProductNotExist = errors.New("product not exists")

//ErrSaveDb failed save to db
var ErrSaveDb = errors.New("failed saved to db")

//ErrAuthentication error authentication invalid username or password
var ErrAuthentication = errors.New("invalid username or password")

//ErrGetToken failed generate token
var ErrGetToken = errors.New("failed generate token")
