-- Add Pizza Margharita to Dishes
INSERT INTO Dishes (name, img) VALUES ('Pizza margharita', 'example.jpg');

-- Add ingredients
INSERT INTO Ingredients (id, ingredientName) VALUES (1, 'Dough');
INSERT INTO Ingredients (id, ingredientName) VALUES (2, 'tomato');
INSERT INTO Ingredients (id, ingredientName) VALUES (3, 'cheese');
INSERT INTO Ingredients (id, ingredientName) VALUES (4, 'olive oil');

-- Associate ingredients with dish
INSERT INTO DishesToIngredients (dishName, ingredientID) VALUES ('Pizza margharita', 1);
INSERT INTO DishesToIngredients (dishName, ingredientID) VALUES ('Pizza margharita', 2);
INSERT INTO DishesToIngredients (dishName, ingredientID) VALUES ('Pizza margharita', 3);
INSERT INTO DishesToIngredients (dishName, ingredientID) VALUES ('Pizza margharita', 4);

-- Add feedback
INSERT INTO Feedbacks (dishName, nationality, like, dislike) VALUES ('Pizza margharita', 'Kazakhstan', 4, 9);
