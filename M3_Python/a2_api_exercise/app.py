from flask import Flask, request, jsonify, abort
from dbconfig import initializeDB
import sqlite3

app = Flask(__name__)

DATABASE = "books.db"

def getDBConnection():
    connection = sqlite3.connect(DATABASE)
    connection.row_factory = sqlite3.Row
    return connection


#Post method to add new book
@app.route("/books", methods=["POST"])
def addBook():
    data = request.get_json()
    if not data or not all(key in data for key in ("title", "author", "published_year", "genre")):
        return jsonify({"error":"Invalid data", "message":"Missing required Fields"}), 400
    
    try:
        with getDBConnection() as connection:
            cursor = connection.cursor()
            cursor.execute("""INSERT INTO books (title, author, published_year, genre) VALUES (?,?,?,?)""", (data["title"], data["author"], data["published_year"], data["genre"]))
            connection.commit()
            id = cursor.lastrowid
            return jsonify({"message":"Book added successfully", "book_id": str(id)}), 201
    except sqlite3.Error as e:
        return jsonify({"error":"Database error", "message":str(e)}), 500
    
#Get method to get all books
@app.route("/books", methods=["GET"])
def getBooks():
    try:
        with getDBConnection() as connection:
            cursor = connection.cursor()
            cursor.execute("""SELECT * FROM books""")
            books = cursor.fetchall()
            return jsonify([dict(book) for book in books])
    except sqlite3.Error as e:
        return jsonify({"error":"Database error", "message":str(e)}), 500


#Get method to get book by id
@app.route("/books/<int:book_id>", methods=["GET"])
def getBookByID(book_id):
    try:
        with getDBConnection() as connection:
            cursor = connection.cursor()
            cursor.execute("""SELECT * FROM books WHERE id = ?""", (book_id,))
            book = cursor.fetchone()
            if not book:
                return jsonify({"error":"Not found", "message":"Book not found"}), 404
            return jsonify(dict(book))
    except sqlite3.Error as e:
        return jsonify({"error":"Database error", "message":str(e)}), 500


#Update book
@app.route("/books/<int:book_id>", methods=["PUT"])
def updateBook(book_id):
    data = request.get_json()
    validGenres = ["Fiction", "Non-Fiction", "Mystery", "Sci-Fi"]
    if not data or not all(key in data for key in ("title", "author", "published_year", "genre")):
        return jsonify({"error":"Invalid data", "message":"Missing required Fields"}), 400
    
    if data["genre"] not in validGenres or data["published_year"] < 0:
        return jsonify({"error":"Invalid data", "message":"Invalid genre or published year"}), 400
    try:
        with getDBConnection() as connection:
            cursor = connection.cursor()
            cursor.execute("""UPDATE books SET title = ?, author = ?, published_year = ?, genre = ? WHERE id = ?""", (data["title"], data["author"], data["published_year"], data["genre"], book_id))
            connection.commit()
            return jsonify({"message":"Book updated successfully"}), 200
    except sqlite3.Error as e:
        return jsonify({"error":"Database error", "message":str(e)}), 500


#Delete book
@app.route("/books/<int:book_id>", methods=["DELETE"])
def deleteBook(book_id):
    try:
        with getDBConnection() as connection:
            cursor = connection.cursor()
            cursor.execute("""DELETE FROM books WHERE id = ?""", (book_id,))
            connection.commit()
            return jsonify({"message":"Book deleted successfully"}), 200
    except sqlite3.Error as e:
        return jsonify({"error":"Database error", "message":str(e)}), 500

@app.route("/books/filter", methods=["GET"])
def filterBooks():
    genre = request.args.get("genre")
    author = request.args.get("author")
    try:
        with getDBConnection() as connection:
            cursor = connection.cursor()
            query = """SELECT * FROM books WHERE 1=1"""
            params = []
            if genre:
                query += " AND genre = ?"
                params.append(genre)
            if author:
                query += " AND author = ?"
                params.append(author)
            cursor.execute(query, params)
            books = cursor.fetchall()
            return jsonify([dict(book) for book in books])
    except sqlite3.Error as e:
        return jsonify({"error":"Database error", "message":str(e)}), 500
    
if __name__ == "__main__":
    app.run(debug=True)