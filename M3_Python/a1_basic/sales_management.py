from models import Transaction
from book_management import books
from customer_management import customers

sales = []

def sell_book(book_title, customer_name, quantity):
    try:
        quantity = int(quantity)
        if quantity <=0:
            raise ValueError("Quantity should not be less than 0")
        customer = None

        for entries in customers:
            if entries.name.lower() == customer_name.lower():
                customer = entries
                break
        if not customer:
            print("Customer not found")
            return

        for book in books:
            if book.title.lower() == book_title.lower():
                if book.quantity>= quantity:
                    book.quantity -= quantity
                    sales.append(Transaction(book_title, customer_name, customer.email, customer.phone, quantity))
                    print(f"{quantity} {book.title} sold to {customer_name}")
                    return
                else:
                    print("Not enough stock")
        print("Book not found")
    except ValueError as e:
        print(f"Error: {e}")

def list_sales():
    if not sales:
        print("No sales available")
        return
    for sale in sales:
        print(sale.__str__())
                