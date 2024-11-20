package utilities

var NoUserError string = "this user does not exist"
var UserDeletedSuccess string = "account successfully deleted"
var UserUpdateError string = "could not update the user data"
var SocialsLoginRequired string = "please login using socials since you signed up with it"
var WrongCredentials string = "wrong password or email"
var AccessTokenUpdateError string = "failed to update access token"
var AccessTokenCreationError string = "failed to create access token"

var ManagerAccountDeleteSuccess string = "your property management account was successfully deleted"
var ManagerAccountUpdateError string = "this property management account failed to update"
var NoManagerAccountError string = "this property management account does not exist"
var ManagerAccountCreatedButNoDataRetrievedError string = "manager was created but failed to get the data back, please close the app and remove it from the background then come back again to see your manager account"

var StandPropertyType string = "stand"
var LandPropertyType string = "land"
var CommercialRentalPropertyType string = "commercialRentalProperty"
var CommercialPropertyForSaleType string = "commercialPropertyForSale"
var ResidentialPropertyForSaleType string = "residentialPropertyForSale"
var ResidentialRentalPropertyType = "residentialRentalProperty"
var PageLimit int = 2
var ImagesOrVideosLimitPerProperty int = 20
var IsPaywallActive bool = false
