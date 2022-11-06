Feature: Booking Cinema

    Books Cinema

HAPPY PATH:
~~~~~~~~~~~~~~~~~~~~
Booking-Service -> 
[BOOKING_CREATED] ->
CINEMA-SERVICE -> [
[BLOCKED_CINEMA_SEATS] -> 
PAYMENT-SERVICE -> 
[PAYMENT_CREDITED] -> 
BOOKING-SERVICE -> 
[BOOKING-CONFIRMED]
~~~~~~~~~~~~~~~~~~~~

FAIL PATH:
~~~~~~~~~~~~~~~~~~~~
[BOOKING-FAILED] ->
PAYMENT-SERVICE ->
[PAYMENT-REFUNDED] ->
CINEMA-SERVICE ->
[UNBLOCKED_CINEMA_SEATS] ->
BOOKING-SERVICE ->
[BOOKING_CANCELLED]
~~~~~~~~~~~~~~~~~~~~