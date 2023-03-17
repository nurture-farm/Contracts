#!/bin/sh
set -x -e -u
mockgen -source=./LoyaltyManagementSystem/LoyaltyManagementService/Gen/GoLoyaltyManagementService/loyalty_management_service.pb.go -destination=./LoyaltyManagementSystem/LoyaltyManagementService/Gen/GoLoyaltyManagementService/mock.go -package=LoyaltyManagementService
mockgen -source=./BookingServices/BookingConfigService/Gen/GoBookingConfigService/booking_config_service.pb.go -destination=./BookingServices/BookingConfigService/Gen/GoBookingConfigService/mock.go -package=BookingConfigService