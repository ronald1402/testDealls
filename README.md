# Technical Test: Login and Signup Functionality

### Prerequisites

- Go 1.19
- MySQL database

## Service Structure

### 1. Overview

Provide a brief overview of the service, its purpose, and its main features.

### 2. Architecture

Outline the overall architecture of the service, including key components and their interactions. You may include a diagram if applicable.

#### 2.1 Components
- **API Layer**: Handles HTTP requests and responses.
- **Service Layer**: Contains business logic and communicates with the data layer.
- **Data Layer**: Interacts with the database to perform CRUD operations.

### 3. Directory Structure

Describe the directory structure of the project, highlighting key files and directories.

## 4. Endpoints

List the available API endpoints, including methods, paths, and descriptions.

| Method | Endpoint         | Description               |
|--------|------------------|---------------------------|
| GET    | `/health`        | Health check endpoint     |
| POST   | `/api/v1/signup` | Create a new user account |
| POST   | `/api/v1/login`  | Authenticate a user       |


## Setting Up the Database

- Start mysql server on your Local

- To create the necessary tables for the application, execute the `ddl.sql` file. This file contains all the required SQL to create all the tables needed.

## Update Config
- Update mysql section with your local setting
```json
{
  "mysql": {
    "host": "127.0.0.1",
    "port":  3306,
    "schema":  "test",
    "username": "root",
    "password": ""
  }
}
```

### Updating Secrets

Make sure to update the secret as needed to ensure secure access to the application.


## Run Application
Run :
```bash
chmod +x run.sh  
```
then run :
```bash
./run.sh  
```

## Testing the Application with Postman

To test the application, you can use the provided Postman collection. Follow these steps to import the collection into Postman:

1. **Download the Postman Collection**:
   The Postman collection file is available in the project directory. Look for a file named `Test.postman_collection.json`.

2. **Open Postman**:
   Launch the Postman application on your computer.

3. **Import the Collection**:
    - Click on the **Import** button in the top-left corner of the Postman interface.
    - Select the **File** tab.
    - Drag and drop the Postman collection file, or click on **Upload Files** to select it from your file system.
    - Click on **Import** to add the collection to Postman.
    - 
4**Run the Requests**:
    - Navigate to the imported collection in the left sidebar.
    - Click on the individual requests to test the endpoints
    - Adjust any parameters or body data as necessary, then click **Send** to execute the request.