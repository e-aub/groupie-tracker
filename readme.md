*# Music Hub Web Application

The Music Hub Web Application is a dynamic platform designed to connect music enthusiasts with their favorite artists and tracks. Leveraging the power of Go for backend development and HTML/CSS for the frontend, this project offers a seamless and interactive user experience. It features artist pages, a welcoming home page, and utilizes Go templates for content rendering, all within a containerized Docker environment for streamlined deployment and scalability.

## Features

- **Artist Pages**: Display artist information with a Spotify-inspired design.
- **Home Page**: A landing page that showcases featured artists or content.
- **Responsive Design**: Utilizes CSS for a responsive layout that adapts to various screen sizes.
- **Docker Integration**: Containerized setup for easy deployment and scalability.

## Technology Stack

- **Backend**: Written in Go, utilizing the standard library for serving web content.
- **Frontend**: HTML pages styled with CSS, following a Spotify-inspired theme.
- **Deployment**: Dockerized environment for easy setup and deployment.

## Project Structure

- `cmd/`: Contains the Go entry point (`main.go`).
- `global/`: Go files for CRUD operations, error handling, server setup, and type definitions.
- `handlers/`: Go files for handling HTTP requests to different endpoints.
- `static/`: Contains CSS files and images for styling the web pages.
- `template/`: HTML templates for dynamically rendering content.
