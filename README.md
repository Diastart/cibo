# CiboCompass

CiboCompass is a mobile application designed to assist international students in Italy with understanding local cuisine from both nutritional and cultural perspectives. The project helps international students navigate unfamiliar restaurant menus by providing details on dish ingredients, cultural taste preferences, and community-based ratings.

---

## Technical Overview

CiboCompass is built with a Go backend and a React Native frontend, ensuring a responsive and seamless experience across mobile platforms.

---

## Backend

The backend is implemented in Go, providing three primary API routes:

- `GET /v1/dishes/{name}`: Fetches detailed information about a specific dish.  
- `POST /v1/dishes/{name}/feedback`: Handles user ratings and feedback submissions.  
- `GET /v1/healthcheck`: Checks the server's operational status.

**Key Technologies:**

- Language: Go  
- Routing: `httprouter`  
- Database: `mattn/go-sqlite3` (SQLite3 driver)  
- Data Structure: Normalized schema with `Dishes`, `Ingredients`, `DishesToIngredients`, and `Feedbacks` tables  
- Personalization: Uses a `nationality` header to provide culturally relevant recommendations

The server defaults to running on port 4000.

---

## Frontend

The frontend is a React Native application built with Expo SDK 53 and React 19, ensuring compatibility with both iOS and Android devices.

**Key Features & Technologies:**

- Framework: React Native  
- SDK: Expo SDK 53  
- React Version: React 19  
- Local Data Persistence: `AsyncStorage` for storing user preferences and rating history  
- Communication: JSON over HTTP with the Go backend

**User Interface Highlights:**

- Nationality setup modal on first launch  
- Real-time dish name search  
- Expandable ingredient lists in tabular format  
- Interactive star-based rating system  
- Dynamic cultural ratings aggregated from similar backgrounds  
- Dietary badges (Vegetarian, Vegan, Halal, Kosher, Gluten-Free) based on ingredient analysis  
- Smooth animations, robust error handling, and scroll state persistence

---

## Screenshots

_Add screenshots or screen recordings of the app here_

![App Screenshot 1](images/screenshot1.png)  
![App Screenshot 2](images/screenshot2.png)  
![App Screenshot 3](images/screenshot3.png)

---

## Future Enhancements

- Auto-suggestion for the search bar  
- "Recent Searches" section  
- Total number of ratings for each dish  
- Comment section for qualitative feedback (requires persistent user sessions on backend)

---

## AI Integration

AI tools supported the project lifecycle with:

- Report writing and documentation refinement  
- Persona generation for user-centered design  
- Survey optimization  
- Storyboard creation  
- Code debugging assistance

---

## Project Team

- [Artur Kadyrzhanov](https://github.com/arturkd)  
- [Dias Nursultan](https://github.com/nursultandias)
- [Beray Nil Atabey](https://github.com/NilAtabey)  
- [Liyu Jin](https://github.com/kk-Syuer)

---
