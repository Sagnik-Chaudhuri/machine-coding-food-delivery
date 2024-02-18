# ClearFood (machine coding food delivery)

ClearFood is a command-line application that simulates a virtual food delivery system. It allows users to register, rate restaurants, place orders, and view order history. The application uses an in-memory database and does not include a graphical user interface.

This Project is a very basic implementation of this [Public Repository](https://github.com/nSpider/machine-coding-food-delivery-app/tree/main). Given the time limit constraints, this is what we both could come up with. 
Please feel free to commit to any repo and add more functionalities, as per your choice of poison (Read: Language!).
Details are mentioned in this README below as well:


## Functional Requirements

### Register Restaurant

```python
Register_restaurant(restaurant_name, list of serviceable pin-codes, food item name, food item price, initial quantity)
```

- Restaurant owners can register their restaurants with specific parameters.
- They can increase the quantity of a food item using `update_quantity(restaurant name, quantity to Add)`.

### User Interaction

```python
rate_restaurant(restaurant name, rating, comment)  # comment optional
```

- Users can rate any registered restaurant on a scale of 1 (Lowest) to 5 (Highest) with an optional comment.

### Listing and Ordering

```python
show_restaurant(rating/price)  # Based on rating or price
place_order(restaurant name, quantity)
```

- Users can view a list of serviceable restaurants with food item names and prices in descending order based on either rating or price.
- Users can place an order from any restaurant with a specified quantity.

### Order History

```python
order_history(user)
```

- For a given user, retrieve the order history.

## Sample Flow

```python
# User Registration
register_user("Pralove", "M", "phoneNumber-1", "HSR")
register_user("Nitesh", "M", "phoneNumber-2", "BTM")
register_user("Vatsal", "M", "phoneNumber-3", "BTM")
login_user("phoneNumber-1")

# Restaurant Registration
register_restaurant("Food Court-1", "BTSIM/HSR", "NI Thali", 100, 5)
register_restaurant("Food Court-2", "BTM/pincode-2", "Burger", 120, 3)
register_restaurant("Food Court-3", "HSR", "SI Thali", 150, 1)

# User Interactions
login_user("phoneNumber-2")
show_restaurant("Price")  # Output: Food Court-2, Burger | Food Court-1, NI Thali
place_order("Food Court-1", 2)  # Output: Order Placed Successfully.
place_order("Food Court-2", 7)  # Output: Cannot place order

rate_restaurant("Food Court-2", 3, "Good Food")
rate_restaurant("Food Court-1", 5, "Nice Food")
show_restaurant("rating")  # Output: Food Court-1, NI Thali | Food Court-2, Burger

login_user("phoneNumber-1")
update_quantity("Food Court-2", 5)  # Output: Food Court-2, BTM, Burger - 8
update_location("Food Court-2", "BTM/HSR")  # Output: Food Court-2, "BTM/HSR", Burger - 8
```

## Note

- The application uses an in-memory database.
- There is no graphical user interface.
- A driver class can be used to simulate all the operations.
