-- 1.Find High-Spending Users


db.orders.aggregate([
  {
    $group: {
      _id: "$userId",
      totalSpent: { $sum: "$totalAmount" }
    }
  },
  {
    $match: {
      totalSpent: { $gt: 500 }
    }
  },
  {
    $lookup: {
      from: "users",
      localField: "_id",
      foreignField: "userId",
      as: "userDetails"
    }
  },
  {
    $unwind: "$userDetails"
  },
  {
    $project: {
      _id: 0,
      userId: "$_id",
      totalSpent: 1,
      name: "$userDetails.name",
      email: "$userDetails.email"
    }
  }
]);


-- 2.List Popular Products by Average Rating

db.products.aggregate([
  {
    $unwind: "$ratings" // Flatten the ratings array
  },
  {
    $group: {
      _id: "$productId",
      averageRating: { $avg: "$ratings.rating" },
      productName: { $first: "$name" },
      category: { $first: "$category" }
    }
  },
  {
    $match: {
      averageRating: { $gte: 4 } // Filter products with average rating >= 4
    }
  },
  {
    $sort: { averageRating: -1 } // Sort by average rating in descending order
  },
  {
    $project: {
      _id: 0,
      productId: "$_id",
      productName: 1,
      category: 1,
      averageRating: 1
    }
  }
]);


-- 3.Search for Orders in a Specific Time Range

db.orders.aggregate([
  {
    $match: {
      orderDate: {
        $gte: ISODate("2024-12-01T00:00:00Z"),
        $lt: ISODate("2025-01-01T00:00:00Z")
      }
    }
  },
  {
    $lookup: {
      from: "users",
      localField: "userId",
      foreignField: "userId",
      as: "userDetails"
    }
  },
  {
    $unwind: "$userDetails" // Flatten the userDetails array
  },
  {
    $project: {
      _id: 0,
      orderId: 1,
      orderDate: 1,
      totalAmount: 1,
      status: 1,
      userName: "$userDetails.name",
      userEmail: "$userDetails.email"
    }
  },
  {
    $sort: { orderDate: 1 } // Sort by orderDate in ascending order
  }
]);


-- 4.Update Stock After Order Completion

db.orders.aggregate([
  { $match: { orderId: "ORD001" } },
  { $unwind: "$items" },
  {
    $project: {
      productId: "$items.productId",
      quantity: "$items.quantity"
    }
  },
  {
    $merge: {
      into: "products",
      whenMatched: [
        { $set: { stock: { $subtract: ["$stock", "$quantity"] } } }
      ],
      whenNotMatched: "discard"
    }
  }
]);

-- 5.Find Nearest Warehouse

db.warehouses.aggregate([
  {
    $geoNear: {
      near: { type: "Point", coordinates: [-74.006, 40.7128] },  // New York City coordinates
      distanceField: "distance",
      maxDistance: 50000,  // 50 km in meters
      spherical: true,
      query: { products: "P001" }  // Filter warehouses that stock "P001"
    }
  },
  {
    $limit: 1  // Get the nearest warehouse
  }
]);


