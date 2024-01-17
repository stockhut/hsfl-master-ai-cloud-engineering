--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1 (Debian 16.1-1.pgdg120+1)
-- Dumped by pg_dump version 16.1 (Debian 16.1-1.pgdg120+1)

-- Started on 2024-01-17 18:21:45 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 16402)
-- Name: Ingredient; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Ingredient" (
    "ingredientName" text NOT NULL,
    "ingredientAmount" double precision NOT NULL,
    "ingredientUnit" text NOT NULL,
    "recipeID" integer NOT NULL
);


ALTER TABLE public."Ingredient" OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16394)
-- Name: Recipe; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Recipe" (
    "recipeID" integer NOT NULL,
    "recipeName" text NOT NULL,
    "timeEstimate" integer NOT NULL,
    difficulty text,
    "feedsPeople" integer NOT NULL,
    directions text NOT NULL,
    author text NOT NULL
);


ALTER TABLE public."Recipe" OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16393)
-- Name: Recipe_recipeID_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Recipe_recipeID_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Recipe_recipeID_seq" OWNER TO postgres;

--
-- TOC entry 3371 (class 0 OID 0)
-- Dependencies: 216
-- Name: Recipe_recipeID_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Recipe_recipeID_seq" OWNED BY public."Recipe"."recipeID";


--
-- TOC entry 215 (class 1259 OID 16384)
-- Name: accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accounts (
    name text NOT NULL,
    email text NOT NULL,
    passwordhash bytea NOT NULL
);


ALTER TABLE public.accounts OWNER TO postgres;

--
-- TOC entry 3211 (class 2604 OID 16397)
-- Name: Recipe recipeID; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Recipe" ALTER COLUMN "recipeID" SET DEFAULT nextval('public."Recipe_recipeID_seq"'::regclass);


--
-- TOC entry 3365 (class 0 OID 16402)
-- Dependencies: 218
-- Data for Name: Ingredient; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Ingredient" ("ingredientName", "ingredientAmount", "ingredientUnit", "recipeID") FROM stdin;
Große Kartoffeln	8	Stück	1
Gemischtes Hackfleisch	500	gramm	1
Sahne	400	ml	1
Gemüsebrühe	400	ml	1
Butter	5	EL	1
Mehl	5	EL	1
Zwiebel	2	Stück	1
Salz und Pfeffer	1	nach Gefühl	1
Gratinkäse	250	gramm	1
Feta	250	gramm	1
Mehl	450	gramm	2
Olivenöl	5	EL	2
Eigelbe	2	Anzahl	2
Salz	2	Prisen	2
lauwarmes Wasser	200	ml	2
Magerquark	500	gramm	2
Schmand	500	gramm	2
Zwiebel	2	Stück	2
Salz und Pfeffer	1	nach Gefühl	2
Gratinkäse	125	gramm	2
Frühlingszwiebeln	4	Stangen	2
Feta	125	gramm	2
Speck und Schinken	175	gramm	2
Toppings nach Wahl	1	nach Geschmack	2
Frischkäse	400	gramm	3
Datteln	6	Stück	3
Olivenöl	1	Schuss	3
Joghurt	100	gramm	3
Knoblauchzehen	2	Stück	3
Agavendicksaft	1	EL	3
Salz und Pfeffer	1	nach Gefühl	3
Petersilie	1	nach Gefühl	3
Mehl	500	gramm	4
lauwarmes Wasser	250	ml	4
Hefe	1	Würfel	4
Zucker	1	Prise	4
Öl	2	EL	4
Salz	1	TL	4
Suppenhuhn	1	Stück	5
Lauchstange	1	Stück	5
Zwiebel	1	Stück	5
Zucker	1	Prise	5
Knoblauchzehen	4	Stück	5
Karotten	4	Stück	5
Algen	1	Handvoll	5
Schweinefilet	500	gramm	5
Sojasouce	100	ml	5
Sake	100	ml	5
Zucker	1	EL	5
Frühlingszwiebel	1	Stück	5
Ingwerwurzel	3	cm	5
Eier	3	Stück	5
Ramen-Nudeln	500	gramm	5
Sojasprossen	100	gramm	5
Salz und Pfeffer	3	Prise	5
\.


