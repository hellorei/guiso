# Guiso Framework

## Overview

Guiso is a customizable end-to-end framework for building high-performance, containerized applications. It leverages a powerful stack of modern technologies to provide a seamless development experience and efficient deployment pipeline.

**Version**: 0.0.1

**Documentation**: [https://www.guiso.io](https://www.guiso.io)

![Screenshot](https://raw.githubusercontent.com/hellorei/guiso/main/public/images/screenshot-guiso-help-cmd.png)

## Tech Stack

- Backend: Go with Echo framework
- Frontend: HTMX, Alpine.js, Templ
- Styling: Tailwind CSS
- Infrastructure: Terraform
- Containerization: Docker
- CI/CD: GitHub Actions

## Features

- High-performance server-side rendering with Go and Templ
- Interactive UIs with HTMX and Alpine.js
- Responsive design with Tailwind CSS
- Infrastructure as Code with Terraform
- Containerized deployment with Docker
- Automated CI/CD pipeline with GitHub Actions
- Internationalization support
- SEO optimization

## Available Modules

- [x] Logger
- [x] I18n / Internationalization
- [ ] Auth
- [x] WebSockets
- [x] Markdown Rendering
- [x] SEO
- [ ] Google Analytics

## Available Infrastructure Templates (via Terraform)

### Google Cloud Platform

- [x] Cloud Run
- [x] Artifact Registry
- [ ] Cloud Storage
- [ ] Cloud SQL
- [ ] Cloud Firestore
- [ ] Cloud Functions

### AWS

- [ ] Fargate
- [ ] Elastic Container Registry
- [ ] S3
- [ ] RDS
- [ ] DynamoDB
- [ ] Lambda

## Project Structure

```
.
├── .vscode/
├── app/
│   ├── framework/
│   ├── middlewares/
│   ├── routes/
│   ├── scripts/
│   ├── templ/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── ts/
│   │   └── types/
│   └── utils/
├── public/
├── terraform/
├── Dockerfile
├── guisofile
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.16+
- Node.js and npm
- Docker
- Terraform

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/guiso.git
   cd guiso
   ```

2. Install dependencies:
   ```
   guiso install
   ```

### Development

Start the development server with hot reload:

```
guiso dev
```

For a clean start:

```
guiso dev-clean
```

### Building

Build the application:

```
guiso build
```

### Testing

Run tests:

```
guiso test
```

### Deployment

1. Build the Docker image:

   ```
   guiso docker-build
   ```

2. Push the Docker image:

   ```
   guiso docker-push
   ```

3. Apply Terraform changes:
   ```
   guiso tf-apply
   ```

## Available Commands

Run `guiso help` or `guiso help` to see a list of available commands. Some key commands include:

- `guiso build`: Build the application
- `guiso dev`: Start development server with hot reload
- `guiso test`: Run tests
- `guiso docker`: Build and run Docker container
- `guiso tf-apply`: Apply Terraform changes

## CI/CD Pipeline

The project uses GitHub Actions for continuous integration and deployment. The pipeline includes:

1. Testing
2. Building Docker images
3. Deploying to Google Cloud Run using Terraform

For more details, refer to the `ci-cd-documentation.md` file in the project root.

## Contributing

Contributions are welcome! Please read our contributing guidelines before submitting pull requests.

## License

[MIT License](LICENSE)

## Support

For support, please open an issue on the GitHub repository or contact the maintainers at support@guiso.io.
