CREATE TABLE IF NOT EXISTS "Recipe" (
	"recipeID"	SERIAL PRIMARY KEY,
	"recipeName"	TEXT NOT NULL,
	"timeEstimate"	INTEGER NOT NULL,
	"difficulty"	TEXT,
	"feedsPeople"	INTEGER NOT NULL,
	"directions"	TEXT NOT NULL,
	"author"	TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "Ingredient" (
	"ingredientName" TEXT NOT NULL,
	"ingredientAmount" FLOAT NOT NULL,
	"ingredientUnit" TEXT NOT NULL,
	"recipeID" INTEGER NOT NULL,
	FOREIGN KEY("recipeID") REFERENCES "Recipe"("recipeID")
);