Restaurant Reservation System
=============================

## Take Home Interview Question

Your task is to implement a simple reservation system for a restaurant. 
The restaurant can only accommodate up to 4 reservations per hour. You 
will need to provide functionalities to add, view, and cancel reservations.

## Requirements
### Add Reservation:
* Input: Customer's name, date and time of the reservation.
* Output: Confirmation of the reservation if the time slot is available or an appropriate 
  message if it is full.

## View Reservations:
* Input: Specific date.
* Output: List of all reservations for that date, organized by time.

## Cancel Reservation:
* Input: Customer's name and the date and time of the reservation.
* Output: Confirmation of the cancellation or an appropriate message if no such reservation exists.

## Constraints
* The restaurant operates from 10:00 AM to 10:00 PM.
* Each hour from 10:00 AM to 10:00 PM can only have up to 4 reservations.

## Deliverables
* Implement the reservation system in a programming language of your choice.
* No need to leverage a database, utilize a map or file to manage reservations

## Follow-Ups
* ***Easy Follow up:*** Implement a feature to automatically cancel reservations that are not confirmed
  within a certain period.

* ***Medium Follow Up:*** The restaurant has expanded into a franchise with multiple locations. 
  What if the reservation system were to be enhanced with adding reservations for a specific 
  location, how would you enhance your current implementation?

    * How would you modify your data structures to accommodate multiple locations?

    * How would you scale your system to support hundreds of locations?

    * How would you handle reservations that need to span multiple days or involve multiple locations


* ***Hard Follow Up:*** Enhance the reservation system to support dynamic capacity management and
  peak time handling. The system should allow different reservation capacities for different 
  time slots and handle peak times efficiently.
    
  * ***Dynamic Reservation Capacity:*** Allow the restaurant to set different reservation capacities 
    for different time slots. For example, during peak hours (e.g., 6:00 PM - 8:00 PM), the 
    restaurant might want to reduce the number of available reservations per hour to ensure a 
    better dining experience.
  
  * ***Waitlist Management:*** Implement a waitlist feature for fully booked time slots. If a reservation 
    is canceled, customers on the waitlist should be notified and given the option to confirm the 
    reservation.
   
  * ***Real-Time Availability Updates:*** Ensure that reservation availability is updated in real-time.
    If two users attempt to book the same time slot simultaneously, the system should handle this
    gracefully and ensure data consistency.
