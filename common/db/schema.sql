CREATE TABLE "Profiles" (
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
CREATE TABLE "Recipes" (
	"recipeID"	INTEGER NOT NULL,
	"recipeName"	TEXT NOT NULL,
	"recipePicture"	BLOB,
	"timeEstimate"	INTEGER,
	"difficulty"	TEXT,
	"feedsPeople"	INTEGER,
	"ingredients"	TEXT NOT NULL,
	"directions"	TEXT NOT NULL,
	"author"	TEXT NOT NULL,
	FOREIGN KEY("author") REFERENCES "Profiles"("username"),
	PRIMARY KEY("recipeID") --autoincrement
);
CREATE TABLE "RecipeCollection" (
	"recipeCollectionID"	INTEGER NOT NULL,
	"recipeCollectionName" TEXT NOT NULL,
	"recipeID"	INTEGER NOT NULL,
	"ownerID"	INTEGER NOT NULL,
	"date"	TEXT,
	"subscriberID" INTEGER,
	FOREIGN KEY("recipeID") REFERENCES "Recipes"("recipeID"),
	FOREIGN KEY("ownerID") REFERENCES "Profiles"("profileID"),
	FOREIGN KEY("subscriberID") REFERENCES "Profiles"("profileID"),
	PRIMARY KEY("recipeCollectionID") --autoincrement
);