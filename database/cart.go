package database

import "errors"

var (
	ErrCantFindProduct = errors.New("product dne")
	ErrCantDecodeProduct = errors.New("cant decode product")
	ErrUserIdIsNotValid = errors.New("not a valid user")
	ErrCantUpdateUser = errors.New("cant update user")
	ErrCantRemoveItemFromCart = errors.New("error removing item from cart")
	ErrCantGetItem = errors.New("err getting item")
	ErrCantBuyCartItem = errors.New("err buying the cart item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}