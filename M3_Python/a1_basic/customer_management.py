
from models import Customer

customers = []

def add_customer(name, email, phone):
    customers.append(Customer(name, email, phone))

def list_customer():
    if not customers:
        print("No customers available")
        return
    for customer in customers:
        print(customer.__str__())
