# Learning Management System Backend
A robust backend system for managing educational resources including notes, books, previous year questions, and virtual classrooms. Built with Go and featuring comprehensive API endpoints for authentication, resource management, and classroom coordination.
## 🚀 Features
- **Authentication System**
  - User signup with role-based access
  - Secure signin with session management
  - HTTP-only cookie implementation for security
- **Notes Management**
  - Upload and store educational notes
  - Fetch notes by subject and year
  - Stream-wise subject organization
  - Secure file storage using AWS S3
- **Previous Year Questions (PYQ) Management**
  - Upload and store previous year question papers
  - Fetch PYQs by subject and year
  - Stream-wise organization
  - Secure file storage using AWS S3
- **Virtual Classroom Management**
  - Create and manage virtual classrooms
  - Subject-wise classroom organization
  - Stream and year-based filtering
  - Integration with video conferencing platforms
  - User-specific classroom tracking
## 🛠️ Tech Stack
- **Backend Framework**: Go
- **Database**: PostgreSQL
- **Storage**: AWS S3
- **Authentication**: Custom implementation with secure cookies
- **API Documentation**: Postman Collection
## 📝 API Documentation
### Authentication APIs
#### Sign Up
```http
POST /signup
```
**Request Body**:
```json
{
    "name": "User Name",
    "email": "user@example.com",
    "password": "password",
    "Role": "User"
}
```
#### Sign In
```http
POST /signin
```
**Request Body**:
```json
{
    "email": "user@example.com",
    "password": "password"
}
```
### Notes APIs
#### Upload Notes
```http
POST /api/v1/file-upload
```
**Query Parameters**:
- `filename`: Name of the file
- `year`: Academic year
- `subjectCode`: Subject code
#### Fetch Notes
```http
POST /api/v1/view-notes
```
**Request Body**:
```json
{
    "year": 2,
    "subjectCode": "UCS124"
}
```
### Previous Year Questions APIs
#### Upload PYQ
```http
POST /api/v1/file-upload
```
**Query Parameters**:
- `filename`: Name of the file
- `year`: Academic year
- `subjectCode`: Subject code

**Response**:
```json
{
    "message": "File ready to be uploaded",
    "upload_url": "https://[bucket-name].s3.[region].amazonaws.com/[file-path]"
}
```

#### Fetch PYQs
```http
POST /api/v1/view-pyqs
```
**Request Body**:
```json
{
    "year": 2,
    "subjectCode": "UCS124"
}
```

#### Fetch Subjects for PYQs
```http
POST /api/v1/fetch-subjects-pyqs
```
**Query Parameters**:
- `year`: Academic year
- `stream`: Stream code (e.g., "COE")

**Response**:
```json
{
    "subjects": ["UCS124"]
}
```

### Classroom APIs
#### Create Class
```http
POST /api/v1/create-class
```
**Request Body**:
```json
{
    "name": "Subject-Class",
    "stream": "COE",
    "subject": "Subject Name",
    "year": 2,
    "URL": "https://meet.google.com/xxx-xxxx-xxx",
    "createdDate": 1732267731862
}
```
#### Fetch User's Classes
```http
GET /api/v1/fetch-class
```
Returns all classes created by the authenticated user.
#### Fetch Classes by Subject
```http
POST /api/v1/fetch-class-subject
```
**Request Body**:
```json
{
    "subject": "Subject Name"
}
```
## 🚦 Environment Variables
```env
API_V1URL=<your-api-url>
AWS_S3_BUCKET=<your-s3-bucket>
AWS_REGION=<aws-region>
# Add other required environment variables
```
## 🛠️ Installation & Setup
1. Clone the repository
```bash
git clone <repository-url>
```
2. Install dependencies
```bash
go mod download
```
3. Set up environment variables
```bash
cp .env.example .env
# Fill in your environment variables
```
4. Run the server
```bash
go run main.go
```
## 🔒 Security Features
- HTTP-only cookies for session management
- Secure password hashing
- CORS configuration
- Request validation
- Access control headers

## 👥 Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
## 📞 Support
For support, email [your-email@example.com] or raise an issue in the repository.
