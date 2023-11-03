CREATE TABLE IF NOT EXISTS "Profile" (
	"profileID"	INTEGER NOT NULL UNIQUE,
	"username"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	"profilePicture"	BLOB,
	"bio"	TEXT,
	"friends"	INTEGER,
	"weekplan"	INTEGER,
	PRIMARY KEY("profileID") --autoincrement 
);  
-- CREATE TABLE "User" (
--     "userID"	INTEGER NOT NULL,
--     "username"	TEXT NOT NULL UNIQUE,
--     "password"	TEXT NOT NULL,
--     "profileID"	INTEGER NOT NULL,
--     FOREIGN KEY("profileID") REFERENCES "Profiles"("profileID"),
--     PRIMARY KEY("userID") --autioincrement
-- );

CREATE TABLE IF NOT EXISTS "Recipe" (
	"recipeID"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"recipeName"	TEXT NOT NULL,
	"recipePicture"	BLOB,
	"timeEstimate"	INTEGER,
	"difficulty"	TEXT,
	"feedsPeople"	INTEGER,
	"directions"	TEXT NOT NULL,
	"author"	TEXT NOT NULL,
	FOREIGN KEY("author") REFERENCES "Profile"("username")
	--PRIMARY KEY("recipeID") --autoincrement
);

CREATE TABLE IF NOT EXISTS "RecipeCollection" (
	"recipeCollectionID"	INTEGER NOT NULL,
	"recipeCollectionName" TEXT NOT NULL,
	"recipeID"	INTEGER NOT NULL,
	"ownerID"	INTEGER NOT NULL,
	"date"	TEXT,
	"subscriberID" INTEGER,
	FOREIGN KEY("recipeID") REFERENCES "Recipe"("recipeID"),
	FOREIGN KEY("ownerID") REFERENCES "Profile"("profileID"),
	FOREIGN KEY("subscriberID") REFERENCES "Profile"("profileID"),
	PRIMARY KEY("recipeCollectionID") --autoincrement
);

CREATE TABLE IF NOT EXISTS "Ingredient" (
	"ingredientName" TEXT NOT NULL,
	"ingredientAmount" INTEGER NOT NULL,
	"ingredientUnit" TEXT NOT NULL,
	"recipeID" INTEGER NOT NULL,
	FOREIGN KEY("recipeID") REFERENCES "Recipe"("recipeID")
);