--
-- TOC entry 3364 (class 0 OID 16394)
-- Dependencies: 217
-- Data for Name: Recipe; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Recipe" ("recipeID", "recipeName", "timeEstimate", difficulty, "feedsPeople", directions, author) FROM stdin;
1	Neles Kartoffelauflauf mit Hackfleisch und Feta	60	schwer	4	Eine große Auflaufform mit etwas Butter einfetten. Ofen auf 200 Grad Umluft vorheizen. Kartoffeln in circa 5mm dicke Scheiben schneiden und in der Form gleichmäßig auslegen. Die Auflaufform bereits in den Ofen schieben,damit die Kartoffeln vorgaren. Zwiebeln klein hacken und mit der Butter in einen kleinen Topf andünsten. In einer Pfanne die zweite gehackte Zwiebel andünsten und danach das Hackfleisch hinzugeben und großzügig mit Salz und Pfeffer würzen. Während das Fleisch brät die 5 EL Mehl nach und nach zu der Butter und den Zwiebeln in den Topf geben und mit einem Schneebesen verrühren. Die Sahne und die Gemüsebrühe zugeben wenn es klumpt. Die entstandene Mehlschwitze mit Salz und Pfeffer und eventuell Muskatnuss würzen. Das gebratene Hackfleisch in die Auflaufform über die Kartoffeln geben. Gefolgt von der Mehlschwitze und den Feta zerbröseln oder zerschneiden und ebenfalls auf den Auflauf geben und den Gratinkäse darüber streuen. Alles im Ofen bei Umluft 200 Grad 45 Minuten lang backen	nele
2	Neles Fkammkuchen ohne Hefe	30	leicht	4	Mehl, zwei Eigelbe, lauwarmes Wasser 5El Öl und zwei Prisen Salz gut miteinander vermengen und Kneten. Den Teig 30 Minuten im Kühlschrank ruhen lassen.Magerquark und Schmand miteinander verrühren und eine Zwibel kleinhacken und dazu geben. Mit Salz und Pfeffer gut würzen. Ofen auf 250Grad Ober-Unterhitze vorheizen. Das Blech beim Vorheizen im Ofen lassen, damit der Flammkuchen krosser wird. Die restlichen Toppings, wie Schinkenscheiben und Frühlingszwiebeln, Feta und die zweite Zwiebel schneiden. Den Flammkuchenteig halbieren. EIne hälfte eintspricht einem Backblech. Teig dünn ausrollen. Die Magerquark-Schmand-Mischung auf dem Teig verteilen und die Toppings verteilen. Mit Gratinkäse abrunden. Den belegten Flammkuchen für 10-12 Minuten bei 250Grad Ober-Unterhitze backen.	nele
3	Dattel-Dip	15	leicht	1	Datteln halbieren und in kleine Stückchen schneiden. Knoblauch fein hacken. Frischkäse, Joghurt, Öl und Agavendicksaft in einer Schüssel verrühren. Datteln, Knoblauch und Kräuter dazugeben und unterrühren. Mit Salz und Pfeffer abschmecken.	nele
4	Pizza	75	schwer	2	250ml lauwarmes Wasser in einen Messbecher füllen. Hefe hineinbröseln und mit einer Prise Zucker und Salz verrühren. 10-15 Minuten gehen lassen.Das Mehl und Salz in eine Schüssel geben. Flüssigkeit und Öl über das Mehl geben und mit den Knethaken des Handrührgeräts mindestens 5 Minuten kneten (von Hand mindestens 10 Minuten lang kneten). Zum Schluss mit den Händen noch einmal wenige Minuten weiterkneten, bis der Teig geschmeidig ist. Die Teigschüssel mit einem Tuch abdecken und an einem warmen Ort (ca. 35°C) etwa 40 Minuten gehen lassen.Teig halbieren und auf bemehlter Arbeitsfläche rund (Ø ca. 28-30 cm)  ausrollen. Ofen auf 240 Grad (Umluft: 220) vorheizen. Zwei Standardbleche (38x45 cm) mit Backpapier belegen. Teige auf die Backbleche legen und etwas in die Ränder zurechtdrücken. Jetzt den Pizzateig nach Belieben mit Tomatensoße und verschiedensten Zutaten belegen. Pizza im vorgeheizten Ofen etwa 15 Minuten backen.	test
5	Japanische Nudelsuppe mit Hühnerbrühe und Schweinefilet-Ramen	360	schwer	4	Das Suppenhuhn mit Lauch, Zwiebel, drei bis vier Zehen Knoblauch, Ingwerwurzel, einer Prise Salz, drei bis vier Karotten und den Algen im kalten Wasser ansetzen. So viel Wasser hinzufügen, dass das Huhn komplett bedeckt ist. Die Topfgröße so wählen, dass alles plus zwei bis drei Liter Wasser gut hineinpasst, ein fünf Liter Topf ist ideal. Alles langsam zum Köcheln bringen und mindestens 3 - 4 Stunden köcheln lassen. Noch leckerer wird die Brühe bei Kochzeiten von 6 - 8 Stunden. Sollte Schaum entstehen, kann man diesen abschöpfen, muss man aber nicht. Die Brühe sollte nicht zu stark kochen, nur leicht simmern. Je länger diese Brühe kocht, desto besser. Die Brühe anschließend abseihen. Das Schweinefilet in einer Pfanne kurz von allen Seiten scharf anbraten, bis es leicht bräunlich ist. Nicht zu lange braten, nur leicht bräunen! Anschließend in einen Topf legen und mit Sojasauce (ich verwende Kikkoman weil sie natürlich gebraut ist) und 50 bis 100 ml Reiswein oder Sake (ich verwende chinesischen Reiswein) aufgießen. Den Zucker, in Scheiben geschnittene Frühlingszwiebel (mit dem Grün und nur wenig von der Zwiebel) und geriebene Ingwerwurzel dazugeben. Mit etwas Wasser aufgießen, sodass das Filet fast ganz mit Flüssigkeit bedeckt ist. Die Flüssigkeit zum Köcheln bringen. Auch hier wieder nur leicht köcheln lassen. Nach 40 Minuten das Filet aus der Flüssigkeit nehmen und zur Seite legen. Das Schweinefilet, bevor es in die Suppe kommt, in ca. 2 - 3 mm dicke Scheiben schneiden. Die Eier hart kochen und pellen. Anschließend in den Sud geben und 10 Minuten mitköcheln lassen. Immer wieder wenden, sodass sie gleichmäßig vom Sud gebräunt werden. Wenn sie fertig sind, halbieren und zur Seite legen. Nudeln in einem Topf kochen.Wenn die Hühnerbrühe fertig ist, diese mit der Brühe vom Fleisch in einen Topf geben und mit einigen Esslöffeln Sojasauce und noch mal einem kleinen Schuss Reiswein oder Sake abschmecken. Eventuell noch nachsalzen, aber eigentlich müsste die Sojasauce genügend Salz liefern. Alles noch mal aufkochen lassen. Man kann auch noch Wasser dazugeben, je nachdem wie stark bzw. verdünnt man die Suppe haben will. Ich lasse die Brühe pur ohne Zugabe von Wasser. Wenn die Suppe kocht, sollten für den nächsten Schritt alle anderen Zutaten schon fertig sein, vor allem auch die Nudeln.Die frischen Sojasprossen mit heißem Wasser in einem Sieb überbrühen. Das Grün von zwei Frühlingszwiebeln in Ringe schneiden. Aus dem Noriblatt kleine Streifen (ca. 2 x 3 cm) schneiden.(https://www.chefkoch.de/rezepte/1804511291817891/Japanische-Nudelsuppe-mit-Huehnerbruehe-und-Schweinefilet-Ramen.html)	test
\.


--
-- TOC entry 3362 (class 0 OID 16384)
-- Dependencies: 215
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.accounts (name, email, passwordhash) FROM stdin;
nele	nele@example.org	\\x243261243130244e384d62796e483168434c6278657a6142754b73646543324a2f366f743030704b376a616f7733352e5668697a7955716775795053
test	test@example.org	\\x2432612431302437366d7664455a504e577761335a4245627746616675322f752e5274754b4a766a4c6e62735630564a64745973305a574e35425479
\.


--
-- TOC entry 3372 (class 0 OID 0)
-- Dependencies: 216
-- Name: Recipe_recipeID_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Recipe_recipeID_seq"', 5, true);


--
-- TOC entry 3217 (class 2606 OID 16401)
-- Name: Recipe Recipe_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Recipe"
    ADD CONSTRAINT "Recipe_pkey" PRIMARY KEY ("recipeID");


--
-- TOC entry 3213 (class 2606 OID 16392)
-- Name: accounts accounts_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_email_key UNIQUE (email);


--
-- TOC entry 3215 (class 2606 OID 16390)
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (name);


--
-- TOC entry 3218 (class 2606 OID 16407)
-- Name: Ingredient Ingredient_recipeID_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Ingredient"
    ADD CONSTRAINT "Ingredient_recipeID_fkey" FOREIGN KEY ("recipeID") REFERENCES public."Recipe"("recipeID");


-- Completed on 2024-01-17 18:21:45 UTC

--
-- PostgreSQL database dump complete
--

