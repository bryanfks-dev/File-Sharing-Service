# File-Sharing-Service

This is a simple file sharing service using go, echo web framework, and REST API, which only reads files from the server storage without database system.

**With this simple service, you can do:**

- Listing files in server storage.
- Save file to server storage.
- Download file from server storage.
- Delete file in the server storage.

Here are the list of usable endpoints.

## Usable Endpoints:

- **GET** `/api/v1/files` - List all files information.
- **POST** `/api/v1/savefile` - Save file to storage.
- **GET** `/api/v1/file/[filename]` - Get file information
- **GET** `/api/v1/file/[filename]/download` - Download file.
- **DELETE** `/api/v1/file/[filename]/delete` - Delete file.

As, I mentioned, this project uses REST API, here is the **example of the response body**:<br>

```json
{
  "statusCode": 200,
  "message": "File uploaded successfully",
  "data": [
    {
      "fileName": "Untitled.jpg",
      "fileSize": 47152,
      "fileType": "Image",
      "uploadTime": "2024-10-11 01:23:31.5738559 +0700 +07",
      "fileURL": "/api/v1/files/Untitled.jpg"
    }
  ]
}
```

**Explaination:**

- **statusCode** - The response status code
- **message** - The message of reponse
- **data** - The data of response

For the data format, it usually contains file information which contain:

- **fileName** - The file name.
- **fileSize** - The file size in **byte** size.
- **fileType** - The type of file (image, video, document, etc..)
- **uploadTime** - The upload file time.
- **fileUrl** - The url of file information.

## Installation:

To start the service,
first make sure you already installed go compiler in your machine and simply run `go run cmd/api/main.go` in your project.

## Customization:

- **Save / Download File Dir**<br>
  To modify the file dir, navigate to `core/constant/dir.go`, and just change the `FILE_DIR` value.

- **Endpoints**<br>
  To add, delete, or modify endpoint, simply navigate to `internal/server/server.go`. To start writting the service logic or functions, navigate `internal/service` dir and create a new file.
