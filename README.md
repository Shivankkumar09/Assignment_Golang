# 🏥 Clinic Management System - Golang Web App

This is a simple clinic management backend system built with Golang, Gin, and PostgreSQL. It provides separate functionalities for **receptionists** and **doctors**, including role-based authentication, patient management, and secure API access using JWT tokens.

---

## 🚀 Features

- 🧾 User Signup & Login (with role: `receptionist` or `doctor`)
- 🔐 JWT-based Authentication & Role Authorization
- 🩺 Receptionist:
  - Register new patients
  - View all patients
  - Update and delete patients
- 👨‍⚕️ Doctor:
  - View all patients
  - Update patient medical data
- ✅ Clean code with modular folder structure
- 🧪 Unit test cases for authentication
- 📄 Postman collection for API testing

---

## 🛠️ Tech Stack

- **Backend**: Golang, Gin
- **Database**: PostgreSQL (via Railway free hosting)
- **Auth**: JWT, bcrypt
- **Testing**: Go’s built-in testing package (`testing`)
- **API Docs**: Postman collection

---

## 📁 Directory Structure

```
Assignment_Golang/
│
├── controllers/        # Auth & Patient logic
├── middlewares/        # JWT auth middleware
├── models/             # DB models
├── routes/             # API routes
├── services/           # Logic for DB + auth
├── utils/              # Token utils
├── config/             # DB connection
├── test/               # Unit test cases
└── main.go             # Entry point
```

---

## 🔐 Auth Collection

The **Auth collection** handles user signup and login. Users have a `username`, `password`, and `role` (`receptionist` or `doctor`). Passwords are securely hashed with bcrypt. A JWT token is issued on login and required for accessing protected routes.

### 🔁 Routes

- **POST** `/signup`
```json
{
  "username": "testuser1",
  "password": "12345678",
  "role": "receptionist"
}
```

- **POST** `/login`
```json
{
  "username": "testuser1",
  "password": "12345678"
}
```

Returns:
```json
{
  "token": "your_jwt_token_here"
}
```

Pass this token in headers for protected routes:
```
Authorization: your_jwt_token_here
```

---

## 👩‍💼 Receptionist Collection

Receptionists can manage patient records.

### 🔁 Routes
- **POST** `/receptionist/patients` – Add new patient
- **GET** `/receptionist/patients` – View all patients
- **PUT** `/receptionist/patients/:id` – Update patient
- **DELETE** `/receptionist/patients/:id` – Delete patient

---

## 👨‍⚕️ Doctor Collection

Doctors can view and update patient details.

### 🔁 Routes
- **GET** `/doctor/patients` – View all patients
- **PUT** `/doctor/patients/:id` – Update patient

---

## ✅ Unit Testing

Unit tests are placed under `/test`. Sample tested: `auth_test.go`  
To run tests:
```bash
go test ./test
```

---

## 🧪 Environment Variables

Create a `.env` file:
```
DB_HOST=your_host
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_db
JWT_SECRET=your_jwt_secret
```

---

## ▶️ Run Locally

```bash
git clone https://github.com/your-username/Assignment_Golang
cd Assignment_Golang

go mod tidy
go run main.go
```

API will be running on: `http://localhost:8080`

---

## 🌐 Deployment

- Use **Render** to deploy the backend
- Use **Railway** for free managed PostgreSQL
- Add environment variables in Render dashboard

---

## 💡 Creative Ideas to Extend

- Add **patient appointment scheduling**
- Integrate **email notifications**
- Create a **React frontend**
- Add **admin panel** to view analytics
- Implement **pagination & filtering**

---

## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first.

---

## 🧑‍💻 Author

**Shivank Kumar**  
GitHub: [shivankkumar](https://github.com/shivankkumar)

